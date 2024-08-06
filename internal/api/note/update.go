package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/converter"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	userModel := converter.ToUserFromDescUpdate(req)

	err := i.noteService.Update(ctx, userModel)
	if err != nil {
		return nil, err
	}

	log.Printf("updated user with id: %d", userModel.ID)

	return &emptypb.Empty{}, nil
}
