package BowlingGame

type Game struct {
	rolls [21]int
	currentRoll int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll ++
}

func (g *Game) Score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame ++ {
		if g.isStrike(frameIndex) {
			// Strike
			score += 10 + g.strikeBonus(frameIndex)
			frameIndex++
		} else if g.isSpare(frameIndex) {
			score += 10 + g.spareBonus(frameIndex)
			frameIndex += 2
		} else {
			score += g.sumOfBallsInFrame(frameIndex)
			frameIndex += 2
		}
	}
	return score
}

func (g *Game) sumOfBallsInFrame(frameIndex int) int {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1]
}

func (g *Game) spareBonus(frameIndex int) int {
	return g.rolls[frameIndex+2]
}

func (g *Game) strikeBonus(frameIndex int) int {
	return g.rolls[frameIndex+1] + g.rolls[frameIndex+2]
}

func (g *Game) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex] + g.rolls[frameIndex+1] == 10
}

func (g *Game) isStrike(frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}