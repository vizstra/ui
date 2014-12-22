package ui

import ()

// Sizer is a generic interface to types which
// represent the rectangular size of something.
type Sizer interface {

	// Size returns a copy as a Size structure.
	Size() Size

	// SetSize set the size from the given Size.
	SetSize(Size)
}

// Size is a type the represents a width and height;
// it is a implements the Sizer interface.
type Size struct {
	W, H float64
}

// Size returns a copy of this Size.
func (self Size) Size() Size {
	return self
}

// SetSize sets the size of the Size type.
func (self *Size) SetSize(s Size) {
	self.W = s.W
	self.H = s.H
}

// Positioner is a generic interface to types which
// represent the 2d location of something.
type Positioner interface {

	// Position returns a copy as a Point.
	Position() Point

	// SetPosition sets the position from the given Point.
	SetPosition(Point)
}

// Point is a type the represents a 2D coordinate;
// it implements the Positioner interface.
type Point struct {
	X, Y float64
}

// Position returns a copy as a Point.
func (self Point) Position() Point {
	return self
}

// SetPosition sets the position from the given Point.
func (self *Point) SetPosition(p Point) {
	self.X, self.Y = p.X, p.Y
}

// Rectangle is a quadrilateral defined, in this type,
// by its upper left point, width and height.
type Rectangle struct {
	X, Y, W, H float64
}

// Contains returns true if the cooridinates provided
// are inside of this rectangle.
func (self Rectangle) Contains(x, y float64) bool {
	if x > self.X && x < self.X+self.W &&
		y > self.Y && y < self.Y+self.H {
		return true
	}
	return false
}

// ContainsPoint returns true if the provided Point
// is inside of this rectangle.
func (self Rectangle) ContainsPoint(p Point) bool {
	return self.Contains(p.X, p.Y)
}

type View struct {
}
