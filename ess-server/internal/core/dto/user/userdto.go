package user

type UserDTO struct {
	Password   string
	Name       string
	Mail      string
	Role      string
}

func NewUserDTO(password, name, mail, role string) *UserDTO {
	return &UserDTO{
		Password: password,
		Name:     name,
		Mail:    mail,
		Role:    role,
	}
}

func (u *UserDTO) GetPassword() string {
	return u.Password
}

func (u *UserDTO) GetName() string {
	return u.Name
}

func (u *UserDTO) GetMail() string {
	return u.Mail
}

func (u *UserDTO) GetRole() string {
	return u.Role
}
