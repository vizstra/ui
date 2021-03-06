package layout

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

type cell struct {
	drawer         ui.Drawer
	row, col, w, h int
	mouseInside    bool
}

const (
	DEFAULT_TABLE_CELL_WIDTH  = 100
	DEFAULT_TABLE_CELL_HEIGHT = 20
)

// Table is a simple layout that organizes child views by way of a
// classic cell system.  The Table differs grom the Grid in several
// ways; most notably Table overflows the view space while a Grid
// does not.
type Table struct {
	ui.Rectangle
	parent            ui.Drawer
	Background        *color.Color
	children          []*cell
	cellHeights       []float64
	cellWidths        []float64
	defaultCellHeight float64
	defaultCellWidth  float64
	cellMargin        ui.Margin
	cellPadding       ui.Padding
	ui.KeyDispatch
	ui.CharDispatch
	ui.MouseDispatch
	ui.ScrollDispatch
}

// NewTable builds a default Table.
func NewTable(parent ui.Drawer) *Table {
	table := &Table{
		ui.NewRectangle(0, 0, 0, 0),
		parent,
		nil,
		make([]*cell, 0),
		make([]float64, 0),
		make([]float64, 0),
		DEFAULT_TABLE_CELL_HEIGHT,
		DEFAULT_TABLE_CELL_WIDTH,
		ui.Margin{0, 0, 0, 0},
		ui.Padding{0, 0, 0, 0},
		ui.NewKeyDispatch(),
		ui.NewCharDispatch(),
		ui.NewMouseDispatch(),
		ui.NewScrollDispatch(),
	}

	table.route(parent)
	return table
}

func (self *Table) SetCellMargin(m ui.Margin) {
	self.cellMargin = m
}

func (self *Table) SetCellPadding(p ui.Padding) {
	self.cellPadding = p
}

func (self *Table) SetDefaultCellDimensions(w, h float64) {
	self.defaultCellHeight = h
	self.defaultCellWidth = w
}

func (self *Table) SetRowHeight(row int, h float64) {
	for row >= len(self.cellHeights) {
		self.cellHeights = append(self.cellHeights, self.defaultCellHeight)
	}
	self.cellHeights[row] = h
}

func (self *Table) SetColWidth(col int, w float64) {
	for col >= len(self.cellWidths) {
		self.cellWidths = append(self.cellWidths, self.defaultCellWidth)
	}
	self.cellWidths[col] = w
}

func (self *Table) AddCell(child ui.Drawer, col, row int) error {
	return self.AddMultiCell(child, col, row, 1, 1)
}

func (self *Table) AddMultiCell(child ui.Drawer, col, row, w, h int) error {
	gvc := &cell{child, row, col, w, h, false}

	self.children = append(self.children, gvc)

	for i := len(self.children); i < row; i++ {
		self.children = append(self.children, gvc)
	}

	for row+h-1 >= len(self.cellHeights) {
		self.cellHeights = append(self.cellHeights, self.defaultCellHeight)
	}

	for col+w-1 >= len(self.cellWidths) {
		self.cellWidths = append(self.cellWidths, self.defaultCellWidth)
	}
	return nil
}

func (self *Table) Draw(ctx vg.Context) {
	for _, child := range self.children {
		r := self.bounds(child)
		child.drawer.SetRectangle(r)
		child.drawer.Draw(ctx)
	}
}

func (self *Table) bounds(child *cell) ui.Rectangle {
	x, y, w, h := self.X, self.Y, 0.0, 0.0
	for i := 0; i < child.col+child.w && i < len(self.cellWidths); i++ {
		if i < child.col {
			x += self.cellPadding.Left
			x += self.cellMargin.Left
			x += self.cellWidths[i]
		} else {
			w += self.cellPadding.Right

			if i == child.col+child.w-1 {
				w -= self.cellMargin.Right
			} else {
				w += self.cellMargin.Left
			}
			w += self.cellWidths[i]
		}
	}

	for i := 0; i < child.row+child.h && i < len(self.cellHeights); i++ {
		if i < child.row {
			y += self.cellPadding.Top
			y += self.cellMargin.Top
			y += self.cellHeights[i]
		} else {
			h += self.cellPadding.Bottom
			if i == child.row+child.h-1 {
				h -= self.cellMargin.Bottom
			} else {
				h += self.cellMargin.Top
			}
			h += self.cellHeights[i]
		}
	}
	return ui.NewRectangle(x, y, w, h)
}

// route will forward events from the parent
// through to the attached handlers.
func (self *Table) route(parent ui.Drawer) {
	// var inside bool
	var mx, my float64

	if p, ok := parent.(ui.MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			mx, my = x, y

			for _, cell := range self.children {
				child := cell.drawer
				inchild := child.Contains(x, y)

				if c, ok := child.(ui.MousePositionDispatcher); ok && inchild {
					c.DispatchMousePosition(x, y)
				}

				if c, ok := child.(ui.MouseEnterDispatcher); ok && cell.mouseInside != inchild {
					cell.mouseInside = inchild
					c.DispatchMouseEnter(inchild)
				}
			}
			self.DispatchMousePosition(x, y)
		})
	}

	if p, ok := parent.(ui.MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			for _, cell := range self.children {
				self.DispatchMouseEnter(in)
				if c, ok := cell.drawer.(ui.MouseEnterDispatcher); ok {
					inside := cell.drawer.Contains(mx, my)
					if inside != cell.mouseInside {
						cell.mouseInside = inside
						c.DispatchMouseEnter(inside)
					}
				}
			}
		})
	}

	if p, ok := parent.(ui.MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m ui.MouseButtonState) {
			self.DispatchMouseClick(m)
			for _, cell := range self.children {
				inchild := cell.drawer.Contains(mx, my)
				if c, ok := cell.drawer.(ui.MouseClickDispatcher); ok && inchild {
					c.DispatchMouseClick(m)
				}
			}
		})
	}

	if p, ok := parent.(ui.ScrollDispatcher); ok {
		p.AddScrollCB(func(xoff, yoff float64) {
			self.DispatchScroll(xoff, yoff)
			for _, cell := range self.children {
				if c, ok := cell.drawer.(ui.ScrollDispatcher); ok && cell.mouseInside {
					c.DispatchScroll(xoff, yoff)
				}
			}
		})
	}
}
