package auth

// LoginInDTO represents the data transfer object for login input.
type LoginInDTO struct {
	UserID   string `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginInDTO(userID, password string) *LoginInDTO {
	return &LoginInDTO{
		UserID:   userID,
		Password: password,
	}
}

func (l *LoginInDTO) GetUserID() string {
	return l.UserID
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
