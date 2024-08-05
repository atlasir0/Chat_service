package login

import (
	"context"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/utils"
)

func (s *serverAuth) GetRefreshToken(ctx context.Context, token string) (string, error) {
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")
	claims, err := utils.VerifyToken(token, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}
	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", nil
	}

	// Преобразование model.UserInfo в model.UserClaims
	userClaims := model.UserClaims{
		Username: claims.Username,
		Role:     r,
	}

	refreshToken, err := utils.GenerateToken(userClaims, []byte(refreshTokenSecretKey), refreshTokenExpiration)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
