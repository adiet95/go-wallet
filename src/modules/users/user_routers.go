package users

import (
	"go-wallet/src/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/user")
	route.Use(middleware.CheckAuth)
	{
		route.GET("/:id", ctrl.SearchId)
		route.PUT("", ctrl.UpdateProfile)
		route.GET("/search", ctrl.SearchName, middleware.CheckAuthor)
	}
}
