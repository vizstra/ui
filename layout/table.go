package layout

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/vg"
)

type cell struct {
	drawer         ui.Drawer
	row, col, w, h int
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
	parent            ui.Drawer
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
	g := &Table{
		parent,
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
	return g
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
	gvc := &cell{child, row, col, w, h}

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
	for _, child := range self.children {
		x, y, w, h := self.bounds(child)
		child.drawer.Draw(x, y, w, h, ctx)
	}
}

func (self *Table) bounds(child *cell) (x, y, w, h float64) {

	for i := 0; i < child.col+child.w && i < len(self.cellWidths); i++ {
		if i < child.col {
			x += self.cellPadding.Left
			x += self.cellMargin.Left
			x += self.cellWidths[i]
		} else {
			w += self.cellPadding.Right
			w -= self.cellMargin.Right
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
			h -= self.cellMargin.Bottom
			h += self.cellHeights[i]
		}
	}
	return x, y, w, h
}
