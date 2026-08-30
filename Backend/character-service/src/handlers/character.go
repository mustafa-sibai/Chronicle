package handlers

import (
	"net/http"

	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/choose"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/create"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/delete"
	"github.com/mustafa-sibai/chronicle/character-service/services/character/list"
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

func ChooseHandler(w http.ResponseWriter, r *http.Request) {
	req, err := common.DecodeRequest[choose.ChooseCharacterBody](r)
	if err != nil {
		config := common.GetConfig()
		common.WriteJSON(w, common.HttpCodes_BadRequest, &choose.ChooseCharacterResponse{
			BaseResponse: common.BaseResponse{
				ApplicationName: config.ApplicationName,
				EnvironmentType: config.EnvironmentType,
				HttpCode:        common.HttpCodes_BadRequest,
				ResponseCode:    common.ResponseCodes_ChooseCharacter,
				Message:         "invalid request body",
			},
			StatusCode: choose.ChooseCharacterCodes_InvalidInput,
		})
		return
	}

	res := choose.ChooseCharacter(r.Context(), req)
	common.WriteJSON(w, res.HttpCode, res)
}
