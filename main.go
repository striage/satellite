package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"satellite/engine"
)

func main() {
	ebiten.SetWindowSize(engine.StageWidth, engine.StageHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(engine.NewGame()); err != nil {
		log.Fatal(err)
	}
}
