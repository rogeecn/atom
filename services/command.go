package services

import (
	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"
)

type CommandService interface {
	Register()
	Execute() error
}

type Command struct {
	dig.In

	Commands []CommandService `group:"command_services"`
}

func ServeCommand(f func() error) error {
	return container.Container.Invoke(func(command Command) error {
		for _, cmd := range command.Commands {
			cmd.Register()
		}

		return f()
	})
}
