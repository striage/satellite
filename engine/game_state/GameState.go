package game_state

import "github.com/hajimehoshi/ebiten/v2"

type GameState interface {
	Update() error
	Draw() (*ebiten.Image, *ebiten.DrawImageOptions)
}
