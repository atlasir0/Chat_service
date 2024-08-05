package converter

import (
	"github.com/atlasir0/Chat_service/Auth_chat/internal/model"
	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/auth_v1"
)

func ToUserClaimsFromLogin(req *desc.LoginRequest) *model.UserClaims {
	return &model.UserClaims{
		Username: req.GetUsername(),
	}
}
