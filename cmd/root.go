package cmd

import (
	_ "atom/providers"

	"atom/container"
	"atom/services/http"
	"atom/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "app",
	Short:   "app",
	Long:    `the app long description`,
	Version: fmt.Sprintf("\nVersion: %s\nGitHash: %s\nBuildAt: %s\n", utils.Version, utils.GitHash, utils.BuildAt),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("using config file: ", utils.ShareConfigFile)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return container.Container.Invoke(http.Serve)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&utils.ShareConfigFile, "config", "c", "config.toml", "config file")
}
