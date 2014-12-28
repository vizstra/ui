package text

import (
	// "fmt"
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"runtime"
	"strings"
)

type Format vg.Align

const (
	LEFT    Format = Format(vg.ALIGN_LEFT)
	RIGHT   Format = Format(vg.ALIGN_RIGHT)
	CENTER  Format = Format(vg.ALIGN_CENTER)
	JUSTIFY Format = 1 << 7
	NOWRAP  Format = 1 << 8
)

type Text struct {
	ui.Element
	text        string
	tokens      []string
	font        string
	fontSize    float64
	format      Format
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
		RIGHT,
		nil,
		make([]*ui.Size, len(tokens)),
	}
	self.Background = White
	self.Foreground = Blue5

	return self
}

func (self *Text) Draw(x, y, w, h float64, ctx vg.Context) {
	ctx.Scissor(x, y, w, h)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillColor(self.Background)
	ctx.Fill()

	ctx.FillColor(self.Foreground)

	self.lastContext = &ctx
	self.forEachDrawnToken(x, y, w, h,
		func(i int, x, y, w, h float64, ctx *vg.Context) {
			ctx.Text(x, y, self.tokens[i])
		})
	ctx.ResetScissor()
}

// type alias for DRY purposes
type tokenIterCB func(i int, x, y, xmax, ymax float64, ctx *vg.Context)

func (self *Text) forEachDrawnToken(x, y, w, h float64, f tokenIterCB) {
	switch self.format {
	case RIGHT:
		self.iterRight(x, y, w, h, f)
	case CENTER:
	case JUSTIFY:
	default:
		self.iterLeft(x, y, w, h, f)
	}
}

func (self *Text) iterLeft(x, y, w, h float64, f tokenIterCB) {
	ctx := self.lastContext
	ctx.SetFontPtSize(self.fontSize)
	_, _, lineh := ctx.TextMetrics()
	tx, ty := x, y+lineh
	spaceWidth := self.spaceWidth(ctx)
	farEdge := x + w
	for i, _ := range self.tokens {
		size := self.tokenSize(i, ctx)
		if tx+size.W > farEdge {
			tx = x
			ty += lineh
		}

		if ty > y+h+lineh {
			break
		}

		f(i, tx, ty, size.W, size.H, ctx)
		tx += size.W + spaceWidth
	}
}

func (self *Text) iterRight(x, y, w, h float64, f tokenIterCB) {
	ctx := self.lastContext
	ctx.SetFontPtSize(self.fontSize)
	_, _, lineh := ctx.TextMetrics()
	lx, ly := x, y+lineh
	spaceWidth := self.spaceWidth(ctx)
	farEdge := x + w
	follower, tx, ty := 0, lx, ly
	runtime.Breakpoint()
	for i, _ := range self.tokens {
		size := self.tokenSize(i, ctx)

		if lx+size.W > farEdge {
			tx = (w - lx)
			ty = ly
			for ; follower < i; follower++ {
				size = self.sizes[follower]
				f(follower, tx, ty, size.W, size.H, ctx)
				tx += size.W + spaceWidth
			}
			lx = x
			ly += lineh
		}

		if ly > y+h+lineh {
			break
		}

		lx += size.W + spaceWidth
	}
}

func drawBounds(xmin, ymin, xmax, ymax float64, ctx vg.Context) {
	ctx.BeginPath()
	ctx.FillColor(Green1)
	ctx.RoundedRect(xmin, ymin, xmax-xmin, ymax-ymin, 3)
	ctx.Fill()
}

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
