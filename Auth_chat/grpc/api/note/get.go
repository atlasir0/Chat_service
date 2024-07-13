package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/converter"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	noteObj, err := i.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	log.Printf("id: %d, title: %s, body: %s, created_at: %v, updated_at: %v\n", noteObj.ID, noteObj.Name, noteObj.Email, noteObj.Password, noteObj.Role, noteObj.CreatedAt, noteObj.UpdatedAt)

	return &desc.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}

