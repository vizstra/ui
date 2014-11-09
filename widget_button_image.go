package ui

import (
	// "fmt"
	"github.com/vizstra/vg"
	"image/color"
	"os"
)

type ImageButton struct {
	Parent          Drawer
	Name            string
	Text            string
	Radius          float64
	Background      color.Color
	HoverBackground color.Color
	ClickBackground color.Color
	displayColor    color.Color
	imagePath       string
	image           *vg.Image
	hoverImagePath  string
	hoverImage      *vg.Image
	displayImage    *vg.Image
	inside          bool
	MouseDispatch
}

func NewImageButton(parent Drawer, name, text string) *ImageButton {
	self := &ImageButton{
		parent,
		name,
		text,
		5,
		Colors[COLOR_IMG_BUTTON_BACKGROUND],
		Colors[COLOR_IMG_BUTTON_HOVER_BACKGROUND],
		Colors[COLOR_IMG_BUTTON_CLICK_BACKGROUND],
		Colors[COLOR_IMG_BUTTON_BACKGROUND],
		"",
		nil,
		"",
		nil,
		nil,
		false,
		NewMouseDispatch(),
	}

	return self
}

func (self *ImageButton) determineBackground() {
	if self.inside {
		self.displayColor = self.HoverBackground
		if self.hoverImage != nil {
			self.displayImage = self.hoverImage
		}
	} else {
		self.displayColor = self.Background
		if self.image != nil {
			self.displayImage = self.image
		}
	}
}

func (self *ImageButton) DispatchMouseEnter(in bool) {
	self.inside = in
	self.determineBackground()
	self.MouseDispatch.DispatchMouseEnter(in)
}

func (self *ImageButton) DispatchMouseClick(m MouseButtonState) {
	if m.MouseButton == MOUSE_BUTTON_LEFT || m.MouseButton == MOUSE_BUTTON_1 {
		self.determineBackground()
		if m.Action == PRESS {
			self.displayColor = self.ClickBackground
		}
	}
	self.MouseDispatch.DispatchMouseClick(m)

}

func (self *ImageButton) SetImagePath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	self.image = nil
	self.imagePath = path
	f.Close()
	return nil
}

func (self *ImageButton) SetHoverImagePath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	self.hoverImage = nil
	self.hoverImagePath = path
	f.Close()
	return nil
}

func (self *ImageButton) Draw(x, y, w, h float64, ctx vg.Context) {
	if self.image == nil {
		self.image = ctx.NewImage(self.imagePath, 0)
		self.displayImage = self.image
	}

	if self.hoverImage == nil {
		self.hoverImage = ctx.NewImage(self.hoverImagePath, 0)
	}

	// draw background
	c := CloneColor(self.displayColor)
	bg := ctx.BoxGradient(x, y, w, h/3, h/2, h, c, self.displayColor)
	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, self.Radius)
	ctx.FillPaint(bg)
	ctx.Fill()

	ww, hh := self.image.Size()
	ctx.Scissor(x, y, w, h)
	self.displayImage.Draw(x+((w-ww)/2), y+((h-hh)/2))
	ctx.ResetScissor()
}
