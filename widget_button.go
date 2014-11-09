package ui

import (
	"github.com/vizstra/vg"
	"image/color"
)

type Button struct {
	Parent          Drawer
	Name            string
	Text            string
	Radius          float64
	Background      color.Color
	HoverBackground color.Color
	ClickBackground color.Color
	displayColor    color.Color
	inside          bool
	MouseDispatch
}

func NewButton(parent Drawer, name, text string) *Button {
	self := &Button{
		parent,
		name,
		text,
		5,
		Colors[COLOR_BUTTON_BACKGROUND],
		Colors[COLOR_BUTTON_HOVER_BACKGROUND],
		Colors[COLOR_BUTTON_CLICK_BACKGROUND],
		Colors[COLOR_BUTTON_BACKGROUND],
		false,
		NewMouseDispatch(),
	}
	return self
}

func (self *Button) determineBackground() {
	if self.inside {
		self.displayColor = self.HoverBackground
	} else {
		self.displayColor = self.Background
	}
}

func (self *Button) DispatchMouseEnter(in bool) {
	self.inside = in
	self.determineBackground()
	self.MouseDispatch.DispatchMouseEnter(in)
}

func (self *Button) DispatchMouseClick(m MouseButtonState) {
	if m.MouseButton == MOUSE_BUTTON_LEFT || m.MouseButton == MOUSE_BUTTON_1 {
		self.determineBackground()
		if m.Action == PRESS {
			self.displayColor = self.ClickBackground
		}
	}
	self.MouseDispatch.DispatchMouseClick(m)

}

func (self *Button) Draw(x, y, w, h float64, ctx vg.Context) {
	x += .5
	y += .5

	c := CloneColor(self.displayColor)
	bg := ctx.BoxGradient(x, y, w, h/3, h/2, h, c, self.displayColor)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.Radius)
	ctx.FillPaint(bg)
	ctx.Fill()

	// ctx.BeginPath()
	// ctx.StrokeWidth(1.0)
	// ctx.RoundedRect(x, y, w, h, self.Radius)
	// ctx.StrokeColor(Color{0, 0, 0, 1})
	// // ctx.StrokePaint()
	// ctx.Stroke()

	ctx.Scissor(x, y, w, h)
	ctx.FillColor(color.RGBA{222, 222, 222, 255})
	ctx.TextAlign(vg.ALIGN_CENTER | vg.ALIGN_MIDDLE)
	ctx.FontSize(17)
	ctx.FindFont(vg.FONT_DEFAULT)
	ctx.WrappedText(x, y+h/2, w, self.Text)
	ctx.ResetScissor()
}
