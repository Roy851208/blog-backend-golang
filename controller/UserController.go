package controller

import (
	"blog/common"
	"blog/model"
	"blog/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	//獲取參數
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//數據驗證
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手機號碼為11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密碼不能少於6位數"})
		return
	}
	//如果名稱沒有傳，給一個10位隨機字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判斷手機號是否存在
	if isTelephoneExist(telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用戶已存在"})
		return
	}
	//創建用戶
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	common.DB.Create(&newUser)
	//返回結果
	c.JSON(200, gin.H{
		"msg": "註冊成功",
	})
}

func isTelephoneExist(telephone string) bool {
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
