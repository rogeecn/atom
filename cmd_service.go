package atom

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"

	"github.com/rogeecn/atom/container"

	"github.com/spf13/cobra"
)

var enableService bool

type ServiceVars struct {
	ExecStart   string
	Description string
	WantedBy    string
}

func withServiceInstall(rootCmd *cobra.Command) *cobra.Command {
	serviceTpl := `[Unit]
Description={{.Description}}

[Service]
Type=simple
ExecStart={{.ExecStart}}

[Install]
WantedBy={{.WantedBy}}`
	var serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "install linux service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return container.Container.Invoke(func() error {
				tpl, err := template.New("service").Parse(serviceTpl)
				if err != nil {
					return err
				}

				exe, err := os.Executable()
				if err != nil {
					return err
				}

				vars := ServiceVars{
					ExecStart:   exe,
					Description: rootCmd.Short,
					WantedBy:    "multi-user.target",
				}

				tplWriter := bytes.NewBuffer(nil)
				if err := tpl.Execute(tplWriter, vars); err != nil {
					return err
				}

				filename := fmt.Sprintf("/etc/systemd/system/%s.service", rootCmd.Use)
				if err := os.WriteFile(filename, tplWriter.Bytes(), os.ModePerm); err != nil {
					return err
				}

				if enableService {
					b, err := exec.Command("sudo", "systemctl", "enable", rootCmd.Use, "--now").CombinedOutput()
					if err != nil {
						return fmt.Errorf("%s: %s", err, string(b))
					}
				}

				return nil
			})
		},
	}

	rootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().BoolVar(&enableService, "enable", true, "enable service after install")

	return rootCmd
}
