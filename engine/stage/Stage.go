package stage

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"satellite/engine/consts"
	"satellite/engine/sprite"
)

type Stage struct {
	background *sprite.Sprite
	sprites    []*sprite.Sprite
	op         *ebiten.DrawImageOptions
}

func NewStage() *Stage {
	return &Stage{
		background: nil,
		sprites:    make([]*sprite.Sprite, 0),
		op:         &ebiten.DrawImageOptions{},
	}
}

func (s *Stage) AppendSprite(sprite *sprite.Sprite) {
	s.sprites = append(s.sprites, sprite)
}

func (s *Stage) TranslateSprite(name string, x, y int) {
	if sprite, found := s.GetSprite(name); found {
		sprite.Options.GeoM.Translate(float64(x), float64(y))
	} else {
		fmt.Printf("[Stage::TranslateSprite] Failed to translate sprite %s as it does not exist\n", name)
	}
}

func (s *Stage) GetSprite(name string) (*sprite.Sprite, bool) {
	for _, sprite := range s.sprites {
		if sprite.Name == name {
			return sprite, true
		}
	}
	return nil, false
}

func (s *Stage) SetBackground(background *sprite.Sprite) {
	s.background = background
}

func (s *Stage) MakeImage() (*ebiten.Image, *ebiten.DrawImageOptions) {
	image := ebiten.NewImage(consts.StageWidth, consts.StageHeight)
	if s.background != nil {
		image.DrawImage(s.background.Img, s.background.Options)
	}
	for i := 0; i < len(s.sprites); i++ {
		image.DrawImage(s.sprites[i].Img, s.sprites[i].Options)
	}
	return image, s.op
}
