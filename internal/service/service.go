package service

import (
	"context"

	model "github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService interface {
	Create(ctx context.Context, info *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.User) error
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}

type AuthService interface {
	Login(ctx context.Context, login *model.Login) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, accessToken string) (string, error)
}

type AccessService interface {
	Check(ctx context.Context, address string) error
}
