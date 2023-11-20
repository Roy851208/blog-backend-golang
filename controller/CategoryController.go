package controller

import (
	"blog/common"
	"blog/model"
	"blog/response"
	"blog/respository"
	"blog/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}
type CategoryController struct {
	Repository respository.CategoryRespository
}

func NewCategoryController() ICategoryController {
	common.DB.AutoMigrate(model.Category{})
	return CategoryController{}
}
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCatogory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCatogory); err != nil {
		response.Fail(ctx, "數據驗證錯誤，分類名稱必填", nil)
		return
	}

	category, err := c.Repository.Create(requestCatogory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "")
}
func (c CategoryController) Update(ctx *gin.Context) {
	//綁定body中的參數
	var requestCatogory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCatogory); err != nil {
		response.Fail(ctx, "數據驗證錯誤，分類名稱必填", nil)
		return
	}
	//獲取path中的參數
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分類不存在", nil)
		return
	}
	//更新分類
	category, err := c.Repository.Update(*updateCategory, requestCatogory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}
func (c CategoryController) Show(ctx *gin.Context) {
	//獲取path中的參數
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分類不存在", nil)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}
func (c CategoryController) Delete(ctx *gin.Context) {
	//獲取path中的參數
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(ctx, "刪除失敗", nil)
		return
	}
	response.Success(ctx, nil, "刪除成功")
}
