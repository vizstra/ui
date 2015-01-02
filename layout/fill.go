package layout

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/vg"
)

type Fill struct {
	parent      ui.Drawer
	child       ui.Drawer
	inside      bool
	insideChild bool
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
		false,
		false,
		ui.NewRectangle(0, 0, 0, 0),
		ui.Margin{0, 0, 0, 0},
		ui.NewMask(true),
		ui.NewKeyDispatch(),
		ui.NewCharDispatch(),
		ui.NewMouseDispatch(),
	}

	f.attachPositionCB(parent)
	f.attachEnterCB(parent)
	f.attachClickCB(parent)
	return f
}

// attachPositionCB will accept mouse position information and
// redirect it to the appropriate listeners; the position
// information may also trigger mouse enter and exit events.
// Please note that because the margin may be set for the Fill
// a second set of checks have to be made to trigger calls for
// the child as the child may not fill the entire space.
func (self *Fill) attachPositionCB(parent ui.Drawer) {
	if p, ok := parent.(ui.MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {

			// Dispatch for fill
			self.DispatchMousePosition(x, y)

			if self.Rectangle.Contains(x, y) && !self.inside {
				self.inside = true
				self.DispatchMouseEnter(self.inside)

			} else if self.inside {
				self.inside = false
				self.DispatchMouseEnter(self.inside)
			}

			// Dispatch for child
			if self.child.Contains(x, y) {
				if c, ok := self.child.(ui.MousePositionDispatcher); ok {
					c.DispatchMousePosition(x, y)
				}

				if c, ok := self.child.(ui.MouseEnterDispatcher); ok {
					if !self.insideChild {
						self.insideChild = true
						c.DispatchMouseEnter(self.insideChild)
					}
				}
			} else {
				if c, ok := self.child.(ui.MouseEnterDispatcher); ok {
					self.insideChild = false
					c.DispatchMouseEnter(self.insideChild)
				}
			}
		})
	}
}

func (self *Fill) attachEnterCB(parent ui.Drawer) {
	if p, ok := parent.(ui.MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			if !in && self.inside {
				self.inside = in
				self.DispatchMouseEnter(self.inside)
			}
		})
	}
}

func (self *Fill) attachClickCB(parent ui.Drawer) {
	if p, ok := parent.(ui.MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m ui.MouseButtonState) {
			if self.inside {
				self.DispatchMouseClick(m)
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
	self.Rectangle = ui.NewRectangle(x, y, w, h)
	if self.child == nil {
		return
	}
	m := self.Margin
	self.child.Draw(x+m.Left, y+m.Top, w-(m.Left+m.Right), h-(m.Top+m.Bottom), ctx)
}
