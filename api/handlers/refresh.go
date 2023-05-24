package handlers

import (
	"net/http"

	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RefreshHandler struct {
	RefreshService models.RefreshService
	Env            *configs.Env
}

func (rh *RefreshHandler) Refresh(c echo.Context) error {
	validate := validator.New()
	var payload models.RefreshInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "Refresh Token Required"})
	}

	id, err := rh.RefreshService.ExtractIDFromToken(payload.RefreshToken, rh.Env.JwtRefreshSecret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "User not found!"})
	}
	user, err := rh.RefreshService.FetchByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "User not found!"})
	}
	at, err := rh.RefreshService.GenerateAccessToken(&user, rh.Env.JwtSecret, rh.Env.JwtExpiresIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	rt, err := rh.RefreshService.GenerateRefreshToken(&user, rh.Env.JwtRefreshSecret, rh.Env.JwtRefreshExpiresIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.RefreshResponse{AccessToken: at, RefreshToken: rt})
}
