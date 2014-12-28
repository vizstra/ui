package text

import (
	// "fmt"
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"strings"
)

type Alignment vg.Align

const (
	LEFT    Alignment = Alignment(vg.ALIGN_LEFT)
	RIGHT   Alignment = Alignment(vg.ALIGN_RIGHT)
	CENTER  Alignment = Alignment(vg.ALIGN_CENTER)
	JUSTIFY Alignment = 1 << 7
	NOWRAP  Alignment = 1 << 8
)

type Text struct {
	ui.Element
	text        string
	tokens      []string
	font        string
	fontSize    float64
	alignment   Alignment
	lastContext *vg.Context
	sizes       []*ui.Size
}

func New(parent ui.Drawer, name, text string) *Text {
	tokens := strings.Fields(text)
	self := &Text{
		ui.NewElement(parent, name),
		text,
		tokens,
		vg.FONT_DEFAULT,
		22,
		NOWRAP,
		nil,
		make([]*ui.Size, len(tokens)),
	}
	self.Background = White
	self.Foreground = Blue5

	return self
}

func (self *Text) Draw(x, y, w, h float64, ctx vg.Context) {
	self.lastContext = &ctx
	ctx.Scissor(x, y, w, h)
	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillColor(self.Background)
	ctx.Fill()
	ctx.FillColor(self.Foreground)
	self.forEachDrawnToken(x, y, w, h,
		func(i int, x, y float64, size *ui.Size, ctx *vg.Context) {
			ctx.Text(x, y, self.tokens[i])
		})
	ctx.ResetScissor()
}

// type alias for DRY purposes
type tokenIterCB func(i int, x, y float64, size *ui.Size, ctx *vg.Context)

func (self *Text) forEachDrawnToken(x, y, w, h float64, f tokenIterCB) {

	ctx := self.lastContext
	ctx.SetFontPtSize(self.fontSize)

	_, _, lineh := ctx.TextMetrics()
	by := y + lineh
	spaceWidth := self.spaceWidth(ctx)
	farEdge := x + w
	linewidth := 0.0
	bottom := y + h + lineh

	for a, b := 0, 0; a < len(self.tokens) && by < bottom; {

		for b < len(self.tokens) {
			size := self.tokenSize(b, ctx)
			width := size.W + spaceWidth

			if width+linewidth > farEdge {
				break
			}
			linewidth += width
			b++
		}

		spread := 0.0
		ax := x
		switch self.alignment {
		case RIGHT:
			ax = x + farEdge - linewidth
		case CENTER:
			ax = (x + farEdge - linewidth) / 2
		case JUSTIFY:
			if b != len(self.tokens) {
				spread = (x + farEdge - linewidth) / float64(b-a-1)
			}
		}

		for ; a < b; a++ {
			size := self.sizes[a]
			f(a, ax, by, size, ctx)
			ax += size.W + spaceWidth + spread
		}

		if self.alignment == NOWRAP {
			size := self.sizes[a]
			f(a, ax, by, size, ctx)
			break
		}

		linewidth = 0
		by += lineh
	}
}

// func (self *Text) iterLeft(x, y, w, h float64, f tokenIterCB) {
// 	ctx := self.lastContext
// 	ctx.SetFontPtSize(self.fontSize)
// 	_, _, lineh := ctx.TextMetrics()
// 	tx, ty := x, y+lineh
// 	spaceWidth := self.spaceWidth(ctx)
// 	farEdge := x + w
// 	for i, _ := range self.tokens {
// 		size := self.tokenSize(i, ctx)
// 		if tx+size.W > farEdge {
// 			tx = x
// 			ty += lineh
// 		}

// 		if ty > y+h+lineh {
// 			break
// 		}

// 		f(i, tx, ty, size, ctx)
// 		tx += size.W + spaceWidth
// 	}
// }

// tokenSize retrieves the metrics from the internal cache, sizes.
// It is expensive to calculate the token sizes each frame and the
// sizes do not change often to justify the cost.  Testing revealed
// this approach yields a slower rate of growth.
//
// TODO: maybe choose to disable cpu scaling for better performance
//         > echo performance > /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor
//       or
//         > sudo update-rc.d ondemand disable
func (self *Text) tokenSize(i int, ctx *vg.Context) *ui.Size {
	size := self.sizes[i]
	if size == nil {
		_, _, w, h := ctx.TextBounds(self.tokens[i], 0, 0)
		size = &ui.Size{w, h}
		self.sizes[i] = size
	}
	return size
}

func (self *Text) spaceWidth(ctx *vg.Context) float64 {
	xmin, _, xmax, _ := ctx.TextBounds("o", 0, 0)
	return xmax - xmin
}
