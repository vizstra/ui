package text

import (
	// "fmt"
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	. "github.com/vizstra/ui/widget"
	"github.com/vizstra/vg"
)

type Text struct {
	Widget
	text string
}

func New(parent ui.Drawer, name string) *Text {
	self := &Text{
		NewWidget(parent, name),
		"Ryan was here.",
	}
	self.Foreground = Palette[Black]
	self.Background = Palette[White]
	return self
}

func (self *Text) Draw(x, y, w, h float64, ctx vg.Context) {
	ctx.Scissor(x, y, w, h)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillColor(self.Background)
	ctx.Fill()

	ctx.SetFontPtSize(14)
	// ascender, descender, lineh := ctx.TextMetrics()
	ascender, descender, lineh := ctx.TextMetrics()
	lx, ly := x, y+lineh
	for i := 0; i < len(self.text); i++ {
		t := self.text[i : i+1]
		// _, _, xmax, _ := ctx.TextBounds(t, lx, ly)
		// drawBounds(xmin, ymin, xmax, ymax, ctx)
		ctx.BeginPath()
		ctx.FillColor(self.Foreground)
		ctx.Text(lx, ly, t)
		lx = lx + descender + ascender
	}
	ctx.Text(x, ly*2, self.text)

	// for i := 0.0; i < 100; i++ {
	// 	ctx.FillColor(self.Foreground)
	// 	ctx.TextAlign(vg.ALIGN_LEFT)
	// 	ctx.FontSize(19)
	// 	ctx.FindFont(vg.FONT_DEFAULT)
	// 	ctx.Text(x+0.5, y+19.5*i, self.text)
	// }
	// fmt.Println(x+0.5, y+39.5)
	// ctx.BeginPath()
	// ctx.StrokeColor(self.Foreground)
	// ctx.MoveTo(x+10.5, y+39.5)
	// ctx.LineTo(x+10.5, y)
	// ctx.Stroke()
	ctx.ResetScissor()
}

func drawBounds(xmin, ymin, xmax, ymax float64, ctx vg.Context) {
	ctx.BeginPath()
	ctx.FillColor(Palette[Green1])
	// ctx.MoveTo(xmin, ymin)
	// ctx.LineTo(xmin, ymax)
	// ctx.LineTo(xmax, ymax)
	// ctx.LineTo(xmax, ymin)
	// ctx.LineTo(xmin, ymin)
	ctx.RoundedRect(xmin, ymin, xmax-xmin, ymax-ymin, 3)
	ctx.Fill()
}
