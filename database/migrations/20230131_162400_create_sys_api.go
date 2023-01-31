package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_162400CreateSysApiMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_162400CreateSysApi struct {
	id string
}

func New20230131_162400CreateSysApiMigration() contracts.Migration {
	return &Migration20230131_162400CreateSysApi{id: "20230131_162400_create_sys_api"}
}

func (m *Migration20230131_162400CreateSysApi) ID() string {
	return m.id
}

func (m *Migration20230131_162400CreateSysApi) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_162400CreateSysApi) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_162400CreateSysApi) table() interface{} {
	type SysApi struct {
		gorm.Model

		Path        string `gorm:"comment:路径"`
		Description string `gorm:"comment:中文描述"`
		ApiGroup    string `gorm:"comment:组"`
		Method      string `gorm:"default:GET;comment:方法"`
	}

	return SysApi{}
}
