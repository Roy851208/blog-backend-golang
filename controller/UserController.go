package controller

import (
	"blog/common"
	"blog/dto"
	"blog/model"
	"blog/response"
	"blog/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	//1.使用map獲取請求的參數
	// var requestMap = make(map[string]string)
	// json.NewDecoder(c.Request.Body).Decode(&requestMap)
	//2.使用結構體獲取請求的參數
	// var requestUser = model.User{}
	// json.NewDecoder(c.Request.Body).Decode(&requestUser)
	//3.使用GIN框架裡的BIND函數
	var requestUser = model.User{}
	c.Bind(&requestUser)
	//獲取參數
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//數據驗證
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"telephone": telephone,
			"code":      422,
			"msg":       "手機號碼為11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密碼不能少於6位數",
		})
		return
	}
	//如果名稱沒有傳，給一個10位隨機字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判斷手機號是否存在
	if isTelephoneExist(telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422,
			"msg": "用戶已存在",
		})
		return
	}
	//創建用戶
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500,
			"msg": "加密錯誤",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	common.DB.Create(&newUser)
	//發放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(
			c,
			http.StatusUnprocessableEntity,
			500,
			nil,
			"系統異常",
		)
		log.Printf("token generate error : %v", err)
	}
	//返回結果
	response.Success(c, gin.H{"token": token}, "註冊成功")
}

func Login(c *gin.Context) {
	//獲取參數
	var requestUser = model.User{}
	c.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password
	//數據驗證
	if len(telephone) != 11 {
		response.Response(
			c,
			http.StatusUnprocessableEntity,
			422,
			nil,
			"手機號碼為11位",
		)
		return
	}
	if len(password) < 6 {
		response.Response(
			c,
			http.StatusUnprocessableEntity,
			422,
			nil,
			"密碼不能少於6位數",
		)
		return
	}
	//手機號是否存在
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(
			c,
			http.StatusUnprocessableEntity,
			422,
			nil,
			"用戶不存在",
		)
		return
	}
	//判斷密碼是否正確
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(
			c,
			http.StatusUnprocessableEntity,
			400,
			nil,
			"密碼錯誤",
		)
		return
	}
	//發放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(
			c,
			http.StatusUnprocessableEntity,
			500,
			nil,
			"系統異常",
		)
		log.Printf("token generate error : %v", err)
	}
	//返回結果
	response.Success(c, gin.H{"token": token}, "登陸成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)), //這裡是類型斷言
		},
	})
}

func isTelephoneExist(telephone string) bool {
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}
