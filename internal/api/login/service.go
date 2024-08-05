package login

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	loginService service.LoginService
}

func NewImplementation(loginService service.LoginService) *Implementation {
	return &Implementation{
		loginService: loginService,
	}
}
