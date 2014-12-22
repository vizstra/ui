package text

import ()

type Document struct {
	Content  History
	ReadOnly bool
}

type History struct {
	Frames []Frame
}

type Frame struct {
	Lines []Line
}

type Line struct {
	Tokens []Token
}
