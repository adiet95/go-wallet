package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"

	"go-wallet/src/libs"
)

func CheckAuthor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerToken := c.Request().Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			return libs.New("invalid header type", 401, true).Send(c)
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkToken, err := libs.CheckToken(token)
		if err != nil {
			return libs.New(err.Error(), 401, true).Send(c)
		}
		if checkToken.Role != "admin" {
			return libs.New("forbidden access", 401, true).Send(c)
		}
		c.Set("role", checkToken.Role)
		if err = next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
