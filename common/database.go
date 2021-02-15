package common

import (
	"com.nicklaus/ginpractice/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	host := "localhost"
	port := "3306"
	database := "test"
	username := "root"
	password := "root"
	charset := "utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database, err: " + err.Error())
	}
	_ = db.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return db
}
