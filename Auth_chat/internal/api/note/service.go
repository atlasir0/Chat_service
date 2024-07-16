package note

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedUserServiceServer
	noteService service.UserService
}

func NewImplementation(noteService service.UserService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
