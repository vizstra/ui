package ui

type Point2d struct {
	X, Y float64
}

type Size2d struct {
	W, H float64
}

type Rectangle struct {
	Position Point2d
	Size     Size2d
}

type View struct {
}
