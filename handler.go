package ui

import (
	"github.com/go-gl/glfw3"
)

type ModifierKey glfw3.ModifierKey
type Key glfw3.Key
type Action glfw3.Action
type KeyHandler interface {
	HandleKey(Key, int, Action, ModifierKey)
}

type CharHandler interface {
	HandleChar(uint)
}

type MouseButton glfw3.MouseButton
type MouseButtonHandler interface {
	HandleMouseButton(MouseButton, Action, ModifierKey)
}

type MousePositionHandler interface {
	HandleMousePosition(x, y float64)
}

type MouseEnterHandler interface {
	HandleMouseEnter(bool)
}
