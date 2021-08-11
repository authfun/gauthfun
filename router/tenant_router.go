package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddGroupTenant(router *gin.Engine) {

	group := router.Group("/api/tenants")
	{
		group.GET("", tenantList)
	}
}

func tenantList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "tenant list"})
}
