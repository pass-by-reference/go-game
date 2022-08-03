package utils

import (
	"testing"
)

func TestGetOpposite(t *testing.T) {
	dir := UP

	expected := GetOpposite(dir)

	if expected != DOWN {
		t.Fatalf(`TestGetOpposite | Expected: DOWN. Actual: %v`, expected)
	}
}