package ui

import (
	"github.com/vizstra/vg"
)

type Fill struct {
	parent Drawer
	child  Drawer
	Margin
	Mask
}

func NewFill(parent Drawer) *Fill {
	return &Fill{parent, nil, Margin{0, 0, 0, 0}, Mask{true}}
}

func (self *Fill) SetChild(child Drawer) {
	self.child = child
}

func (self *Fill) Child() Drawer {
	return self.child
}

func (self *Fill) Draw(x, y, w, h float64, ctx vg.Context) {
	if self.child == nil {
		return
	}
	m := self.Margin
	self.child.Draw(x+m.Left, y+m.Top, w-(m.Left+m.Right), h-(m.Top+m.Bottom), ctx)
}
