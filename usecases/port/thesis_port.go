package port

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
)

//データ受け取り
type ThesisInputPort interface {
	ThesisGets(ctx context.Context, theses *[]entities.Thesis, userID int)
	ThesisCreate(ctx context.Context, thesis *entities.Thesis)
	ThesisUpdate(ctx context.Context, thesis *entities.Thesis, id string)
	ThesisDelete(ctx context.Context, thesis *entities.Thesis, thesisID string)
}

//データ吐き出し
type ThesisOutputPort interface {
	OutputTheses(theses *[]entities.Thesis)
	OutputThesis(thesis *entities.Thesis)
	Output200()
	ErrorsPort
}

//受け取ったデータ処理
type ThesisRepository interface {
	ThesisGets(ctx context.Context, theses *[]entities.Thesis, userID int) (*[]entities.Thesis, error)
	ThesisCreate(ctx context.Context, thesis *entities.Thesis) (*entities.Thesis, error)
	ThesisUpdate(ctx context.Context, thesis *entities.Thesis, id string) (*entities.Thesis, error)
	ThesisDelete(ctx context.Context, thesis *entities.Thesis, thesisID string) error
	ExistsThesisByID(ctx context.Context, thesis *entities.Thesis, thesisID string) (bool, error)
}

type ThesisValidator interface {
	ThesisCreateValidator(request entities.Thesis) error
	ThesisUpdateValidator(request entities.Thesis) error
	ThesisIDValidator(ThesisID string) error
	ThesisDeleteValidator(ThesisID string) error
}
