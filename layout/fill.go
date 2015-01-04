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

	f.attachCB(parent)
	return f
}

// attachPositionCB will accept mouse position information and
// redirect it to the appropriate listeners; the position
// information may also trigger mouse enter and exit events.
// Please note that because the margin may be set for the Fill
// a second set of checks have to be made to trigger calls for
// the child as the child may not fill the entire space.
func (self *Fill) attachCB(parent ui.Drawer) {
	var mx, my float64
	if p, ok := parent.(ui.MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			mx, my = x, y

			// Dispatch for fill
			self.DispatchMousePosition(x, y)

			inside := self.Rectangle.Contains(x, y)
			if inside != self.inside {
				self.inside = inside
				self.DispatchMouseEnter(self.inside)
			}

			// Dispatch for child
			inchild := self.child.Contains(x, y)
			if c, ok := self.child.(ui.MousePositionDispatcher); ok && inchild {
				c.DispatchMousePosition(x, y)
			}

			if c, ok := self.child.(ui.MouseEnterDispatcher); ok && self.insideChild != inchild {
				self.insideChild = inchild
				c.DispatchMouseEnter(self.insideChild)
			}
		})
	}

	if p, ok := parent.(ui.MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			if in != self.inside {
				self.inside = in
				self.DispatchMouseEnter(self.inside)
				if c, ok := self.child.(ui.MouseEnterDispatcher); ok {
					if in && self.child.Contains(mx, my) || !in && self.insideChild {
						self.insideChild = in
						c.DispatchMouseEnter(self.inside)
					}
				}
			}
		})
	}

	if p, ok := parent.(ui.MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m ui.MouseButtonState) {
			if self.Contains(mx, my) {
				self.DispatchMouseClick(m)
				if c, ok := self.child.(ui.MouseClickDispatcher); ok && self.child.Contains(mx, my) {
					c.DispatchMouseClick(m)
				}
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

func (self *Fill) Draw(ctx vg.Context) {
	x, y, w, h := self.Bounds()
	if self.child == nil {
		return
	}
	m := self.Margin
	self.child.SetBounds(x+m.Left, y+m.Top, w-(m.Left+m.Right), h-(m.Top+m.Bottom))
	self.child.Draw(ctx)
}
