package utils

import (
	"image/color"
)

func GetGreenColor() color.Color {
	return color.RGBA {
		R: 70,
		G: 255,
		B: 0,
		A: 255,
	}
}

func GetRedColor() color.Color {
	return color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
}

func GetWhiteColor() color.Color {
	return color.RGBA {
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
}

func GetBlackColor() color.Color {
	return color.RGBA {
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
}

func GetNearestTen(number int) float64 {
	remainder := number % 10

	return float64(number - remainder)
}