package atom

import (
	"log"

	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type SeedersContainer struct {
	dig.In

	DB      *gorm.DB
	Faker   *gofakeit.Faker
	Seeders []contracts.Seeder `group:"seeders"`
}

func withSeederCommand(rootCmd *cobra.Command) *cobra.Command {
	// seedCmd represents the seed command
	seedCmd := &cobra.Command{
		Use:   "seed",
		Short: "seed databases",
		Long:  `seed your database with data using seeders.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return container.Container.Invoke(func(c SeedersContainer) error {
				if len(c.Seeders) == 0 {
					log.Print("no seeder exists")
					return nil
				}

				for _, seeder := range c.Seeders {
					seeder.Run(c.Faker, c.DB)
				}
				return nil
			})
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			log.Println("BINGO! seeding done")
		},
	}
	rootCmd.AddCommand(seedCmd)
	return rootCmd
}
