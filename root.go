package atom

import (
	"log"

	"github.com/pkg/errors"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/providers/config"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

var cfgFile string

var (
	GroupRoutes     = dig.Group("routes")
	GroupGrpcServer = dig.Group("grpc_server_services")
)

func Serve(providers container.Providers, opts ...Option) error {
	var rootCmd = &cobra.Command{}
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.toml", "config file path")

	for _, opt := range opts {
		opt(rootCmd)
	}
	if err := LoadProviders(cfgFile, providers); err != nil {
		return err
	}

	withMigrationCommand(rootCmd)
	withModelCommand(rootCmd)
	withSeederCommand(rootCmd)

	return rootCmd.Execute()
}

func LoadProviders(cfgFile string, providers container.Providers) error {
	// parse config files
	configure, err := config.Load(cfgFile)
	if err != nil {
		return errors.Wrapf(err, "load config file: %s", cfgFile)
	}

	if err := providers.Provide(configure); err != nil {
		return err
	}
	return nil
}

type Option func(*cobra.Command)

func Name(name string) Option {
	return func(cmd *cobra.Command) {
		cmd.Use = name
	}
}

func Short(short string) Option {
	return func(cmd *cobra.Command) {
		cmd.Short = short
	}
}

func Long(long string) Option {
	return func(cmd *cobra.Command) {
		cmd.Long = long
	}
}
func Run(run func(cmd *cobra.Command, args []string)) Option {
	return func(cmd *cobra.Command) {
		cmd.Run = run
	}
}

func RunE(run func(cmd *cobra.Command, args []string) error) Option {
	return func(cmd *cobra.Command) {
		cmd.RunE = run
	}
}

func PostRun(run func(cmd *cobra.Command, args []string)) Option {
	return func(cmd *cobra.Command) {
		cmd.PostRun = run
	}
}

func PostRunE(run func(cmd *cobra.Command, args []string) error) Option {
	return func(cmd *cobra.Command) {
		cmd.PostRunE = run
	}
}

func PreRun(run func(cmd *cobra.Command, args []string)) Option {
	return func(cmd *cobra.Command) {
		cmd.PreRun = run
	}
}

func PreRunE(run func(cmd *cobra.Command, args []string) error) Option {
	return func(cmd *cobra.Command) {
		cmd.PreRunE = run
	}
}

func Config(file string) Option {
	return func(cmd *cobra.Command) {
		_ = cmd.PersistentFlags().Set("config", file)
	}
}

func Seeders(seeders ...contracts.SeederProvider) Option {
	return func(cmd *cobra.Command) {
		for _, seeder := range seeders {
			if err := container.Container.Provide(seeder, dig.Group("seeder")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func Migrations(migrations ...contracts.MigrationProvider) Option {
	return func(cmd *cobra.Command) {
		for _, migration := range migrations {
			if err := container.Container.Provide(migration, dig.Group("migrations")); err != nil {
				log.Fatal(err)
			}
		}
	}
}
