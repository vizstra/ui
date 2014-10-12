package ui

import (
	"image/color"
)

const (
	COLOR_BUTTON_BACKGROUND          = "button background"
	COLOR_BUTTON_HOVER_BACKGROUND    = "button hover background"
	COLOR_BUTTON_CLICK_BACKGROUND    = "button click background"
	COLOR_BUTTON_DISABLED_BACKGROUND = "button disabled background"

	COLOR_WIDGET_BACKGROUND          = "widget background"
	COLOR_WIDGET_HOVER_BACKGROUND    = "widget hover background"
	COLOR_WIDGET_CLICK_BACKGROUND    = "widget click background"
	COLOR_WIDGET_DISABLED_BACKGROUND = "widget inactive background"
)

var DefaultColors map[string]color.Color
var Colors map[string]color.Color

func init() {
	DefaultColors = make(map[string]color.Color)
	Colors = DefaultColors
	DefaultColors[COLOR_BUTTON_BACKGROUND] = color.RGBA{31, 31, 31, 235}
	DefaultColors[COLOR_BUTTON_HOVER_BACKGROUND] = color.RGBA{10, 51, 105, 235}
	DefaultColors[COLOR_BUTTON_CLICK_BACKGROUND] = color.RGBA{10, 31, 65, 235}
	DefaultColors[COLOR_BUTTON_DISABLED_BACKGROUND] = color.RGBA{170, 170, 170, 235}

	DefaultColors[COLOR_WIDGET_BACKGROUND] = color.RGBA{171, 171, 171, 235}
	DefaultColors[COLOR_WIDGET_HOVER_BACKGROUND] = color.RGBA{161, 171, 171, 235}
	DefaultColors[COLOR_WIDGET_CLICK_BACKGROUND] = color.RGBA{141, 151, 151, 235}
	DefaultColors[COLOR_WIDGET_DISABLED_BACKGROUND] = color.RGBA{70, 70, 70, 235}
}
