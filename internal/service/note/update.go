package note

import (
	"context"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
)

func (s *serv) Update(ctx context.Context, user *model.User) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.noteRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
