package login

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
)

type serverAuth struct {
	loginRepository repository.LoginRepository
	txManager       db.TxManager
}

func NewService(
	loginRepository repository.LoginRepository,
	txManager db.TxManager,
) service.LoginService {
	return &serverAuth{
		loginRepository: loginRepository,
		txManager:       txManager,
	}
}
