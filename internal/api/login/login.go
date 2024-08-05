package login

import (
	"context"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/converter"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	userClaims := converter.ToUserClaimsFromLogin(req)
	obj, err := i.loginService.Login(ctx, userClaims)
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: obj,
	}, nil
}
