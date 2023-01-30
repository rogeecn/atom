package migrations

import (
	"atom/container"
	"atom/contracts"
	"atom/providers/log"
	"time"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(New20230130_193508CreateSysOperationRecordMigration, dig.Group("migrations")); err != nil {
		log.Fatal(err)
	}
}

type Migration20230130_193508CreateSysOperationRecord struct {
	id string
}

func New20230130_193508CreateSysOperationRecordMigration() contracts.Migration {
	return &Migration20230130_193508CreateSysOperationRecord{id: "20230130_193508_create_sys_operation_record"}
}

func (m *Migration20230130_193508CreateSysOperationRecord) ID() string {
	return m.id
}

func (m *Migration20230130_193508CreateSysOperationRecord) Up(tx *gorm.DB) error {
	table := m.table()
	return tx.AutoMigrate(&table)
}

func (m *Migration20230130_193508CreateSysOperationRecord) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
}

func (m *Migration20230130_193508CreateSysOperationRecord) table() interface{} {
	type SysOperationRecord struct {
		gorm.Model
		Ip           string        `gorm:"comment:请求ip"`
		Method       string        `gorm:"comment:请求方法"`
		Path         string        `gorm:"comment:请求路径"`
		Status       int           `gorm:"comment:请求状态"`
		Latency      time.Duration `gorm:"comment:延迟"`
		Agent        string        `gorm:"comment:代理"`
		ErrorMessage string        `gorm:"comment:错误信息"`
		Body         string        `gorm:"comment:请求Body"`
		Resp         string        `gorm:"comment:响应Body"`
		UserID       int           `gorm:"comment:用户id"`
	}

	return SysOperationRecord{}
}
