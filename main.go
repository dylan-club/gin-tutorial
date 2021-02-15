package main

import (
	"com.nicklaus/ginpractice/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
