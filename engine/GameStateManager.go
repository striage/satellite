package engine

import "github.com/hajimehoshi/ebiten/v2"

type GameStateManager struct {
	currentState GameState
}

func NewGameStateManager(initialState GameState) *GameStateManager {
	return &GameStateManager{
		currentState: initialState,
	}
}

func (gsm *GameStateManager) SetState(state GameState) {
	gsm.currentState = state
}

func (gsm *GameStateManager) Update() error {
	if gsm.currentState != nil {
		return gsm.currentState.Update()
	}
	return nil
}

func (gsm *GameStateManager) Draw(screen *ebiten.Image) {
	if gsm.currentState != nil {
		gsm.currentState.Draw(screen)
	}
}
