package config

import (
	"github.com/Mitmadhu/mysqlDB/database"
	"gorm.io/gorm"
)

var Cnf = Config{}

var X = 100

type Config struct{
	 DB *gorm.DB
}

func GetDB() *gorm.DB{
	return Cnf.DB
}

func InitConfig() error{
	// initDB
	db, err := database.DBFactory("", "")
	Cnf.DB = db
	
	if err != nil {
		return err
	}
	return nil
}

