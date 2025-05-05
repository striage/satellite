package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"satellite/engine/consts"
	"satellite/engine/game_state"
)

type Game struct {
	gameStateManager *game_state.GameStateManager
	assetManager     *AssetManager
	width            int
	height           int
}

func NewGame() *Game {
	gameStateManager := game_state.NewGameStateManager(nil)
	initialStage := game_state.NewMainMenuStage(gameStateManager)
	assetManager := NewAssetManager()

	initialStage.AppendSprite(assetManager.getAsset("rika"))
	initialStage.AppendSprite(assetManager.getAsset("dir/button"))
	initialStage.AppendSprite(assetManager.getAsset("dir/newdir/button"))

	initialStage.TranslateSprite("rika", 900, 0)
	initialStage.TranslateSprite("dir/button", 200, 0)
	initialStage.TranslateSprite("dir/newdir/button", 400, 0)
	initialStage.SetBackground(assetManager.getAsset("house"))

	gameStateManager.SetState(initialStage)

	game := Game{
		gameStateManager: gameStateManager,
		width:            consts.StageWidth,
		height:           consts.StageHeight,
	}
	return &game
}

func (g *Game) Update() error {
	return g.gameStateManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	stageImage, op := g.gameStateManager.Draw()
	screen.DrawImage(stageImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return consts.StageWidth, consts.StageHeight
}
