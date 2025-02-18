package config

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"

	"go-wallet/src/database"
	"go-wallet/src/modules/payment"
	"go-wallet/src/modules/topup"
	"go-wallet/src/modules/transfer"
	"go-wallet/src/routers"

	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	db, err := database.New()
	if err != nil {
		return errors.New("failed init database")
	}
	rd := database.RedisClient()

	go payment.NewWorker(db, rd)
	go transfer.NewWorker(db, rd)
	go topup.NewWorker(db, rd)
	fmt.Println("Worker is running")

	e := echo.New()
	if mainRoute, err := routers.New(e, db, rd); err == nil {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedHeaders:   []string{"Content-Type", "Authorization"},
			AllowedMethods:   []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production
			Debug: false,
		})

		handlerCors := c.Handler(mainRoute)

		var addrs = ":8080"
		if port := os.Getenv("PORT"); port != "" {
			addrs = port
		}

		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      handlerCors,
		}
		fmt.Println("Server is running on port ", addrs)
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
		return nil
	} else {
		return err
	}
}
