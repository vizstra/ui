package widget

import (
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

type ProgressBar struct {
	Widget
	Value float64
	Max   float64
}

func NewProgressBar(parent ui.Drawer, name string, max float64) *ProgressBar {
	self := &ProgressBar{
		NewWidget(parent, name),
		0,
		max,
	}

	self.Widget.ClickBackground = Palette[Blue4]

	self.DrawCB = func(x, y, w, h float64, ctx vg.Context) {
		fg := ctx.BoxGradient(x, y, w, h/3, h/2, h, self.ClickBackground, self.ClickBackground)
		ctx.BeginPath()
		ctx.RoundedRect(x+1, y+1, (w-2)*(self.Value/self.Max), h-2, self.CornerRadius)
		ctx.FillPaint(fg)
		ctx.Fill()
	}
	return self
}
