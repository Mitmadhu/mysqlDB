package api

import (
	"errors"
	"net/http"

	"github.com/Mitmadhu/broker/constants"
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

	if !validated {
		println(err.Error())
		err = errors.New(constants.UserNotFound)
	}
	if err != nil {
		helper.SendErrorResponse(w, req.MsgID, err.Error(), http.StatusUnauthorized)
		return
	}

	result := true
	helper.SendSuccessResponse(w, req.MsgID, dto.ValidateUserResponse{
		IsValid: &result,
	}, http.StatusOK)
}
