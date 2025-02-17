package database

import (
	"github.com/spf13/cobra"
)

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "start seeding",
	RunE:  seeder,
}

func seeder(cmd *cobra.Command, args []string) error {
	// db, err := New()
	// if err != nil {
	// 	return err
	// }
	// admin, _ := libs.HashPassword("admin12345678")
	// user, _ := libs.HashPassword("user12345678")

	// var datas = []models.User{
	// 	{Email: "admin@gmail.com", Password: admin, Role: "admin"},
	// 	{Email: "user@gmail.com", Password: user, Role: "user"},
	// }

	// if res := db.Create(&datas); res.Error != nil {
	// 	return res.Error
	// }
	// log.Println("Seeding successful")
	return nil
}
