package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/create"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[create.CreateCharacterBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &create.CreateCharacterResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_CreateCharacter,
				Message:         "invalid request body",
			},
			StatusCode: create.CreateCharacterCodes_InvalidInput,
		})
		return
	}

	res := create.CreateCharacter(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
