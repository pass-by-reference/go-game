package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"snake/sprite"
	"snake/vector"
	"snake/utils"
)

type MainMenuScreen struct {
	texts [3]*sprite.SpriteText
}

func (m *MainMenuScreen) Update(g *GameState) {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		m.OnClickEvents(g)
	}

	m.OnHoverEvents()

}

func (m *MainMenuScreen) DrawOn(screen *ebiten.Image) {
	for _, text := range(m.texts) {
		text.DrawOn(screen)
	}
}

func (m *MainMenuScreen) OnClickEvents(g *GameState) {
	playText := m.texts[1]
	quitText := m.texts[2]

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		if playText.DoesCursorIntersect() {
			g.IsMenuActive = false
		}

		if quitText.DoesCursorIntersect() {
			g.Quit = true
		}
	}
}

func (m *MainMenuScreen) OnHoverEvents() {

	for index, text := range(m.texts) {
		if text.DoesCursorIntersect() && index != 0 {
			text.OnHover()
		} else {
			text.OnNotHover()
		}
	}

}

func SetTextPosition(texts [3]*sprite.SpriteText) {

	yCumulative := utils.ScreenHeight/10
	previousTextLength := yCumulative
	titleButtonYOffset := int(utils.ScreenHeight * 0.1)

	for index, text := range(texts) {
		width, length := text.Image().Size()

		x := (utils.ScreenWidth - width)/2
		y := 0
		if index == 1 {
			y = yCumulative + previousTextLength + titleButtonYOffset
			yCumulative += previousTextLength + titleButtonYOffset
		} else {
			y = yCumulative + previousTextLength
			yCumulative += previousTextLength
		}

		
		
		previousTextLength = length

		text.SetPositionComponents(float64(x), float64(y))
	}
	
}

func MakeMainMenuScreen() *MainMenuScreen {
	
	titleText := sprite.MakeSpriteText("Snake", 30, vector.MakeVector2D(0, 0))
	playText := sprite.MakeSpriteText("Play", 20, vector.MakeVector2D(0, 0))
	quitText := sprite.MakeSpriteText("Quit", 20, vector.MakeVector2D(0, 0))

	texts := [3]*sprite.SpriteText{titleText, playText, quitText}

	SetTextPosition(texts)
	
	return &MainMenuScreen {
		texts: texts,
	}
}

