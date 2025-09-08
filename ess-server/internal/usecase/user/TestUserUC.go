package user

import (
	abstract "ess-server/internal/core/domain/service/usecase"
	dto "ess-server/internal/core/dto/user"
)

type TestUserUC struct {}

func NewTestUserUC() *TestUserUC {
	return &TestUserUC{}
}

var _ abstract.IUserUsecase = (*TestUserUC)(nil)

func (u *TestUserUC) CreateUser(user *dto.UserDTO) error {
	return nil
}

func (u *TestUserUC) GetUser(userId string) ([]*dto.UserDTO, error) {
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

func (u *TestUserUC) UpdateUser(user *dto.UserDTO) error {
	return nil
}

func (u *TestUserUC) DeleteUser(userId string) error {
	return nil
}