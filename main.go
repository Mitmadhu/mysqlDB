package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Mitmadhu/mysqlDB/config"
	"github.com/Mitmadhu/mysqlDB/database/model"
	"github.com/Mitmadhu/mysqlDB/pb/grpc_test"
	"github.com/Mitmadhu/mysqlDB/pb/grpc_user"
	"github.com/Mitmadhu/mysqlDB/server"
	"google.golang.org/grpc"
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
	go func ()  {
		server.InitRouter()
	}()

	go func ()  {
		server.StartGRPCServer()
	}()

	time.Sleep(time.Second)
	
	

	// ---------------------------
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error in grpc dial", err.Error())
	}
	defer conn.Close()
	client := grpc_user.NewMysqlClient(conn)
	
	request := grpc_user.ValidateUserRequest{
		MsgID: "1230",
		Password: "123",
		Username: "123",
		Token: "123",
	}
	resp, err := client.UserExists(context.Background(), &request)
	if err != nil {
		log.Fatal("error in login call", err.Error())
	}

	fmt.Printf("%+v\n", resp)

	// ------------------------
	// time.Sleep(time.Second)

	testclient := grpc_test.NewTestClient(conn)
	
	req := grpc_test.TestRequest{
		MsgID: "123",
	}
	respTest, err := testclient.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal("error in login call", err.Error())
	}

	fmt.Printf("%+v\n", respTest)



	// _____

	resp, err = client.UserExists(context.Background(), &request)
	if err != nil {
		log.Fatal("error in login call", err.Error())
	}

	fmt.Printf("%+v\n", resp)


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
