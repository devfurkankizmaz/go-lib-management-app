package utils

import (
	"fmt"
	"time"

	"github.com/devfurkankizmaz/go-lib-management-app/models"
	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(user *models.User, key string, ex int) (accessToken string, err error) {
	expTime := time.Now().Add(time.Hour * time.Duration(ex)).Unix()
	claims := &models.JwtCustomClaims{
		FullName: user.FullName,
		ID:       (user.ID).String(),
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return t, err
}

func GenerateRefreshToken(user *models.User, key string, ex int) (refreshToken string, err error) {
	expTime := time.Now().Add(time.Hour * time.Duration(ex)).Unix()
	claims := &models.JwtCustomRefreshClaims{
		ID: (user.ID).String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return rt, err
}

func IsAuthorized(inputToken string, key string) (bool, error) {
	_, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(inputToken string, key string) (string, error) {
	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}
