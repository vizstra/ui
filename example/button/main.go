package main

import (
	"fmt"
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/layout"
)

func main() {
	window := ui.NewWindow("", "Button Example", 240, 60, 1570, 60)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{10, 10, 10, 10})
	button := ui.NewButton(fill, "", "Click Here!")
	button.AddMousePositionCB(func(x, y float64) {
		fmt.Println(x, y)
	})
	button.AddMouseClickCB(func(m ui.MouseButtonState) {
		fmt.Println(m)
	})
	fill.SetChild(button)
	window.SetChild(fill)
	end := window.Start()
	<-end
}
