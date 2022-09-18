package interactor

import (
	"context"
	"fmt"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type LoginInteractor struct {
	OutputPort port.LoginOutputPort
	Repository port.LoginRepository
	DbHandler  gateways.DbHandler
	Auth       port.AuthService
}

func NewLoginInputPort(output port.LoginOutputPort, repository port.LoginRepository, dbHandler gateways.DbHandler, auth port.AuthService) port.LoginInputPort {
	return &LoginInteractor{
		OutputPort: output,
		Repository: repository,
		DbHandler:  dbHandler,
		Auth:       auth,
	}
}

func (i *LoginInteractor) Login(ctx context.Context, user *entities.User) {
	user, err := i.Repository.Login(ctx, user)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}

	tokenString, err := i.Auth.CreateToken(int(user.Id))

	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}

	i.Auth.SetCookie(tokenString)
	i.OutputPort.Output200()
}

func (i *LoginInteractor) Register(ctx context.Context, User *entities.User) {
	userID, err := i.Repository.CreateUser(ctx, User)
	if err != nil {

		i.OutputPort.Output500Error(err)
		return
	}
	fmt.Println("register_controller", userID, err)
	tokenString, err := i.Auth.CreateToken(int(userID))
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}

	i.Auth.SetCookie(tokenString)

	i.OutputPort.Output200()
}

func (i *LoginInteractor) Logout(ctx context.Context) {
	i.OutputPort.Output200()
}
