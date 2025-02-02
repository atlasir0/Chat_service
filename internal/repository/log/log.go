package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
)

const (
	tableName = "log_transaction"

	colLogID     = "id"
	colInfo      = "info"
	colCreatedAt = "created_at"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// CreateLog - ...
func (r *Repo) CreateLog(ctx context.Context, log *model.Log) error {
	// Делаем запрос на вставку записи в таблицу auth
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colInfo).
		Values(log.Text).
		Suffix("RETURNING " + colLogID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "log_repository.CreateLog",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
