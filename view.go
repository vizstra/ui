package ui

import ()

type View interface {
	Draw(x, y, w, h float64, ctx Context)
}
