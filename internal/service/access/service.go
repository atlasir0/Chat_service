package access

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	def "github.com/atlasir0/Chat_service/Auth_chat/internal/service"
)

const (
	authPrefix      = "Bearer "
	accessTokenName = "access"
)

var _ def.AccessService = (*service)(nil)

type service struct {
	permRepository   repository.PermRepository
	secretRepository repository.SecretRepository
}

// NewService - ...
func NewService(permRepository repository.PermRepository, secretRepository repository.SecretRepository) *service {
	return &service{
		permRepository:   permRepository,
		secretRepository: secretRepository,
	}
}
