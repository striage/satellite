package game_state

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"satellite/engine/sprite"
	"satellite/engine/stage"
)

type MainMenuStage struct {
	gameStateManager *GameStateManager
	stage            *stage.Stage
}

func NewMainMenuStage(gsm *GameStateManager) *MainMenuStage {
	return &MainMenuStage{
		gameStateManager: gsm,
		stage:            stage.NewStage(),
	}
}

func (m *MainMenuStage) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		fmt.Println("enter")
		m.gameStateManager.SetState(NewInGameStage(m.gameStateManager))
	}
	return nil
}

func (m *MainMenuStage) SetBackground(sprite *sprite.Sprite) {
	m.stage.SetBackground(sprite)
}

func (m *MainMenuStage) AppendSprite(sprite *sprite.Sprite) {
	m.stage.AppendSprite(sprite)
}

func (m *MainMenuStage) TranslateSprite(name string, x, y int) {
	m.stage.TranslateSprite(name, x, y)
}

func (m *MainMenuStage) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return m.stage.MakeImage()
}
