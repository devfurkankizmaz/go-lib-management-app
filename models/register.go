package models

type RegisterInput struct {
	FullName string `json:"full_name" validate:"required,min=4,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=25"`
}

type RegisterService interface {
	Create(user *User) error
	FetchByEmail(email string) (User, error)
}
