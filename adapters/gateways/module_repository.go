package gateways

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
	"gorm.io/gorm"
)

type ModuleRepository struct {
	DbHandler
}

func NewModuleRepository(dbHandler DbHandler) port.ModuleRepository {
	return &ModuleRepository{
		DbHandler: dbHandler,
	}
}

// 検索
func (repo *ModuleRepository) ModuleGets(ctx context.Context, modules *[]entities.Module, userID int) (*[]entities.Module, error) {

	result := repo.GetDB().Where("user_id = ?", userID).Order("id desc").Find(modules)

	if result.Error != nil {
		return modules, result.Error
	}

	return modules, nil
}

func (repo *ModuleRepository) ModuleCreate(ctx context.Context, module *entities.Module) (*entities.Module, error) {

	result := repo.GetDB().Create(module)

	if result.Error != nil {
		fmt.Println("fail")
		return module, result.Error
	}

	return module, nil
}

func (repo *ModuleRepository) ModuleUpdate(ctx context.Context, module *entities.Module, id string) (*entities.Module, error) {

	var selectedModuleForGorm entities.Module
	result := repo.GetDB().First(&selectedModuleForGorm, id)

	if result.Error != nil {
		return module, result.Error
	}

	module.CreatedAt = selectedModuleForGorm.CreatedAt
	module.UpdatedAt = time.Now()
	module.UserId = selectedModuleForGorm.UserId
	module.Id = selectedModuleForGorm.Id

	repo.GetDB().Save(module)

	if result.Error != nil {
		return module, result.Error
	}

	return module, nil
}

func (repo *ModuleRepository) ModuleDelete(ctx context.Context, module *entities.Module, moduleId string) error {

	result := repo.GetDB().Where("id = ?", moduleId).Delete(module)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ModuleRepository) ExistsModuleByID(ctx context.Context, module *entities.Module, ModuleID string) (bool, error) {
	result := r.GetDB().First(module, ModuleID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, result.Error
}
