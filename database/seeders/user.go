package seeders

import (
	"atom/container"
	"atom/contracts"
	"fmt"
	"log"
	"time"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

func init() {
	if err := container.Container.Provide(NewUserSeeder, dig.Group("seeders")); err != nil {
		log.Fatal(err)
	}
}

type UserSeeder struct {
}

func NewUserSeeder() contracts.Seeder {
	return &UserSeeder{}
}

type User struct {
	gorm.Model

	Username string
	Password string
	Birthday time.Time
	Status   bool
}

func (s *UserSeeder) Run(db *gorm.DB) {
	times := 10
	for i := 0; i < times; i++ {
		data := s.Generate(i)
		if i == 0 {
			stmt := &gorm.Statement{DB: db}
			_ = stmt.Parse(&data)
			log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
		}
		db.Create(&data)
	}
}

func (s *UserSeeder) Generate(idx int) User {
	return User{
		Username: fmt.Sprintf("username: %d", idx),
		Password: fmt.Sprintf("password: %d", idx),
		Birthday: time.Now().Add(time.Hour * 24 * time.Duration(idx)),
		Status:   idx%2 == 0,
	}
}
