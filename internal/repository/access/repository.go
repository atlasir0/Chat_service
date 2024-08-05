package auth

import (
	"context"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository/access/model"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.AccessRepository {
	return &repo{db: dbClient}
}

func (r *repo) Roles(ctx context.Context) (map[string]string, error) {
	accessibleRoles := make(map[string]string)
	accessibleRoles[model.ExamplePath] = "admin"
	return accessibleRoles, nil
}
