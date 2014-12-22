package button

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/vg"
	"os"
)

type ImageButton struct {
	ui.Element
	Text           string
	imagePath      string
	image          *vg.Image
	hoverImagePath string
	hoverImage     *vg.Image
	displayImage   *vg.Image
}

func NewImageButton(parent ui.Drawer, name, text string) *ImageButton {
	self := &ImageButton{
		ui.NewElement(parent, name),
		text,
		"",
		nil,
		"",
		nil,
		nil,
	}

	// handle alternate images on hover
	self.AddMouseEnterCB(func(in bool) {
		if self.MouseInside() && self.hoverImage != nil {
			self.displayImage = self.hoverImage
		} else if self.image != nil {
			self.displayImage = self.image
		}
	})

	self.DrawCB = func(x, y, w, h float64, ctx vg.Context) {
		ui.DrawDefaultElement(x, y, w, h, self.DisplayColor(), ctx)
		if self.image == nil {
			self.image = ctx.NewImage(self.imagePath, 0)
			self.displayImage = self.image
		}

		if self.hoverImage == nil && len(self.hoverImagePath) > 0 {
			self.hoverImage = ctx.NewImage(self.hoverImagePath, 0)
		}

		ww, hh := self.image.Size()
		ctx.Scissor(x, y, w, h)
		self.displayImage.Draw(x+((w-ww)/2), y+((h-hh)/2))
		ctx.ResetScissor()
	}
	return self
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
