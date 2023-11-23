package controller

import (
	"blog/common"
	"blog/model"
	"blog/response"
	"blog/respository"
	"blog/vo"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	RestController
	PageList(c *gin.Context)
}
type PostController struct {
	Repository respository.CategoryRespository
}

func NewPostController() IPostController {
	common.DB.AutoMigrate(model.Post{})
	return PostController{}
}
func (p PostController) Create(ctx *gin.Context) {
	//綁定body中的參數
	var requestPost vo.CreatePostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "數據驗證錯誤", nil)
		return
	}
	//獲取登陸用戶
	user, _ := ctx.Get("user")
	//創建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := common.DB.Create(&post).Error; err != nil {
		response.Fail(ctx, "創建失敗", nil)
		panic(err)
	}

	response.Success(ctx, gin.H{"post": post}, "創建成功")
}
func (p PostController) Update(ctx *gin.Context) {
	//綁定body中的參數
	var requestPost vo.CreatePostRequest
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, "數據驗證錯誤", nil)
		return
	}
	//獲取path中的ID
	postId := ctx.Params.ByName("id")
	var post model.Post
	if err := common.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	//判斷當前用戶是否為用戶作者
	//獲取登陸用戶
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, "非文章作者", nil)
		return
	}
	//更新文章
	if err := common.DB.Model(&post).Updates(requestPost).Error; err != nil {
		response.Fail(ctx, "更新失敗", nil)
	}

	response.Success(ctx, gin.H{"post": post}, "更新成功")
}
func (p PostController) Show(ctx *gin.Context) {
	//獲取path中的ID
	postId := ctx.Params.ByName("id")
	var post model.Post
	if err := common.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	response.Success(ctx, gin.H{"post": post}, "")
}
func (p PostController) Delete(ctx *gin.Context) {
	//獲取path中的參數
	postId := ctx.Params.ByName("id")
	//判斷文章是否存在
	var post model.Post
	if err := common.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, "文章不存在", nil)
		return
	}
	//判斷用戶是否為文章作者
	//獲取登陸用戶
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, "非文章作者", nil)
		return
	}

	if err := common.DB.Delete(&post).Error; err != nil {
		response.Success(ctx, nil, "刪除失敗")
		return
	}
	response.Success(ctx, nil, "刪除成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	//獲取分頁參數
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	//分頁
	var post []model.Post
	common.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&post)

	//前端渲染分頁需要知道總數
	var total int64
	common.DB.Model(model.Post{}).Count(&total)

	response.Success(ctx, gin.H{"data": post, "total": total}, "成功")
}
