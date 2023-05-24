package models

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshService interface {
	FetchByID(id string) (User, error)
	GenerateAccessToken(user *User, key string, ex int) (token string, err error)
	GenerateRefreshToken(user *User, key string, ex int) (token string, err error)
	ExtractIDFromToken(inputToken string, key string) (string, error)
}
