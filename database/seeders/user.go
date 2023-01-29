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

func (s *UserSeeder) Times() uint {
	return 10
}

func (s *UserSeeder) Run(db *gorm.DB) {
	var i uint
	for i = 0; i < s.Times(); i++ {
		data := s.Generate(i)
		db.Table(s.Table()).Create(&data)
	}
}

func (s *UserSeeder) Table() string {
	return "users"
}

func (s *UserSeeder) Generate(idx uint) User {
	return User{
		Username: fmt.Sprintf("username: %d", idx),
		Password: fmt.Sprintf("password: %d", idx),
		Birthday: time.Now().Add(time.Hour * 24 * time.Duration(idx)),
		Status:   idx%2 == 0,
	}
}
