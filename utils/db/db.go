package db

import (
	"database/sql"
	"fmt"
)

func TruncateTable(db *sql.DB, table string) {
	_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table))
	_ = err
}
