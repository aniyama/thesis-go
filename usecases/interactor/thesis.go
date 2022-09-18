package interactor

import (
	"context"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type ThesisInteractor struct {
	OutputPort port.ThesisOutputPort
	Repository port.ThesisRepository
}

func NewThesisInputPort(output port.ThesisOutputPort, repository port.ThesisRepository, dbHandler gateways.DbHandler) port.ThesisInputPort {
	return &ThesisInteractor{
		OutputPort: output,
		Repository: repository,
	}
}

func (i *ThesisInteractor) ThesisGets(ctx context.Context, theses *[]entities.Thesis, userID int) {
	theses, err := i.Repository.ThesisGets(ctx, theses, userID)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputTheses(theses)
}

func (i *ThesisInteractor) ThesisCreate(ctx context.Context, thesis *entities.Thesis) {
	createdThesis, err := i.Repository.ThesisCreate(ctx, thesis)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputThesis(createdThesis)
}

func (i *ThesisInteractor) ThesisUpdate(ctx context.Context, thesis *entities.Thesis, id string) {
	updatedThesis, err := i.Repository.ThesisUpdate(ctx, thesis, id)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputThesis(updatedThesis)
}

func (i *ThesisInteractor) ThesisDelete(ctx context.Context, thesis *entities.Thesis, thesisID string) {
	err := i.Repository.ThesisDelete(ctx, thesis, thesisID)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}

	i.OutputPort.Output200()
}
