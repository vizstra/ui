package ui

import ()

type ScrollDispatcher interface {
	AddScrollCB(func(float64, float64)) ScrollHandler
	AddScrollHandler(ScrollHandler)
	DispatchScroll(xoff, yoff float64)
}

type ScrollDispatch struct {
	ScrollHandlers map[ScrollHandler]bool
	ScrollChans    map[chan ScrollState]bool
}

func (self *ScrollDispatch) DispatchScroll(xoff, yoff float64) {
	for k, _ := range self.ScrollHandlers {
		k.HandleScroll(xoff, yoff)
		return
	}

	for k, _ := range self.ScrollChans {
		k <- ScrollState{xoff, yoff}
	}
}

func (self *ScrollDispatch) AddScrollHandler(h ScrollHandler) {
	if _, ok := self.ScrollHandlers[h]; !ok {
		self.ScrollHandlers[h] = true
	}
}

func (self *ScrollDispatch) AddScrollCB(scb func(float64, float64)) ScrollHandler {
	r := NewScrollHandler(scb)
	self.AddScrollHandler(r)
	return r
}

func (self *ScrollDispatch) AddDrawerScrollHandler(d Drawer) {
	if h, ok := d.(ScrollHandler); ok {
		self.AddScrollHandler(h)
	}
}

type ScrollState struct {
	XOffset, YOffset float64
}

type ScrollHandler interface {
	HandleScroll(xoff, yoff float64)
}

func NewScrollDispatch() ScrollDispatch {
	return ScrollDispatch{
		make(map[ScrollHandler]bool, 0),
		make(map[chan ScrollState]bool, 0),
	}
}

type scrollBridge struct {
	f func(float64, float64)
}

func (self scrollBridge) HandleScroll(xoff, yoff float64) {
	self.f(xoff, yoff)
}

func NewScrollHandler(h func(float64, float64)) ScrollHandler {
	return &scrollBridge{h}
}
