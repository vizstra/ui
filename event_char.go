package ui

import ()

type CharHandler interface {
	HandleChar(uint)
}

type CharDispatcher interface {
	AddCharHandler(CharHandler)
	RemoveCharHandler(CharHandler)
	AddCharChan(chan uint)
	RemoveCharChan(chan uint)
	Dispatch(uint)
}

type CharDispatch struct {
	charHandlers map[CharHandler]bool
	charChans    map[chan uint]bool
}

func NewCharDispatch() CharDispatch {
	return CharDispatch{
		make(map[CharHandler]bool, 0),
		make(map[chan uint]bool, 0),
	}
}

// AddDrawerCharHandler will add the drawer to the char
// handler if it satisfies the CharHandler interface.
func (self *CharDispatch) AddDrawerCharHandler(d Drawer) {
	if h, ok := d.(CharHandler); ok {
		self.AddCharHandler(h)
	}
}

func (self *CharDispatch) AddCharHandler(h CharHandler) {
	if _, ok := self.charHandlers[h]; !ok {
		self.charHandlers[h] = true
	}
}

func (self *CharDispatch) RemoveCharHandler(h CharHandler) {
	if _, ok := self.charHandlers[h]; ok {
		delete(self.charHandlers, h)
	}
}

func (self *CharDispatch) AddCharChan(c chan uint) {
	if _, ok := self.charChans[c]; !ok {
		self.charChans[c] = true
	}
}

func (self *CharDispatch) RemoveCharChan(c chan uint) {
	if _, ok := self.charChans[c]; ok {
		delete(self.charChans, c)
	}
}

func (self *CharDispatch) Dispatch(c uint) {
	for h, _ := range self.charHandlers {
		h.HandleChar(c)
	}
	for h, _ := range self.charChans {
		h <- c
	}
}
