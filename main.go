package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"satellite/engine"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&engine.Game{}); err != nil {
		log.Fatal(err)
	}
}
