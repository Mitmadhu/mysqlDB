package model

import (

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (u User) CreateTable(db *gorm.DB) error{
	return db.AutoMigrate(&User{})
}