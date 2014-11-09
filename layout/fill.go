package layout

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/vg"
)

type Fill struct {
	parent ui.Drawer
	child  ui.Drawer
	ui.Rectangle
	ui.Margin
	ui.Mask
	ui.KeyDispatch
	ui.CharDispatch
	ui.MouseDispatch
}

func NewFill(parent ui.Drawer) *Fill {
	f := &Fill{
		parent,
		nil,
		ui.Rectangle{0, 0, 0, 0},
		ui.Margin{0, 0, 0, 0},
		ui.NewMask(true),
		ui.NewKeyDispatch(),
		ui.NewCharDispatch(),
		ui.NewMouseDispatch(),
	}
	f.attach(parent)
	return f
}

func (f *Fill) attach(parent ui.Drawer) {
	inside := false
	if p, ok := parent.(ui.MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			if x > f.Rectangle.X+f.Margin.Left &&
				y > f.Rectangle.Y+f.Margin.Top &&
				x < f.Rectangle.X+f.Rectangle.W-f.Margin.Right &&
				y < f.Rectangle.Y+f.Rectangle.H-f.Margin.Bottom {
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

	if p, ok := parent.(ui.MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			if !in {
				inside = in
				f.DispatchMouseEnter(inside)
			}
		})
	}

	if p, ok := parent.(ui.MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m ui.MouseButtonState) {
			if inside {
				f.DispatchMouseClick(m)
			}
		})
	}
}

func (self *Fill) SetChild(child ui.Drawer) {
	self.child = child
	self.AddDrawerKeyHandler(child)
	self.AddDrawerCharHandler(child)
	self.AddDrawerMouseClickHandler(child)
	self.AddDrawerMousePositionHandler(child)
	self.AddDrawerMouseEnterHandler(child)
}

func (self *Fill) Child() ui.Drawer {
	return self.child
}

func (self *Fill) Draw(x, y, w, h float64, ctx vg.Context) {
	self.Rectangle = ui.Rectangle{x, y, w, h}
	if self.child == nil {
		return
	}
	m := self.Margin
	self.child.Draw(x+m.Left, y+m.Top, w-(m.Left+m.Right), h-(m.Top+m.Bottom), ctx)
}
