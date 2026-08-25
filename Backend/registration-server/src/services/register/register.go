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
			ResponseCode:    common.ResponseCodeRegister,
		},
		StatusCode: RegisterCodeUnknown,
	}

	req.Normalize()

	if err := req.Validate(); err != nil {
		res.HttpCode = common.HttpCodeBadRequest
		res.StatusCode = RegisterCodeInvalidInput
		res.Message = err.Error()
		return res
	}

	accounts := collections.GetAccountsCollection()

	err := accounts.FindOne(ctx, bson.M{"email": req.Email}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodeInternalServerError
		res.Message = "Failed to register account"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodeConflict
		res.StatusCode = RegisterCodeEmailTaken
		res.Message = "Email already registered"
		return res
	}

	err = accounts.FindOne(ctx, bson.M{"username": req.Username}).Err()
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		res.HttpCode = common.HttpCodeInternalServerError
		res.Message = "Failed to register account"
		return res
	}
	if err == nil {
		res.HttpCode = common.HttpCodeConflict
		res.StatusCode = RegisterCodeUsernameTaken
		res.Message = "Username already taken"
		return res
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpCode = common.HttpCodeInternalServerError
		res.Message = "Failed to register account"
		return res
	}

	now := time.Now().UTC()
	account := &models.Account{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashed),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if _, err := accounts.InsertOne(ctx, account); err != nil {
		res.HttpCode = common.HttpCodeInternalServerError
		res.Message = "Failed to register account"
		return res
	}

	res.HttpCode = common.HttpCodeCreated
	res.StatusCode = RegisterCodeUserRegisteredSuccessfully
	res.Message = "Account registered successfully"
	return res
}
