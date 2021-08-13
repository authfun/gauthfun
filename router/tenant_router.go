package router

import (
	"errors"
	"net/http"

	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/schema"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, tenants)
}
