package chart

import (
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"math"
)

type LineChartModeler interface {
	Title() string
	Series() []Series
	Limits() (min, max, count float64)
}

type LineChartModel struct {
	ChartTitle string
	Data       []Series
}

func (self *LineChartModel) Title() string {
	return self.ChartTitle
}

func (self *LineChartModel) Series() []Series {
	return self.Data
}

func (self *LineChartModel) Limits() (min, max, count float64) {
	min = math.MaxFloat64
	max = -min

	for _, s := range self.Data {
		a, b := s.MinMax()
		if a < min {
			min = a
		}

		if b > max {
			max = b
		}

		l := float64(len(s.Data))
		if count < l {
			count = l
		}
	}
	return min, max, count
}

type Series struct {
	Data        []float64
	Color       Color
	StrokeWidth float64
}

// MinMax returns the minimum and maximum values from the series data.
func (self *Series) MinMax() (min, max float64) {
	min = math.MaxFloat64
	max = -min
	for _, f := range self.Data {
		if f > max {
			max = f
		}

		if f < min {
			min = f
		}
	}
	return
}

type LineChart struct {
	ui.Element
	Model        LineChartModeler
	Title        Title
	Background   Color
	displayColor Color
	Fill         bool
}

func NewLineChart(parent ui.Drawer, name string, mdl LineChartModeler) *LineChart {
	self := &LineChart{
		ui.NewElement(parent, name),
		mdl,
		Title{Model: mdl, FontSize: 17},
		Palette[CHART_BACKGROUND],
		Palette[CHART_BACKGROUND],
		true,
	}

	return self
}

func (self *LineChart) Draw(x, y, w, h float64, ctx vg.Context) {
	ui.DrawDefaultElement(x, y, w, h, Palette[CHART_BACKGROUND], ctx)
	// th := self.Title.draw(x, y, w, h, ctx)
	ctx.BeginPath()
	ctx.Fill()
	self.drawSeries(x, y, w, h, ctx)
	ctx.ResetScissor()
}

func (self *LineChart) drawSeries(x, y, w, h float64, ctx vg.Context) {
	// Boundries
	baseline := y + h
	_, max, count := self.Model.Limits()
	horIncr := w / (count - 1)
	vratio := (h - 1) / max
	for _, s := range self.Model.Series() {
		c1 := CloneColor(s.Color)
		c1.A = 150
		ctx.StrokeColor(c1)
		ctx.StrokeWidth(s.StrokeWidth)
		ctx.BeginPath()
		for i, f := range s.Data {
			if i == 0 {
				ctx.MoveTo(x+float64(i)*horIncr, baseline-(f*vratio))
			} else {
				ctx.LineTo(x+float64(i)*horIncr, baseline-(f*vratio))
			}
		}
		ctx.Stroke()

		if self.Fill {
			ctx.BeginPath()
			// ctx.MoveTo(x, baseline)
			c2 := CloneColor(s.Color)
			c2.A = 200
			ctx.FillColor(c2)
			for i, f := range s.Data {
				if i == 0 {
					ctx.MoveTo(x+float64(i)*horIncr, baseline-(f*vratio))
				}

				ctx.LineTo(x+float64(i)*horIncr, baseline-(f*vratio))

				if i == len(s.Data)-1 {
					ctx.LineTo(x+float64(i)*horIncr, baseline)
					ctx.LineTo(x, baseline)
					ctx.MoveTo(x, baseline-(s.Data[0]*vratio))
					break
				}
			}
			ctx.Fill()
		}
	}
}
