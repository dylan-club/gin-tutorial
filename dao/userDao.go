package dao

import (
	"com.nicklaus/ginpractice/model"
	"gorm.io/gorm"
	"log"
)

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	log.Println(user)
	if user.ID == 0 {
		return false
	}
	return true
}
