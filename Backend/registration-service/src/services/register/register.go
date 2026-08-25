package register

import (
	"context"
	"errors"
	"time"

	"github.com/mustafa-sibai/chronicle/backend-lib/collections"
	"github.com/mustafa-sibai/chronicle/backend-lib/common"
	models "github.com/mustafa-sibai/chronicle/backend-lib/types"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterAccount(ctx context.Context, req RegisterRequest) *RegisterResponse {
	config := common.GetConfig()

	res := &RegisterResponse{
		BaseResponse: common.BaseResponse{
			ApplicationName: config.ApplicationName,
			EnvironmentType: config.EnvironmentType,
			ResponseCode:    common.ResponseCodes_Register,
		},
		StatusCode: RegisterCodes_Unknown,
	}

	req.Body.Normalize()

	if err := req.Body.Validate(); err != nil {
		res.HttpCode = common.HttpCodes_BadRequest
		res.StatusCode = RegisterCodes_InvalidInput
		res.Message = err.Error()
		return res
	}

	accountsDBCollection := collections.GetAccountsCollection()

	err := accountsDBCollection.FindOne(ctx, bson.M{"email": req.Body.Email}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to register account"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodes_Conflict
		res.StatusCode = RegisterCodes_EmailTaken
		res.Message = "Email already registered"
		return res
	}

	err = accountsDBCollection.FindOne(ctx, bson.M{"username": req.Body.Username}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to register account"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodes_Conflict
		res.StatusCode = RegisterCodes_UsernameTaken
		res.Message = "Username already taken"
		return res
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Body.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to register account"
		return res
	}

	now := time.Now().UTC()
	account := &models.Account{
		Username:  req.Body.Username,
		Email:     req.Body.Email,
		Password:  string(hashed),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if _, err := accountsDBCollection.InsertOne(ctx, account); err != nil {
		res.HttpCode = common.HttpCodes_InternalServerError
		res.Message = "Failed to register account"
		return res
	}

	res.HttpCode = common.HttpCodes_Created
	res.StatusCode = RegisterCodes_UserRegisteredSuccessfully
	res.Message = "Account registered successfully"
	return res
}
