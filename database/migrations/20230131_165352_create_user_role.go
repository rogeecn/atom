package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_165352CreateUserRoleMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_165352CreateUserRole struct {
	id string
}

func New20230131_165352CreateUserRoleMigration() contracts.Migration {
	return &Migration20230131_165352CreateUserRole{id: "20230131_165352_create_user_role"}
}

func (m *Migration20230131_165352CreateUserRole) ID() string {
	return m.id
}

func (m *Migration20230131_165352CreateUserRole) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_165352CreateUserRole) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_165352CreateUserRole) table() interface{} {
	type UserRole struct {
		UserID uint `gorm:"column:user_id"`
		RoleID uint `gorm:"column:role_id"`
	}

	return UserRole{}
}
