package port

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
)

//データ受け取り
type CategoryInputPort interface {
	CategoryGets(ctx context.Context, categories *[]entities.Category)
}

//データ吐き出し
type CategoryOutputPort interface {
	OutputCategories(categories *[]entities.Category)
	ErrorsPort
}

//受け取ったデータ処理
type CategoryRepository interface {
	CategoryGets(ctx context.Context, categorys *[]entities.Category) (*[]entities.Category, error)
}

type CategoryValidator interface {
	CategoryCreateValidator(request entities.Category) error
}
