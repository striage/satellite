package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	stage        *Stage
	assetManager *AssetManager
	width        int
	height       int
}

func NewGame() *Game {
	game := Game{
		assetManager: newAssetManager(),
		stage:        newStage(),
		width:        StageWidth,
		height:       StageHeight,
	}
	game.stage.appendSprite(game.assetManager.getAsset("rika"))
	game.stage.setBackground(game.assetManager.getAsset("house"))
	return &game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	stageImage, op := g.stage.makeImage()
	windowWidth, windowHeight := ebiten.WindowSize()
	scaleX := float64(windowWidth) / float64(g.width)
	scaleY := float64(windowHeight) / float64(g.height)
	scale := min(scaleX, scaleY)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(stageImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
