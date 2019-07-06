// Basic imports
package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupSuite() {

	fmt.Println("SetupSuite")
}

func (suite *ExampleTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite")
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	fmt.Println("TestExample")
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestExample1() {
	fmt.Println("TestExample1")
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
