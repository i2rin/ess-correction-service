package auth


// LoginInDTO represents the data transfer object for login input.
type LoginInDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginInDTO(username, password string) *LoginInDTO {
	return &LoginInDTO{
		Username: username,
		Password: password,
	}
}

func (l *LoginInDTO) GetUsername() string {
	return l.Username
}

func (l *LoginInDTO) GetPassword() string {
	return l.Password
}


// LoginOutDTO represents the data transfer object for login output.
type LoginOutDTO struct {
	Role string `json:"role"`
}

func NewLoginOutDTO(role string) *LoginOutDTO {
	return &LoginOutDTO{
		Role: role,
	}
}

func (l *LoginOutDTO) GetRole() string {
	return l.Role
}
