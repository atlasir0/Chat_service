package converter

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
	auth "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
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
