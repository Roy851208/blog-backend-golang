package common

import (
	"blog/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
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
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: driverName,
		DSN:        args,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	DB.AutoMigrate(&model.User{})
}
