package dao

import (
	"context"
	"log"
	"testing"

	// 这里的依赖需要被导入，否则会报错
	"atom/container"
	"atom/database/models"
	"atom/database/query"
	_ "atom/providers"
	"atom/utils"

	"github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type UserInjectParams struct {
	dig.In

	DB    *gorm.DB
	Dao   UserDao
	Query *query.Query
	Faker *gofakeit.Faker
}

type UserSuite struct {
	suite.Suite
	UserInjectParams
}

func init() {
	if err := container.Container.Provide(NewUserDao); err != nil {
		log.Fatal(err)
	}
}

func Test_UserSuite(t *testing.T) {
	err := container.Container.Invoke(func(p UserInjectParams) {
		s := &UserSuite{}
		s.UserInjectParams = p

		suite.Run(t, s)
	})
	assert.NoError(t, err)
}

func (s *UserSuite) BeforeTest(suiteName, testName string) {
	log.Println("BeforeTest: ", testName)

	utils.TruncateTable(s.DB, s.Query.User.TableName())

	switch testName {
	case "":
		log.Println("BeforeTest: insert test data")
		_, _ = s.Dao.Create(context.Background(), &models.User{
			UUID:     s.Faker.UUID(),
			Username: s.Faker.Username(),
			Password: s.Faker.Password(true, true, true, true, false, 16),
			Nickname: s.Faker.Name(),
			Avatar:   s.Faker.ImageURL(100, 100),
			RoleID:   0,
			Phone:    s.Faker.Phone(),
			Email:    s.Faker.Email(),
			Status:   s.Faker.RandomString([]string{"enable", "disabled"}),
		})
	}
}

func (s *UserSuite) Test_Create() {
	Convey("Test_Create", s.T(), func() {
		Convey("create", func() {
			model, err := s.Dao.Create(context.Background(), &models.User{
				UUID:     s.Faker.UUID(),
				Username: s.Faker.Username(),
				Password: s.Faker.Password(true, true, true, true, false, 16),
				Nickname: s.Faker.Name(),
				Avatar:   s.Faker.ImageURL(100, 100),
				RoleID:   0,
				Phone:    s.Faker.Phone(),
				Email:    s.Faker.Email(),
				Status:   s.Faker.RandomString([]string{"enable", "disabled"}),
			})
			So(err, ShouldBeNil)
			So(model.ID, ShouldEqual, 1)
		})
	})
}
