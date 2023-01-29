package cmd

import (
	"log"
	"sort"

	// init dependencies
	_ "atom/database/migrations"
	_ "atom/providers"

	"atom/container"
	"atom/contracts"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database tables",
	Long:  `migrate database tables`,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	migrateCmd.PersistentFlags().StringVar(&migrateToId, "to", "", "migration to id")
}

var migrateToId string

// MigrationInfo http service container
type MigrationInfo struct {
	dig.In

	DB         *gorm.DB
	Migrations []contracts.Migration `group:"migrations"`
}

// migrateUpCmd represents the migrateUp command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up database tables",
	Long:  `migrate up database tables`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(func(mi MigrationInfo) error {
			m := gormigrate.New(mi.DB, gormigrate.DefaultOptions, sortedMigrations(mi.Migrations))

			if len(migrateToId) > 0 {
				log.Printf("migrate up to [%s]\n", migrateToId)
				return m.MigrateTo(migrateToId)
			}
			return m.Migrate()
		})
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("BINGO! migrate up done")
	},
}

// migrateDownCmd represents the migrateDown command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate down database tables",
	Long:  `migrate down database tables`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(func(mi MigrationInfo) error {
			m := gormigrate.New(mi.DB, gormigrate.DefaultOptions, sortedMigrations(mi.Migrations))

			if len(migrateToId) > 0 {
				log.Printf("migrate down to [%s]\n", migrateToId)
				return m.RollbackTo(migrateToId)
			}
			return m.RollbackLast()
		})
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		log.Println("BINGO! migrate down done")
	},
}

func sortedMigrations(ms []contracts.Migration) []*gormigrate.Migration {
	migrationKeys := []string{}
	migrationMaps := make(map[string]*gormigrate.Migration)
	for _, m := range ms {
		migrationKeys = append(migrationKeys, m.ID())
		migrationMaps[m.ID()] = &gormigrate.Migration{
			ID:       m.ID(),
			Migrate:  m.Up,
			Rollback: m.Down,
		}
	}
	sort.Strings(migrationKeys)

	migrations := []*gormigrate.Migration{}
	for _, key := range migrationKeys {
		migrations = append(migrations, migrationMaps[key])
	}

	return migrations
}
