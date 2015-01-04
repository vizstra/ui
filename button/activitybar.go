package button

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"math"
	"time"
)

type ActivityBar struct {
	ui.Element
	MaxValue float64
	Data     []float64
	tween    *color.Tween
}

// NewActivityBar returns an ActivityBar widget.
// maxValue: helps determines drawn data height; must
//           be greater than 0.0, otherwise set to 1.0
// data:     data slice to use to render
func NewActivityBar(parent ui.Drawer, name string, maxValue float64, data []float64) *ActivityBar {
	self := &ActivityBar{
		ui.NewElement(parent, name),
		maxValue,
		data,
		color.NewTween(color.Gray13, color.Red1, 250*time.Millisecond),
	}

	self.AddMouseEnterCB(func(b bool) {
		if b {
			self.tween = color.NewTween(color.Gray13, color.Red1, 250*time.Millisecond)
			self.tween.Start()
		} else {
			c := self.tween.Color()
			self.tween = color.NewTween(c, color.Gray13, 250*time.Millisecond)
			self.tween.Start()
		}
	})

	self.DrawCB = func(ctx vg.Context) {
		x, y, w, h := self.Bounds()
		barw := 5.0 // width of the bar
		sepw := 1.0
		totalw := (barw + sepw)
		leftover := w - (math.Floor(w/totalw) * totalw)
		x += leftover / 2
		w -= 2 + totalw
		c := self.tween.Color()

		l := float64(len(self.Data))
		if l == 0 {
			return
		}

		maxcount := w / totalw
		s := 0
		if l > maxcount {
			s = int(l - 1 - maxcount)
		}

		iy := y + h

		for i := s; i < len(self.Data); i++ {
			v := self.Data[i]
			ctx.BeginPath()
			ix := math.Floor(x + (float64(i-s) * totalw))
			ih := math.Floor(iy - (h * (v / self.MaxValue)))
			ctx.FillColor(c)
			ax := ix + sepw
			bx := ix + barw
			ctx.MoveTo(ax, iy)
			ctx.LineTo(ax, ih)
			ctx.LineTo(bx, ih)
			ctx.LineTo(bx, iy)
			ctx.Fill()
		}
	}

	return self
}
