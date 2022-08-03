package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
	"snake/sprite"
	"snake/vector"
	"snake/utils"
)

type Board struct {
	snake *SnakeBody
	bounds []*sprite.Sprite
	boardPosition *vector.Vector2D
	food []*sprite.Sprite
	IsWin bool
}

func (b *Board) DrawOn(screen *ebiten.Image) {
	for _, sprite := range(b.bounds) {
		sprite.DrawOn(screen)
	}

	for _, food := range(b.food) {
		food.DrawOn(screen)
	}

	b.snake.DrawOn(screen)
}

func (b *Board) TranslateBoardPosition(translate *vector.Vector2D) {
	b.boardPosition = b.boardPosition.Add(*translate)

	for _, bound := range(b.bounds) {
		newBounds := bound.GetPosition().Add(*translate)

		bound.SetPosition(newBounds)
	}

	newSnakePos := b.snake.GetHead().GetPosition().Add(*translate)
	b.snake.SetPosition(newSnakePos)
}

func (b *Board) GetCorners() []*vector.Vector2D {
	var corners []*vector.Vector2D

	baseX, baseY := b.boardPosition.GetComponents()

	topLeft := vector.MakeVector2D(baseX, baseY)
	topRight := vector.MakeVector2D(baseX + utils.BoardLength, baseY)
	bottomLeft := vector.MakeVector2D(baseX, baseY + utils.BoardLength)
	bottomRight := vector.MakeVector2D(baseX + utils.BoardLength, baseY)

	corners = append(corners, topLeft, topRight, bottomLeft, bottomRight)

	return corners
}

func (b *Board) AddFood() {
	food := sprite.MakeSpriteEmpty(utils.GridSize, utils.GridSize)
	color := utils.GetRedColor()

	food.Image().Fill(color)

	x,y := b.GetNewFoodPosition()
	food.SetPositionComponents(x, y)

	b.food = append(b.food, food)
}

func (b *Board) RemoveFood(removeIndex int) {
	var newFood = make([]*sprite.Sprite, 0, 10)

	for foodIndex, food := range(b.food) {
		if foodIndex == removeIndex {
			continue
		}

		newFood = append(newFood, food)
	}

	b.food = newFood
}

func (b *Board) HasSnakeCollidedFood() (bool, int) {

	hasCollided := false
	removeFoodIndex := -1

	for foodIndex, food := range(b.food) {
		if b.snake.HasCollided(food) {
			hasCollided = true
			removeFoodIndex = foodIndex
		}
	}

	return hasCollided, removeFoodIndex

}

func (b *Board) HasSnakeCollidedBorder() bool {
	for _, border := range(b.bounds) {
		if b.snake.HasCollided(border) {
			return true
		}
	}

	return false
}

func (b *Board) GetNewFoodPosition() (float64, float64) {
	listFreeFood := b.GetFreeSpaceForFood()

	if len(listFreeFood) > 0 {
		randomInt := rand.Intn(len(listFreeFood))

		return listFreeFood[randomInt].GetComponents()
	}

	b.IsWin = true

	return -100.0, -100.0
}

func (b *Board) Snake() *SnakeBody {
	return b.snake
}

func (b *Board) GetFreeSpaceForFood() []*vector.Vector2D {
	boardX, boardY := b.boardPosition.GetComponents()
	startX := boardX + utils.BorderSize
	endX := boardX + utils.BoardLength - utils.BorderSize

	startY := boardY + utils.BorderSize
	endY := boardY + utils.BoardLength - utils.BorderSize

	canSpawnList := make([]*vector.Vector2D, 0, 100)

	for column := startX; column < endX; column += utils.GridSize {
		for row := startY; row < endY; row += utils.GridSize {

			if b.snake.CanFoodSpawnAt(column, row) {
				vector := vector.MakeVector2D(column, row)
				canSpawnList = append(canSpawnList, vector)
			}
		}
	}

	return canSpawnList
}

func (b *Board) Bounds() []*sprite.Sprite {
	return b.bounds
}

func MakeBoard() *Board {

	snake := MakeSnakeBody(2, utils.RIGHT)
	vectorPos := vector.MakeVector2D(utils.GridSize, utils.GridSize)
	snake.SetPosition(vectorPos)

	var bounds []*sprite.Sprite
	bounds = setupBorderPositions(bounds)
	bounds = fillColor(bounds)

	boardPosition := vector.MakeVector2D(0,0)

	var emptyFoodList = make([]*sprite.Sprite, 0, 10)

	board := &Board {
		snake: snake,
		bounds: bounds,
		boardPosition: boardPosition,
		food: emptyFoodList,
		IsWin: false,
	}

	return board
}

func setupBorderPositions(bounds []*sprite.Sprite) []*sprite.Sprite {

	boundHorizontalTop := sprite.MakeSpriteEmpty(utils.BoardLength, utils.BorderSize)
	boundHorizontalBottom := sprite.MakeSpriteEmpty(utils.BoardLength, utils.BorderSize)
	boundVerticalLeft := sprite.MakeSpriteEmpty(utils.BorderSize, utils.BoardLength)
	boundVerticalRight := sprite.MakeSpriteEmpty(utils.BorderSize, utils.BoardLength)

	boundHorizontalTop.SetPosition(vector.MakeVector2D(0,0))
	boundHorizontalBottom.SetPosition(vector.MakeVector2D(0, utils.BoardLength - utils.BorderSize))
	boundVerticalRight.SetPosition(vector.MakeVector2D(utils.BoardLength - utils.BorderSize, 0))
	boundVerticalLeft.SetPosition(vector.MakeVector2D(0, 0))

	return append(bounds, boundHorizontalTop, boundHorizontalBottom, boundVerticalLeft, boundVerticalRight)
}

func fillColor(bounds []*sprite.Sprite) []*sprite.Sprite {
	color := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	for _, boundsSprite := range(bounds) {
		boundsSprite.Image().Fill(color)
	}

	return bounds
}
