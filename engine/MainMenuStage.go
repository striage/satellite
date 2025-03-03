package engine

import "github.com/hajimehoshi/ebiten/v2"

type MainMenuStage struct {
	gameStateManager *GameStateManager
	stage            *Stage
}

func NewMainMenuStage(gsm *GameStateManager) *MainMenuStage {
	return &MainMenuStage{
		gameStateManager: gsm,
		stage:            newStage(),
	}
}

func (m *MainMenuStage) Update() error {
	// Example: Simulating a button press to start the game
	if /* some condition like a button click */ false {
		m.gameStateManager.SetState(NewInGameStage(m.gameStateManager))
	}
	return nil
}

func (m *MainMenuStage) Draw(screen *ebiten.Image) {
	img, op := m.stage.makeImage()
	screen.DrawImage(img, op)
}
