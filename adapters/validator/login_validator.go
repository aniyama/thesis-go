package validator

import (
	"errors"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginValidator struct {
	ctx       *gin.Context
	dbHandler gateways.DbHandler
	userID    string
}

func NewLoginValidator(ctx *gin.Context, dbHandler gateways.DbHandler) port.LoginValidator {
	return &LoginValidator{
		ctx:       ctx,
		dbHandler: dbHandler,
		userID:    "",
	}
}

func (cm *LoginValidator) LoginValidator(request entities.User) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Name,
			validation.Required.Error("ユーザー名は必須です"),
			validation.Length(1, 20).Error("ユーザー名は1文字から、20文字の間で指定してください"),
		),
		validation.Field(
			&request.Password,
			validation.Required.Error("パスワードは必須です"),
			validation.Length(1, 50).Error("パスワードは1文字から、20文字の間で指定してください"),
		),
	)
}

func (cm *LoginValidator) RegisterValidator(request entities.User) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Name,
			validation.Required.Error("ユーザー名は必須です"),
			validation.Length(1, 20).Error("ユーザー名は1文字から、20文字の間で指定してください"),
		),
		validation.Field(
			&request.Password,
			validation.Required.Error("パスワードは必須です"),
			validation.Length(1, 50).Error("パスワードは1文字から、20文字の間で指定してください"),
		),
	)
}

func (cm *LoginValidator) ExistsUserByUserName(userName string) error {
	user, err := gateways.NewLoginRepository(cm.dbHandler).GetUserByName(cm.ctx.Request.Context(), userName)
	// ユーザーがいたらアウト
	if err == nil || user != nil {
		return errors.New("同じユーザーが存在します...")
	}

	return nil
}
