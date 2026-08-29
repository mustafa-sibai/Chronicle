package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/list"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[list.ListCharactersBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &list.ListCharactersResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_ListCharacters,
				Message:         "invalid request body",
			},
			StatusCode: list.ListCharactersCodes_InvalidInput,
		})
		return
	}

	res := list.ListCharacters(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
