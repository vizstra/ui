package main

import (
	// "fmt"
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/layout"
	"github.com/vizstra/ui/widget"
)

func main() {
	window := ui.NewWindow("", "Calculator", 1440, 0, 475, 475)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{30, 30, 30, 30})

	table := layout.NewTable(fill)
	table.SetDefaultCellDimensions(50, 50)
	table.SetCellMargin(ui.Margin{2, 2, 2, 2})

	fill.SetChild(table)
	window.SetChild(fill)

	makebutton := func(text string, x, y, w, h int) *widget.Button {
		button := widget.NewButton(table, text, text)
		table.AddMultiCell(button, x, y, w, h)
		return button
	}

	makebutton("7", 0, 2, 1, 1)
	makebutton("4", 0, 3, 1, 1)
	makebutton("1", 0, 4, 1, 1)

	makebutton("8", 1, 2, 1, 1)
	makebutton("5", 1, 3, 1, 1)
	makebutton("2", 1, 4, 1, 1)

	makebutton("9", 2, 2, 1, 1)
	makebutton("6", 2, 3, 1, 1)
	makebutton("3", 2, 4, 1, 1)

	makebutton("Ce", 3, 2, 1, 1)
	makebutton("÷", 3, 3, 1, 1)
	makebutton("*", 3, 4, 1, 1)

	makebutton("-", 0, 1, 1, 1)
	makebutton("+", 1, 1, 1, 1)
	makebutton("=", 2, 1, 2, 1)

	end := window.Start()
	<-end
}

func addTwoNumber(a, b int) int {
	return a + b
}
