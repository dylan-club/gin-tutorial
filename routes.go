package main

import (
	"com.nicklaus/ginpractice/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.UserRegister)
	return r
}
