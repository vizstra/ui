package ui

import (
	"github.com/go-gl/glfw3"
)

const (
	MOUSE_BUTTON_1      MouseButton = MouseButton(glfw3.MouseButton1)
	MOUSE_BUTTON_2      MouseButton = MouseButton(glfw3.MouseButton2)
	MOUSE_BUTTON_3      MouseButton = MouseButton(glfw3.MouseButton3)
	MOUSE_BUTTON_4      MouseButton = MouseButton(glfw3.MouseButton4)
	MOUSE_BUTTON_5      MouseButton = MouseButton(glfw3.MouseButton5)
	MOUSE_BUTTON_6      MouseButton = MouseButton(glfw3.MouseButton6)
	MOUSE_BUTTON_7      MouseButton = MouseButton(glfw3.MouseButton7)
	MOUSE_BUTTON_8      MouseButton = MouseButton(glfw3.MouseButton8)
	MOUSE_BUTTON_LAST   MouseButton = MouseButton(glfw3.MouseButtonLast)
	MOUSE_BUTTON_LEFT   MouseButton = MouseButton(glfw3.MouseButtonLeft)
	MOUSE_BUTTON_RIGHT  MouseButton = MouseButton(glfw3.MouseButtonRight)
	MOUSE_BUTTON_MIDDLE MouseButton = MouseButton(glfw3.MouseButtonMiddle)

	RELEASE Action = Action(glfw3.Release)
	PRESS   Action = Action(glfw3.Press)
	REPEAT  Action = Action(glfw3.Repeat)
)

type MouseClickDispatcher interface {
	AddMouseClickCB(func(MouseButtonState)) MouseClickHandler
	AddMouseClickHandler(MouseClickHandler)
}

type MouseButton glfw3.MouseButton
type MouseClickHandler interface {
	HandleMouseClick(m MouseButtonState)
}

type MouseButtonState struct {
	MouseButton
	Action
	ModifierKey
}

func (self *MouseDispatch) DispatchMouseClick(m MouseButtonState) {
	for k, _ := range self.mouseClickHandlers {
		k.HandleMouseClick(m)
		return
	}

	for k, _ := range self.mouseClickChans {
		k <- m
	}
}

type mouseClickBridge struct {
	f func(MouseButtonState)
}

func (self mouseClickBridge) HandleMouseClick(m MouseButtonState) {
	self.f(m)
}

func NewMouseClickHandler(h func(MouseButtonState)) MouseClickHandler {
	return &mouseClickBridge{h}
}

func (self *MouseDispatch) AddMouseClickHandler(h MouseClickHandler) {
	if _, ok := self.mouseClickHandlers[h]; !ok {
		self.mouseClickHandlers[h] = true
	}
}

func (self *MouseDispatch) AddMouseClickCB(mpcb func(MouseButtonState)) MouseClickHandler {
	r := NewMouseClickHandler(mpcb)
	self.AddMouseClickHandler(r)
	return r
}

func (self *MouseDispatch) AddDrawerMouseClickHandler(d Drawer) {
	if h, ok := d.(MouseClickHandler); ok {
		self.AddMouseClickHandler(h)
	}
}
