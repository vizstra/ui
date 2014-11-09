package ui

import ()

type MouseDispatcher interface {
	MousePositionDispatcher
	MouseEnterDispatcher
	MouseClickDispatcher
}

type MousePositionDispatcher interface {
	AddMousePositionCB(func(float64, float64)) MousePositionHandler
	AddMousePositionHandler(MousePositionHandler)
	DispatchMousePosition(x, y float64)
}

type MouseDispatch struct {
	MouseClickHandlers    map[MouseClickHandler]bool
	MousePositionHandlers map[MousePositionHandler]bool
	MouseEnterHandlers    map[MouseEnterHandler]bool
	MouseClickChans       map[chan MouseButtonState]bool
	MousePositionChans    map[chan MousePositionState]bool
	MouseEnterChans       map[chan bool]bool
}

func NewMouseDispatch() MouseDispatch {
	return MouseDispatch{
		make(map[MouseClickHandler]bool, 0),
		make(map[MousePositionHandler]bool, 0),
		make(map[MouseEnterHandler]bool, 0),
		make(map[chan MouseButtonState]bool, 0),
		make(map[chan MousePositionState]bool, 0),
		make(map[chan bool]bool, 0),
	}
}

func (self *MouseDispatch) DispatchMousePosition(x, y float64) {
	for k, _ := range self.MousePositionHandlers {
		k.HandleMousePosition(x, y)
		return
	}

	for k, _ := range self.MousePositionChans {
		k <- MousePositionState{x, y}
	}
}

type MousePositionHandler interface {
	HandleMousePosition(x, y float64)
}

type mousePositionBridge struct {
	f func(float64, float64)
}

func (self mousePositionBridge) HandleMousePosition(x, y float64) {
	self.f(x, y)
}

func NewMousePositionHandler(h func(float64, float64)) MousePositionHandler {
	return &mousePositionBridge{h}
}

func (self *MouseDispatch) AddMousePositionHandler(h MousePositionHandler) {
	if _, ok := self.MousePositionHandlers[h]; !ok {
		self.MousePositionHandlers[h] = true
	}
}

func (self *MouseDispatch) AddMousePositionCB(mpcb func(float64, float64)) MousePositionHandler {
	r := NewMousePositionHandler(mpcb)
	self.AddMousePositionHandler(r)
	return r
}

func (self *MouseDispatch) AddDrawerMousePositionHandler(d Drawer) {
	if h, ok := d.(MousePositionHandler); ok {
		self.AddMousePositionHandler(h)
	}
}

type MousePositionState struct {
	X, Y float64
}
