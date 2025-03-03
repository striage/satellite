package engine

import "github.com/hajimehoshi/ebiten/v2"

type GameState interface {
	Update() error
	Draw(screen *ebiten.Image)
}
