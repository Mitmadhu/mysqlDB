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
		IsValid: result,
	}
	helper.SendSuccessResponse(w, resp, http.StatusOK)
}
