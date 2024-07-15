package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/converter"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userObj, err := i.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %v, created_at: %v, updated_at: %v\n", userObj.ID, userObj.Name, userObj.Email, userObj.Role, userObj.CreatedAt, userObj.UpdatedAt)

	return &desc.GetResponse{
		Info: converter.ToNoteFromService(userObj),
	}, nil
}
