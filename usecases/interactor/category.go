package interactor

import (
	"context"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type CategoryInteractor struct {
	OutputPort port.CategoryOutputPort
	Repository port.CategoryRepository
}

func NewCategoryInputPort(output port.CategoryOutputPort, repository port.CategoryRepository, dbHandler gateways.DbHandler) port.CategoryInputPort {
	return &CategoryInteractor{
		OutputPort: output,
		Repository: repository,
	}
}

func (i *CategoryInteractor) CategoryGets(ctx context.Context, categories *[]entities.Category) {
	categories, err := i.Repository.CategoryGets(ctx, categories)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputCategories(categories)
}
