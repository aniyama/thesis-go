package validator

import (
	"errors"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ModuleValidator struct {
	ctx       *gin.Context
	dbHandler gateways.DbHandler
	ModuleID  string
}

func NewModuleValidator(ctx *gin.Context, dbHandler gateways.DbHandler) port.ModuleValidator {
	return &ModuleValidator{
		ctx:       ctx,
		dbHandler: dbHandler,
		ModuleID:  "",
	}
}

func (cm *ModuleValidator) ModuleIDValidator(ModuleID string) error {
	return validation.Validate(ModuleID,
		validation.Required.Error("リクエストが正しくありません"),
	)
}

func (cm *ModuleValidator) ModuleDeleteValidator(ModuleID string) error {
	return validation.Validate(ModuleID,
		validation.Required.Error("リクエストが正しくありません"),
		validation.By(cm.validateModuleID),
	)
}

func (cm *ModuleValidator) ModuleCreateValidator(request entities.Module) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Title,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 30).Error("タイトルは1文字から、30文字の間で指定してください"),
		),
		validation.Field(
			&request.Content,
			validation.Required.Error("論文内容は必須です"),
			validation.Length(1, 2000).Error("論文内容は1文字から、2000文字の間で指定してください"),
		),
	)
}

func (cm *ModuleValidator) ModuleUpdateValidator(request entities.Module) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Title,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 30).Error("タイトルは1文字から、30文字の間で指定してください"),
		),
		validation.Field(
			&request.Content,
			validation.Required.Error("論文内容は必須です"),
			validation.Length(1, 1000).Error("論文内容は1文字から、1000文字の間で指定してください"),
		),
	)
}

func (cm *ModuleValidator) validateModuleID(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return errors.New("リクエストが正しくありません")
	}

	var module *entities.Module

	exists, err := gateways.NewModuleRepository(cm.dbHandler).ExistsModuleByID(cm.ctx.Request.Context(), module, val)
	if err != nil || !exists {
		return errors.New("リクエストが正しくありません")
	}

	return nil
}
