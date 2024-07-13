package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/converter"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.Create) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
