package access

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
)

type serverAccess struct {
	accessRepository repository.AccessRepository
	txManager        db.TxManager
}

func NewService(
	accessRepository repository.AccessRepository,
	txManager db.TxManager,
) service.AccessService {
	return &serverAccess{
		accessRepository: accessRepository,
		txManager:        txManager,
	}
}
