package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Sprite struct {
	Name    string
	Img     *ebiten.Image
	Options *ebiten.DrawImageOptions
}

func NewSpriteFromImage(name string, img *ebiten.Image, x, y int) *Sprite {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	return &Sprite{
		Name:    name,
		Img:     img,
		Options: &op,
	}
}
