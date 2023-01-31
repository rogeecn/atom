package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_165218CreateSysDictionaryDetailMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_165218CreateSysDictionaryDetail struct {
	id string
}

func New20230131_165218CreateSysDictionaryDetailMigration() contracts.Migration {
	return &Migration20230131_165218CreateSysDictionaryDetail{id: "20230131_165218_create_sys_dictionary_detail"}
}

func (m *Migration20230131_165218CreateSysDictionaryDetail) ID() string {
	return m.id
}

func (m *Migration20230131_165218CreateSysDictionaryDetail) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_165218CreateSysDictionaryDetail) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_165218CreateSysDictionaryDetail) table() interface{} {
	type SysDictionaryDetail struct {
		gorm.Model

		Label           string `gorm:"column:label;comment:展示值"`              // 展示值
		Value           int    `gorm:"column:value;comment:字典值"`              // 字典值
		Status          *bool  `gorm:"column:status;comment:启用状态"`            // 启用状态
		Sort            int    `gorm:"column:sort;comment:排序标记"`              // 排序标记
		SysDictionaryID int    `gorm:"column:sys_dictionary_id;comment:关联标记"` // 关联标记
	}

	return SysDictionaryDetail{}
}
