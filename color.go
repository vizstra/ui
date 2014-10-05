package ui

import (
	"image/color"
)

const (
	COLOR_BUTTON_BACKGROUND = "button background"
)

var DefaultColors map[string]color.Color
var Colors map[string]color.Color

func init() {
	DefaultColors = make(map[string]color.Color)
	Colors = DefaultColors

	DefaultColors[COLOR_BUTTON_BACKGROUND] = color.RGBA{31, 31, 31, 235}
}
