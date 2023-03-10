/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"atom/container"
	"atom/providers/config"
	"atom/providers/database"
	"errors"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// MigrationInfo http service container
type GenQueryGenerator struct {
	dig.In

	Config *config.Config
	DB     *gorm.DB
}

// // Dynamic SQL
// type Querier interface {
// 	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
// 	FilterWithNameAndRole(name, role string) ([]gen.T, error)
// }

// modelCmd represents the gen command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "gorm model&query generator",
	Long:  `gorm model&query generator. more info, see https://gorm.io/gen`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(func(gq GenQueryGenerator) error {
			var tables []string

			switch gq.Config.Database.Driver {
			case database.DriverMySQL:
				err := gq.DB.Raw("show tables").Scan(&tables).Error
				if err != nil {
					log.Fatal(err)
				}
			case database.DriverPostgres:
				err := gq.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables).Error
				if err != nil {
					log.Fatal(err)
				}
			case database.DriverSQLite:
				err := gq.DB.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tables).Error
				if err != nil {
					log.Fatal(err)
				}
			}

			if len(tables) == 0 {
				return errors.New("no tables in database, run migrate first")
			}

			g := gen.NewGenerator(gen.Config{
				OutPath:          "database/query",
				OutFile:          "query.gen.go",
				ModelPkgPath:     "database/models",
				FieldSignable:    true,
				FieldWithTypeTag: true,
				Mode:             gen.WithDefaultQuery | gen.WithQueryInterface,
			})

			g.UseDB(gq.DB) // reuse your gorm db

			models := []interface{}{}
			for _, table := range tables {
				models = append(models, g.GenerateModel(table))
			}

			// Generate basic type-safe DAO API for struct `model.User` following conventions
			g.ApplyBasic(models...)

			// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
			// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

			// Generate the code
			g.Execute()
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
