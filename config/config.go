package config

import (
	"fmt"

	"github.com/Mitmadhu/mysqlDB/database"
	"gorm.io/gorm"
)

var Cnf = Config{}

type Config struct{
	 DB *gorm.DB
}

func GetDB() *gorm.DB {
	if Cnf.DB == nil{
		err := initDB()
		if err != nil || Cnf.DB == nil{
			panic(fmt.Sprintf("error while connection to db, err: %v", err))
		}
	}
	return Cnf.DB
}

func initDB()error{
	// initDB
	db, err := database.DBFactory("", "")
	Cnf.DB = db
	
	if err != nil {
		return err
	}
	return nil
}

func InitConfig() error{
	return nil
}

