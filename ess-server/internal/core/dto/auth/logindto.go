package auth

// LoginInDTO represents the data transfer object for login input.
type LoginInDTO struct {
	UserID   string `json:"userId" binding:"required"`
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
	Admin bool `json:"admin"`
}

func NewLoginOutDTO(admin bool) *LoginOutDTO {
	return &LoginOutDTO{
		Admin: admin,
	}
}

func (l *LoginOutDTO) GetAdmin() bool {
	return l.Admin
}
