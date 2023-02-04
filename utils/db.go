package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func TruncateTable(db *gorm.DB, table string) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
}
