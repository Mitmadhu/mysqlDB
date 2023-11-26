package server

import (
	"net/http"

	"github.com/Mitmadhu/broker/constants"
	router "github.com/Mitmadhu/commons/server"
	"github.com/Mitmadhu/mysqlDB/api"
	"github.com/Mitmadhu/mysqlDB/dto"
)

func InitRouter() {
	router.RouterMap = map[string]router.RouterRequest{
		"/user-exists":{
			DTO: &dto.ValidateUserRequest{},
			Method: http.MethodPost,
			Handler: api.ValidateUser,
			ValidationType: constants.JWTValidation,
		},
	}
	router.Routers()
}
