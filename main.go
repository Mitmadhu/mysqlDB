package main

import (
	"fmt"

	"github.com/Mitmadhu/mysqlDB/config"
	"github.com/Mitmadhu/mysqlDB/database/model"
	"gorm.io/gorm"
)

type table interface{
	CreateTable(db *gorm.DB) error
}

func migrateTables(){
	db := config.GetDB()

	fmt.Println("Migrating DB...")
	arr := []table{
		model.User{},
		model.Friend{},
	}
	
	for _, t := range arr {
		err := t.CreateTable(db)
		if err != nil {
			panic(fmt.Sprintf("error while creating table, err: %v", err))
		}
	}
	fmt.Println("Migration done!")
}

func main() {
	config.InitConfig()
	// migrateTables()
}
