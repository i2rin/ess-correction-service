package auth

import (
	"ess-server/internal/core/dto/auth"
	abstract "ess-server/internal/core/domain/service/usecase"
)

type AuthUsecase struct {}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

var _ abstract.IAuthUsecase = (*AuthUsecase)(nil)

func (u *AuthUsecase) Login(i *auth.LoginInDTO) (*auth.LoginOutDTO, error) {
	return &auth.LoginOutDTO{Role: "user"}, nil
}