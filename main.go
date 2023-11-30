package main

import (
	"fmt"

	"github.com/Mitmadhu/mysqlDB/config"
	"github.com/Mitmadhu/mysqlDB/database/model"
	"github.com/Mitmadhu/mysqlDB/server"
	"gorm.io/gorm"
)

type table interface {
	CreateTable(db *gorm.DB) error
}

func migrateTables() {
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
	server.InitRouter()
	// migrateTables()
	

	// err:= model.User{}.Register("ayush", "123", "a", "v", 2)
	// if(err != nil){
	// 	fmt.Println(err.Error())
	// }

	// b, err := model.User{}.ValidateUser("ayushi", "12345")
	// if(err != nil){
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(b)

}
