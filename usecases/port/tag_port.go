package port

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
)

//データ受け取り
type TagInputPort interface {
	TagGets(ctx context.Context, tags *[]entities.Tag, userID int)
	TagCreate(ctx context.Context, Tag *entities.Tag)
	TagUpdate(ctx context.Context, Tag *entities.Tag, id string)
	TagDelete(ctx context.Context, Tag *entities.Tag, TagID string)
}

//データ吐き出し
type TagOutputPort interface {
	OutputTags(tags *[]entities.Tag)
	OutputTag(tag *entities.Tag)
	Output200()
	ErrorsPort
}

//受け取ったデータ処理
type TagRepository interface {
	TagGets(ctx context.Context, tags *[]entities.Tag, userID int) (*[]entities.Tag, error)
	TagCreate(ctx context.Context, tag *entities.Tag) (*entities.Tag, error)
	TagUpdate(ctx context.Context, tag *entities.Tag, id string) (*entities.Tag, error)
	TagDelete(ctx context.Context, tag *entities.Tag, tagID string) error
	ExistsTagByID(ctx context.Context, tag *entities.Tag, tagID string) (bool, error)
}

type TagValidator interface {
	TagCreateValidator(request entities.Tag) error
	TagUpdateValidator(request entities.Tag) error
	TagIDValidator(tagID string) error
	TagDeleteValidator(tagID string) error
}
