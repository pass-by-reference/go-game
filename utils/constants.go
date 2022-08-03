package utils

type Direction int

const (
	DEFAULT Direction = iota
	LEFT
	RIGHT
	UP
	DOWN
)

const (
	GridSize = 20.0
	BoardLength = GridSize * 10
	BorderSize = GridSize
)

const ASSET_PATH = "/home/panda/Desktop/Repository/go-game/assets/"

const (
	ScreenWidth = 400
	ScreenHeight = 400
)

func GetOpposite(dir Direction) Direction {
	switch dir {
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	case UP:
		return DOWN
	case DOWN:
		return UP
	default:
		return DEFAULT
	}
}
