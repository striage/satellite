package engine

import "github.com/hajimehoshi/ebiten/v2"

type Stage struct {
	background *Sprite
	sprites    []*Sprite
}

func newStage() *Stage {
	return &Stage{
		background: nil,
		sprites:    make([]*Sprite, 0),
	}
}

func (stage *Stage) makeImage() *ebiten.Image {
	image := ebiten.NewImage(StageWidth, StageHeight)
	if stage.background != nil {
		image.DrawImage(stage.background.img, stage.background.options)
	}
	for i := 0; i < len(stage.sprites); i++ {
		image.DrawImage(stage.sprites[i].img, stage.sprites[i].options)
	}
	return image
}
