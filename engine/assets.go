package engine

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

//go:embed assets/*
var assets embed.FS

var Rika = mustLoadSprite("assets/rika.png")
var House = mustLoadSprite("assets/house.png")

func mustLoadSprite(name string) *Sprite {
	img := mustLoadImage(name)
	return newSpriteFromImage(name, img, 0, 0)
}

func mustLoadImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
