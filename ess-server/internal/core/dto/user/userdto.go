package user

type UserInDTO struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Mail     string `json:"mailaddress"`
	Role     string `json:"role"`
}

func NewUserInDTO(password, name, mail string, role string) *UserInDTO {
	return &UserInDTO{
		Password: password,
		Name:     name,
		Mail:     mail,
		Role:     role,
	}
}

func (u *UserInDTO) GetPassword() string {
	return u.Password
}

func (u *UserInDTO) GetName() string {
	return u.Name
}

func (u *UserInDTO) GetMail() string {
	return u.Mail
}

func (u *UserInDTO) GetRole() string {
	return u.Role
}

type UserOutDTO struct {
	Nickname string `json:"nickname"`
	UserId   string `json:"userid"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Mail     string `json:"mailaddress"`
	Role     string `json:"role"`
}

func NewUserOutDTO(userId, password, name, mail string, role string) *UserOutDTO {
	return &UserOutDTO{
		UserId:   userId,
		Password: password,
		Name:     name,
		Mail:     mail,
		Role:     role,
	}
}

func (u *UserOutDTO) GetUserId() string {
	return u.UserId
}

func (u *UserOutDTO) GetPassword() string {
	return u.Password
}

func (u *UserOutDTO) GetName() string {
	return u.Name
}

func (u *UserOutDTO) GetMail() string {
	return u.Mail
}

func (u *UserOutDTO) GetRole() string {
	return u.Role
}
