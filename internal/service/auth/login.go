package auth

import (
	"context"
	"errors"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	modelRepo "github.com/atlasir0/Chat_service/Auth_chat/internal/repository/note/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// Login-..
func (s *service) Login(ctx context.Context, login *model.Login) (string, error) {
    // Лезем в базу за данными пользователя
    user, err := s.userRepository.Get(ctx, modelRepo.UserFilter{Name: &login.Username})
    if err != nil {
        return "", err
    }

    // Сверяем хэши пароля
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
    if err != nil {
        return "", errors.New("invalid password")
    }

    refreshTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, refreshTokenName)
    if err != nil {
        return "", errors.New("key receipt error")
    }

    refreshToken, err := utils.GenerateToken(model.UserClaims{
        Username: user.Name,
        Role:     int32(user.Role), // Преобразование int в int32
    },
        []byte(refreshTokenSecretKey),
        refreshTokenExpiration,
    )
    if err != nil {
        return "", errors.New("failed to generate token")
    }

    return refreshToken, nil
}