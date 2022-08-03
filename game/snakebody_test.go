package game

import (
	"testing"
	"snake/utils"
	"snake/vector"
	"snake/sprite"
)

func TestMakeSnake(t *testing.T) {
	snakeBody := MakeSnakeBody(2, utils.UP)

	headPos := snakeBody.GetHead().GetPosition()
	tailPos := snakeBody.GetTail().GetPosition()

	expectedHeadPos := vector.MakeVector2D(0, 0)
	expectedTailPos := vector.MakeVector2D(0, utils.GridSize)

	if !headPos.IsEqual(expectedHeadPos) {
		t.Errorf("MakeSnakeBody(): Head position not set correctly")
	}

	if !tailPos.IsEqual(expectedTailPos) {
		t.Errorf("MakeSnakeBody(): Tail position not set correctly")
	}

}

var getAppendDirectionTests = []struct {
	snakeBody *SnakeBody
	out utils.Direction
}{
	{MakeSnakeBody(2, utils.UP), utils.DOWN},
	{MakeSnakeBody(2, utils.DOWN), utils.UP},
	{MakeSnakeBody(2, utils.RIGHT), utils.LEFT},
	{MakeSnakeBody(2, utils.LEFT), utils.RIGHT},
}

func TestGetAppendDirectionInitial(t *testing.T) {
	var bounds = []*sprite.Sprite{} // Empty because there is no bounds

	for _, tt := range getAppendDirectionTests {
		t.Run("GetAppendDirection() | ", func(t *testing.T) {

			appendDirection := tt.snakeBody.getAppendDirection(bounds)

			if appendDirection != tt.out {
				t.Errorf("Actual: %v, Expected: %v", appendDirection, tt.out)
			}
		})
	}

}

func TestAppend(t *testing.T) {
	snakeBody := MakeSnakeBody(2, utils.UP)
	var bounds = []*sprite.Sprite{} // Empty because there is no bounds
	snakeBody.Append(bounds)

	appendedChunkPos := snakeBody.GetTail().GetPosition()
	expectedPosition := vector.MakeVector2D(0,utils.GridSize * 2)

	if !appendedChunkPos.IsEqual(expectedPosition) {
		t.Errorf("Appended Chunk Is Not Equal. Expected: %v. Actual: %v", expectedPosition, appendedChunkPos)
	}
}

func TestSetPosition(t *testing.T) {
	snakeBody := MakeSnakeBody(2, utils.UP)
	newPosition := vector.MakeVector2D(30, 30)

	snakeBody.SetPosition(newPosition)

	headPos := snakeBody.GetHead().GetPosition()
	tailPos := snakeBody.GetTail().GetPosition()

	expectedHeadPos := vector.MakeVector2D(30, 30)
	expectedTailPos := vector.MakeVector2D(30, 30 + utils.GridSize)


	if !headPos.IsEqual(expectedHeadPos) {
		t.Errorf("TestSetPosition(): Head position not set correctly")
	}

	if !tailPos.IsEqual(expectedTailPos) {
		t.Errorf("TestSetPosition(): Tail position not set correctly")
	}
}

func TestMove(t *testing.T) {
	snakeBody := MakeSnakeBody(2, utils.DOWN)

	snakeBody.Move()

	head := snakeBody.GetHead()
	headPosition := head.GetPosition()

	expectedHeadPosition := vector.MakeVector2D(0,utils.GridSize)

	if !headPosition.IsEqual(expectedHeadPosition) {
		t.Errorf("TestMove() | New head is not at right position." +
		"Expected: %v. Actual: %v", 
		expectedHeadPosition, headPosition)
	}
}

func TestHasSelfCollided(t *testing.T) {
	snakeBody := MakeSnakeBody(6, utils.DOWN)

	snakeBody.SetDirection(utils.RIGHT)
	snakeBody.Move()
	snakeBody.SetDirection(utils.UP)
	snakeBody.Move()
	snakeBody.SetDirection(utils.LEFT)
	snakeBody.Move()

	hasSelfCollided := snakeBody.HasSelfCollided()

	if hasSelfCollided != true {
		t.Errorf("TestHasSelfCollided() | SnakeBody should have self collided")
	}
}

func TestHasCollided(t *testing.T) {
	snakeBody := MakeSnakeBody(2, utils.DOWN)

	obstacle := sprite.MakeSpriteEmpty(utils.GridSize, utils.GridSize)
	obstacle.SetPositionComponents(0,0)

	hasCollided := snakeBody.HasCollided(obstacle)

	if hasCollided != true {
		t.Errorf("TestHasCollided() | SnakeBody should have collided with obstacle." +
		"Obstacle: %v. Head: %v", obstacle, snakeBody.GetHead().GetPosition())
	}
}