package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var sprite *Sprite

func init() {
	sprite = newSprite("rika", "assets/rika.png", 120, 120)
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	sprite.render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
