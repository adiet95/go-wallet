package routers

import (
	"errors"

	"github.com/labstack/echo/v4"

	"go-wallet/src/database"
	auth "go-wallet/src/modules/auth"
	"go-wallet/src/modules/users"
)

func New(mainRoute *echo.Echo) (*echo.Echo, error) {
	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	auth.New(mainRoute, db)
	users.New(mainRoute, db)

	return mainRoute, nil
}
