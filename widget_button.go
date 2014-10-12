package ui

import (
	"github.com/vizstra/vg"
	"image/color"
)

type Button struct {
	Parent          Drawer
	Name            string
	Text            string
	Radius          float64
	Background      color.Color
	HoverBackground color.Color
	ClickBackground color.Color
	displayColor    color.Color
	MouseDispatch
}

func NewButton(parent Drawer, name, text string) *Button {
	self := &Button{
		parent,
		name,
		text,
		5,
		Colors[COLOR_BUTTON_BACKGROUND],
		Colors[COLOR_BUTTON_HOVER_BACKGROUND],
		Colors[COLOR_BUTTON_CLICK_BACKGROUND],
		Colors[COLOR_BUTTON_BACKGROUND],
		NewMouseDispath(),
	}

	if p, ok := parent.(MousePositionDispatcher); ok {
		p.AddMousePositionCB(func(x, y float64) {
			self.DispatchMousePosition(x, y)
		})
	}

	inside := false
	colorbg := func() {
		if inside {
			self.displayColor = self.HoverBackground
		} else {
			self.displayColor = self.Background
		}
	}

	if p, ok := parent.(MouseEnterDispatcher); ok {
		p.AddMouseEnterCB(func(in bool) {
			inside = in
			colorbg()
			self.DispatchMouseEnter(in)
		})
	}

	if p, ok := parent.(MouseClickDispatcher); ok {
		p.AddMouseClickCB(func(m MouseButtonState) {
			if m.MouseButton == MOUSE_BUTTON_LEFT || m.MouseButton == MOUSE_BUTTON_1 {
				colorbg()
				if m.Action == PRESS {
					self.displayColor = self.ClickBackground
				}
			}
			self.DispatchMouseClick(m)
		})
	}
	return self
}

func (self *Button) Draw(x, y, w, h float64, ctx vg.Context) {
	c := CloneColor(self.displayColor)
	// c.A = .3
	bg := ctx.BoxGradient(x, y, w, h/3, h/2, h, c, self.displayColor)

	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.Radius)
	ctx.FillPaint(bg)
	ctx.Fill()

	ctx.Scissor(x, y, w, h)
	ctx.FillColor(color.RGBA{222, 222, 222, 255})
	ctx.TextAlign(vg.ALIGN_CENTER | vg.ALIGN_MIDDLE)
	ctx.FontSize(17)
	ctx.FindFont(vg.FONT_DEFAULT)
	ctx.WrappedText(x, y+h/2, w, self.Text)
	ctx.ResetScissor()
}
