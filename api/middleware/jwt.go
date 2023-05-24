package middleware

import (
	"net/http"
	"strings"

	"github.com/devfurkankizmaz/go-lib-management-app/configs"
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/devfurkankizmaz/go-lib-management-app/utils"
	"github.com/labstack/echo/v4"
)

func MiddlewareJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 👇 check for auth token
		authorization := c.Request().Header.Get("Authorization")
		t := strings.Split(authorization, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := utils.IsAuthorized(authToken, configs.App().Env.JwtSecret)
			if authorized {
				userID, err := utils.ExtractIDFromToken(authToken, configs.App().Env.JwtSecret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, models.Response{Message: err.Error()})
					c.Error(echo.ErrUnauthorized)
					return nil
				}
				c.Set("x-user-id", userID)
				next(c)
				return nil
			}
			c.JSON(http.StatusUnauthorized, models.Response{Message: err.Error()})
			c.Error(echo.ErrUnauthorized)
			return nil
		}
		c.JSON(http.StatusUnauthorized, models.Response{Message: "Not Authorized"})
		c.Error(echo.ErrUnauthorized)
		return nil
	}
}
