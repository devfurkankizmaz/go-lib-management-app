package models

type Profile struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type ProfileService interface {
	FetchProfileByID(id string) (*Profile, error)
}
