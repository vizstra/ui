package ui

import ()

type Fill struct {
	child View
}

func NewFill() *Fill {
	return &Fill{nil}
}

func (self *Fill) SetChild(child View) {
	self.child = child
}

func (self *Fill) Child() View {
	return self.child
}

func (self *Fill) Draw(x, y, w, h float64, ctx Context) {
	if self.child == nil {
		return
	}

	self.child.Draw(x, y, w, h, ctx)
}
