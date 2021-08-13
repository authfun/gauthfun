package router

import (
	"net/http"

	"github.com/authfun/gauthfun/casbin"
	"github.com/gin-gonic/gin"
)

func AddGroupPermission(router *gin.Engine) {

	group := router.Group("/api/permissions")
	{
		group.GET("/validate", validatePermission)
	}
}

// /api/permissions/validate?sub=admin&dom=domain1&obj=data1&act=read
func validatePermission(c *gin.Context) {
	sub := c.Query("sub")
	dom := c.Query("dom")
	obj := c.Query("obj")
	act := c.Query("act")

	isValid, _ := casbin.Enforcer.Enforce(sub, dom, obj, act)
	c.JSON(http.StatusOK, isValid)
}
