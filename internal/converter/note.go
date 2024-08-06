package converter

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	auth "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteFromService(note *model.User) *desc.UserInfo {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.UserInfo{
		Id:        note.ID,
		Name:      note.Name,
		Email:     note.Email,
		Role:      desc.UserRole(note.Role),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToServiceLogin(login *auth.LoginRequest) *model.Login {
	return &model.Login{
		Username: login.Username,
		Password: login.Password,
	}
}

func ToUserFromDescCreate(user *desc.User) *model.User {
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     int(user.Role),
	}
}

func ToUserFromDescUpdate(req *desc.UpdateRequest) *model.User {
	user := &model.User{
		ID: req.Id,
	}

	if req.Name != nil {
		user.Name = req.Name.Value
	}

	if req.Email != nil {
		user.Email = req.Email.Value
	}
	user.Role = int(req.Role)

	return user
}
