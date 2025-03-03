package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	stage *Stage
}

func NewGame() *Game {
	game := Game{
		stage: newStage(),
	}
	game.stage.sprites = append(game.stage.sprites, Rika)
	game.stage.background = House
	return &game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	stageImage := g.stage.makeImage()
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(stageImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
