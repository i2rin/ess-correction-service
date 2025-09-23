package usecase

import (
	"ess-server/internal/core/dto/auth"
)

type IAuthUsecase interface {
	Login(i *auth.LoginInDTO) (*auth.LoginOutDTO, error)
}
