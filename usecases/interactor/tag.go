package interactor

import (
	"context"
	"fmt"

	"github.com/aniyama/thesis-go/adapters/gateways"
	"github.com/aniyama/thesis-go/entities"
	"github.com/aniyama/thesis-go/usecases/port"
)

type TagInteractor struct {
	OutputPort port.TagOutputPort
	Repository port.TagRepository
}

func NewTagInputPort(output port.TagOutputPort, repository port.TagRepository, dbHandler gateways.DbHandler) port.TagInputPort {
	return &TagInteractor{
		OutputPort: output,
		Repository: repository,
	}
}

func (i *TagInteractor) TagGets(ctx context.Context, tags *[]entities.Tag, userID int) {
	tags, err := i.Repository.TagGets(ctx, tags, userID)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputTags(tags)
}

func (i *TagInteractor) TagCreate(ctx context.Context, Tag *entities.Tag) {
	createdTag, err := i.Repository.TagCreate(ctx, Tag)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputTag(createdTag)
}

func (i *TagInteractor) TagUpdate(ctx context.Context, Tag *entities.Tag, id string) {
	fmt.Println("fail233333", Tag)
	updatedTag, err := i.Repository.TagUpdate(ctx, Tag, id)
	fmt.Println("fail233333", updatedTag, err)

	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}
	i.OutputPort.OutputTag(updatedTag)
}

func (i *TagInteractor) TagDelete(ctx context.Context, Tag *entities.Tag, TagID string) {
	err := i.Repository.TagDelete(ctx, Tag, TagID)
	if err != nil {
		i.OutputPort.Output500Error(err)
		return
	}

	i.OutputPort.Output200()
}
