package controller

import (
	"blog/common"
	"blog/model"
	"blog/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}
type CategoryController struct {
}

func NewCategoryController() ICategoryController {
	common.DB.AutoMigrate(model.Category{})
	return CategoryController{}
}
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCatogory model.Category
	ctx.Bind(&requestCatogory)
	if requestCatogory.Name == "" {
		response.Fail(ctx, "數據驗證錯誤，分類名稱必填", nil)
		return
	}
	common.DB.Create(&requestCatogory)
	response.Success(ctx, gin.H{"category": requestCatogory}, "")
}
func (c CategoryController) Update(ctx *gin.Context) {
	//綁定body中的參數
	var requestCatogory model.Category
	ctx.Bind(&requestCatogory)
	if requestCatogory.Name == "" {
		response.Fail(ctx, "數據驗證錯誤，分類名稱必填", nil)
		return
	}
	//獲取path中的參數
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if err := common.DB.First(&updateCategory, categoryId).Error; err != nil {
		response.Fail(ctx, "分類不存在", nil)
		return
	}
	//更新分類
	common.DB.Model(&updateCategory).Update("name", requestCatogory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}
func (c CategoryController) Show(ctx *gin.Context) {
	//獲取path中的參數
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if err := common.DB.First(&category, categoryId).Error; err != nil {
		response.Fail(ctx, "分類不存在", nil)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}
func (c CategoryController) Delete(ctx *gin.Context) {
	//獲取path中的參數
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := common.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, "刪除失敗", nil)
		return
	}
	response.Success(ctx, nil, "刪除成功")
}
