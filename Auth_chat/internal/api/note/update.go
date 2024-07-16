package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	userModel := &model.User{
		ID: req.Id,
	}

	if req.Name != nil {
		userModel.Name = req.Name.Value
	}

	if req.Email != nil {
		userModel.Email = req.Email.Value
	}

	updatedUser := &model.User{
		ID:    req.Id,
		Name:  userModel.Name,
		Email: userModel.Email,
		Role:  userModel.Role,
	}

	_, err := i.noteService.Update(ctx, userModel, updatedUser)
	if err != nil {
		return nil, err
	}

	log.Printf("updated user with id: %d", req.Id)

	return &emptypb.Empty{}, nil
}
