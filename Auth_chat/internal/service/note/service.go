package note

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
)

type serv struct {
	noteRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}

func NewMockService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.noteRepository = s
		}
	}

	return &srv
}
