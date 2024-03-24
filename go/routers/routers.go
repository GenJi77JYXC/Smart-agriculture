package routers

import (
	"demo/controller"
	"demo/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.GET("/ip", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.ClientIP())
	})

	r.POST("/login", controller.Login)
	r.POST("/regist", controller.Regist)
	r.POST("/logout", middleware.AuthMiddleWare(), controller.LogOut)

	// mqtt
	r.POST("/publish", controller.PublicMsg)
	r.POST("/listen", controller.Listen)

	return r
}
