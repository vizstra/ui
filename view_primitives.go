package ui

import ()

type Point2d struct {
	X, Y float64
}

type Size2d struct {
	W, H float64
}

type Rectangle struct {
	X, Y float64
	W, H float64
}

func (self Rectangle) Contains(x, y float64) bool {

	if x > self.X && x < self.X+self.W &&
		y > self.Y && y < self.Y+self.H {
		return true
	}

	return false
}

type View struct {
}
