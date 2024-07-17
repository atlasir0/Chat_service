package auth

import (
	"context"
	"errors"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	"github.com/atlasir0/Chat_service/Auth_chat/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAccessToken-...
func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	refreshTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, refreshTokenName)
	if err != nil {
		return "", errors.New("key receipt error")
	}

	claims, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	accessTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, accessTokenName)
	if err != nil {
		return "", errors.New("key receipt error")
	}

	accessToken, err := utils.GenerateToken(model.UserClaims{
		Username: claims.Username,
		Role:     claims.Role,
	},
		[]byte(accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
