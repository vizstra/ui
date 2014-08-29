package main

import (
	"github.com/vizstra/ui"
)

func main() {
	window := ui.NewWindow("", "AA", 320, 600, 1570, 60)
	window.SetTitle("Test App")
	button := ui.NewButton("", "This is a Button")
	fill := ui.NewFill()
	fill.SetChild(button)
	window.SetChild(fill)
	end := window.Start()
	<-end
}
