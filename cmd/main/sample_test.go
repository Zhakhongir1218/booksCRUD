package main

import (
	"booksCRUD/internal/user"
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

var u user.User = user.User{
	Name:     "john",
	Password: "123",
	Status:   1,
}

func TestSignUp(t *testing.T) {

}

func (suite *ExampleTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}
