package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/delete"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[delete.DeleteCharacterBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &delete.DeleteCharacterResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_DeleteCharacter,
				Message:         "invalid request body",
			},
			StatusCode: delete.DeleteCharacterCodes_InvalidInput,
		})
		return
	}

	res := delete.DeleteCharacter(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
