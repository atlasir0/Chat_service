package repository

import (
	"context"

	model "github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	modelRepo "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.User) (int64, error)
	Get(ctx context.Context, filter modelRepo.UserFilter) (*model.User, error)
	Update(ctx context.Context, req *model.User) error
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}

type AccessRepository interface {
	Roles(ctx context.Context) (map[string]string, error)
}

type LoginRepository interface {
	Login(ctx context.Context, info *model.UserClaims) (string, error)
	GetAccessToken(ctx context.Context, token string) (string, error)
	GetRefreshToken(ctx context.Context, token string) (string, error)
	GetUserRole(ctx context.Context) (string, error)
}
