package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddGroupMock(router *gin.Engine) {

	group := router.Group("/api/mock")
	{
		group.POST("/login", login)
		group.GET("/info", info)
		group.POST("/logout", logout)
	}
}

func login(c *gin.Context) {

	result := LoginResult{
		Code: 20000,
		Data: LoginResultData{
			Token: "admin-token",
		},
	}
	c.JSON(http.StatusOK, result)
}

func info(c *gin.Context) {

	result := InfoResult{
		Code: 20000,
		Data: InfoResultData{
			Roles:        []string{"admin"},
			Introduction: "I am a super administrator",
			Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			Name:         "Admin",
		},
	}
	c.JSON(http.StatusOK, result)
}

func logout(c *gin.Context) {

	result := LogoutResult{
		Code: 20000,
		Data: "success",
	}
	c.JSON(http.StatusOK, result)
}

type LoginResult struct {
	Code int             `json:"code"`
	Data LoginResultData `json:"data"`
}

type LoginResultData struct {
	Token string `json:"token"`
}

type LogoutResult struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type InfoResult struct {
	Code int            `json:"code"`
	Data InfoResultData `json:"data"`
}

type InfoResultData struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}
