package ui

import ()

type MouseEnterDispatcher interface {
	AddMouseEnterCB(func(bool)) MouseEnterHandler
	AddMouseEnterHandler(MouseEnterHandler)
}

func (self *MouseDispatch) DispatchMouseEnter(in bool) {
	for k, _ := range self.mouseEnterHandlers {
		k.HandleMouseEnter(in)
		return
	}

	for k, _ := range self.mouseEnterChans {
		k <- in
	}
}

type MouseEnterHandler interface {
	HandleMouseEnter(in bool)
}

type mouseEnterBridge struct {
	f func(bool)
}

func (self mouseEnterBridge) HandleMouseEnter(in bool) {
	self.f(in)
}

func NewMouseEnterHandler(h func(bool)) MouseEnterHandler {
	return &mouseEnterBridge{h}
}

func (self *MouseDispatch) AddMouseEnterHandler(h MouseEnterHandler) {
	if _, ok := self.mouseEnterHandlers[h]; !ok {
		self.mouseEnterHandlers[h] = true
	}
}

func (self *MouseDispatch) AddMouseEnterCB(mecb func(bool)) MouseEnterHandler {
	r := NewMouseEnterHandler(mecb)
	self.AddMouseEnterHandler(r)
	return r
}

func (self *MouseDispatch) AddDrawerMouseEnterHandler(d Drawer) {
	if h, ok := d.(MouseEnterHandler); ok {
		self.AddMouseEnterHandler(h)
	}
}
