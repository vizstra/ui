package text

import (
	"fmt"
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"strings"
)

type Text struct {
	ui.Element
	text     string
	tokens   []string
	WrapText bool
}

func New(parent ui.Drawer, name, text string) *Text {
	self := &Text{
		ui.NewElement(parent, name),
		text,
		strings.Fields(text),
		true,
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

	ctx.SetFontPtSize(23)
	_, _, lineh := ctx.TextMetrics()
	tx, ty := x, y+lineh
	ctx.FillColor(self.Foreground)

	xmin, _, xmax, _ := ctx.TextBounds("o", 0, 0)
	spaceWidth := (xmax - xmin)
	farEdge := x + w
	for i, s := range self.tokens {
		xmin, _, xmax, _ := ctx.TextBounds(s, tx, ty)
		if i == -1 {
			fmt.Println(xmax, xmin, x)
		}

		if xmax > farEdge {
			if self.WrapText {
				tx = x
				ty += lineh
			} else if xmin > farEdge {
				break
			}
		}

		if ty > y+h+lineh {
			break
		}

		ctx.Text(tx, ty, s)
		tx += (xmax - xmin) + spaceWidth

	}
	ctx.ResetScissor()
}

func drawBounds(xmin, ymin, xmax, ymax float64, ctx vg.Context) {
	ctx.BeginPath()
	ctx.FillColor(Green1)
	ctx.RoundedRect(xmin, ymin, xmax-xmin, ymax-ymin, 3)
	ctx.Fill()
}
