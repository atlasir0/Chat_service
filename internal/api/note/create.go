package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/converter"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)


func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToUserFromDescCreate(req.GetUser()))
	if err != nil {
		return nil, err
	}
	log.Printf("inserted auth with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
