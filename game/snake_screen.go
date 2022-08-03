package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"snake/utils"
	"snake/vector"
	"snake/sprite"
	"fmt"
)

type Status int

const (
	DEFAULT Status = iota
	NONE
	WIN
	LOSE
)

type SnakeScreen struct {
	board *Board
	scoreText *sprite.SpriteText
	statusText *sprite.SpriteText
	Score int
	Status Status
}

func (sc *SnakeScreen) DrawOn(screen *ebiten.Image) {
	sc.board.DrawOn(screen)
	sc.scoreText.DrawOn(screen)
	sc.statusText.DrawOn(screen)
}

func (sc *SnakeScreen) Update(g *GameState) {
	board := sc.board

	if board.IsWin {
		board.Snake().ShouldMove = false
		sc.Status = WIN
		sc.DrawStatus()

		g.WillQuit = true
	}

	hasFoodCollided, index := board.HasSnakeCollidedFood()
	if hasFoodCollided {
		board.RemoveFood(index)
		board.Snake().Append(board.Bounds())
		board.AddFood()
		
		sc.Score += 1
	}

	if board.HasSnakeCollidedBorder() {
		board.Snake().ShouldMove = false
		
		sc.Status = LOSE
		sc.DrawStatus()

		g.WillQuit = true
	}

	if board.Snake().HasSelfCollided() {
		board.Snake().ShouldMove = false

		sc.Status = LOSE
		sc.DrawStatus()

		g.WillQuit = true
	}
	
	if g.Counter % 20 == 0 && board.Snake().ShouldMove {
		board.Snake().Move()
	
		if g.KeyRecentlyPressed {
			g.KeyRecentlyPressed = false
		}
	}

	if !g.KeyRecentlyPressed {
		g.KeyRecentlyPressed = HandleKeyPressed(sc.board)
	}

	sc.DrawStatus()
	sc.DrawScoreText()
}

func (sc *SnakeScreen) DrawScoreText() {
	score := fmt.Sprintf("Score: %d ", sc.Score)

	sc.scoreText.ChangeText(score)
}

func (sc *SnakeScreen) DrawStatus() {
	statusMessage := ""
	switch sc.Status {
	case NONE:
		statusMessage = "WASD To Move"
	case WIN:
		statusMessage = "     You Win"
	case LOSE:
		statusMessage = "    You Lose"
	}

	sc.statusText.ChangeText(statusMessage)
}

func HandleKeyPressed(board *Board) bool {

	recentlyPressed := false

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		board.Snake().SetDirection(utils.LEFT)
		recentlyPressed = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		board.Snake().SetDirection(utils.DOWN)
		recentlyPressed = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		board.Snake().SetDirection(utils.UP)
		recentlyPressed = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		board.Snake().SetDirection(utils.RIGHT)
		recentlyPressed = true
	}

	return recentlyPressed
}

func SetBoardCenter(board *Board) {
	boardStartX := (utils.ScreenWidth - utils.BoardLength)/2
	boardStartY := (utils.ScreenHeight - utils.BoardLength)/2
	board.TranslateBoardPosition(vector.MakeVector2D(boardStartX, boardStartY))
}

func SetScoreTextPos(scoreText *sprite.SpriteText) {
	width, _ := scoreText.Image().Size()

	x := float64((utils.ScreenWidth - width)/2)
	y := utils.ScreenHeight * 0.1

	scoreText.SetPositionComponents(x,y)
}

func SetStatusPos(scoreText *sprite.SpriteText) {
	width, _ := scoreText.Image().Size()

	x := float64((utils.ScreenWidth - width)/2)
	y := utils.ScreenHeight * 0.2

	scoreText.SetPositionComponents(x,y)
}


func MakeSnakeScreen() *SnakeScreen {
	board := MakeBoard()
	SetBoardCenter(board)
	board.AddFood()

	scoreText := sprite.MakeSpriteText("Score: 00", 15, vector.MakeVector2D(0,0))
	statusText := sprite.MakeSpriteText("WASD To Move", 10, vector.MakeVector2D(0,0))

	SetScoreTextPos(scoreText)
	SetStatusPos(statusText)

	return &SnakeScreen {
		board: board,
		scoreText: scoreText,
		statusText: statusText,
		Score: 0,
		Status: NONE,
	}
}