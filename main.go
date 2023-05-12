package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	entities []Entity
}

func (g *Game) Update() error {
	for _, e := range g.entities {
		e.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Mouse horizontal position: wind direction")
	ebitenutil.DebugPrint(screen, "Mouse button 1: rain velocity")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("go-rain")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
