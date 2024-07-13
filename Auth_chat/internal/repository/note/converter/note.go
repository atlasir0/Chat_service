package converter

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	modelRepo "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.User) *model.User {
	return &model.User{
		ID:        note.ID,
		Name:      note.Name,
		Email:     note.Email,
		Password:  note.Password,
		Role:      note.Role,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}
