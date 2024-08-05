package login

import (
	"context"

	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	obj, err := i.loginService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: obj,
	}, nil
}
