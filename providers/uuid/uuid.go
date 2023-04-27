package uuid

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/providers"

	"github.com/gofrs/uuid"
)

type Generator struct {
	generator uuid.Generator
}

func Provide(o *providers.Options) error {
	return container.Container.Provide(func() (*Generator, error) {
		return &Generator{
			generator: uuid.DefaultGenerator,
		}, nil
	}, o.DiOptions()...)
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
