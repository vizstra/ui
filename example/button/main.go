package main

import (
	"fmt"
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/button"
	"github.com/vizstra/ui/layout"
)

func main() {
	window := ui.NewWindow("", "Button Example", 240, 60, 1570, 60)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{10, 10, 10, 10})
	b := button.NewButton(fill, "", "Click Here!")
	b.AddMousePositionCB(func(x, y float64) {
		fmt.Println(x, y)
	})
	b.AddMouseClickCB(func(m ui.MouseButtonState) {
		fmt.Println(m)
	})
	fill.SetChild(b)
	window.SetChild(fill)
	end := window.Start()
	<-end
}
