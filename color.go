package ui

import (
	"image/color"
)

const (
	COLOR_BUTTON_BACKGROUND          = "button background"
	COLOR_BUTTON_HOVER_BACKGROUND    = "button hover background"
	COLOR_BUTTON_CLICK_BACKGROUND    = "button click background"
	COLOR_BUTTON_DISABLED_BACKGROUND = "button disabled background"

	COLOR_IMG_BUTTON_BACKGROUND          = "image button background"
	COLOR_IMG_BUTTON_HOVER_BACKGROUND    = "image button hover background"
	COLOR_IMG_BUTTON_CLICK_BACKGROUND    = "image button click background"
	COLOR_IMG_BUTTON_DISABLED_BACKGROUND = "image button disabled background"

	COLOR_DATA_BACKGROUND = "data background"

	COLOR_WIDGET_BACKGROUND          = "widget background"
	COLOR_WIDGET_HOVER_BACKGROUND    = "widget hover background"
	COLOR_WIDGET_CLICK_BACKGROUND    = "widget click background"
	COLOR_WIDGET_DISABLED_BACKGROUND = "widget inactive background"

	COLOR_TITLEBAR_BACKGROUND          = "titlebar background"
	COLOR_TITLEBAR_HOVER_BACKGROUND    = "titlebar hover background"
	COLOR_TITLEBAR_CLICK_BACKGROUND    = "titlebar click background"
	COLOR_TITLEBAR_DISABLED_BACKGROUND = "titlebar inactive background"
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

	DefaultColors[COLOR_IMG_BUTTON_BACKGROUND] = color.RGBA{31, 31, 31, 35}
	DefaultColors[COLOR_IMG_BUTTON_HOVER_BACKGROUND] = color.RGBA{10, 31, 65, 35}
	DefaultColors[COLOR_IMG_BUTTON_CLICK_BACKGROUND] = color.RGBA{10, 51, 105, 35}
	DefaultColors[COLOR_IMG_BUTTON_DISABLED_BACKGROUND] = color.RGBA{170, 170, 170, 10}

	DefaultColors[COLOR_DATA_BACKGROUND] = color.RGBA{255, 255, 255, 235}

	DefaultColors[COLOR_WIDGET_BACKGROUND] = color.RGBA{200, 200, 200, 255}
	DefaultColors[COLOR_WIDGET_HOVER_BACKGROUND] = color.RGBA{190, 200, 200, 255}
	DefaultColors[COLOR_WIDGET_CLICK_BACKGROUND] = color.RGBA{170, 180, 180, 255}
	DefaultColors[COLOR_WIDGET_DISABLED_BACKGROUND] = color.RGBA{70, 70, 70, 255}

	DefaultColors[COLOR_TITLEBAR_BACKGROUND] = color.RGBA{100, 150, 200, 255}
	DefaultColors[COLOR_TITLEBAR_HOVER_BACKGROUND] = color.RGBA{90, 140, 200, 255}
	DefaultColors[COLOR_TITLEBAR_CLICK_BACKGROUND] = color.RGBA{70, 170, 170, 255}
	DefaultColors[COLOR_TITLEBAR_DISABLED_BACKGROUND] = color.RGBA{70, 70, 70, 255}
}
