package order

import (
	"go-wallet/src/middleware"
	"go-wallet/src/modules/users"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func New(rt *echo.Echo, db *gorm.DB, rd *redis.Client) {
	repo := NewRepo(db, rd)
	userRepo := users.NewRepo(db)
	svc := NewService(repo, userRepo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/pay")
	route.Use(middleware.CheckAuth)
	{
		route.POST("", ctrl.PostPayment)
	}
}
