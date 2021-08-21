package router

import (
	"errors"
	"net/http"

	"github.com/authfun/gauthfun/database"
	"github.com/authfun/gauthfun/model"
	"github.com/authfun/gauthfun/schema"
	util "github.com/authfun/gauthfun/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func AddGroupApi(router *gin.Engine) {

	group := router.Group("/api/apis")
	{
		group.GET("", apiList)
		group.GET("/:id", getApi)
		group.POST("", addApi)
		group.PUT("/:id", updateApi)
		group.DELETE("/:id", deleteApi)
	}
}

// @Summary Get api list
// @Tags Api
// @Accept json
// @Produce json
// @Success 200 {array} schema.Api
// @Router /api/apis/ [get]
func apiList(c *gin.Context) {
	var apis []schema.Api
	result := database.AuthDatabase.Find(&apis)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, apis)
}

// @Summary Get api
// @Tags Api
// @Accept json
// @Produce json
// @Param id path string true "Api Id"
// @Success 200 {object} schema.Api
// @Router /api/apis/{id} [get]
func getApi(c *gin.Context) {
	id := c.Param("id")
	var api schema.Api
	result := database.AuthDatabase.First(&api, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	c.JSON(http.StatusOK, api)
}

// @Summary Add api
// @Tags Api
// @Accept json
// @Produce json
// @Param account body model.ApiForm true "Api info"
// @Success 201 {object} schema.Api
// @Router /api/apis [post]
func addApi(c *gin.Context) {
	var form model.ApiForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}
	var api schema.Api

	copier.Copy(&api, &form)
	api.Id = util.GenerateUUID()
	result := database.AuthDatabase.Create(&api)
	if result.Error != nil {
		InternalServerError(c, result.Error)
		return
	}

	c.JSON(http.StatusCreated, api)
}

// @Summary Update api
// @Tags Api
// @Accept json
// @Produce json
// @Param id path string true "Api Id"
// @Param account body model.ApiForm true "Api info"
// @Success 200
// @Router /api/apis/{id} [put]
func updateApi(c *gin.Context) {
	id := c.Param("id")
	var form model.ApiForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}

	var api schema.Api

	copier.Copy(&api, &form)
	api.Id = id
	result := database.AuthDatabase.Save(&api)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	c.JSON(http.StatusOK, api)
}

// @Summary Delete api
// @Tags Api
// @Accept json
// @Produce json
// @Param id path string true "Api Id"
// @Success 204
// @Router /api/apis/{id} [delete]
func deleteApi(c *gin.Context) {
	id := c.Param("id")

	result := database.AuthDatabase.Where("id = ?", id).Delete(&schema.Api{})
	if result.Error != nil {
		InternalServerError(c, result.Error)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
