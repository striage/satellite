package engine

import "github.com/hajimehoshi/ebiten/v2"

type Stage struct {
	background *Sprite
	sprites    []*Sprite
	op         *ebiten.DrawImageOptions
}

func newStage() *Stage {
	return &Stage{
		background: nil,
		sprites:    make([]*Sprite, 0),
		op:         &ebiten.DrawImageOptions{},
	}
}

func (s *Stage) appendSprite(sprite *Sprite) {
	s.sprites = append(s.sprites, sprite)
}

func (s *Stage) setBackground(background *Sprite) {
	s.background = background
}

func (s *Stage) makeImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	image := ebiten.NewImage(StageWidth, StageHeight)
	if s.background != nil {
		image.DrawImage(s.background.img, s.background.options)
	}
	for i := 0; i < len(s.sprites); i++ {
		image.DrawImage(s.sprites[i].img, s.sprites[i].options)
	}
	return image, s.op
}
