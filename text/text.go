package text

import (
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

type Text struct {
	ui.Element
	text string
}

func New(parent ui.Drawer, name, text string) *Text {
	self := &Text{
		ui.NewElement(parent, name),
		text,
	}
	return self
}

func (self *Text) Draw(x, y, w, h float64, ctx vg.Context) {
	ctx.Scissor(x, y, w, h)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillColor(self.Background)
	ctx.Fill()

	ctx.SetFontPtSize(23)
	_, _, lineh := ctx.TextMetrics()
	_, ly := x, y+lineh
	ctx.FillColor(self.Foreground)

	ctx.WrappedText(x, ly, w, self.text)
	ctx.ResetScissor()
}

func drawBounds(xmin, ymin, xmax, ymax float64, ctx vg.Context) {
	ctx.BeginPath()
	ctx.FillColor(Green1)
	ctx.RoundedRect(xmin, ymin, xmax-xmin, ymax-ymin, 3)
	ctx.Fill()
}
