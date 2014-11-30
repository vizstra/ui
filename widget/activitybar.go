package widget

import (
	"github.com/vizstra/ui"
	// "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

type ActivityBar struct {
	Widget
	maxValue float64
	Data     []float64
}

// NewActivityBar returns an ActivityBar widget.
// maxValue: helps determines drawn data height; must
//           be greater than 0.0, otherwise set to 1.0
// data:     data slice to use to render
func NewActivityBar(parent ui.Drawer, name string, maxValue float64, data []float64) *ActivityBar {
	self := &ActivityBar{
		NewWidget(parent, name),
		maxValue,
		data,
	}

	self.DrawCB = func(x, y, w, h float64, ctx vg.Context) {
		barw := 5.0 // width of the bar
		x += 1
		w -= 2 + barw
		c := self.Foreground
		l := float64(len(self.Data))
		if l == 0 {
			return
		}

		maxcount := w / barw
		s := 0
		if l > maxcount {
			s = int(l - maxcount)
		}

		iy := y + h

		for i := s; i < len(self.Data); i++ {
			v := self.Data[i]
			ctx.BeginPath()
			ix := float64(int(x + (float64(i-s) * barw)))
			ih := float64(int(iy - (h * (v / maxValue))))
			ctx.FillColor(c)
			ctx.MoveTo(ix+1, iy)
			ctx.LineTo(ix+1, ih)
			ctx.LineTo(ix+barw, ih)
			ctx.LineTo(ix+barw, iy)
			ctx.Fill()
		}
	}

	return self
}
