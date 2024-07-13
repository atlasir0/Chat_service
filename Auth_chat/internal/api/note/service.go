package note

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedUserServiceServer
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
