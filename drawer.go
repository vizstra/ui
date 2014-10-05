package ui

import (
	"github.com/vizstra/vg"
)

type Drawer interface {
	Draw(x, y, w, h float64, ctx vg.Context)
}
