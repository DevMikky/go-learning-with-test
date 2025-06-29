package struct_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape        Shape
		hasPerimeter float64
	}{
		{Rectangle{10, 20}, 60.0},
		{Circle{5}, 31.41592653589793},
		{Triangle{5, 10, 20}, 35},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
		}
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{5, 10, 15}, 25},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
