package router

import (
	"errors"
	"net/http"

	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/schema"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddGroupFeature(router *gin.Engine) {

	group := router.Group("/api/features")
	{
		group.GET("", featureList)
		// group.GET("/options", menuOptions)
		// group.GET("/:id", getMenu)
		// group.POST("", addMenu)
		// group.PUT("/:id", updateMenu)
		// group.DELETE("/:id", deleteMenu)
	}
}

// @Summary Get feature list
// @Tags Feature
// @Accept json
// @Produce json
// @Success 200 {array} schema.Feature
// @Router /api/features/ [get]
func featureList(c *gin.Context) {
	var features []schema.Feature
	result := database.AuthDatabase.Find(&features)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, features)
}
