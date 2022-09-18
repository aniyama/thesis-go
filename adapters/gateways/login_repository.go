package gateways

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type LoginRepository struct {
	DbHandler
}

func NewLoginRepository(dbHandler DbHandler) port.LoginRepository {
	return &LoginRepository{
		DbHandler: dbHandler,
	}
}

func (repo *LoginRepository) Login(ctx context.Context, user *entities.User) (*entities.User, error) {

	db := repo.GetDB()

	var userForGorm entities.User
	result := db.Where("name = ?", user.Name).Where("password = ?", user.Password).First(&userForGorm)

	if result.Error != nil {
		return &userForGorm, result.Error
	}

	return &userForGorm, nil
}

func (repo *LoginRepository) GetUserByName(ctx context.Context, name string) (user *entities.User, err error) {

	db := repo.GetDB()

	var userForGorm entities.User
	result := db.Where(&entities.User{Name: name}).First(&userForGorm)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil

}

func (repo *LoginRepository) CreateUser(ctx context.Context, user *entities.User) (userId int64, err error) {

	db := repo.GetDB()

	result := db.Create(&user)

	if result.Error != nil {
		return int64(user.Id), result.Error
	}

	return int64(user.Id), nil

}
