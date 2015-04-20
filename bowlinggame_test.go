package BowlingGame

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct {
	suite.Suite
	g Game
}

func (suite *GameTestSuite) SetupTest() {
	suite.g = Game{}
}

func (suite *GameTestSuite) rollMany(rolls, pins int) {
	for i := 0; i < rolls; i ++ {
		suite.g.Roll(pins)
	}
}

func (suite *GameTestSuite) rollSpare() {
	suite.g.Roll(5)
	suite.g.Roll(5)
}

func (suite *GameTestSuite) rollStrike() {
	suite.g.Roll(10)
}

func (suite *GameTestSuite) TestGutterGame() {
	suite.rollMany(20, 0)
	assert.Equal(suite.T(), 0, suite.g.Score())
}

func (suite *GameTestSuite) TestAllOnes() {
	suite.rollMany(20, 1)
	assert.Equal(suite.T(), 20, suite.g.Score())
}

func (suite *GameTestSuite) TestOneSpare() {
	suite.rollSpare()
	suite.g.Roll(3)
	suite.rollMany(17, 0)
	assert.Equal(suite.T(), 16, suite.g.Score())
}

func (suite *GameTestSuite) TestOneStrike() {
	suite.rollStrike()
	suite.g.Roll(3)
	suite.g.Roll(4)
	suite.rollMany(16, 0)
	assert.Equal(suite.T(), 24, suite.g.Score())
}

func (suite *GameTestSuite) TestPerfectGame() {
	suite.rollMany(12, 10)
	assert.Equal(suite.T(), 300, suite.g.Score())
}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}