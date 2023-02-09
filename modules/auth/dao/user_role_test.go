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

type UserRoleInjectParams struct {
	dig.In

	DB    *gorm.DB
	Dao   *UserRoleDao
	Query *query.Query
	Faker *gofakeit.Faker
}

type UserRoleSuite struct {
	suite.Suite
	UserRoleInjectParams
}

func init() {
	if err := container.Container.Provide(NewUserRoleDao); err != nil {
		log.Fatal(err)
	}
}

func Test_UserRoleSuite(t *testing.T) {
	err := container.Container.Invoke(func(p UserRoleInjectParams) {
		s := &UserRoleSuite{}
		s.UserRoleInjectParams = p

		suite.Run(t, s)
	})
	assert.NoError(t, err)
}

func (s *UserRoleSuite) BeforeTest(suiteName, testName string) {
	utils.TruncateTable(s.DB, s.Query.UserRole.TableName())

}

func (s *UserRoleSuite) Test_GetByUserID() {
	Convey("Test_GetByUserID", s.T(), func() {
		Reset(func() {
			s.BeforeTest("_", "Test_GetByUserID")
		})

		Convey("not exists", func() {
			has := s.Dao.Exists(context.Background(), 1)
			So(has, ShouldBeFalse)
		})

		Convey("exists", func() {
			_ = s.Query.UserRole.WithContext(context.Background()).Create(&models.UserRole{
				UserID: 1,
				RoleID: 1,
			})

			has := s.Dao.Exists(context.Background(), 1)
			So(has, ShouldBeTrue)
		})
	})
}
