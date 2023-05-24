package models

import "github.com/golang-jwt/jwt/v4"

type Response struct {
	Message interface{}
}

type JwtCustomClaims struct {
	FullName string `json:"full_name"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
