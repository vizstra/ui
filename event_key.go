package ui

import (
	"github.com/go-gl/glfw3"
)

type ModifierKey glfw3.ModifierKey

type Key glfw3.Key

type Action glfw3.Action

type KeyEvent struct {
	Key      Key
	Scancode int
	Action   Action
	Modifier ModifierKey
}

type KeyHandler interface {
	HandleKey(KeyEvent)
}

type KeyDispatch struct {
	keyHandlers map[KeyHandler]bool
	keyChans    map[chan KeyEvent]bool
}

func NewKeyDispatch() KeyDispatch {
	return KeyDispatch{
		make(map[KeyHandler]bool, 0),
		make(map[chan KeyEvent]bool, 0),
	}
}

// Dispatch events to all handlers and channels.
func (self *KeyDispatch) Dispatch(ke KeyEvent) {
	for h, _ := range self.keyHandlers {
		h.HandleKey(ke)
	}
	for h, _ := range self.keyChans {
		h <- ke
	}
}

// AddDrawerKeyHandler adds the drawer to the key handler
// list if it satisfies the KeyHandler interface.
func (self *KeyDispatch) AddDrawerKeyHandler(d Drawer) {
	if h, ok := d.(KeyHandler); ok {
		self.AddKeyHandler(h)
	}
}

// AddKeyHandler adds the key handler from the dispatch.
func (self *KeyDispatch) AddKeyHandler(h KeyHandler) {
	if _, ok := self.keyHandlers[h]; !ok {
		self.keyHandlers[h] = true
	}
}

// RemoveKeyHandler removes the key handler from the dispatch.
func (self *KeyDispatch) RemoveKeyHandler(h KeyHandler) {
	if _, ok := self.keyHandlers[h]; ok {
		delete(self.keyHandlers, h)
	}
}

// AddKeyChan adds the key channel from the dispatch.
func (self *KeyDispatch) AddKeyChan(h chan KeyEvent) {
	if _, ok := self.keyChans[h]; !ok {
		self.keyChans[h] = true
	}
}

// RemoveKeyChan removes the key channel from the dispatch.
func (self *KeyDispatch) RemoveKeyChan(h chan KeyEvent) {
	if _, ok := self.keyChans[h]; ok {
		delete(self.keyChans, h)
	}
}
