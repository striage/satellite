package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"satellite/engine"
	"satellite/engine/consts"
)

func main() {
	ebiten.SetWindowSize(consts.StageWidth, consts.StageHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(engine.NewGame()); err != nil {
		log.Fatal(err)
	}
}
