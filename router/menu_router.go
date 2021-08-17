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

func AddGroupMenu(router *gin.Engine) {

	group := router.Group("/api/menus")
	{
		group.GET("", menuList)
		group.GET("/options", menuOptions)
		group.GET("/:id", getMenu)
		group.POST("", addMenu)
		group.PUT("/:id", updateMenu)
		group.DELETE("/:id", deleteMenu)
	}
}

// @Summary Get menu list
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {array} schema.Menu
// @Router /api/menus/ [get]
func menuList(c *gin.Context) {
	var menus []schema.Menu
	result := database.AuthDatabase.Find(&menus)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}
	c.JSON(http.StatusOK, menus)
}

// @Summary Get menu option
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {array} model.Option
// @Router /api/menus/options [get]
func menuOptions(c *gin.Context) {
	var menus []schema.Menu
	result := database.AuthDatabase.Find(&menus)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}

	var options []model.Option
	copier.Copy(&options, &menus)
	c.JSON(http.StatusOK, options)
}

// @Summary Get menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Menu Id"
// @Success 200 {object} schema.Menu
// @Router /api/menus/{id} [get]
func getMenu(c *gin.Context) {
	id := c.Param("id")
	var menu schema.Menu
	result := database.AuthDatabase.First(&menu, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	c.JSON(http.StatusOK, menu)
}

// @Summary Add menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param account body model.MenuForm true "Menu info"
// @Success 201 {object} schema.Menu
// @Router /api/menus [post]
func addMenu(c *gin.Context) {
	var form model.TenantForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}
	var menu schema.Menu

	copier.Copy(&menu, &form)
	menu.Id = util.GenerateUUID()
	result := database.AuthDatabase.Create(&menu)
	if result.Error != nil {
		InternalServerError(c, result.Error)
		return
	}

	c.JSON(http.StatusCreated, menu)
}

// @Summary Update menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Menu Id"
// @Param account body model.MenuForm true "Menu info"
// @Success 200
// @Router /api/menus/{id} [put]
func updateMenu(c *gin.Context) {
	id := c.Param("id")
	var form model.MenuForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}

	var menu schema.Menu

	copier.Copy(&menu, &form)
	menu.Id = id
	result := database.AuthDatabase.Save(&menu)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	c.JSON(http.StatusOK, menu)
}

// @Summary Delete menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Menu Id"
// @Success 204
// @Router /api/menus/{id} [delete]
func deleteMenu(c *gin.Context) {
	id := c.Param("id")

	result := database.AuthDatabase.Where("id = ?", id).Delete(&schema.Menu{})
	if result.Error != nil {
		InternalServerError(c, result.Error)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}