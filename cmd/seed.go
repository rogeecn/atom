/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "atom/database/seeders"
	_ "atom/providers"

	"atom/container"
	"atom/contracts"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed databases",
	Long:  `seed your database with data using seeders.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(func(seedersContainer SeedersContainer) error {
			if len(seedersContainer.Seeders) == 0 {
				log.Print("no seeder exists")
				return nil
			}

			for _, seeder := range seedersContainer.Seeders {
				log.Printf("seed table (%s) for %d times", seeder.Table(), seeder.Times())
				seeder.Run(seedersContainer.DB)
			}
			return nil
		})
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("BINGO! seeding done")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

type SeedersContainer struct {
	dig.In

	DB      *gorm.DB
	Seeders []contracts.Seeder `group:"seeders"`
}
