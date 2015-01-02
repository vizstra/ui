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

func (self *Element) DisplayColor() Color {
	return self.displayColor
}

func (self *Element) determineBackground() {
	if self.inside {
		self.displayColor = self.HoverBackground
	} else {
		self.displayColor = self.Background
	}
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
			self.displayColor = self.ClickBackground
		}
	}
	self.MouseDispatch.DispatchMouseClick(m)
}

func (self *Element) Clamp(x, y float64) (X, Y float64) {
	return math.Floor(x) + .5, math.Floor(y) + .5
}

func (self *Element) Draw(x, y, w, h float64, ctx vg.Context) {
	x, y = self.Clamp(x, y)

	// draw background
	c := CloneColor(self.displayColor)
	bg := ctx.BoxGradient(x, y, w, h/3, h/2, h, c, self.displayColor)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.CornerRadius)
	ctx.FillPaint(bg)
	ctx.Fill()
	ctx.StrokeColor(self.displayColor.Lighten(-0.20))
	ctx.Stroke()

	if self.DrawCB != nil {
		self.DrawCB(x, y, w, h, ctx)
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
