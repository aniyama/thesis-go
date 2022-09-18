package validator

import (
	"errors"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ThesisValidator struct {
	ctx       *gin.Context
	dbHandler gateways.DbHandler
	thesisID  string
}

func NewThesisValidator(ctx *gin.Context, dbHandler gateways.DbHandler) port.ThesisValidator {
	return &ThesisValidator{
		ctx:       ctx,
		dbHandler: dbHandler,
		thesisID:  "",
	}
}

func (cm *ThesisValidator) ThesisIDValidator(thesisID string) error {
	return validation.Validate(thesisID,
		validation.Required.Error("リクエストが正しくありません"),
	)
}

func (cm *ThesisValidator) ThesisDeleteValidator(ThesisID string) error {
	return validation.Validate(ThesisID,
		validation.Required.Error("リクエストが正しくありません"),
		validation.By(cm.validateThesisID),
	)
}

func (cm *ThesisValidator) ThesisCreateValidator(request entities.Thesis) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.ThesisTitle,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 30).Error("タイトルは1文字から、30文字の間で指定してください"),
		),
		validation.Field(
			&request.ThesisContent,
			validation.Required.Error("論文内容は必須です"),
			validation.Length(1, 2000).Error("論文内容は1文字から、2000文字の間で指定してください"),
		),
	)
}

func (cm *ThesisValidator) ThesisUpdateValidator(request entities.Thesis) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.ThesisTitle,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 30).Error("タイトルは1文字から、30文字の間で指定してください"),
		),
		validation.Field(
			&request.ThesisContent,
			validation.Required.Error("論文内容は必須です"),
			validation.Length(1, 2500).Error("論文内容は1文字から、2500文字の間で指定してください"),
		),
	)
}

func (cm *ThesisValidator) validateThesisID(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return errors.New("リクエストが正しくありません")
	}

	var thesis *entities.Thesis

	exists, err := gateways.NewThesisRepository(cm.dbHandler).ExistsThesisByID(cm.ctx.Request.Context(), thesis, val)
	if err != nil || !exists {
		return errors.New("リクエストが正しくありません")
	}

	return nil
}
