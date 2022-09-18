package port

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
)

//データ受け取り
type LoginInputPort interface {
	Login(ctx context.Context, user *entities.User)
	Register(ctx context.Context, user *entities.User)
	Logout(ctx context.Context)
}

//データ吐き出し
type LoginOutputPort interface {
	Output200()
	ErrorsPort
}

//受け取ったデータ処理
type LoginRepository interface {
	Login(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByName(ctx context.Context, name string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) (int64, error)
}

type LoginValidator interface {
	LoginValidator(request entities.User) error
	RegisterValidator(request entities.User) error
	ExistsUserByUserName(request string) error
}

type AuthService interface {
	CreateToken(userId int) (tokenString string, err error)
	SetCookie(tokenString string)
}
