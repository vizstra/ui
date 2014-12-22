package ui

import (
	"github.com/vizstra/vg"
	"image/color"
)

func DrawDefaultElement(x, y, w, h float64, bg color.Color, ctx vg.Context) {

	// Shadow
	ctx.BeginPath()
	ctx.StrokeColor(color.RGBA{100, 100, 100, 100})
	ctx.RoundedRect(x, y, w, h, 3)
	ctx.StrokeWidth(1)
	ctx.Stroke()
	ctx.FillPaint(ctx.BoxGradient(x, y, w, h, h*.5, h, bg, bg))
	ctx.Fill()
}
