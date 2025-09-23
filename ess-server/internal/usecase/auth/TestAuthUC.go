package auth

import (
	"ess-server/internal/core/dto/auth"
)

type TestAuthUC struct{}

func NewTestAuthUC() *TestAuthUC {
	return &TestAuthUC{}
}

func (u *TestAuthUC) Login(i *auth.LoginInDTO) (*auth.LoginOutDTO, error) {

	if i.GetUserID() == "" || i.GetPassword() == "" {
		return nil, nil
	}
	return &auth.LoginOutDTO{Role: "admin"}, nil
}
