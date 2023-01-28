package system

import (
	"atom/container"
	"context"
	"fmt"
	"testing"

	_ "atom/modules/system/container"
	"atom/modules/system/service"
	_ "atom/providers/config" // 这里的依赖需要被导入，否则会报错
	_ "atom/providers/logger"
	_ "atom/providers/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/dig"
)

type InjectServiceParams struct {
	dig.In

	Service service.SystemService // 注意这里注入的参数需要大写
}

type TServiceSuite struct {
	suite.Suite

	InjectServiceParams
}

func Test_ServiceSuiteSuite(t *testing.T) {
	err := container.Container.Invoke(func(p InjectServiceParams) {
		s := &TServiceSuite{}
		s.InjectServiceParams = p

		suite.Run(t, s)
	})
	assert.NoError(t, err)
}

func (s *TServiceSuite) Test_GetName() {
	name, err := s.Service.GetName(context.Background())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "System.GetName", name.Name)
}

// ////
func (s *TServiceSuite) SetupSuite() {
	fmt.Println("SetupSuite")
}
func (s *TServiceSuite) SetupTest() {
	fmt.Println("SetupTest")
}
func (s *TServiceSuite) BeforeTest(suiteName, testName string) {
	fmt.Println("BeforeTest:", suiteName, testName)
}
func (s *TServiceSuite) AfterTest(suiteName, testName string) {
	fmt.Println("AfterTest:", suiteName, testName)
}
func (s *TServiceSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Println("HandleStats:", suiteName, stats)
}
func (s *TServiceSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}
func (s *TServiceSuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
}
