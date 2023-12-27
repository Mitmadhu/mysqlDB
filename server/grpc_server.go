package server

import (
	"context"
	"log"
	"net"

	"github.com/Mitmadhu/mysqlDB/pb/grpc_test"
	"github.com/Mitmadhu/mysqlDB/pb/grpc_user"
	"google.golang.org/grpc"
)

type GRPCServerImpl struct {
	grpc_user.MysqlServer
}


type GRPCTestImpl struct{
	grpc_test.TestServer
}

func (*GRPCTestImpl) SayHello(context.Context, *grpc_test.TestRequest) (*grpc_test.TestResponse, error) {
	// TODO: impl
	return &grpc_test.TestResponse{
		Msg: "hello",
	}, nil
} 

func (*GRPCServerImpl) UserExists(context.Context, *grpc_user.ValidateUserRequest) (*grpc_user.ValidateUserResponse, error) {
	// TODO: impl
	isValid := true
	return &grpc_user.ValidateUserResponse{
		IsValid:    &isValid,
		MsgID:      "123",
		Success:    true,
		StatusCode: 200,
		Message:    "hello madhu",
	}, nil
}

func StartGRPCServer() {
	// grpc server
	lis, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		log.Fatal("error in listening", err.Error())
	}

	grpcServer := grpc.NewServer()

	// register server for grpc
	grpc_user.RegisterMysqlServer(grpcServer, &GRPCServerImpl{})
	grpc_test.RegisterTestServer(grpcServer, &GRPCTestImpl{})

	println("grpc server starting...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("grpc serve error ", err.Error())
	}
}
