package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (u User) CreateTable(db *gorm.DB){
	if db == nil {
		panic("db is nil")
	}
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("%+v", err)
	}
}