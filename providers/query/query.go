package query

import (
	"atom/container"
	"atom/database/query"
	"log"

	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewQuery); err != nil {
		log.Fatal(err)
	}
}

func NewQuery(db *gorm.DB) *query.Query {
	return query.Use(db)
}
