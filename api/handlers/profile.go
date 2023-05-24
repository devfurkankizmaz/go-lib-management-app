package handlers

import (
	"fmt"
	"net/http"

	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	ProfileService models.ProfileService
}

func (ph *ProfileHandler) Fetch(c echo.Context) error {
	userID := c.Get("x-user-id")
	newUID := fmt.Sprint(userID)
	profile, err := ph.ProfileService.FetchProfileByID(newUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, profile)
}
