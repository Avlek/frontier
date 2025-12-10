package app

type GameState struct {
	CurrentDay int
}

func NewGameState() *GameState {
	return &GameState{
		CurrentDay: 1,
	}
}

func (g *GameState) EndDay() {
	g.CurrentDay++
}
