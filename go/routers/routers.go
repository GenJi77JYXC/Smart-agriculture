package routers

import (
	"demo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ip", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.ClientIP())
	})

	r.POST("/login", controller.Login)

	return r
}
