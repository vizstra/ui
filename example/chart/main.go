package main

import (
	"fmt"
	"time"

	"github.com/vizstra/ui"
	"github.com/vizstra/ui/chart"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/ui/layout"
)

func main() {
	window := ui.NewWindow("", "Chart Example", 500, 300, 1570, 60)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{15, 15, 15, 15})
	s := make([]chart.Series, 3)
	s[0] = chart.Series{[]float64{1, 10, 3, 2, 6, 60, 30, 25, 26, 100, 90, 40, 600, 700, 800, 90}, color.Color{200, 100, 100, 255}, 1}
	s[1] = chart.Series{[]float64{10, 3, 2, 6, 60, 30, 25, 3, 100, 200, 90, 40, 60, 70, 80, 900}, color.Color{100, 200, 100, 255}, 1}
	s[2] = chart.Series{[]float64{10, 3, 2, 6, 60, 30, 25, 26, 100, 300, 90, 40, 60, 70, 80, 40}, color.Color{100, 100, 200, 255}, 1}
	c := chart.NewLineChart(fill, "", &chart.LineChartModel{"Example Line Chart", s})
	c.AddMousePositionCB(func(x, y float64) {
		fmt.Println(x, y)
	})
	c.AddMouseClickCB(func(m ui.MouseButtonState) {
		fmt.Println(m)
	})
	fill.SetChild(c)
	window.SetChild(fill)
	go PTime()
	end := window.Start()
	<-end
}

func PTime() {
	t := time.Now()
	for {
		fmt.Println(time.Since(t))
		time.Sleep(10 * time.Second)
	}
}
