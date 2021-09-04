package router

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/authfun/gauthfun/casbin"
	"github.com/authfun/gauthfun/consts"
	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
	service "github.com/authfun/gauthfun/service"
	util "github.com/authfun/gauthfun/util"
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
		group.POST("", addFeature)
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

// @Summary Add feature
// @Tags Feature
// @Accept json
// @Produce json
// @Param account body model.FeatureForm true "Feature info"
// @Success 201 {object} schema.Feature
// @Router /api/features [post]
func addFeature(c *gin.Context) {
	var form model.FeatureForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}

	var feature schema.Feature
	copier.Copy(&feature, &form)

	var apis []schema.Api
	database.AuthDatabase.Find(apis, form.ApiIds)

	feature.Id = util.GenerateUUID()
	menuRules := <-generateMenuRules(feature.Id, form.MenuIds)
	apiRules := <-generateApiRules(feature.Id, apis)
	featureRules := <-generateFeatureRules(feature.Id, form.FeatureIds)

	err := database.AuthDatabase.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&feature).Error; err != nil {
			return err
		}

		if menuRules != nil {
			if err := tx.Create(&menuRules).Error; err != nil {
				return err
			}
		}

		if apiRules != nil {
			if err := tx.Create(&apiRules).Error; err != nil {
				return err
			}
		}

		if featureRules != nil {
			if err := tx.Create(&featureRules).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		InternalServerError(c, err)
		return
	}

	err = casbin.Enforcer.LoadPolicy()
	if err != nil {
		InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, feature)
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

func generateMenuRules(featureId string, menuIds []string) <-chan []schema.CasbinRule {
	menuRuleChan := make(chan []schema.CasbinRule, 1)

	go func() {
		var menuRules []schema.CasbinRule
		linq.From(menuIds).SelectT(func(id string) schema.CasbinRule {
			rule := schema.CasbinRule{
				Ptype: consts.PType_Policy,
				V0:    consts.FeaturePrefix + fmt.Sprintf("%v", featureId),
				V1:    consts.Domain_Pattern_All,
				V2:    consts.MenuPrefix + fmt.Sprintf("%v", id),
				V3:    consts.Act_Pattern_All,
			}
			return rule
		}).ToSlice(&menuRules)

		menuRuleChan <- menuRules
		close(menuRuleChan)
	}()

	return menuRuleChan
}

func generateApiRules(featureId string, apis []schema.Api) <-chan []schema.CasbinRule {
	apiRuleChan := make(chan []schema.CasbinRule, 1)

	go func() {
		var apiRules []schema.CasbinRule
		linq.From(apis).SelectT(func(api schema.Api) schema.CasbinRule {
			rule := schema.CasbinRule{
				Ptype: consts.PType_Policy,
				V0:    consts.FeaturePrefix + fmt.Sprintf("%v", featureId),
				V1:    consts.Domain_Pattern_All,
				V2:    consts.ApiPrefix + api.Route,
				V3:    api.Method,
			}
			return rule
		}).ToSlice(&apiRules)

		apiRuleChan <- apiRules
		close(apiRuleChan)
	}()

	return apiRuleChan
}

func generateFeatureRules(featureId string, featureIds []string) <-chan []schema.CasbinRule {
	featureRuleChan := make(chan []schema.CasbinRule, 1)

	go func() {
		var featureRules []schema.CasbinRule
		linq.From(featureIds).SelectT(func(id string) schema.CasbinRule {
			rule := schema.CasbinRule{
				Ptype: consts.PType_Group,
				V0:    consts.FeaturePrefix + fmt.Sprintf("%v", featureId),
				V1:    consts.FeaturePrefix + fmt.Sprintf("%v", id),
				V2:    consts.Domain_Pattern_All,
			}
			return rule
		}).ToSlice(&featureRules)

		featureRuleChan <- featureRules
		close(featureRuleChan)
	}()

	return featureRuleChan
}
