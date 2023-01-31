package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"

	"github.com/gofrs/uuid"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230131_160904CreateUserMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230131_160904CreateUser struct {
	id string
}

func New20230131_160904CreateUserMigration() contracts.Migration {
	return &Migration20230131_160904CreateUser{id: "20230131_160904_create_user"}
}

func (m *Migration20230131_160904CreateUser) ID() string {
	return m.id
}

func (m *Migration20230131_160904CreateUser) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230131_160904CreateUser) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230131_160904CreateUser) table() interface{} {
	type User struct {
		gorm.Model

		UUID     uuid.UUID `gorm:"index;comment:UUID"`
		Username string    `gorm:"index;comment:登录名"`
		Password string    `gorm:"comment:登录密码"`
		Nickname string    `gorm:"default:'';comment:昵称"`
		Avatar   string    `gorm:"default:'';comment:头像"`
		RoleID   uint      `gorm:"default:1;comment:角色ID"`
		Phone    string    `gorm:"comment:手机号"`
		Email    string    `gorm:"comment:邮箱"`
		Status   string    `gorm:"default:ok;comment:用户状态"` // OK,BLOCKED,DISABLED,
	}

	return User{}
}
