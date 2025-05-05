package game_state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"satellite/engine/stage"
)

type InGameStage struct {
	gameStateManager *GameStateManager
	stage            *stage.Stage
}

func NewInGameStage(gameStateManager *GameStateManager) *InGameStage {
	return &InGameStage{
		gameStateManager: gameStateManager,
		stage:            stage.NewStage(),
	}
}

func (igs *InGameStage) Update() error {
	return nil
}

func (igs *InGameStage) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return igs.stage.MakeImage()
}
