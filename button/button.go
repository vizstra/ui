package button

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/vg"
)

type Button struct {
	ui.Element
	Text string
}

func NewButton(parent ui.Drawer, name, text string) *Button {
	self := &Button{
		ui.NewElement(parent, name),
		text,
	}

	self.DrawCB = func(ctx vg.Context) {
		x, y, w, h := self.Bounds()
		ctx.Scissor(x, y, w, h)
		ctx.FillColor(self.Foreground)
		ctx.TextAlign(vg.ALIGN_CENTER | vg.ALIGN_MIDDLE)

		ctx.SetFontSize(35)
		ctx.FindFont(vg.FONT_DEFAULT)
		ctx.WrappedText(x, y+h/2, w, self.Text)
		ctx.ResetScissor()
	}

	return self
}
