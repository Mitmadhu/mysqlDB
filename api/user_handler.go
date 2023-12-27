package api

import (
	"net/http"

	"github.com/Mitmadhu/commons/constants"
	"github.com/Mitmadhu/commons/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
	"github.com/Mitmadhu/mysqlDB/dto"
)

func ValidateUser(w http.ResponseWriter, rawReq interface{}) {
	req, ok := rawReq.(*dto.ValidateUserRequest)
	if !ok {
		helper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// validate password
	u := model.User{}
	validated, err := u.ValidateUser(req.Username, req.Password)

	if err != nil && err.Error() != constants.UserNotFound{
		println(err.Error())
		helper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}

	result := true
	if !validated {
		result = false
	}

	resp := dto.ValidateUserResponse{
		BaseResponse: dto.BaseResponse{
			MsgID:      req.MsgID,
			StatusCode: http.StatusOK,
			Success:    true,
		},
		IsValid: &result,
	}
	helper.SendSuccessResponse(w, resp, http.StatusOK)
}

func RegisterUser(w http.ResponseWriter, rawReq interface{}){
	req, ok := rawReq.(*dto.RegisterUserRequest)
	if(!ok){
		helper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
	}

	// check if user already exist
	u := model.User{}
	_ , err := u.GetUserByUsername(req.Username)
	if err == nil {
		helper.SendErrorResponse(w, req.MsgID, constants.UsernameExists, http.StatusIMUsed)
		return
	}
	if err.Error() != constants.UserNotFound {
		helper.SendErrorResponse(w, req.MsgID, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// create user if not exist
	err = u.Register(req.Username, req.Password, req.FirstName, req.LastName, req.Age)

	if err != nil {
		println(err.Error())
		helper.SendErrorResponse(w, req.MsgID, err.Error(), http.StatusInternalServerError)
		return
	}
	// send response

	helper.SendSuccessResponse(w, dto.RegisterUserResponse{
		BaseResponse: dto.BaseResponse{
			MsgID: req.MsgID,
			Success: true,
			StatusCode: http.StatusAccepted,
			Message: "constants.UserCreated",
		},
	}, http.StatusAccepted)


}
