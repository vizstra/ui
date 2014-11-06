package chart

import (
	"github.com/vizstra/vg"
	"image/color"
)

type Title struct {
	Model    LineChartModeler
	Font     string
	FontSize float64
}

func (self *Title) Draw(x, y, w, h float64, ctx vg.Context) {
	self.draw(x, y, w, h, ctx)
}

// draw draws the title and returns the height.
func (self *Title) draw(x, y, w, h float64, ctx vg.Context) float64 {
	th := 0.0
	name := self.Model.Title()
	if len(name) > 0 {
		th = self.FontSize + 16
		// ctx.BeginPath()
		// ctx.StrokeColor(color.RGBA{0, 50, 100, 25})
		// ctx.StrokeWidth(1)
		// ctx.LineCap(vg.ROUND)
		// ctx.LineJoin(vg.ROUND)
		// ctx.MoveTo(x, y+th)
		// ctx.LineTo(x+1.5, y+th)
		// ctx.LineTo(x+w+1.5, y+th)
		// ctx.Stroke()

		ctx.FillColor(color.RGBA{255, 251, 251, 255})
		ctx.TextAlign(vg.ALIGN_LEFT | vg.ALIGN_MIDDLE)
		ctx.FontSize(self.FontSize + 3)
		ctx.FindFont(vg.FONT_DEFAULT)
		ctx.WrappedText(x+10, y+self.FontSize, w, name)

	}
	return th
}
