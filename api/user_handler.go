package api

import (
	"net/http"

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

	if err != nil || !validated {
		helper.SendErrorResponse(w, req.MsgId, err.Error(), http.StatusUnauthorized)
		return
	}

	helper.SendSuccessResponse(w, req.MsgId, dto.ValidateUserResponse{
		IsValid: true,
	}, http.StatusOK)
}
