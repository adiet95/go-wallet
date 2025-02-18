package database

import (
	"log"

	"go-wallet/src/models/entity"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var migUp bool
var migDown bool

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "start migration",
	RunE:  dbMigrate,
}

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration rollback")

}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&entity.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&entity.User{})
			},
		},
	})
	if migUp {
		if err := m.Migrate(); err != nil {
			return err
		}
		log.Println("Migration up done")
		return nil
	}
	if migDown {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		log.Println("Migration rollback done")
		return nil
	}
	log.Println("init schema database done")
	return nil
}
