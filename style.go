package ui

import ()

type Margin struct {
	Top, Bottom, Left, Right float64
}

func (self *Margin) SetMargin(m Margin) {
	self.Top = m.Top
	self.Bottom = m.Bottom
	self.Left = m.Left
	self.Right = m.Right
}

func (self Margin) Margin() Margin {
	return self
}

type Padding struct {
	Top, Bottom, Left, Right float64
}

type Mask struct {
	mask bool
}

func (self *Mask) SetMask(b bool) {
	self.mask = b
}

func (self *Mask) Mask() bool {
	return self.mask
}
