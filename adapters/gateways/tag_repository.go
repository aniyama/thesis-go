package gateways

import (
	"context"
	"errors"
	"fmt"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/jinzhu/gorm"
)

type TagRepository struct {
	DbHandler
}

func NewTagRepository(dbHandler DbHandler) port.TagRepository {
	return &TagRepository{
		DbHandler: dbHandler,
	}
}

// 検索
func (repo *TagRepository) TagGets(ctx context.Context, tags *[]entities.Tag, userID int) (*[]entities.Tag, error) {

	result := repo.GetDB().Where("user_id = ?", userID).Order("id desc").Find(tags)

	if result.Error != nil {
		return tags, result.Error
	}

	return tags, nil
}

func (repo *TagRepository) TagCreate(ctx context.Context, Tag *entities.Tag) (*entities.Tag, error) {

	result := repo.GetDB().Create(Tag)

	if result.Error != nil {
		fmt.Println("fail")
		return Tag, result.Error
	}

	return Tag, nil
}

func (repo *TagRepository) TagUpdate(ctx context.Context, tag *entities.Tag, id string) (*entities.Tag, error) {

	var selectedTagForGorm entities.Tag
	result := repo.GetDB().First(&selectedTagForGorm, id)

	if result.Error != nil {
		return tag, result.Error
	}

	repo.GetDB().Save(tag)

	if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

func (repo *TagRepository) TagDelete(ctx context.Context, Tag *entities.Tag, TagId string) error {
	result := repo.GetDB().Where("id = ?", TagId).Delete(Tag)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *TagRepository) ExistsTagByID(ctx context.Context, Tag *entities.Tag, TagID string) (bool, error) {
	result := r.GetDB().First(Tag, TagID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, result.Error
}
