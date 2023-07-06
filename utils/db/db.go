package db

import (
	"fmt"

	"gorm.io/gorm"
)

func TruncateTable(db *gorm.DB, table string) {
	if db.Dialector.Name() == "postgres" {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table))
		return
	}
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
}
