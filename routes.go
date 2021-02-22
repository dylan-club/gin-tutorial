package main

import (
	"com.nicklaus/ginpractice/controller"
	"com.nicklaus/ginpractice/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.UserLogin)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.UserInfo)

	categoryGroup := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryGroup.POST("/create", categoryController.Create)
	categoryGroup.DELETE("/delete/:id", categoryController.Delete)
	categoryGroup.GET("/show/:id", categoryController.Show)
	categoryGroup.PUT("/update/:id", categoryController.Update)
	return r
}
