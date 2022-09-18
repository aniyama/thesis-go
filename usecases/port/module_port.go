package port

import (
	"context"

	"github.com/aniyama/thesis-go/entities"
)

//データ受け取り
type ModuleInputPort interface {
	ModuleGets(ctx context.Context, modules *[]entities.Module, userID int)
	ModuleCreate(ctx context.Context, module *entities.Module)
	ModuleUpdate(ctx context.Context, module *entities.Module, id string)
	ModuleDelete(ctx context.Context, module *entities.Module, ModuleID string)
}

//データ吐き出し
type ModuleOutputPort interface {
	OutputModules(modules *[]entities.Module)
	OutputModule(module *entities.Module)
	Output200()
	ErrorsPort
}

//受け取ったデータ処理
type ModuleRepository interface {
	ModuleGets(ctx context.Context, modules *[]entities.Module, userID int) (*[]entities.Module, error)
	ModuleCreate(ctx context.Context, module *entities.Module) (*entities.Module, error)
	ModuleUpdate(ctx context.Context, module *entities.Module, id string) (*entities.Module, error)
	ModuleDelete(ctx context.Context, module *entities.Module, moduleID string) error
	ExistsModuleByID(ctx context.Context, module *entities.Module, moduleID string) (bool, error)
}

type ModuleValidator interface {
	ModuleCreateValidator(request entities.Module) error
	ModuleUpdateValidator(request entities.Module) error
	ModuleIDValidator(moduleID string) error
	ModuleDeleteValidator(moduleID string) error
}
