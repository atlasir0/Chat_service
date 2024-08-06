package note

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/client/db"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note/converter"
	modelRepo "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	//Хэш пароля по DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)

	// Делаем запрос на вставку записи в таблицу auth
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, passwordColumn, emailColumn, roleColumn).
		Values(user.Name, user.Password, user.Email, user.Role).
		Suffix("RETURNING " + idColumn)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		log.Printf("failed to created user: %v", err)
		return 0, err
	}

	return userID, nil
}

func (r *repo) Get(ctx context.Context, filter modelRepo.UserFilter) (*model.User, error) {
	builderSelectOne := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn, passwordColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	if filter.ID != nil {
		builderSelectOne = builderSelectOne.Where(sq.Eq{idColumn: filter.ID})
	}

	if filter.Name != nil {
		builderSelectOne = builderSelectOne.Where(sq.Eq{nameColumn: filter.Name})
	}

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var getUser modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &getUser, q, args...)
	if err != nil {
		log.Printf("failed to ScanOneContext: %v", err)
		return nil, err
	}
	return converter.ToNoteFromRepo(&getUser), nil
}

func (r *repo) Update(ctx context.Context, user *model.User) error {
	// Делаем запрос на обновление записи в таблице auth
	builderUpdate := sq.Update(tableName).PlaceholderFormat(sq.Dollar)
	if len(user.Name) > 0 {
		builderUpdate = builderUpdate.Set(nameColumn, user.Name)
	}
	if len(user.Email) > 0 {
		builderUpdate = builderUpdate.Set(emailColumn, user.Email)
	}
	if user.Role != 0 {
		builderUpdate = builderUpdate.Set(roleColumn, user.Role)
	}

	builderUpdate = builderUpdate.Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to updated user: %v", err)
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) (*emptypb.Empty, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "note_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
