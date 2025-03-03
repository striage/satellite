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
		stage:            newStage(), // Create the in-game scene
	}
}

func (igs *InGameStage) Update() error {
	return nil
}

func (igs *InGameStage) Draw(screen *ebiten.Image) {
	stageImage, op := igs.stage.makeImage()
	screen.DrawImage(stageImage, op)
}
