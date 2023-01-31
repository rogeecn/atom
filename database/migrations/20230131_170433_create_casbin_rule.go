package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	adapter "github.com/casbin/gorm-adapter/v3"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_170433CreateCasbinRuleMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_170433CreateCasbinRule struct {
	id string
}

func New20230131_170433CreateCasbinRuleMigration() contracts.Migration {
	return &Migration20230131_170433CreateCasbinRule{id: "20230131_170433_create_casbin_rule"}
}

func (m *Migration20230131_170433CreateCasbinRule) ID() string {
	return m.id
}

func (m *Migration20230131_170433CreateCasbinRule) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_170433CreateCasbinRule) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_170433CreateCasbinRule) table() interface{} {
	return adapter.CasbinRule{}
}
