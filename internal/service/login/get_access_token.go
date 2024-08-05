package login

import (
	"context"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/utils"
)

func (s *serverAuth) GetAccessToken(ctx context.Context, token string) (string, error) {
	accessTokenSecretKey := os.Getenv("accessTokenSecretKey")
	refreshTokenSecretKey := os.Getenv("refreshTokenSecretKey")
	claims, err := utils.VerifyToken(token, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}
	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", nil
	}

	userClaims := model.UserClaims{
		Username: claims.Username,
		Role:     r,
	}

	accessToken, err := utils.GenerateToken(userClaims, []byte(accessTokenSecretKey), accessTokenExpiration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
