package handlers

import (
	"net/http"
	"strings"

	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterHandler struct {
	RegisterService models.RegisterService
	Env             *configs.Env
}

func (rh *RegisterHandler) Register(c echo.Context) error {
	validate := validator.New()
	var payload *models.RegisterInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	newUser := models.User{
		FullName: payload.FullName,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
	}

	err = rh.RegisterService.Create(&newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, newUser)
}
