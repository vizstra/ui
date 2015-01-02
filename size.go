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
