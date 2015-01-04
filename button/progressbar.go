package button

import (
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

type ProgressBar struct {
	ui.Element
	Value float64
	Max   float64
}

func NewProgressBar(parent ui.Drawer, name string, max float64) *ProgressBar {
	self := &ProgressBar{
		ui.NewElement(parent, name),
		0,
		max,
	}

	self.Element.ClickBackground = Blue4

	self.DrawCB = func(ctx vg.Context) {
		x, y, w, h := self.Bounds()
		fg := ctx.BoxGradient(x, y, w, h/3, h/2, h, self.ClickBackground, self.ClickBackground)
		ctx.BeginPath()
		ctx.RoundedRect(x+1, y+1, (w-2)*(self.Value/self.Max), h-2, self.CornerRadius)
		ctx.FillPaint(fg)
		ctx.Fill()
	}
	return self
}
