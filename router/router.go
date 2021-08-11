package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()

	AddGroupTenant(router)

	return router
}
