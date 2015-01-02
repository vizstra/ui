package ui

// Positioner is a generic interface to types which
// represent the 2d location of something.
type Positioner interface {

	// Position returns a copy as a Point.
	Position() Position

	// SetPosition sets the position from the given Point.
	SetPosition(Position)
}

// Position is a type the represents a 2D coordinate;
// it implements the Positioner interface.
type Position struct {
	X, Y float64
}

// Position returns a copy as a Position.
func (self Position) Position() Position {
	return self
}

// SetPosition sets the position from the given Position.
func (self *Position) SetPosition(p Position) {
	self.X, self.Y = p.X, p.Y
}
