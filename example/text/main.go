package main

import (
	"fmt"

	"github.com/vizstra/ui"
	"github.com/vizstra/ui/layout"
	"github.com/vizstra/ui/widget/text"
)

func main() {
	window := ui.NewWindow("", "Text Example", 1350, 25, 600, 500)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{3, 3, 3, 3})
	txt := text.New(fill, "")
	txt.AddMousePositionCB(func(x, y float64) {
		fmt.Println(x, y)
	})
	txt.AddMouseClickCB(func(m ui.MouseButtonState) {
		fmt.Println(m)
	})
	fill.SetChild(txt)
	window.SetChild(fill)
	end := window.Start()
	<-end
}
