package note

import (
	"context"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	modelRepo "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.noteRepository.Get(ctx, modelRepo.UserFilter{ID: &id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
