package note

import (
	"context"
	"log"

	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	id := req.GetId()

	_, err := i.noteService.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	log.Printf("deleted user with id: %d", id)
	return &emptypb.Empty{}, nil
}
