package common

import (
	"blog/model"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	worDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(worDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("host.driverName")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
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

// func InitConfig() {
// 	worDir, _ := os.Getwd()
// 	viper.SetConfigName("application")
// 	viper.SetConfigType("yml")
// 	viper.AddConfigPath(worDir + "/config")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }
