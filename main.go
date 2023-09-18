package main

import (
	"github.com/Mitmadhu/mysqlDB/config"
	"github.com/Mitmadhu/mysqlDB/database/model"
)

func main(){
	config.InitConfig()

	u := model.User{}
	
	u.CreateTable(config.Cnf.DB)
}