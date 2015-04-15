package BowlingGame

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func rollMany(g *Game, n, pins int) {
	for i := 0; i < n; i ++ {
		g.Roll(pins)
	}
}

func rollSpare(g *Game) {
	g.Roll(5)
	g.Roll(5)
}

func rollStrike(g *Game) {
	g.Roll(10)
}

func TestGutterGame(t *testing.T) {
	g := Game{}
	rollMany(&g, 20, 0)
	assert.Equal(t, 0, g.Score())
}

func TestAllOnes(t *testing.T) {
	g := Game{}
	rollMany(&g, 20, 1)
	assert.Equal(t, 20, g.Score())
}

func TestOneSpare(t *testing.T) {
	g := Game{}
	rollSpare(&g)
	g.Roll(3)
	rollMany(&g, 17, 0)
	assert.Equal(t, 16, g.Score())
}

func TestOneStrike(t *testing.T) {
	g := Game{}
	rollStrike(&g)
	g.Roll(3)
	g.Roll(4)
	rollMany(&g, 16, 0)
	assert.Equal(t, 24, g.Score())
}

func TestPerfectGame(t *testing.T) {
	g := Game{}
	rollMany(&g, 12, 10)
	assert.Equal(t, 300, g.Score())
}