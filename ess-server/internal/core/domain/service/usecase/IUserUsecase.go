package usecase

import (
	"ess-server/internal/core/dto/user"
)

type IUserUsecase interface {
	CreateUser(user *user.UserDTO) error
	GetUser(userId string) ([]*user.UserDTO, error)
	UpdateUser(user *user.UserDTO) error
	DeleteUser(userId string) error
}