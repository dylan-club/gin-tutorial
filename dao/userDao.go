package dao

import (
	"com.nicklaus/ginpractice/model"
	"gorm.io/gorm"
)

func FindUserByPhone(db *gorm.DB, telephone string) model.User {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	return user
}
