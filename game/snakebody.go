package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"snake/sprite"
	"snake/utils"
	"snake/vector"
)

type StartingLocation func() (float64, float64)

type SnakeBody struct {
	body []*sprite.Sprite
	direction utils.Direction
	ShouldMove bool
}

func (s *SnakeBody) DrawOn(screen *ebiten.Image) {
	for _, sprite := range(s.body) {
		sprite.DrawOn(screen)
	}
}

func (s *SnakeBody) Move() {
	startingPos := s.GetHead().GetPosition()

	newHeadPosition := s.getAppendPosition(startingPos, s.direction)
	
	tail := s.GetTail()
	tail.SetPosition(newHeadPosition)

	newBody := []*sprite.Sprite {tail}

	bodyWithoutLastElement := s.body[:(len(s.body)-1)]
	s.body = append(newBody, bodyWithoutLastElement...)
}

func (s *SnakeBody) Append(bounds []*sprite.Sprite) {
	newBody := sprite.MakeSpriteEmpty(utils.GridSize, utils.GridSize)
	color := utils.GetGreenColor()

	newBody.Image().Fill(color)

	startingPosition := s.GetTail().GetPosition()
	appendDirection := s.getAppendDirection(bounds)

	position := s.getAppendPosition(startingPosition, appendDirection)

	newBody.SetPosition(position)
	s.body = append(s.body, newBody)
}

func (s *SnakeBody) getAppendDirection(bounds []*sprite.Sprite) utils.Direction {
	initialDirection := s.getInitialAppendDirection()
	collideWithBorder := s.doesAppendCollideWithBorder(initialDirection, bounds)

	if collideWithBorder {
		var directions = []utils.Direction{utils.UP, utils.DOWN, utils.RIGHT, utils.LEFT}

		for _, direction := range(directions) {
			if !s.doesAppendCollideWithBorder(direction, bounds) {
				return direction
			}
		}
	} 

	return initialDirection
}

func (s *SnakeBody) getInitialAppendDirection() utils.Direction {
	snakeLength := len(s.body)
	lastChunkX, lastChunkY := s.body[snakeLength - 1].GetPosition().GetComponents()
	secLastChunkX, secLastChunkY := s.body[snakeLength - 2].GetPosition().GetComponents()

	if 
		lastChunkX == secLastChunkX &&
		lastChunkY < secLastChunkY {
		return utils.UP
	}

	if lastChunkX == secLastChunkX &&
		 lastChunkY > secLastChunkY {
		return utils.DOWN
	}

	if lastChunkY == secLastChunkY &&
	   lastChunkX < secLastChunkX {
		return utils.LEFT
	}

	if lastChunkY == secLastChunkY &&
		 lastChunkX > secLastChunkX {
		return utils.RIGHT
	}

	return utils.DEFAULT
}

func (s *SnakeBody) doesAppendCollideWithBorder(direction utils.Direction, bounds []*sprite.Sprite) bool {
	startingX, startingY := s.GetTail().GetPosition().GetComponents()
	position := vector.MakeVector2D(0.0, 0.0)

	switch direction {
	case utils.DOWN:
		position.SetComponents(startingX, startingY + utils.GridSize)
	case utils.UP:
		position.SetComponents(startingX, startingY - utils.GridSize)
	case utils.LEFT:
		position.SetComponents(startingX - utils.GridSize, startingY)
	case utils.RIGHT:
		position.SetComponents(startingX + utils.GridSize, startingY)
	case utils.DEFAULT:
	}

	for _, bound := range(bounds) {
		dummy_sprite := sprite.MakeSpriteEmpty(utils.GridSize, utils.GridSize)
		dummy_sprite.SetPosition(position)

		if dummy_sprite.IsWithin(bound) {
			return true
		}
	}

	return false
}

func (s *SnakeBody) getAppendPosition(startingPos *vector.Vector2D, direction utils.Direction) *vector.Vector2D {
	
	position := vector.MakeVector2D(0.0, 0.0)
	startingX, startingY := startingPos.GetComponents()

	switch direction {
	case utils.DOWN:
		position.SetComponents(startingX, startingY + utils.GridSize)
	case utils.UP:
		position.SetComponents(startingX, startingY - utils.GridSize)
	case utils.LEFT:
		position.SetComponents(startingX - utils.GridSize, startingY)
	case utils.RIGHT:
		position.SetComponents(startingX + utils.GridSize, startingY)
	case utils.DEFAULT:
		return position
	}

	return position
}

func (s *SnakeBody) SetPosition(headPosition *vector.Vector2D) {
	for _, sprite := range(s.body) {

		spriteNewPosition := sprite.GetPosition().Add(*headPosition)
		sprite.SetPosition(spriteNewPosition)
	}
}

func (s *SnakeBody) GetTail() *sprite.Sprite {
	snakeLength := len(s.body)

	return s.body[snakeLength - 1]
}

func (s *SnakeBody) GetHead() *sprite.Sprite {
	return s.body[0]
}

func (s *SnakeBody) SetDirection(direction utils.Direction) {
	oppositeDirection := utils.GetOpposite(s.direction)
	if direction == oppositeDirection {
		return
	}

	s.direction = direction
}

func (s *SnakeBody) HasSelfCollided() bool {
	head := s.GetHead()
	bodyNoHead := s.body[1:]
	for _, sprite := range(bodyNoHead) {
		if head.IsWithin(sprite) {
			return true
		}
	}

	return false
}

func (s *SnakeBody) HasCollided(sprite *sprite.Sprite) bool {
	head := s.GetHead()

	return head.IsWithin(sprite)
}

func (s *SnakeBody) CanFoodSpawnAt(x,y float64) bool {
	for _, spriteBody := range(s.body) {
		spriteX, spriteY := spriteBody.GetPosition().GetComponents()

		if spriteX == x && spriteY == y {
			return false
		}
	}

	return true
}

func MakeSnakeBody(initialLength int, startDirection utils.Direction) *SnakeBody {
	var body []*sprite.Sprite

	appendDirection := utils.GetOpposite(startDirection)

	for i := 0; i < initialLength; i++ {
		x, y := getNextChunkPosition(i, appendDirection)

		sprite := sprite.MakeSpriteEmpty(utils.GridSize, utils.GridSize)
		color := utils.GetGreenColor()
		sprite.Image().Fill(color)
		sprite.SetPositionComponents(x, y)

		body = append(body, sprite)
	}

	return &SnakeBody {
		direction: startDirection,
		body: body,
		ShouldMove: true,
	}
}

func getNextChunkPosition(chunkNum int, appendDirection utils.Direction) (float64, float64) {
	
	startingX, startingY := 0.0, 0.0

	x, y := 0.0, 0.0
	switch appendDirection {
	case utils.LEFT:
		x = startingX - float64(chunkNum) * float64(utils.GridSize)
		y = startingY
	case utils.RIGHT:
		x = startingX + float64(chunkNum) * float64(utils.GridSize)
		y = startingY
	case utils.UP:
		x = startingX
		y = startingY - float64(chunkNum) * float64(utils.GridSize)
	case utils.DOWN:
		x = startingX
		y = startingY + float64(chunkNum) * float64(utils.GridSize)
	}

	return x, y
}
