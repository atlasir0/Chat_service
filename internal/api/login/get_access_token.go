package login

import (
	"context"

	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	obj, err := i.loginService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: obj,
	}, nil
}
