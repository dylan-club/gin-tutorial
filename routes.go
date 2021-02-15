package main

import (
	"com.nicklaus/ginpractice/controller"
	"com.nicklaus/ginpractice/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.UserLogin)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.UserInfo)
	return r
}
