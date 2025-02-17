package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	validate := validator.New()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc, validate)

	route := rt.Group("")
	{
		route.POST("/login", ctrl.SignIn)
		route.POST("/register", ctrl.Register)
	}
}
