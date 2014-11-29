package widget

import (
	"github.com/vizstra/ui"
)

type Text struct {
	Widget
}

func NewText(parent ui.Drawer, name string) *Text {
	return &Text{
		NewWidget(parent, name),
	}
}
