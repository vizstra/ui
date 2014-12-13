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
	parent ui.Drawer

	// last drawn location
	x, y       float64
	Background *color.Color

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
}

func NewTable(parent ui.Drawer) *Table {
	table := &Table{
		parent,
		0,
		0,
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
	}

	table.configRouter(parent)
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

func (self *Table) Draw(x, y, w, h float64, ctx vg.Context) {
	self.x, self.y = x, y
	if self.Background != nil {
		ui.DrawDefaultWidget(x, y, w, h, self.Background, ctx)
	}

	for _, child := range self.children {
		r := self.bounds(child)
		child.drawer.Draw(r.X, r.Y, r.W, r.H, ctx)
	}
}

func (self *Table) bounds(child *cell) ui.Rectangle {
	x, y, w, h := self.x, self.y, 0.0, 0.0
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
	return ui.Rectangle{x, y, w, h}
}

// configRouter will forward events from the parent
// through to the attached handlers.
func (self *Table) configRouter(parent ui.Drawer) {
	// var inside bool
	var mx, my float64

	if p, ok := parent.(ui.MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			mx, my = x, y
			for _, cell := range self.children {
				child := cell.drawer
				inchild := self.bounds(cell).Contains(x, y)

				if !inchild {
					if c, ok := child.(ui.MouseEnterDispatcher); ok && cell.mouseInside {
						c.DispatchMouseEnter(false)
						cell.mouseInside = false
					}
					continue
				}

				if c, ok := child.(ui.MouseEnterDispatcher); ok && !cell.mouseInside {
					c.DispatchMouseEnter(true)
					cell.mouseInside = true
				}

				if c, ok := child.(ui.MousePositionDispatcher); ok && cell.mouseInside {
					c.DispatchMousePosition(x, y)
				}
			}
			self.DispatchMousePosition(x, y)
		})
	}

	if p, ok := parent.(ui.MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			for _, cell := range self.children {
				if cell.mouseInside {
					cell.drawer.DispatchMouseEnter(in)
				}
			}
			self.DispatchMouseEnter(in)
		})
	}

	if p, ok := parent.(ui.MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m ui.MouseButtonState) {
			for _, cell := range self.children {
				child := cell.drawer
				if c, ok := child.(ui.MouseClickDispatcher); ok && self.bounds(cell).Contains(mx, my) {
					c.DispatchMouseClick(m)
				}
			}
			self.DispatchMouseClick(m)
		})
	}
}
