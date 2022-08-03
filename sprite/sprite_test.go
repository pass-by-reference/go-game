package sprite

import (
	"testing"
	"snake/vector"
)

func TestHasCollided(t *testing.T) {
	snakeBodySize := 10

	figureOne := MakeSpriteEmpty(snakeBodySize, snakeBodySize)
	figureOnePosition := vector.MakeVector2D(0,0)
	figureOne.SetPosition(figureOnePosition)

	figureTwo := MakeSpriteEmpty(snakeBodySize, snakeBodySize)
	figureTwoPosition := vector.MakeVector2D(9.9,9.9)
	figureTwo.SetPosition(figureTwoPosition)

	hasCollided := figureOne.HasCollided(figureTwo)

	if hasCollided == false {
		t.Fatalf(`TestHasCollided() | Expected: true. Actual: %v`, hasCollided)
	}
}