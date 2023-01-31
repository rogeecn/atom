package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_165134CreateSysDictionaryMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_165134CreateSysDictionary struct {
	id string
}

func New20230131_165134CreateSysDictionaryMigration() contracts.Migration {
	return &Migration20230131_165134CreateSysDictionary{id: "20230131_165134_create_sys_dictionary"}
}

func (m *Migration20230131_165134CreateSysDictionary) ID() string {
	return m.id
}

func (m *Migration20230131_165134CreateSysDictionary) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_165134CreateSysDictionary) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_165134CreateSysDictionary) table() interface{} {
	type SysDictionary struct {
		gorm.Model

		Name   string `gorm:"column:name;comment:字典名（中）"` // 字典名（中）
		Type   string `gorm:"column:type;comment:字典名（英）"` // 字典名（英）
		Status *bool  `gorm:"column:status;comment:状态"`   // 状态
		Desc   string `gorm:"column:desc;comment:描述"`     // 描述
	}

	return SysDictionary{}
}
