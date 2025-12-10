package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var bgColor = color.RGBA{80, 60, 30, 255}

type Game struct {
	State *GameState
}

func NewGame() *Game {
	return &Game{
		State: NewGameState(),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
