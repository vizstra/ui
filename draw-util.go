package ui

import (
	"github.com/vizstra/vg"
	"image/color"
)

func DrawDefaultWidget(x, y, w, h float64, ctx vg.Context) {

	// Shadow
	ctx.BeginPath()
	ctx.StrokeColor(color.RGBA{151, 151, 151, 150})
	ctx.RoundedRect(x, y, w, h, 3)
	ctx.StrokeWidth(1)
	ctx.Stroke()
	ctx.FillPaint(ctx.BoxGradient(x, y, w, h, h*.5, h, color.RGBA{255, 255, 255, 255}, color.RGBA{200, 200, 200, 255}))
	ctx.Fill()

}
