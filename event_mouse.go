package ui

import (
	"github.com/go-gl/glfw3"
)

type MousePositionHandler interface {
	HandleMousePosition(x, y float64)
}

type MouseEnterHandler interface {
	HandleMouseEnter(bool)
}

type MouseDispatch struct {
	mouseButtonHandlers   map[MouseButtonHandler]bool
	mousePositionHandlers map[MousePositionHandler]bool
	mouseEnterHandlers    map[MouseEnterHandler]bool
}

func (self *Window) AddMouseButtonHandler(h MouseButtonHandler) {
	if _, ok := self.mouseButtonHandlers[h]; !ok {
		self.mouseButtonHandlers[h] = true
	}
}

func (self *Window) AddMousePositionHandler(h MousePositionHandler) {
	if _, ok := self.mousePositionHandlers[h]; !ok {
		self.mousePositionHandlers[h] = true
	}
}

func (self *Window) AddMouseEnterHandler(h MouseEnterHandler) {
	if _, ok := self.mouseEnterHandlers[h]; !ok {
		self.mouseEnterHandlers[h] = true
	}
}

type MouseButton glfw3.MouseButton
type MouseButtonHandler interface {
	HandleMouseButton(MouseButton, Action, ModifierKey)
}
