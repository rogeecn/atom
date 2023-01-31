package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_165509CreateSysRoleMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_165509CreateSysRole struct {
	id string
}

func New20230131_165509CreateSysRoleMigration() contracts.Migration {
	return &Migration20230131_165509CreateSysRole{id: "20230131_165509_create_sys_role"}
}

func (m *Migration20230131_165509CreateSysRole) ID() string {
	return m.id
}

func (m *Migration20230131_165509CreateSysRole) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_165509CreateSysRole) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_165509CreateSysRole) table() interface{} {
	type SysRole struct {
		gorm.Model

		Alias         string `gorm:"not null;unique;primary_key;comment:角色Alias;size:90"` // 角色ID
		Name          string `gorm:"comment:角色名"`                                         // 角色名
		ParentId      *uint  `gorm:"comment:父角色ID"`                                       // 父角色ID
		DefaultRouter string `gorm:"comment:默认菜单;default:dashboard"`                      // 默认菜单(默认dashboard)
	}

	return SysRole{}
}
