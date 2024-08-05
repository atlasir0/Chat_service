package login

import (
	"context"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/utils"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 1 * time.Minute
)

func (s *serverAuth) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")
	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", nil
	}

	// Преобразование model.UserInfo в model.UserClaims
	userClaims := model.UserClaims{
		Username: info.Username,
		Role:     r,
	}

	refreshToken, err := utils.GenerateToken(userClaims, []byte(refreshTokenSecretKey), refreshTokenExpiration)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
