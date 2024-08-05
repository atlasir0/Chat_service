package access

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/service"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
