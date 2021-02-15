package controller

import (
	"com.nicklaus/ginpractice/common"
	"com.nicklaus/ginpractice/dao"
	"com.nicklaus/ginpractice/model"
	"com.nicklaus/ginpractice/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	//获取参数
	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须是11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if dao.IsTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}

	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}
