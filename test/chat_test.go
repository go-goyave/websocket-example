package test

import (
	"testing"

	"goyave.dev/goyave/v3"
)

type ChatTestSuite struct {
	goyave.TestSuite
}

func (suite *ChatTestSuite) TestHello() {
	// TODO chat tests
	suite.True(true)
}

func TestChatSuite(t *testing.T) { // Run the test suite
	goyave.RunTest(t, new(ChatTestSuite))
}
