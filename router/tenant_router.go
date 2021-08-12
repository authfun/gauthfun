package router

import (
	"net/http"

	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/schema"
	"github.com/gin-gonic/gin"
)

func AddGroupTenant(router *gin.Engine) {

	group := router.Group("/api/tenants")
	{
		group.GET("", tenantList)
	}
}

func tenantList(c *gin.Context) {
	var tenants []schema.Tenant
	db := database.AuthDatabase
	result := db.Find(&tenants)
	if result.Error != nil {

	}
	c.JSON(http.StatusOK, tenants)
}
