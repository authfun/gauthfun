package router

import (
	"errors"
	"net/http"

	"sort"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/authfun/gauthfun/consts"
	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
	service "github.com/authfun/gauthfun/service"
	rbac_errors "github.com/casbin/casbin/v2/errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func AddGroupOrganization(router *gin.Engine) {

	group := router.Group("/api/organizations")
	{
		group.GET("/:id", getOrganization)
		group.GET("/tree", tree)
	}
}

// @Summary Get organization
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Success 200 {object} model.OrganizationDetail
// @Router /api/organizations/{id} [get]
func getOrganization(c *gin.Context) {
	id := c.Param("id")

	var organization schema.Organization
	result := database.AuthDatabase.First(&organization, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	detail := model.OrganizationDetail{
		Roles: []schema.Role{},
		Users: []schema.User{},
	}

	copier.Copy(&detail, &organization)

	userIdChan, userIdErr := getUsersForOrganization(id, organization.TenantId)
	if err, existed := <-userIdErr; existed && !errors.Is(err, rbac_errors.ERR_NAME_NOT_FOUND) {
		InternalServerError(c, err)
		return
	}

	roleIdChan, roleIdErr := getRolesForOrganization(id, organization.TenantId)
	if err, existed := <-roleIdErr; existed && !errors.Is(err, rbac_errors.ERR_NAME_NOT_FOUND) {
		InternalServerError(c, err)
		return
	}

	userIds, roleIds := <-userIdChan, <-roleIdChan

	if len(userIds) > 0 {
		database.AuthDatabase.Find(&detail.Users, userIds)
	}

	if len(roleIds) > 0 {
		database.AuthDatabase.Find(&detail.Roles, roleIds)
	}

	c.JSON(http.StatusOK, detail)
}

// @Summary Get organization tree
// @Tags Organization
// @Accept json
// @Produce json
// @Param tenantId query string true "Tenant ID"
// @Success 200 {array} model.OrganizationNode
// @Router /api/organizations/tree [get]
func tree(c *gin.Context) {
	tenantId := c.Query("tenantId")
	var organizations []schema.Organization
	result := database.AuthDatabase.Where("tenant_id = ?", tenantId).Order("id").Find(&organizations)
	if result.Error != nil {
		InternalServerError(c, result.Error)
	}

	var nodes []*model.OrganizationNode
	linq.From(organizations).SelectT(func(organization schema.Organization) *model.OrganizationNode {
		node := model.OrganizationNode{
			Id:       organization.Id,
			ParentId: organization.ParentId,
			Name:     organization.Name,
			Children: []*model.OrganizationNode{},
		}
		return &node
	}).ToSlice(&nodes)

	tree := buildTree(nodes)
	c.JSON(http.StatusOK, tree)
}

func buildTree(nodes []*model.OrganizationNode) []model.OrganizationNode {

	nodeMap := map[string]*model.OrganizationNode{}
	for _, menu := range nodes {
		nodeMap[menu.Id] = menu
	}

	resultMap := map[string]*model.OrganizationNode{}
	for _, node := range nodeMap {
		if node.ParentId == nil {
			resultMap[node.Id] = node
		} else {
			if nodeMap[*node.ParentId] == nil {
				resultMap[node.Id] = node
			} else {
				nodeMap[*node.ParentId].Children = append(nodeMap[*node.ParentId].Children, node)
				sort.Slice(nodeMap[*node.ParentId].Children, func(i, j int) bool {
					return nodeMap[*node.ParentId].Children[i].Id < nodeMap[*node.ParentId].Children[j].Id
				})
			}
		}
	}

	tree := []model.OrganizationNode{}
	for _, node := range resultMap {
		tree = append(tree, *node)
	}

	sort.Slice(tree, func(i, j int) bool {
		return tree[i].Id < tree[j].Id
	})

	return tree
}

func getUsersForOrganization(organizationId string, tenantId string) (<-chan []string, <-chan error) {
	userIdChan := make(chan []string, 1)
	errChan := make(chan error, 1)

	go func() {
		users, err := service.GetUsersForOrganization(organizationId, tenantId)
		if err != nil {
			errChan <- err
		} else {
			userIds := service.GetIds(users, consts.UserPrefix)
			userIdChan <- userIds
		}

		close(userIdChan)
		close(errChan)
	}()

	return userIdChan, errChan
}

func getRolesForOrganization(organizationId string, tenantId string) (<-chan []string, <-chan error) {
	roleIdChan := make(chan []string, 1)
	errChan := make(chan error, 1)

	go func() {
		roles, err := service.GetRolesForOrganization(organizationId, tenantId)
		if err != nil {
			errChan <- err
		} else {
			roleIds := service.GetIds(roles, consts.RolePrefix)
			roleIdChan <- roleIds
		}

		close(roleIdChan)
		close(errChan)
	}()

	return roleIdChan, errChan
}
