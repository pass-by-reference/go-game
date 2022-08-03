package game

type GameState struct {
	Counter int
	KeyRecentlyPressed bool
	IsMenuActive bool
	WillQuit bool
	Quit bool
}

func MakeGameState() *GameState {
	return &GameState {
		Counter: 0,
		KeyRecentlyPressed: false,
		IsMenuActive: true,
		WillQuit: false,
		Quit: false,
	}
}