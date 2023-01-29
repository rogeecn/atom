package migrations

import (
	"atom/container"
	"atom/contracts"
	"log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20140202InitMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20140202Init struct {
	id string
}

func New20140202InitMigration() contracts.Migration {
	return &Migration20140202Init{id: "20140202_init"}
}

func (m *Migration20140202Init) ID() string {
	return m.id
}

func (m *Migration20140202Init) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20140202Init) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20140202Init) table() interface{} {
	type TableName struct {
		FieldName string
	}

	return TableName{}
}
