package ui

import (
	. "github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"math"
)

type Element struct {
	Parent Drawer
	MouseDispatch
	ScrollDispatch
	Rectangle
	Name             string
	inside           bool
	CornerRadius     float64
	BorderWidth      float64
	Foreground       Color
	Background       Color
	HoverBackground  Color
	ClickBackground  Color
	ActiveBackground Color
	DrawCB           func(ctx vg.Context)
}

// NewElement builds and returns a Element.
// parent: the type forwarding drawing calls to this Element
// name: reference string for testing and accessibility purposes.
func NewElement(parent Drawer, name string) Element {
	return Element{
		parent,
		NewMouseDispatch(),
		NewScrollDispatch(),
		Rectangle{Position{0, 0}, Size{0, 0}},
		name,
		false,
		3,
		1,
		Palette[ELEMENT_FOREGROUND],
		Palette[ELEMENT_BACKGROUND],
		Palette[ELEMENT_HOVER_BACKGROUND],
		Palette[ELEMENT_CLICK_BACKGROUND],
		Palette[ELEMENT_BACKGROUND],
		nil,
	}
}

func (self *Element) MouseInside() bool {
	return self.inside
}

func (self *Element) determineBackground() {
	if self.inside {
		self.ActiveBackground = self.HoverBackground
	} else {
		self.ActiveBackground = self.Background
	}
}

func (self *Element) SetCommonBackground(c Color) {
	self.Background = c
	self.HoverBackground = c
	self.ClickBackground = c
	self.ActiveBackground = c
}

func (self *Element) DispatchMouseEnter(in bool) {
	self.inside = in
	self.determineBackground()
	self.MouseDispatch.DispatchMouseEnter(in)
}

func (self *Element) DispatchMouseClick(m MouseButtonState) {
	if m.MouseButton == MOUSE_BUTTON_LEFT || m.MouseButton == MOUSE_BUTTON_1 {
		self.determineBackground()
		if m.Action == PRESS {
			self.ActiveBackground = self.ClickBackground
		}
	}
	self.MouseDispatch.DispatchMouseClick(m)
}

func (self *Element) Clamp(x, y float64) (X, Y float64) {
	return math.Floor(x) + .5, math.Floor(y) + .5
}

func (self *Element) Draw(ctx vg.Context) {
	x, y, w, h := self.Bounds()
	x, y = self.Clamp(x, y)

	// draw background
	c := CloneColor(self.ActiveBackground)
	bg := ctx.BoxGradient(x, y, w, h, h, h, c, c)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillPaint(bg)
	ctx.Fill()
	ctx.StrokeWidth(self.BorderWidth)
	ctx.StrokeColor(c.Lighten(-0.5))
	ctx.Stroke()

	if self.DrawCB != nil {
		self.DrawCB(ctx)
	}
}

const (
	ELEMENT_FOREGROUND = 0x2e68051c + iota
	ELEMENT_BACKGROUND
	ELEMENT_HOVER_BACKGROUND
	ELEMENT_CLICK_BACKGROUND
)

func init() {
	Palette[ELEMENT_FOREGROUND] = White
	Palette[ELEMENT_BACKGROUND] = Gray10
	Palette[ELEMENT_HOVER_BACKGROUND] = Gray6
	Palette[ELEMENT_CLICK_BACKGROUND] = Gray12
}
