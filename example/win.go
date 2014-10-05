package main

import (
	"github.com/vizstra/ui"
)

func main() {
	window := ui.NewWindow("", "AA", 320, 600, 1570, 60)
	window.SetTitle("Button Render Test")
	fill := ui.NewFill(window)
	fill.SetMargin(ui.Margin{40, 200, 100, 10})
	button := ui.NewButton(fill, "", "This is a Button")
	fill.SetChild(button)
	window.SetChild(fill)
	end := window.Start()
	<-end
}
