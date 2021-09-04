package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/authfun/gauthfun/consts"
	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
	service "github.com/authfun/gauthfun/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func AddGroupFeature(router *gin.Engine) {

	group := router.Group("/api/features")
	{
		group.GET("", featureList)
		group.GET("/options", featureOptions)
		group.GET("/:id", getFeature)
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

// @Summary Get feature by id
// @Tags Feature
// @Accept json
// @Produce json
// @Param id path string true "Feature ID"
// @Param implicit query bool false "Whether to get the implicit info"
// @Success 200 {object} model.FeatureDetail
// @Router /api/features/{id} [get]
func getFeature(c *gin.Context) {
	id := c.Param("id")

	var feature schema.Feature
	result := database.AuthDatabase.First(&feature, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	implicit, err := strconv.ParseBool(c.Query("implicit"))
	if err != nil {
		implicit = false
	}

	objects, err := service.GetObjectsForFeature(id, implicit)
	if err != nil {
		InternalServerError(c, err)
		return
	}

	detail := model.FeatureDetail{
		Menus:    []schema.Menu{},
		Apis:     []schema.Api{},
		Features: []schema.Feature{},
	}
	copier.Copy(&detail, &feature)

	featureIdChan, featureIdErr := getFeaturesForFeature(id, implicit)
	if err, existed := <-featureIdErr; existed {
		InternalServerError(c, err)
		return
	}

	featureIds, menuIds, apiIds := <-featureIdChan, <-getMenusForFeature(objects), <-getApisForFeature(objects)

	if menuIds != nil {
		database.AuthDatabase.Find(&detail.Menus, menuIds)
	}

	if apiIds != nil {
		database.AuthDatabase.Find(&detail.Apis, apiIds)
	}

	if featureIds != nil {
		database.AuthDatabase.Find(&detail.Features, featureIds)
	}

	c.JSON(http.StatusOK, detail)
}

func getMenusForFeature(objects [][]string) <-chan []string {
	menuIdChan := make(chan []string, 1)

	go func() {
		menuIdChan <- service.FilterObjects(objects, consts.MenuPrefix)
		close(menuIdChan)
	}()

	return menuIdChan
}

func getApisForFeature(objects [][]string) <-chan []string {
	apiIdChan := make(chan []string, 1)

	go func() {
		apiIdChan <- service.FilterObjects(objects, consts.ApiPrefix)
		close(apiIdChan)
	}()

	return apiIdChan
}

func getFeaturesForFeature(featureId string, implicit bool) (<-chan []string, <-chan error) {
	featureIdChan := make(chan []string, 1)
	errChan := make(chan error, 1)

	go func() {
		features, err := service.GetFeaturesForFeature(featureId, implicit)
		if err != nil {
			errChan <- err
		} else {
			featureIds := service.FilterIds(features, consts.FeaturePrefix)
			featureIdChan <- featureIds
		}

		close(featureIdChan)
		close(errChan)
	}()

	return featureIdChan, errChan
}
