package note

import (
	"context"
	"fmt"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
)

func (s *serv) Update(ctx context.Context, user *model.User) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.noteRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.CreateLog(ctx, &model.Log{
			Text: fmt.Sprintf("User updated: %d", user.ID),
		})
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
