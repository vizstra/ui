package ui

type Rectangular interface {
	Contains(x, y float64) bool
	ContainsPosition(p Position) bool
	SetSize(s Size)
	SetPosition(p Position)
}

// Rectangle is a quadrilateral defined, in this type,
// by its upper left position, width and height.
type Rectangle struct {
	Position
	Size
}

func NewRectangle(x, y, w, h float64) Rectangle {
	return Rectangle{Position{x, y}, Size{w, h}}
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

// ContainsPosition returns true if the provided Position
// is inside of this rectangle.
func (self Rectangle) ContainsPosition(p Position) bool {
	return self.Contains(p.X, p.Y)
}
