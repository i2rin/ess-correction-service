package usecase

import (
	"ess-server/internal/core/dto/user"
)

type IUserUsecase interface {
	CreateUser(user *user.UserInDTO) error
	GetUser(page, size int) ([]*user.UserOutDTO, error)
	UpdateUser(user *user.UserInDTO, userId string) error
	DeleteUser(userId string) error
}
