package controller

import (
	"demo/database"
	"demo/middleware"
	"demo/model"
	"demo/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	DB := database.GetDB()

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 判断密码合法性
	if len(password) < 6 {
		response.Response(ctx, 403, 403, "密码长度太低", "密码至少需要6位")
		return
	}
	var user model.User
	DB.Table("users").Where("username = ?", username).First(&user)
	if user.ID == 0 {
		response.Response(ctx, 403, 403, "用户不存在", "请重新输入信息")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(ctx, 403, 403, "密码错误", "请重新输入密码")
		return
	}
	// 发token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, 500, 500, "token发放错误", "token发放错误")
		return
	}
	response.Response(ctx, 200, 200, token, "登录成功")
}
