package handlers

import (
	"net/http"

	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	LoginService models.LoginService
	Env          *configs.Env
}

func (lh *LoginHandler) Login(c echo.Context) error {
	validate := validator.New()
	var payload *models.LoginInput
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	user, err := lh.LoginService.FetchByEmail(payload.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{Message: "User not found with the given email"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{Message: "Invalid credentials"})
	}

	accessToken, err := lh.LoginService.GenerateAccessToken(&user, lh.Env.JwtSecret, lh.Env.JwtExpiresIn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}
	refreshToken, err := lh.LoginService.GenerateRefreshToken(&user, lh.Env.JwtRefreshSecret, lh.Env.JwtRefreshExpiresIn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
