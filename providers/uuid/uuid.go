package uuid

import (
	"atom/container"
	"log"

	"github.com/gofrs/uuid"
)

func init() {
	if err := container.Container.Provide(NewUUID); err != nil {
		log.Fatal(err)
	}
}

type Generator struct {
	generator uuid.Generator
}

func NewUUID() (*Generator, error) {
	return &Generator{
		generator: uuid.DefaultGenerator,
	}, nil
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
