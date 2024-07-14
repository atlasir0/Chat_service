package note

import (
	"context"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serv) Update(ctx context.Context, info *model.User, req *model.User) (*emptypb.Empty, error) {
	_, err := s.noteRepository.Update(ctx, info)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
