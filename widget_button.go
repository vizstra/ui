package ui

import (
	"image/color"
)

type Button struct {
	text string
}

func NewButton(name, text string) *Button {
	return &Button{text}
}

func (self *Button) Draw(x, y, w, h float64, ctx Context) {
	bg := ctx.BoxGradient(10, 10, 300, 50, 50/2, 50, color.NRGBA{90, 90, 90, 1}, color.NRGBA{70, 70, 70, 1})
	ctx.BeginPath()
	ctx.RoundedRect(10, 10, 300, 50, 5)
	ctx.FillPaint(bg)
	ctx.Fill()
}
