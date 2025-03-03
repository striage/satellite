package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Sprite struct {
	name    string
	img     *ebiten.Image
	options *ebiten.DrawImageOptions
}

func newSpriteFromImage(name string, img *ebiten.Image, x, y int) *Sprite {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	return &Sprite{
		name:    name,
		img:     img,
		options: &op,
	}
}
