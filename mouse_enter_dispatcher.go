package ui

import ()

type MouseEnterDispatcher interface {
	AddMouseEnterCB(func(bool)) MouseEnterHandler
	AddMouseEnterHandler(MouseEnterHandler)
	DispatchMouseEnter(in bool)
}

func (self *MouseDispatch) DispatchMouseEnter(in bool) {
	for k, _ := range self.MouseEnterHandlers {
		k.HandleMouseEnter(in)
		return
	}

	for k, _ := range self.MouseEnterChans {
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
	if _, ok := self.MouseEnterHandlers[h]; !ok {
		self.MouseEnterHandlers[h] = true
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
