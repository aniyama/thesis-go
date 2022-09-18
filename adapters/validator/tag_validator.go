package validator

import (
	"errors"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type TagValidator struct {
	ctx       *gin.Context
	dbHandler gateways.DbHandler
	TagID     string
}

func NewTagValidator(ctx *gin.Context, dbHandler gateways.DbHandler) port.TagValidator {
	return &TagValidator{
		ctx:       ctx,
		dbHandler: dbHandler,
		TagID:     "",
	}
}

func (cm *TagValidator) TagIDValidator(TagID string) error {
	return validation.Validate(TagID,
		validation.Required.Error("リクエストが正しくありません"),
	)
}

func (cm *TagValidator) TagDeleteValidator(TagID string) error {
	return validation.Validate(TagID,
		validation.Required.Error("リクエストが正しくありません"),
		validation.By(cm.validateTagID),
	)
}

func (cm *TagValidator) TagCreateValidator(request entities.Tag) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Title,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 30).Error("タイトルは1文字から、30文字の間で指定してください"),
		),
	)
}

func (cm *TagValidator) TagUpdateValidator(request entities.Tag) error {
	return validation.ValidateStruct(&request,
		validation.Field(
			&request.Title,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 30).Error("タイトルは1文字から、30文字の間で指定してください"),
		),
	)
}

func (cm *TagValidator) validateTagID(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return errors.New("リクエストが正しくありません")
	}

	var Tag *entities.Tag

	exists, err := gateways.NewTagRepository(cm.dbHandler).ExistsTagByID(cm.ctx.Request.Context(), Tag, val)
	if err != nil || !exists {
		return errors.New("リクエストが正しくありません")
	}

	return nil
}
