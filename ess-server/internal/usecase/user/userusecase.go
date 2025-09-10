package user

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/user"
)

type UserUsecase struct {}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

var _ abstract.IUserUsecase = (*UserUsecase)(nil)

func (u *UserUsecase) CreateUser(user *dto.UserDTO) error {
	return nil
}

func (u *UserUsecase) GetUser(userId string) ([]*dto.UserDTO, error) {
    return []*dto.UserDTO{
        {
            Name:     "John Doe",
            Mail:    "john.doe@example.com",
            Password: "password1",
            Role:     "admin",
        },
        {
            Name:     "Jane Smith",
            Mail:    "jane.smith@example.com",
            Password: "password2",
            Role:     "user",
        },
    }, nil
}

func (u *UserUsecase) UpdateUser(user *dto.UserDTO) error {
	return nil
}

func (u *UserUsecase) DeleteUser(userId string) error {
	return nil
}