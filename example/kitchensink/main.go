package main

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/chart"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/ui/layout"
	"github.com/vizstra/ui/widget"
)

func main() {
	window := ui.NewWindow("", "Vizstra Kitchen Sink", 1570, 60, 350, 350)

	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{10, 10, 10, 10})
	table := layout.NewTable(fill)
	fill.SetChild(table)

	table.SetDefaultCellDimensions(75, 75)
	table.SetCellMargin(ui.Margin{3, 3, 3, 3})

	b1 := widget.NewImageButton(table, "", "Click Here!")
	b1.HoverBackground = color.RGBA(10, 10, 10, 50)
	b1.SetImagePath("src/github.com/vizstra/ui/res/img/b.png")
	b1.SetHoverImagePath("src/github.com/vizstra/ui/res/img/a.png")
	table.AddCell(b1, 0, 0)

	b2 := widget.NewImageButton(table, "", "Click Here!")
	b2.HoverBackground = color.RGBA(10, 10, 10, 50)
	b2.SetImagePath("src/github.com/vizstra/ui/res/img/candy-apple-icon.png")
	table.AddCell(b2, 1, 0)

	button := widget.NewButton(table, "", "Normal Button")
	table.AddMultiCell(button, 2, 0, 2, 1)

	table2 := layout.NewTable(table)
	table2.SetDefaultCellDimensions(75, 30)
	table2.SetCellMargin(ui.Margin{1, 1, 1, 1})
	table.AddMultiCell(table2, 0, 1, 4, 1)

	pb := widget.NewProgressBar(table2, "", 100)
	pb.HoverBackground = color.RGBA(10, 10, 10, 50)
	pb.Value = 70
	table2.AddMultiCell(pb, 0, 1, 4, 1)

	table.AddMultiCell(BuildChart(table), 0, 2, 4, 2)
	window.SetChild(fill)
	end := window.Start()
	<-end
}

func BuildChart(parent ui.Drawer) *chart.LineChart {
	s := make([]chart.Series, 3)
	s[0] = chart.Series{[]float64{1, 10, 3, 2, 6, 60, 30, 25, 26, 100, 90, 40, 600, 700, 800, 90}, color.RGBA(200, 100, 100, 255), 1}
	s[1] = chart.Series{[]float64{10, 3, 2, 6, 60, 30, 25, 3, 100, 200, 90, 40, 60, 70, 80, 900}, color.RGBA(100, 200, 100, 255), 1}
	s[2] = chart.Series{[]float64{10, 3, 2, 6, 60, 30, 25, 26, 100, 300, 90, 40, 60, 70, 80, 40}, color.RGBA(100, 100, 200, 255), 1}
	return chart.NewLineChart(parent, "", &chart.LineChartModel{"Example Line Chart", s})
}
