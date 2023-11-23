package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(20); not null"`
	Telephone string `json:"telephone" gorm:"varchar(110); not null; unique"`
	Password  string `json:"password" gorm:"size:255; not null"`
}
