package note

import (
	"context"
	"log"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	user := req.GetUser()
	userModel := &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     int(user.Role),
	}

	id, err := i.noteService.Create(ctx, userModel)
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

/// dsafasdfg
