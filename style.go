package ui

import ()

/*
 * This is a good example of a data trait
 */

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

func (self *Padding) SetPadding(p Padding) {
	self.Top = p.Top
	self.Bottom = p.Bottom
	self.Left = p.Left
	self.Right = p.Right
}

func (self Padding) Padding() Padding {
	return self
}

type Mask struct {
	mask bool
}

func NewMask(m bool) Mask {
	return Mask{m}
}

func (self *Mask) SetMask(b bool) {
	self.mask = b
}

func (self *Mask) Mask() bool {
	return self.mask
}
