package ui

import (
	"image/color"
)

const (
	COLOR_BUTTON_BACKGROUND          = "button background"
	COLOR_BUTTON_HOVER_BACKGROUND    = "button hover background"
	COLOR_BUTTON_CLICK_BACKGROUND    = "button click background"
	COLOR_BUTTON_INACTIVE_BACKGROUND = "button inactive background"
)

var DefaultColors map[string]color.Color
var Colors map[string]color.Color

func init() {
	DefaultColors = make(map[string]color.Color)
	Colors = DefaultColors

	DefaultColors[COLOR_BUTTON_BACKGROUND] = color.RGBA{31, 31, 31, 235}
	DefaultColors[COLOR_BUTTON_HOVER_BACKGROUND] = color.RGBA{31, 131, 31, 235}
	DefaultColors[COLOR_BUTTON_CLICK_BACKGROUND] = color.RGBA{10, 51, 200, 235}
}
