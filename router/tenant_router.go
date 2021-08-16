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

func AddGroupTenant(router *gin.Engine) {

	group := router.Group("/api/tenants")
	{
		group.GET("", tenantList)
		group.GET("/options", tenantOptions)
		group.GET("/:id", getTenant)
		group.POST("", addTenant)
		group.PUT("/:id", updateTenant)
		group.DELETE("/:id", deleteTenant)
	}
}

// @Summary Get tenant list
// @Tags Tenant
// @Accept json
// @Produce json
// @Success 200 {array} schema.Tenant
// @Router /api/tenants/ [get]
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

// @Summary Get tenant option
// @Tags Tenant
// @Accept json
// @Produce json
// @Success 200 {array} model.Option
// @Router /api/tenants/options [get]
func tenantOptions(c *gin.Context) {
	var tenants []schema.Tenant
	db := database.AuthDatabase
	result := db.Find(&tenants)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		InternalServerError(c, result.Error)
		return
	}

	var options []model.Option
	copier.Copy(&options, &tenants)
	c.JSON(http.StatusOK, options)
}

// @Summary Get tenant
// @Tags Tenant
// @Accept json
// @Produce json
// @Param id path string true "Tenant Id"
// @Success 200 {object} schema.Tenant
// @Router /api/tenants/{id} [get]
func getTenant(c *gin.Context) {
	id := c.Param("id")
	var tenant schema.Tenant
	result := database.AuthDatabase.First(&tenant, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	c.JSON(http.StatusOK, tenant)
}

// @Summary Add tenant
// @Tags Tenant
// @Accept json
// @Produce json
// @Param account body model.TenantForm true "Tenant info"
// @Success 201 {object} schema.Tenant
// @Router /api/tenants [post]
func addTenant(c *gin.Context) {
	var form model.TenantForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}
	var tenant schema.Tenant

	copier.Copy(&tenant, &form)
	tenant.Id = util.GenerateUUID()
	result := database.AuthDatabase.Create(&tenant)
	if result.Error != nil {
		InternalServerError(c, result.Error)
		return
	}

	c.JSON(http.StatusCreated, tenant)
}

// @Summary Update tenant
// @Tags Tenant
// @Accept json
// @Produce json
// @Param id path string true "Tenant Id"
// @Param account body model.TenantForm true "Tenant info"
// @Success 200
// @Router /api/tenants/{id} [put]
func updateTenant(c *gin.Context) {
	id := c.Param("id")
	var form model.TenantForm
	if err := c.ShouldBindJSON(&form); err != nil {
		BadRequest(c)
		return
	}

	var tenant schema.Tenant

	copier.Copy(&tenant, &form)
	tenant.Id = id
	result := database.AuthDatabase.Save(&tenant)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			NotFound(c)
		} else {
			InternalServerError(c, result.Error)
		}
		return
	}

	c.JSON(http.StatusOK, tenant)
}

// @Summary Delete tenant
// @Tags Tenant
// @Accept json
// @Produce json
// @Param id path string true "Tenant Id"
// @Success 204
// @Router /api/tenants/{id} [delete]
func deleteTenant(c *gin.Context) {
	id := c.Param("id")

	result := database.AuthDatabase.Where("id = ?", id).Delete(&schema.Tenant{})
	if result.Error != nil {
		InternalServerError(c, result.Error)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
