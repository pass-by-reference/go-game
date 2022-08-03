package vector

import (
	"math"
)

type Vector interface {
	Magnitude() float64
	Normalized() Vector
	Add(Vector) Vector
	Subtract(Vector) Vector
	Multiply(float64) Vector
	Divide(float64) Vector
}

type Vector2D struct {
	x float64
	y float64
}

func (v *Vector2D) Magnitude() float64 {
	x_squared := math.Pow(float64(v.x), 2)
	y_squared := math.Pow(float64(v.y), 2)

	return math.Sqrt(x_squared + y_squared)
}

// Normalized is different than normal vector
func (v *Vector2D) Normalized() Vector2D {

	magnitude := v.Magnitude()

	return Vector2D {
		x: float64(v.x)/magnitude,
		y: float64(v.y)/magnitude,
	}
}

func (v *Vector2D) Add(vector Vector2D) *Vector2D {
	return &Vector2D {
		x: v.x + vector.x,
		y: v.y + vector.y,
	}
}

func (v *Vector2D) Subtract(vector Vector2D) *Vector2D {
	return &Vector2D {
		x: v.x - vector.x,
		y: v.y - vector.y,
	}
}

func (v *Vector2D) Multiply(scalar float64) *Vector2D {
	return &Vector2D {
		x: v.x * scalar,
		y: v.y * scalar,
	}
}

func (v *Vector2D) Divide(scalar float64) *Vector2D {
	return &Vector2D {
		x: v.x / scalar,
		y: v.y / scalar,
	}
}

func (v *Vector2D) Set(vector Vector2D) {
	v.x = vector.x
	v.y = vector.y
}

func (v *Vector2D) SetComponents(x,y float64) {
	v.x = x
	v.y = y
}

func (v *Vector2D) GetComponents() (float64, float64) {
	return v.x, v.y
}

func (v *Vector2D) IsEqual(vector *Vector2D) bool {
	if v.x == vector.x &&
		 v.y == vector.y {
		return true
	}

	return false
}

func MakeVector2D(x,y float64) *Vector2D {
	return &Vector2D {
		x: x,
		y: y,
	}
}