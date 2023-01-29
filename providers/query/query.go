package query

import (
	"atom/container"
	"log"

	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewQuery); err != nil {
		log.Fatal(err)
	}
}

func NewQuery(db *gorm.DB) *Query {
	return Use(db)
}
