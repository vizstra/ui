package ui

import (
	"github.com/vizstra/vg"
)

type Fill struct {
	parent Drawer
	child  Drawer
	Rectangle
	Margin
	Mask
	KeyDispatch
	CharDispatch
	MouseDispatch
}

func NewFill(parent Drawer) *Fill {
	f := &Fill{
		parent,
		nil,
		Rectangle{Point2d{0, 0}, Size2d{0, 0}},
		Margin{0, 0, 0, 0},
		Mask{true},
		NewKeyDispatch(),
		NewCharDispatch(),
		NewMouseDispath(),
	}

	inside := false
	if p, ok := parent.(MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			if x > f.Rectangle.Position.X+f.Margin.Left &&
				y > f.Rectangle.Position.Y+f.Margin.Top &&
				x < f.Rectangle.Position.X+f.Rectangle.Size.W-f.Margin.Right &&
				y < f.Rectangle.Position.Y+f.Rectangle.Size.H-f.Margin.Bottom {
				if !inside {
					inside = true
					f.DispatchMouseEnter(inside)
				}
				f.DispatchMousePosition(x, y)
			} else if inside {
				inside = false
				f.DispatchMouseEnter(inside)
			}
		})
	}

	if p, ok := parent.(MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			if !in {
				inside = in
				f.DispatchMouseEnter(inside)
			}
		})
	}

	if p, ok := parent.(MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m MouseButtonState) {
			if inside {
				f.DispatchMouseClick(m)
			}
		})
	}

	return f
}

func (self *Fill) SetChild(child Drawer) {
	self.child = child
	self.AddDrawerKeyHandler(child)
	self.AddDrawerCharHandler(child)
	self.AddDrawerMouseClickHandler(child)
	self.AddDrawerMousePositionHandler(child)
	self.AddDrawerMouseEnterHandler(child)
}

func (self *Fill) Child() Drawer {
	return self.child
}

func (self *Fill) Draw(x, y, w, h float64, ctx vg.Context) {
	self.Rectangle = Rectangle{Point2d{x, y}, Size2d{w, h}}
	if self.child == nil {
		return
	}
	m := self.Margin
	self.child.Draw(x+m.Left, y+m.Top, w-(m.Left+m.Right), h-(m.Top+m.Bottom), ctx)
}
