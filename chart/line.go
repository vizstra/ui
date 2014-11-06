package chart

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/vg"
	"image/color"
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
	Color       color.Color
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
	Parent       ui.Drawer
	Name         string
	Model        LineChartModeler
	Title        Title
	Radius       float64
	Background   color.Color
	displayColor color.Color
	Fill         bool
	ui.MouseDispatch
}

func NewLineChart(parent ui.Drawer, name string, mdl LineChartModeler) *LineChart {
	self := &LineChart{
		parent,
		name,
		mdl,
		Title{Model: mdl, FontSize: 17},
		3,
		ui.Colors[ui.COLOR_DATA_BACKGROUND],
		ui.Colors[ui.COLOR_DATA_BACKGROUND],
		true,
		ui.NewMouseDispatch(),
	}

	if p, ok := parent.(ui.MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			self.DispatchMousePosition(x, y)
		})
	}

	if p, ok := parent.(ui.MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			self.DispatchMouseEnter(in)
		})
	}

	if p, ok := parent.(ui.MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m ui.MouseButtonState) {
			if m.MouseButton == ui.MOUSE_BUTTON_LEFT || m.MouseButton == ui.MOUSE_BUTTON_1 {
				if m.Action == ui.PRESS {

				}
			}
			self.DispatchMouseClick(m)
		})
	}
	return self
}

func (self *LineChart) Draw(x, y, w, h float64, ctx vg.Context) {
	ui.DrawDefaultWidget(x, y, w, h, ctx)
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
		c1 := ui.CloneColor(s.Color)
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
			c2 := ui.CloneColor(s.Color)
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
