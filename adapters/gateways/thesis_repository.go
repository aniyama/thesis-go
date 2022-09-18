package gateways

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"github.com/jinzhu/gorm"
)

type ThesisRepository struct {
	DbHandler
}

func NewThesisRepository(dbHandler DbHandler) port.ThesisRepository {
	return &ThesisRepository{
		DbHandler: dbHandler,
	}
}

// 検索
func (repo *ThesisRepository) ThesisGets(ctx context.Context, theses *[]entities.Thesis, userID int) (*[]entities.Thesis, error) {

	var a []entities.Thesis
	result := repo.GetDB().Where("user_id = ?", userID).Order("id desc").Find(theses)
	fmt.Printf("%T\n", theses)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	if result.Error != nil {
		fmt.Println("#########", &a, theses)
		return theses, result.Error
	}

	return theses, nil
}

func (repo *ThesisRepository) ThesisCreate(ctx context.Context, thesis *entities.Thesis) (*entities.Thesis, error) {

	result := repo.GetDB().Create(thesis)

	if result.Error != nil {
		fmt.Println("fail")
		return thesis, result.Error
	}

	return thesis, nil
}

func (repo *ThesisRepository) ThesisUpdate(ctx context.Context, thesis *entities.Thesis, id string) (*entities.Thesis, error) {

	var selectedThesisForGorm entities.Thesis

	result := repo.GetDB().First(&selectedThesisForGorm, id)

	if result.Error != nil {
		return thesis, result.Error
	}

	thesis.CreatedAt = selectedThesisForGorm.CreatedAt
	thesis.UpdatedAt = time.Now()
	thesis.UserId = selectedThesisForGorm.UserId
	thesis.Id = selectedThesisForGorm.Id

	repo.GetDB().Save(thesis)

	fmt.Println(result.Error, thesis)

	if result.Error != nil {
		return thesis, result.Error
	}

	return thesis, nil
}

// func (repo *ThesisRepository) ThesisUpdate(ctx context.Context, id string, thesis *entities.Thesis, userId int, c *gin.Context) (*entities.Thesis, error) {

// 	result := repo.GetDB().First(&thesis, id)

// 	thesis.UserId = userId

// 	thesis.UpdatedAt = time.Now()
// 	err := c.BindJSON(&thesis)
// 	if err != nil {
// 		panic("unMarchal")
// 	}

// 	repo.GetDB().Save(&thesis)

// 	if result.Error != nil {
// 		return thesis, result.Error
// 	}

// 	return thesis, nil
// }

func (repo *ThesisRepository) ThesisDelete(ctx context.Context, thesis *entities.Thesis, thesisID string) error {

	result := repo.GetDB().Where("id = ?", thesisID).Delete(thesis)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ThesisRepository) ExistsThesisByID(ctx context.Context, thesis *entities.Thesis, thesisID string) (bool, error) {
	result := r.GetDB().First(thesis, thesisID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, result.Error
}
