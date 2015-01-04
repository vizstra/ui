package ui

import (
	"github.com/vizstra/vg"
)

type Drawer interface {
	Rectangular
	Draw(ctx vg.Context)
}
