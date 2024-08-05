package access

import (
	"context"
	"errors"
	"strings"

	"github.com/atlasir0/Chat_service/Auth_chat/internal/utils"
	"google.golang.org/grpc/metadata"
)

const (
	grpcPort   = 50051
	authPrefix = "Bearer "

	refreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="
)

var accessibleRoles map[string]string

func (s *serverAccess) Check(ctx context.Context, endpointAddress string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[endpointAddress]
	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}

// Возвращает мапу с адресом эндпоинта и ролью, которая имеет доступ к нему
func (s *serverAccess) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		Roles, err := s.accessRepository.Roles(ctx)
		if err != nil {
			return nil, nil
		}
		accessibleRoles = Roles
	}

	return accessibleRoles, nil
}
