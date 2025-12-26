package main

import (
	"log"

	"github.com/NautiluX/gofirst/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// A typical mobile aspect ratio (9:16) at low-res
	virtualWidth  = 200
	virtualHeight = 320
)

func main() {
	ebiten.SetWindowTitle("GoFirst - Startspieler")

	// This makes the window look like a phone on your desktop
	ebiten.SetWindowSize(virtualWidth*2, virtualHeight*2)

	g := game.NewGame(virtualWidth, virtualHeight)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
