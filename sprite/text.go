package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image/color"
)

import (
	"snake/vector"
	"snake/utils"
)

type SpriteText struct {
	*Sprite
	fontSize float64
	text string
	fontColor color.Color
	font font.Face
}

func (st *SpriteText) DoesCursorIntersect() bool {

	xInt,yInt := ebiten.CursorPosition()
	x := float64(xInt)
	y := float64(yInt)

	xStart, xEnd := st.GetXMinMax()
	yStart, yEnd := st.GetYMinMax()

	return x > xStart && 
	x < xEnd &&
	y > yStart &&
	y < yEnd
}

func (st *SpriteText) ChangeText(newText string) {
	st.text = newText

	st.image.Clear()

	text.Draw(st.image, st.text, st.font, 5, int(st.fontSize*1.5), st.fontColor)
}

func (st *SpriteText) OnNotHover() {
	fontColor := utils.GetWhiteColor()
	st.image.Fill(utils.GetBlackColor())

	text.Draw(st.image, st.text, st.font, 5, int(st.fontSize*1.5), fontColor)
}

func (st *SpriteText) OnHover() {
	fontColor := utils.GetBlackColor()
	st.image.Fill(utils.GetWhiteColor())

	text.Draw(st.image, st.text, st.font, 5, int(st.fontSize*1.5), fontColor)
}

func GetLength(font font.Face, inputText string) int {
	cumulative_advance := fixed.I(0)
	cumulative_kern := fixed.I(0)

	previous_char := '\n'

	for _, char := range(inputText) {

		_, advance, _ := font.GlyphBounds(char)

		if previous_char != '\n' {
			cumulative_kern += (font.Kern(previous_char, char))
		}

		previous_char = char
		cumulative_advance += advance 
	}

	return cumulative_advance.Round() + cumulative_kern.Round()
}

func GetFont(inputText string, fontSize float64) font.Face {

	dpi := 120.0
	tt, _ := opentype.Parse(fonts.MPlus1pRegular_ttf)
	font, _ := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	return font
}

func MakeSpriteText(inputText string, fontSize float64, pos *vector.Vector2D) *SpriteText {
	

	tt := GetFont(inputText, fontSize)
	font_pixel_length := GetLength(tt, inputText)

	imgLengthOffset := 10
	imgHeightOffset := fontSize * 1.05
	imgLength := font_pixel_length + imgLengthOffset
	imgHeight := int(fontSize) + int(imgHeightOffset)

	img := ebiten.NewImage(imgLength, imgHeight)

	fontColor := utils.GetWhiteColor()
	img.Fill(utils.GetBlackColor())

	text.Draw(img, inputText, tt, 5, int(fontSize*1.5), fontColor)

	sprite := &Sprite {
		image: img,
		position: pos,
	}

	return &SpriteText {
		Sprite: sprite,
		fontSize: fontSize,
		text: inputText,
		fontColor: fontColor,
		font: tt,
	}
}

