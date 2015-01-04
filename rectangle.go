package ui

type Rectangular interface {
	Bounds() (x, y, w, h float64)
	Contains(x, y float64) bool
	ContainsPosition(p Position) bool
	SetBounds(x, y, w, h float64)
	SetPosition(p Position)
	SetRectangle(r Rectangle)
	SetSize(s Size)
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

func (self *Rectangle) SetRectangle(r Rectangle) {
	self.X = r.X
	self.Y = r.Y
	self.W = r.W
	self.H = r.H
}

func (self *Rectangle) SetBounds(x, y, w, h float64) {
	self.X = x
	self.Y = y
	self.W = w
	self.H = h
}

func (self Rectangle) Bounds() (x, y, w, h float64) {
	return self.X, self.Y, self.W, self.H
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
