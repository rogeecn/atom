package atom

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/rogeecn/atom-addons/providers/config"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

var cfgFile string

var (
	GroupInitialName           = "initials"
	GroupRoutesName            = "routes"
	GroupGrpcServerServiceName = "grpc_server_services"
	GroupCommandName           = "command_services"
	GroupQueueName             = "queue_handler"

	GroupInitial    = dig.Group(GroupInitialName)
	GroupRoutes     = dig.Group(GroupRoutesName)
	GroupGrpcServer = dig.Group(GroupGrpcServerServiceName)
	GroupCommand    = dig.Group(GroupCommandName)
	GroupQueue      = dig.Group(GroupQueueName)
)

func Serve(providers container.Providers, opts ...Option) error {
	rootCmd := &cobra.Command{Use: "app"}
	for _, opt := range opts {
		opt(rootCmd)
	}

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Println(err)
		cmd.Println(cmd.UsageString())
		return err
	})

	defaultCfgFile := fmt.Sprintf(".%s.toml", rootCmd.Use)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path, lookup in dir: $HOME, $PWD, /etc, /usr/local/etc, filename: "+defaultCfgFile)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return LoadProviders(cfgFile, rootCmd.Use, providers)
	}

	return rootCmd.Execute()
}

func LoadProviders(cfgFile, appName string, providers container.Providers) error {
	// parse config files
	configure, err := config.Load(cfgFile, appName)
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

func CmdSeeders(seeders ...contracts.SeederProvider) Option {
	return func(cmd *cobra.Command) {
		withSeederCommand(cmd)

		for _, seeder := range seeders {
			if err := container.Container.Provide(seeder, dig.Group("seeders")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func CmdMigrations(migrations ...contracts.MigrationProvider) Option {
	return func(cmd *cobra.Command) {
		withMigrationCommand(cmd)

		for _, migration := range migrations {
			if err := container.Container.Provide(migration, dig.Group("migrations")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func CmdModel() Option {
	return func(cmd *cobra.Command) {
		withModelCommand(cmd)
	}
}

func CmdService() Option {
	return func(cmd *cobra.Command) {
		withServiceInstall(cmd)
	}
}

func CmdQueue() Option {
	return func(cmd *cobra.Command) {
		withQueueCommand(cmd)
	}
}
