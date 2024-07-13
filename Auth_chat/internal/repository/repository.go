package repository

import (
	"context"

	model "github.com/atlasir0/Chat_service/Auth_chat/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, req *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	// Update(ctx context.Context, req *model.User) (*emptypb.Empty, error)
	// Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
