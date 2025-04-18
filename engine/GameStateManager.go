package engine

import "github.com/hajimehoshi/ebiten/v2"

type GameStateManager struct {
	currentState GameState
	gameStates   []*GameState
}

func NewGameStateManager(initialState GameState) *GameStateManager {
	//	gameStates := make(map[string]*GameState)
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

func (gsm *GameStateManager) Draw() (*ebiten.Image, *ebiten.DrawImageOptions) {
	return gsm.currentState.Draw()
}
