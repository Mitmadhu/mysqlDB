package dto

import (
	"net/http"

	"github.com/Mitmadhu/commons/helper"
)

type ValidateUserRequest struct {
	BaseRequest
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type ValidateUserResponse struct {
	BaseResponse
	IsValid *bool `json:"is_valid"`
}

type RegisterUserRequest struct{
	BaseRequest
	Username string `json:"username"`
	LastName string `json:"last_name"`
	FirstName string `json:"first_name"`
	Password string `json:"password"`
}

type RegisterUserResponse struct{
	BaseResponse
}

func (v ValidateUserRequest) HasError(w http.ResponseWriter) bool {
	errs := []error{}
	helper.CheckEmpty(v.MsgID, &errs, "msg_id")
	helper.CheckEmpty(v.Username, &errs, "username")
	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}
