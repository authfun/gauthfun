package router

import (
	"net/http"

	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
	"github.com/gin-gonic/gin"
	"sort"
	linq "github.com/ahmetb/go-linq/v3"
)

func AddGroupOrganization(router *gin.Engine) {

	group := router.Group("/api/organizations")
	{
		group.GET("/tree", tree)
	}
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