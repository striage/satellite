package engine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Stage struct {
	background *Sprite
	sprites    []*Sprite
	op         *ebiten.DrawImageOptions
}

func NewStage() *Stage {
	return &Stage{
		background: nil,
		sprites:    make([]*Sprite, 0),
		op:         &ebiten.DrawImageOptions{},
	}
}

func (s *Stage) AppendSprite(sprite *Sprite) {
	s.sprites = append(s.sprites, sprite)
}

func (s *Stage) TranslateSprite(name string, x, y int) {
	if sprite, found := s.GetSprite(name); found {
		sprite.options.GeoM.Translate(float64(x), float64(y))
	} else {
		fmt.Printf("[Stage::TranslateSprite] Failed to translate sprite %s as it does not exist\n", name)
	}
}

func (s *Stage) GetSprite(name string) (*Sprite, bool) {
	for _, sprite := range s.sprites {
		if sprite.name == name {
			return sprite, true
		}
	}
	return nil, false
}

func (s *Stage) SetBackground(background *Sprite) {
	s.background = background
}

func (s *Stage) MakeImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	image := ebiten.NewImage(StageWidth, StageHeight)
	if s.background != nil {
		image.DrawImage(s.background.img, s.background.options)
	}
	for i := 0; i < len(s.sprites); i++ {
		image.DrawImage(s.sprites[i].img, s.sprites[i].options)
	}
	return image, s.op
}
