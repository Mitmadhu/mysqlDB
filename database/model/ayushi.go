package model

import "gorm.io/gorm"

type Ayushi struct{
	gorm.Model
	UserName string

}