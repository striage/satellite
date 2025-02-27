package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Sprite struct {
	name    string
	img     *ebiten.Image
	x       int
	y       int
	options *ebiten.DrawImageOptions
}

func newSprite(name string, path string, x, y int) *Sprite {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return &Sprite{
		name:    name,
		img:     img,
		x:       x,
		y:       y,
		options: &ebiten.DrawImageOptions{},
	}
}

func (sprite *Sprite) render(screen *ebiten.Image) {
	screen.DrawImage(sprite.img, sprite.options)
}
