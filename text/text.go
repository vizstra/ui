package text

import (
	"fmt"
	"strings"

	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

const DEBUG = false

type Alignment vg.Align

const (
	LEFT    Alignment = Alignment(vg.ALIGN_LEFT)
	RIGHT   Alignment = Alignment(vg.ALIGN_RIGHT)
	CENTER  Alignment = Alignment(vg.ALIGN_CENTER)
	JUSTIFY Alignment = 1 << 7
	NOWRAP  Alignment = 1 << 8
)

// Text serves as the basic text rendering element for most UI components.
type Text struct {
	ui.Element
	ui.Scroll
	text        string
	tokens      []string
	font        string
	fontSize    float64
	Alignment   Alignment
	lastContext *vg.Context
	bounds      []*ui.Rectangle
	selections  []*ui.Rectangle
}

func New(parent ui.Drawer, name, text string) *Text {
	tokens := strings.Fields(text)
	self := &Text{
		ui.NewElement(parent, name),
		ui.NewScroll(),
		text,
		tokens,
		vg.FONT_DEFAULT,
		21,
		LEFT,
		nil,
		make([]*ui.Rectangle, len(tokens)),
		make([]*ui.Rectangle, len(tokens)),
	}

	self.CornerRadius = 0

	self.AddScrollCB(func(xoff, yoff float64) {
		vertical := self.YOffset()
		vertical -= yoff * self.Increment()
		if vertical < 0 {
			vertical = 0
		}
		self.SetYOffset(vertical)
	})

	X, Y := 0.0, 0.0
	self.AddMousePositionCB(func(x, y float64) {
		X, Y = x, y
	})

	self.AddMouseClickCB(func(state ui.MouseButtonState) {
		fmt.Println("CLICK!")
		xx, yy, w, h := self.Bounds()

		self.forEachDrawnToken(xx, yy, w, h,
			func(i int, x, y, lineHeight float64, bounds *ui.Rectangle, ctx *vg.Context) {
				r := ui.Rectangle{ui.Position{x, y - lineHeight/2}, ui.Size{bounds.Size.W, bounds.Size.H + lineHeight}}

				if r.Contains(X, Y) {
					fmt.Println(tokens[i])
				}
			},
		)
	})

	self.Background = White
	self.Foreground = Gray10
	return self
}

func (self *Text) Draw(ctx vg.Context) {
	x, y, w, h := self.Bounds()
	self.lastContext = &ctx
	ctx.Scissor(x, y, w, h)
	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillColor(self.Background)
	ctx.Fill()
	ctx.FillColor(self.Foreground)

	self.forEachDrawnToken(x, y, w, h,
		func(i int, x, y, lineHeight float64, bounds *ui.Rectangle, ctx *vg.Context) {
			ctx.Text(x, y, self.tokens[i])
		},
	)

	// for i := 0; i < len(self.selections); i++ {
	// 	s := self.selections[i]
	// }

	ctx.ResetScissor()
}

func (self *Text) forEachDrawnToken(x, y, w, h float64,
	f func(i int, x, y, lineHeight float64, bounds *ui.Rectangle, ctx *vg.Context)) {

	ctx := self.lastContext
	ctx.SetFontPtSize(self.fontSize)

	_, _, lineh := ctx.TextMetrics()
	self.SetIncrement(lineh)

	by := y - self.YOffset() + lineh/1.25
	spaceWidth := self.spaceWidth(ctx)
	farEdge := x + w
	linewidth := 0.0
	bottom := y + h + lineh

	for a, b := 0, 0; a < len(self.tokens) && by < bottom; {

		for ; b < len(self.tokens); b++ {
			bounds := self.tokenBounds(b, ctx)
			if x+bounds.W+linewidth > farEdge {
				break
			}
			linewidth += bounds.W + spaceWidth
		}

		justificationSpread := 0.0
		ax := x

		switch self.Alignment {
		case RIGHT:
			ax = farEdge - linewidth + spaceWidth
		case CENTER:
			ax = (x + farEdge - linewidth + spaceWidth) / 2
		case JUSTIFY:
			if b != len(self.tokens) {
				justificationSpread = (w - linewidth + spaceWidth) / float64(b-a-1)
			}
		}

		for ; a < b; a++ {
			bounds := self.bounds[a]
			// if by < y {
				f(a, ax, by, lineh, bounds, ctx)
			// }
			ax += bounds.W + spaceWidth + justificationSpread
		}

		if self.Alignment == NOWRAP && a < len(self.tokens) {
			bounds := self.bounds[a]
			if by < y {
				f(a, ax, by, lineh, bounds, ctx)
			}
			break
		}

		linewidth = 0
		by += lineh
	}
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
func (self *Text) tokenBounds(i int, ctx *vg.Context) *ui.Rectangle {
	bounds := self.bounds[i]
	if bounds == nil {
		x, y, w, h := ctx.TextBounds(self.tokens[i], 0, 0)

		bounds = &ui.Rectangle{ui.Position{x, y}, ui.Size{w, h}}
		self.bounds[i] = bounds
	}
	return bounds
}

func (self *Text) spaceWidth(ctx *vg.Context) float64 {
	xmin, _, xmax, _ := ctx.TextBounds(".", 0, 0)
	return xmax - xmin
}
