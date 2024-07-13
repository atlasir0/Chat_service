package note

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
)

type serv struct {
	noteRepository repository.UserRepository
}

func NewService(noteRepository repository.UserRepository) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}
