package controller

import (
	"com.nicklaus/ginpractice/common"
	"com.nicklaus/ginpractice/dao"
	"com.nicklaus/ginpractice/dto"
	"com.nicklaus/ginpractice/model"
	"com.nicklaus/ginpractice/response"
	"com.nicklaus/ginpractice/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "手机号必须是11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if dao.FindUserByPhone(db, telephone).ID != 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "用户已经存在")
		return
	}

	//密码加盐
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{}, "密码加密错误")
		return
	}

	//创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	db.Create(&newUser)

	response.Success(ctx, gin.H{}, "注册成功")
}

func UserLogin(ctx *gin.Context) {
	//获取数据
	db := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "手机号必须是11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "密码不能少于6位")
		return
	}
	//比较用户名和密码输入
	user := dao.FindUserByPhone(db, telephone)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, gin.H{}, "用户名不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, gin.H{}, "密码有误")
		return
	}
	//登录成功，返回token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, gin.H{}, "系统崩溃")
		log.Printf("generate token failed error: %v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "用户信息")
}
