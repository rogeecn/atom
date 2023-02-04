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

type RoleInjectParams struct {
	dig.In
	DB    *gorm.DB
	Dao   RoleDao
	Query *query.Query
	Faker *gofakeit.Faker
}

type RoleSuite struct {
	suite.Suite
	RoleInjectParams
}

func init() {
	if err := container.Container.Provide(NewRoleDao); err != nil {
		log.Fatal(err)
	}
}

func Test_RoleSuite(t *testing.T) {
	err := container.Container.Invoke(func(p RoleInjectParams) {
		s := &RoleSuite{}
		s.RoleInjectParams = p

		suite.Run(t, s)
	})
	assert.NoError(t, err)
}

func (s *RoleSuite) BeforeTest(suiteName, testName string) {
	log.Println("BeforeTest: ", testName)
	utils.TruncateTable(s.DB, s.Query.SysRole.TableName())
	switch testName {
	case "Test_FindByID", "Test_UpdateByID", "Test_DeleteByID", "Test_DeletePermanentlyByID":
		log.Println("BeforeTest: insert test data")
		_, _ = s.Dao.Create(context.Background(), &models.SysRole{
			UUID:          s.Faker.UUID(),
			Name:          s.Faker.Name(),
			ParentID:      s.Faker.Uint64(),
			DefaultRouter: s.Faker.Animal(),
		})
	}
}

func (s *RoleSuite) AfterTest(suiteName, testName string) {

}

///////////////////
// start testing cases
//////////////////

func (s *RoleSuite) Test_Create() {
	Convey("Test_Create", s.T(), func() {
		Reset(func() {
			s.BeforeTest("_", "Test_Create")
		})

		Convey("create", func() {
			model, err := s.Dao.Create(context.Background(), &models.SysRole{
				UUID:          s.Faker.UUID(),
				Name:          s.Faker.Name(),
				ParentID:      s.Faker.Uint64(),
				DefaultRouter: s.Faker.Animal(),
			})
			So(err, ShouldBeNil)
			So(model.ID, ShouldEqual, 1)
		})
	})
}

func (s *RoleSuite) Test_FindByID() {
	Convey("Test_FindByID", s.T(), func() {
		model, err := s.Dao.FindByID(context.Background(), 1)
		So(err, ShouldBeNil)
		So(model.ID, ShouldEqual, 1)
	})
}

func (s *RoleSuite) Test_UpdateByID() {
	Convey("Test_UpdateByID", s.T(), func() {
		model, err := s.Dao.FindByID(context.Background(), 1)
		So(err, ShouldBeNil)
		So(model.ID, ShouldEqual, 1)

		name := "TEST_UpdateByID"
		model.Name = name
		newModel, err := s.Dao.UpdateByID(context.Background(), model)
		So(err, ShouldBeNil)
		So(newModel.Name, ShouldEqual, name)
	})
}

func (s *RoleSuite) Test_DeleteByID() {
	Convey("Test_DeleteByID", s.T(), func() {
		model, err := s.Dao.FindByID(context.Background(), 1)
		So(err, ShouldBeNil)
		So(model.ID, ShouldEqual, 1)

		err = s.Dao.DeleteByID(context.Background(), model.ID)
		So(err, ShouldBeNil)

		model, err = s.Query.SysRole.
			WithContext(context.TODO()).
			Unscoped().
			Where(s.Query.SysRole.ID.Eq(1)).
			First()
		So(err, ShouldBeNil)
		So(model.DeletedAt, ShouldNotBeNil)
	})
}

func (s *RoleSuite) Test_DeletePermanentlyByID() {
	Convey("Test_DeletePermanentlyByID", s.T(), func() {
		model, err := s.Dao.FindByID(context.Background(), 1)
		So(err, ShouldBeNil)
		So(model.ID, ShouldEqual, 1)

		err = s.Dao.DeletePermanentlyByID(context.Background(), model.ID)
		So(err, ShouldBeNil)

		_, err = s.Query.SysRole.
			WithContext(context.TODO()).
			Unscoped().
			Where(s.Query.SysRole.ID.Eq(1)).
			First()
		So(err, ShouldNotBeNil)
	})
}
