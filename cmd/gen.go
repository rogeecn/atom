/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"atom/cmd/model"
	"atom/container"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// MigrationInfo http service container
type GenQueryGenerator struct {
	dig.In

	DB *gorm.DB
}

// // Dynamic SQL
// type Querier interface {
// 	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
// 	FilterWithNameAndRole(name, role string) ([]gen.T, error)
// }

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "gorm query generator",
	Long:  `gorm query generator. more info, see https://gorm.io/gen`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(func(gq GenQueryGenerator) error {
			g := gen.NewGenerator(gen.Config{
				OutPath: "providers/query",
				Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
			})

			g.UseDB(gq.DB) // reuse your gorm db

			// Generate basic type-safe DAO API for struct `model.User` following conventions
			g.ApplyBasic(model.User{})

			// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
			// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

			// Generate the code
			g.Execute()
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
