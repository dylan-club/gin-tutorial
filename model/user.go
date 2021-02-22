package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null" json:"name" form:"name"`
	Telephone string `gorm:"type:varchar(110);not null;unique" json:"telephone" form:"telephone"`
	Password  string `gorm:"size:255;not null" json:"password" form:"password"`
}
