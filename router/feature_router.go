package router

import (
	"errors"
	"net/http"

	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func AddGroupFeature(router *gin.Engine) {

	group := router.Group("/api/features")
	{
		group.GET("", featureList)
		group.GET("/options", featureOptions)
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

// @Summary Get feature option
// @Tags Feature
// @Accept json
// @Produce json
// @Success 200 {array} model.Option
// @Router /api/features/options [get]
func featureOptions(c *gin.Context) {
	var features []schema.Feature
	result := database.AuthDatabase.Find(&features)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}

	var options []model.Option
	copier.Copy(&options, &features)
	c.JSON(http.StatusOK, options)
}
