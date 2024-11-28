package db

import (
	"database/sql"
	"fmt"
)

func TruncateTable(db *sql.DB, table string) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table))
}
