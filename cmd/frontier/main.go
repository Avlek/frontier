package main

import (
	"github.com/avlek/frontier/internal/app"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := app.NewGame()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Frontier")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
