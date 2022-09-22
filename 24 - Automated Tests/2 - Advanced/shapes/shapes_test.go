package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rect", func(t *testing.T) {
		rect := Rect{10, 12}
		expected := float64(120)
		received := rect.Area()

		if expected != received {
			t.Fatalf("Received area %f is different from expected %f",
				received,
				expected)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}

		expected := float64(math.Pi * 100)
		received := circle.Area()

		if expected != received {
			t.Fatalf("Received area %f is different from expected %f",
				received,
				expected)
		}
	})
}
