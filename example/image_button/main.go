package main

import (
	"fmt"
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/chart"
	"github.com/vizstra/ui/layout"
	"image/color"
)

func BuildTable(parent ui.Drawer) *chart.LineChart {
	s := make([]chart.Series, 3)
	s[0] = chart.Series{[]float64{1, 10, 3, 2, 6, 60, 30, 25, 26, 100, 90, 40, 600, 700, 800, 90}, color.RGBA{200, 100, 100, 255}, 1}
	s[1] = chart.Series{[]float64{10, 3, 2, 6, 60, 30, 25, 3, 100, 200, 90, 40, 60, 70, 80, 900}, color.RGBA{100, 200, 100, 255}, 1}
	s[2] = chart.Series{[]float64{10, 3, 2, 6, 60, 30, 25, 26, 100, 300, 90, 40, 60, 70, 80, 40}, color.RGBA{100, 100, 200, 255}, 1}
	c := chart.NewLineChart(parent, "", &chart.LineChartModel{"Example Line Chart", s})
	c.AddMousePositionCB(func(x, y float64) {
		fmt.Println(x, y)
	})
	c.AddMouseClickCB(func(m ui.MouseButtonState) {
		fmt.Println(m)
	})
	return c
}

func main() {

	window := ui.NewWindow("", "Image Button Example", 300, 300, 1570, 60)

	table := layout.NewTable(window)
	table.SetDefaultCellDimensions(75, 75)
	table.SetCellMargin(ui.Margin{1, 1, 1, 1})

	b1 := ui.NewImageButton(table, "", "Click Here!")
	b1.HoverBackground = color.RGBA{10, 10, 10, 50}
	b1.SetImagePath("src/github.com/vizstra/ui/res/img/b.png")
	b1.SetHoverImagePath("src/github.com/vizstra/ui/res/img/a.png")

	b2 := ui.NewImageButton(table, "", "Click Here!")
	b2.HoverBackground = color.RGBA{10, 10, 10, 50}
	b2.SetImagePath("src/github.com/vizstra/ui/res/img/candy-apple-icon.png")

	button := ui.NewButton(table, "", "Click Here!")

	table.AddCell(b1, 0, 0)
	table.AddCell(b2, 1, 0)
	table.AddCell(b1, 2, 0)
	table.AddCell(button, 3, 0)
	table.AddMultiCell(BuildTable(table), 0, 1, 4, 3)
	window.SetChild(table)
	end := window.Start()
	<-end
}
