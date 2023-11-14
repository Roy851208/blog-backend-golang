package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20); not null"`
	Telephone string `gorm:"varchar(110); not null; unique"`
	Password  string `gorm:"size:255; not null"`
}

func main() {
	db := InitDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
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
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		//判斷手機號是否存在
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用戶已存在"})
			return
		}
		//創建用戶
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)
		//返回結果
		c.JSON(200, gin.H{
			"msg": "註冊成功",
		})
	})
	panic(r.Run())
}

func RandomString(n int) string {
	var letters = []byte("asdfsadfasdfASDFQWERF")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))-1]
	}
	return string(result)
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	return user.ID != 0
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        args,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}
