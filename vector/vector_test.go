package vector

import (
	"testing"
	"math"
)

func roundFloat(v float64, precision int) float64 {
	baseTenPrecision := math.Pow(10, float64(precision))
	roundedInt := math.Round(v * baseTenPrecision)

	return roundedInt / baseTenPrecision 
}

func TestMakeVector2D(t *testing.T) {
	expected_x := 10.0
	expected_y := 15.0

	vector := MakeVector2D(expected_x, expected_y)

	if vector.x != expected_x {
		t.Fatalf(`TestMakeVector2D() | Actual x: %v. Expected x: %v`, vector.x, expected_x)
	}

	if vector.y != 15 {
		t.Fatalf(`TestMakeVector2D() | Actual y: %v. Expected y: %v`, vector.y, expected_y)
	}
}

func TestMagnitude(t *testing.T) {
	vector := MakeVector2D(2.0,5.0)

	mag_expected := 5.3851
	mag_actual := vector.Magnitude()

	if mag_actual == mag_expected {
		t.Fatalf(`TestMagnitude() | 
			Actual magnitude: %v. 
			Expected magnitude: %v`, 
		mag_actual, mag_expected)
	}
}

func TestNormalized(t *testing.T) {
	vector := MakeVector2D(2,5)

	expectedXComp := 0.37139
	expectedYComp := 0.92848
	normal_vector := vector.Normalized()

	actual_x := roundFloat(normal_vector.x, 5)
	actual_y := roundFloat(normal_vector.y, 5)

	if actual_x != expectedXComp {
		t.Fatalf(`TestNormalized. Actual x: %v. Expected x: %v`, actual_x, expectedXComp)
	}

	if actual_y != expectedYComp {
		t.Fatalf(`TestNormalized. Actual y: %v. Expected y: %v`, actual_y, expectedYComp)
	}
}

func TestAdd(t *testing.T) {
	baseVector := MakeVector2D(10,10)
	addingVector := MakeVector2D(20,12)

	resultVector := baseVector.Add(*addingVector)

	if resultVector.x != 30.0 {
		t.Fatalf(`TestAdd() | Expected 30. Actual %v`, resultVector.x)
	}

	if resultVector.y != 22.0 {
		t.Fatalf(`TestAdd() | Expected 22. Actual %v`, resultVector.y)
	}
}

func TestSubtract(t *testing.T) {
	baseVector := MakeVector2D(10,10)
	subtractVector := MakeVector2D(20,32)

	resultVector := baseVector.Subtract(*subtractVector)

	if resultVector.x != -10 {
		t.Fatalf(`TestSubtract() | Expected 30. Actual %v`, resultVector.x)
	}

	if resultVector.y != -22 {
		t.Fatalf(`TestSubtract() | Expected -22. Actual %v`, resultVector.y)
	}
}

func TestMultiply(t *testing.T) {
	vector := Vector2D {
		x: 10,
		y: 10,
	}
	scalar := 5.0

	resultVector := vector.Multiply(scalar)

	if resultVector.x != 50 {
		t.Fatalf(`TestMultiply() | Expected 50. Actual %v`, resultVector.x)
	}

	if resultVector.y != 50 {
		t.Fatalf(`TestMultiply() | Expected 50. Actual %v`, resultVector.y)
	}
}

func TestDivide(t *testing.T) {
	vector := Vector2D {
		x: 10,
		y: 10,
	}

	scalar := 2.0

	resultVector := vector.Divide(scalar)

	if resultVector.x != 5 {
		t.Fatalf(`TestMultiply() | Expected 5. Actual %v`, resultVector.x)
	}

	if resultVector.y != 5 {
		t.Fatalf(`TestMultiply() | Expected 5. Actual %v`, resultVector.y)
	}
}

func TestSet(t *testing.T) {
	vector := MakeVector2D(0,0)
	set_vector := MakeVector2D(5,8)

	vector.Set(*set_vector)

	x, y := vector.GetComponents()

	if x != 5 {
		t.Fatalf(`TestSet(): Expected x: %v. Actual x: %v`, 5, x)
	}

	if y != 8 {
		t.Fatalf(`TestSet(): Expected x: %v. Actual x: %v`, 8, y)
	}
}

func TestSetComponents(t *testing.T) {
	vector := MakeVector2D(0,0)

	vector.SetComponents(4.0, 8.0)

	x, y := vector.GetComponents()

	if x != 4.0 {
		t.Fatalf(`TestSet(): Expected x: %v. Actual x: %v`, 4.0, x)
	}

	if y != 8.0 {
		t.Fatalf(`TestSet(): Expected x: %v. Actual x: %v`, 8.0, y)
	}
}