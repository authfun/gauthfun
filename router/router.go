package router

import (
	"log"
	"net/http"

	_ "github.com/authfun/gauthfun/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	AddGroupTenant(router)
	AddGroupMenu(router)
	AddGroupApi(router)
	AddGroupFeature(router)
	AddGroupPermission(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func BadRequest(c *gin.Context, msg ...string) {
	if len(msg) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input, please check!"})
	}
}

func NotFound(c *gin.Context, msg ...string) {
	if len(msg) > 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": msg})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found!"})
	}
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error, please try again later or contact the administrator!"})
	log.Panicln(err)
}
