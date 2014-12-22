package text

import ()

type Token struct {
	Text string
	Type TokenType
}

type TokenType int16

const (
	NORMAL = iota
)
