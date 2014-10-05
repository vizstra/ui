package ui

import (
	"github.com/vizstra/vg"
	"image/color"
)

type Button struct {
	Parent     Drawer
	Name       string
	Text       string
	Radius     float64
	Background color.Color
}

func NewButton(parent Drawer, name, text string) *Button {
	return &Button{parent, name, text, 5, Colors[COLOR_BUTTON_BACKGROUND]}
}

func (self *Button) Draw(x, y, w, h float64, ctx vg.Context) {
	c := CloneColor(self.Background)
	c.A = .1
	bg := ctx.BoxGradient(x, y, w, h/3, h/2, h, c, self.Background)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.Radius)
	ctx.FillPaint(bg)
	ctx.Fill()

	ctx.Scissor(x, y, w, h)
	ctx.FillColor(color.RGBA{222, 222, 222, 255})
	ctx.TextAlign(vg.ALIGN_CENTER | vg.ALIGN_MIDDLE)
	ctx.FontSize(20)
	ctx.NewFont("Arial", "src/github.com/vizstra/ui/res/fonts/cmu/cmunbmr.ttf")
	ctx.WrappedText(x, y+h/2, w, self.Text)
	ctx.ResetScissor()
}
