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

// LogRepository - ...
type LogRepository interface {
	CreateLog(ctx context.Context, log *model.Log) error
}

// PermRepository - ...
type PermRepository interface {
	GetPermission(ctx context.Context) ([]*model.Permission, error)
}

// SecretRepository - ...
type SecretRepository interface {
	GetKeyTokens(ctx context.Context, tokenName string) (string, error)
}
