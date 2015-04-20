package BowlingGame

type Game struct {
	rolls [21]int
	numRolls int
}

func (g *Game)Roll(pins int) {
	g.rolls[g.numRolls] = pins
	g.numRolls++
}

func (g *Game) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1] == 10
}

func (g *Game) spareBonus(frameIndex int) int {
	return g.rolls[frameIndex+2]
}

func (g *Game) isStrike(frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}

func (g *Game) strikeBonus(frameIndex int) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func (g *Game)Score() int {
	score := 0
	frameIndex := 0
	for i := 0; i < 10; i ++ {
		if g.isSpare(frameIndex) {
			score += g.rolls[frameIndex] + g.rolls[frameIndex+1] + g.spareBonus(frameIndex)
			frameIndex += 2
		} else if g.isStrike(frameIndex) {
			score += g.rolls[frameIndex] + g.strikeBonus(frameIndex)
			frameIndex += 1
		} else {
			score += g.rolls[frameIndex] + g.rolls[frameIndex+1]
			frameIndex += 2
		}
	}
	return score
}