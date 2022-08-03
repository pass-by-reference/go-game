package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"errors"
	"time"
)

import (
	"snake/game"
	"snake/utils"
)

type SnakeGame struct {
	gameState *game.GameState
	snakeScreen *game.SnakeScreen
	mainMenuScreen *game.MainMenuScreen
}

func (g *SnakeGame) WaitNSecondBeforeQuit(seconds time.Duration) {
	time.Sleep(seconds * time.Second)

	g.gameState.Quit = true
}

func (g *SnakeGame) Update() error {

	if g.gameState.IsMenuActive {
		g.mainMenuScreen.Update(g.gameState)
	} else {
		g.snakeScreen.Update(g.gameState)
	}

	if g.gameState.WillQuit {
		go g.WaitNSecondBeforeQuit(2)
	}

	if g.gameState.Quit {
		return errors.New("terminated")
	}

	g.gameState.Counter += 1
	return nil
}

func (g *SnakeGame) Draw(screen *ebiten.Image) {
	
	if g.gameState.IsMenuActive {
		g.mainMenuScreen.DrawOn(screen)
	} else {
		g.snakeScreen.DrawOn(screen)
	}

}

func (g *SnakeGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &SnakeGame{
		gameState: game.MakeGameState(),
		snakeScreen: game.MakeSnakeScreen(),
		mainMenuScreen: game.MakeMainMenuScreen(),
	}

	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}