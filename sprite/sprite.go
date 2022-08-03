package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"snake/vector"
	"image"
	_ "image/png"
	"os"
	"log"
)

type SpriteInterface interface {
	DrawOn(*ebiten.Image)
	SetPosition(*vector.Vector2D)
	GetPosition() *vector.Vector2D
	Image() *ebiten.Image
	HasCollided(SpriteInterface) bool
	GetXMinMax() (float64, float64)
	GetYMinMax() (float64, float64)
}

type Sprite struct {
	image *ebiten.Image
	position *vector.Vector2D
}  

func (s *Sprite) DrawOn(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	x, y := s.position.GetComponents()

	options.GeoM.Translate(x, y)

	screen.DrawImage(s.image, options)
}

func (s *Sprite) SetPosition(vector *vector.Vector2D) {
	s.position.Set(*vector)
}

func (s *Sprite) GetPosition() *vector.Vector2D {
	return s.position
}

func (s *Sprite) Image() *ebiten.Image {
	return s.image
}

func (s *Sprite) HasCollided(sprite SpriteInterface) bool {
	x,y := s.GetPosition().GetComponents()
	widthInt, heightInt := s.Image().Size()
	width := float64(widthInt)
	height := float64(heightInt)

	cornerTopLeft := vector.MakeVector2D(x, y)
	cornerTopRight := vector.MakeVector2D(x + width, y)
	cornerBottomLeft := vector.MakeVector2D(x, y+height)
	cornerBottomRight := vector.MakeVector2D(x+width, y+height)

	return s.isCornerIn(sprite, cornerTopLeft) ||
		s.isCornerIn(sprite, cornerTopRight) ||
		s.isCornerIn(sprite, cornerBottomLeft) ||
		s.isCornerIn(sprite, cornerBottomRight)
}

func (s *Sprite) isCornerIn(sprite SpriteInterface, corner *vector.Vector2D) bool {
	x, y := sprite.GetPosition().GetComponents()
	widthInt, heightInt := sprite.Image().Size()
	width := float64(widthInt)
	height := float64(heightInt)

	cornerX, cornerY := corner.GetComponents()

	return cornerX >= x &&
		cornerX <= x + width &&
		cornerY >= y &&
		cornerY <= y + height
}

func (s *Sprite) IsWithin(sprite SpriteInterface) bool {
	currXMin, currXMax := s.GetXMinMax()
	currYMin, currYMax := s.GetYMinMax()

	spriteXMin, spriteXMax := sprite.GetXMinMax()
	spriteYMin, spriteYMax := sprite.GetYMinMax()

	return (spriteXMin <= currXMin && spriteXMin <= currXMax) &&
	(spriteXMax >= currXMin && spriteXMax >= currXMax) &&
	(spriteYMin <= currYMin && spriteYMin <= currYMax) &&
	(spriteYMax >= currYMin && spriteYMax >= currYMax)
}

func (s *Sprite) GetXMinMax() (float64, float64) {
	x, _:= s.position.GetComponents()
	width, _ := s.Image().Size()

	return x, x + float64(width)
}

func (s *Sprite) GetYMinMax() (float64, float64) {
	_, y := s.position.GetComponents()
	_, height := s.Image().Size()

	return y, y + float64(height)
}

func (s *Sprite) SetPositionComponents(x,y float64) {
	s.position.SetComponents(x, y)
}

func MakeSprite(fileName string) *Sprite {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer f.Close()
	image, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Sprite {
		image: ebiten.NewImageFromImage(image),
		position: vector.MakeVector2D(0,0),
	}
}

func MakeSpriteEmpty(width, height int) *Sprite {
	return &Sprite {
		image: ebiten.NewImage(width, height),
		position: vector.MakeVector2D(0,0),
	}
}
