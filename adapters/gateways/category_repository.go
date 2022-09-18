package gateways

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type CategoryRepository struct {
	DbHandler
}

func NewCategoryRepository(dbHandler DbHandler) port.CategoryRepository {
	return &CategoryRepository{
		DbHandler: dbHandler,
	}
}

// 検索
func (repo *CategoryRepository) CategoryGets(ctx context.Context, categories *[]entities.Category) (*[]entities.Category, error) {

	result := repo.GetDB().Find(categories)

	if result.Error != nil {
		return categories, result.Error
	}

	return categories, nil
}
