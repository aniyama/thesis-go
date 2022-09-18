package interactor

import (
	"context"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type ModuleInteractor struct {
	OutputPort port.ModuleOutputPort
	Repository port.ModuleRepository
}

func NewModuleInputPort(output port.ModuleOutputPort, repository port.ModuleRepository, dbHandler gateways.DbHandler) port.ModuleInputPort {
	return &ModuleInteractor{
		OutputPort: output,
		Repository: repository,
	}
}

func (i *ModuleInteractor) ModuleGets(ctx context.Context, modules *[]entities.Module, userID int) {
	modules, err := i.Repository.ModuleGets(ctx, modules, userID)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputModules(modules)
}

func (i *ModuleInteractor) ModuleCreate(ctx context.Context, Module *entities.Module) {
	createdModule, err := i.Repository.ModuleCreate(ctx, Module)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputModule(createdModule)
}

func (i *ModuleInteractor) ModuleUpdate(ctx context.Context, Module *entities.Module, id string) {
	updatedModule, err := i.Repository.ModuleUpdate(ctx, Module, id)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputModule(updatedModule)
}

func (i *ModuleInteractor) ModuleDelete(ctx context.Context, module *entities.Module, ModuleID string) {
	err := i.Repository.ModuleDelete(ctx, module, ModuleID)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}

	i.OutputPort.Output200()
}
