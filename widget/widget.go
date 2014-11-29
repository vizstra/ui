package widget

import (
	"github.com/vizstra/ui"
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
)

type Widget struct {
	Parent ui.Drawer
	ui.MouseDispatch
	Name            string
	inside          bool
	CornerRadius    float64
	Foreground      Color
	Background      Color
	HoverBackground Color
	ClickBackground Color
	displayColor    Color
	DrawCB          func(x, y, w, h float64, ctx vg.Context)
}

// NewWidget builds and returns a Widget.
// parent: the type forwarding drawing calls to this Widget
// name: reference string for testing and accessibility purposes.
func NewWidget(parent ui.Drawer, name string) Widget {
	return Widget{
		parent,
		ui.NewMouseDispatch(),
		name,
		false,
		3,
		Palette[WIDGET_FOREGROUND],
		Palette[WIDGET_BACKGROUND],
		Palette[WIDGET_HOVER_BACKGROUND],
		Palette[WIDGET_CLICK_BACKGROUND],
		Palette[WIDGET_BACKGROUND],
		nil,
	}
}

func (self *Widget) determineBackground() {
	if self.inside {
		self.displayColor = self.HoverBackground
	} else {
		self.displayColor = self.Background
	}
}

func (self *Widget) DispatchMouseEnter(in bool) {
	self.inside = in
	self.determineBackground()
	self.MouseDispatch.DispatchMouseEnter(in)
}

func (self *Widget) DispatchMouseClick(m ui.MouseButtonState) {
	if m.MouseButton == ui.MOUSE_BUTTON_LEFT || m.MouseButton == ui.MOUSE_BUTTON_1 {
		self.determineBackground()
		if m.Action == ui.PRESS {
			self.displayColor = self.ClickBackground
		}
	}
	self.MouseDispatch.DispatchMouseClick(m)
}

func (self *Widget) Draw(x, y, w, h float64, ctx vg.Context) {
	// draw background
	c := CloneColor(self.displayColor)
	bg := ctx.BoxGradient(x, y, w, h/3, h/2, h, c, self.displayColor)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillPaint(bg)
	ctx.Fill()

	if self.DrawCB != nil {
		self.DrawCB(x, y, w, h, ctx)
	}
}

const (
	WIDGET_FOREGROUND = 0x2e68051c + iota
	WIDGET_BACKGROUND
	WIDGET_HOVER_BACKGROUND
	WIDGET_CLICK_BACKGROUND
)

func init() {
	DefaultPalette[WIDGET_FOREGROUND] = DefaultPalette[White]
	DefaultPalette[WIDGET_BACKGROUND] = DefaultPalette[Gray12]
	DefaultPalette[WIDGET_HOVER_BACKGROUND] = DefaultPalette[Gray13]
	DefaultPalette[WIDGET_CLICK_BACKGROUND] = DefaultPalette[Gray12]
}
