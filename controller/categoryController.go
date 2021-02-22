package controller

import (
	"com.nicklaus/ginpractice/common"
	"com.nicklaus/ginpractice/repository"
	"com.nicklaus/ginpractice/response"
	"com.nicklaus/ginpractice/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	IRestController
}

type CategoryController struct {
	Repository repository.ICategoryRepository
}

func NewCategoryController() ICategoryController {
	categoryRepository := repository.NewCategoryRepository()
	return &CategoryController{Repository: categoryRepository}
}

func (c *CategoryController) Create(ctx *gin.Context) {
	categoryVO := new(vo.CategoryVO)
	//判断分类名称是否为空
	if err := ctx.ShouldBind(categoryVO); err != nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgEmptyCategoryError)
		return
	}
	//判断category是否存在
	if _, err := c.Repository.SelectByName(categoryVO.Name); err == nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgExistCategoryError)
		return
	}

	category, err := c.Repository.Create(categoryVO.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, common.ConstMsgCreateCategorySuccess)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	//判断id输入格式是否有误
	if err != nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgConvertCategoryIdError)
		return
	}

	err = c.Repository.DeleteById(id)
	if err != nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgDeleteCategoryFail)
		return
	}
	response.Success(ctx, gin.H{}, common.ConstMsgDeleteCategorySuccess)
}

func (c *CategoryController) Show(ctx *gin.Context) {
	//获取id
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(id)
	if err != nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgNotExistCategoryError)
		return
	}
	response.Success(ctx, gin.H{"category": category}, common.ConstMsgShowCategorySuccess)
}

func (c *CategoryController) Update(ctx *gin.Context) {
	//获取id
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategoryVO = new(vo.CategoryVO)
	//判断分类名称是否为空
	if err := ctx.ShouldBind(updateCategoryVO); err != nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgEmptyCategoryError)
		return
	}

	updateCategory, err := c.Repository.SelectById(id)
	if err != nil {
		response.Fail(ctx, gin.H{}, common.ConstMsgNotExistCategoryError)
		return
	}
	//更新category
	updateCategory, err = c.Repository.UpdateForName(updateCategory, updateCategoryVO.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": updateCategory}, common.ConstMsgUpdateCategorySuccess)
}
