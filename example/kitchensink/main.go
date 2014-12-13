package main

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/chart"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/ui/layout"
	"github.com/vizstra/ui/widget"
	"math/rand"
	"time"
)

func main() {
	window := ui.NewWindow("", "Vizstra Kitchen Sink", 1570, 60, 450, 450)

	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{3, 3, 3, 3})
	table := layout.NewTable(fill)
	fill.SetChild(table)

	table.SetDefaultCellDimensions(50, 50)
	table.SetCellMargin(ui.Margin{2, 2, 2, 2})

	b1 := widget.NewImageButton(table, "none", "Click Here!")
	b1.HoverBackground = color.RGBA(10, 10, 10, 50)
	b1.SetImagePath("src/github.com/vizstra/ui/res/img/b.png")
	b1.SetHoverImagePath("src/github.com/vizstra/ui/res/img/a.png")
	table.AddMultiCell(b1, 0, 0, 2, 1)

	b2 := widget.NewImageButton(table, "none", "Click Here!")
	b2.HoverBackground = color.RGBA(10, 10, 10, 50)
	b2.SetImagePath("src/github.com/vizstra/ui/res/img/candy-apple-icon.png")
	table.AddMultiCell(b2, 2, 0, 2, 1)

	button := widget.NewButton(table, "none", "Normal Button")
	table.AddMultiCell(button, 4, 0, 3, 1)

	text := widget.NewText(table, "")
	table.AddMultiCell(text, 4, 1, 4, 7)

	table2 := layout.NewTable(table)
	bg := color.Palette[color.Orange1]
	table2.Background = &bg
	table2.SetDefaultCellDimensions(30, 30)
	table.AddMultiCell(table2, 0, 1, 0, 0)

	activity := widget.NewActivityBar(table, "", 100.0, []float64{})
	go func() {
		for {
			activity.Data = append(activity.Data, rand.Float64()*100)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	activity.Foreground = color.Palette[color.Purple2]
	table.AddMultiCell(activity, 0, 5, 8, 1)

	pb := widget.NewProgressBar(table2, "", 100)
	pb.HoverBackground = color.RGBA(10, 10, 10, 50)
	pb.Value = 70
	table2.AddMultiCell(pb, 0, 1, 4, 1)

	table.AddMultiCell(BuildChart(table), 0, 3, 4, 2)
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
