package uuid

import (
	"github.com/rogeecn/atom/container"
	"go.uber.org/dig"

	"github.com/gofrs/uuid"
)

type Generator struct {
	generator uuid.Generator
}

func Provide(opts ...dig.ProvideOption) error {
	return container.Container.Provide(func() (*Generator, error) {
		return &Generator{
			generator: uuid.DefaultGenerator,
		}, nil
	})
}

func (u *Generator) MustGenerate() string {
	uuid, _ := u.Generate()
	return uuid
}

func (u *Generator) Generate() (string, error) {
	uuid, err := u.generator.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), err
}
