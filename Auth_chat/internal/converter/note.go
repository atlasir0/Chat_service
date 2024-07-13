package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

func ToNoteFromService(note *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.User{
		Id:    note.ID,
		Name:  note.Name,
		Email: note.Email,

		Role:      desc.Role(note.Role),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
