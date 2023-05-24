package models

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type LoginService interface {
	FetchByEmail(email string) (User, error)
	GenerateAccessToken(user *User, key string, ex int) (token string, err error)
	GenerateRefreshToken(user *User, key string, ex int) (token string, err error)
}
