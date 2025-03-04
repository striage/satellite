package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type InGameStage struct {
	gameStateManager *GameStateManager
	stage            *Stage
}

func NewInGameStage(gameStateManager *GameStateManager) *InGameStage {
	return &InGameStage{
		gameStateManager: gameStateManager,
		stage:            newStage(),
	}
}

func (igs *InGameStage) Update() error {
	return nil
}

func (igs *InGameStage) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return igs.stage.makeImage()
}
