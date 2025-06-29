package struct_methods_interfaces

import "math"

// Shape interface for different shapes methods and structs
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle shape Structs and methods
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {

	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {

	return 2 * (r.Width + r.Height)
}

// Circle shape Structs and methods
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {

	return 2 * math.Pi * c.Radius
}

// Triangle shape Structs and methods
type Triangle struct {
	Base, Height, Side float64
}

func (t Triangle) Area() float64 {

	return (t.Base * t.Height) * 0.5
}

func (t Triangle) Perimeter() float64 {

	return t.Base + t.Side + t.Height
}
