package router

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"sort"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/authfun/gauthfun/casbin"
	"github.com/authfun/gauthfun/consts"
	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
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

	name := consts.OrganizationPrefix + fmt.Sprintf("%v", id)
	domain := consts.TenantPrefix + fmt.Sprintf("%v", organization.TenantId)

	users, err := casbin.Enforcer.GetUsersForRole(name, domain)
	if err != nil && !errors.Is(err, rbac_errors.ERR_NAME_NOT_FOUND) {
		InternalServerError(c, err)
		return
	}

	roles, err := casbin.Enforcer.GetRolesForUser(name, domain)
	if err != nil && !errors.Is(err, rbac_errors.ERR_NAME_NOT_FOUND) {
		InternalServerError(c, err)
		return
	}

	userIdChan := make(chan []int64)
	roleIdChan := make(chan []int64)

	go getUserIds(users, userIdChan)
	go getRoleIds(roles, roleIdChan)

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

func getUserIds(users []string, userIdChan chan []int64) {
	var ids []int64
	linq.From(users).SelectT(func(user string) int64 {
		s := strings.ReplaceAll(user, consts.UserPrefix, "")
		userId, _ := strconv.ParseInt(s, 10, 64)
		return userId
	}).ToSlice(&ids)

	userIdChan <- ids
}

func getRoleIds(roles []string, roleIdChan chan []int64) {
	var ids []int64
	linq.From(roles).SelectT(func(role string) int64 {
		s := strings.ReplaceAll(role, consts.RolePrefix, "")
		roleId, _ := strconv.ParseInt(s, 10, 64)
		return roleId
	}).ToSlice(&ids)

	roleIdChan <- ids
}
