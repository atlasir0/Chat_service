
package note

import (
	"context"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.User) (int64, error) {
	id, err := s.noteRepository.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
