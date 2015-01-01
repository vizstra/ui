package ui

import ()

type Scroller interface {
	IncrementYOffset(float64)
	IncrementXOffset(float64)
	YOffset() float64
	XOffset() float64
	SetIncrementValue(float64)
	IncrementValue() float64
}

type Scroll struct {
	yoff float64
	xoff float64
	incr float64
}

func NewScroll() Scroll {
	return Scroll{0, 0, 0}
}

func (self *Scroll) SetYOffset(yoff float64) {
	self.yoff = yoff
}

func (self Scroll) YOffset() float64 {
	return self.yoff
}

func (self *Scroll) SetXOffset(xoff float64) {
	self.xoff = xoff
}

func (self Scroll) XOffset() float64 {
	return self.xoff
}

func (self *Scroll) SetIncrement(incr float64) {
	self.incr = incr
}

func (self Scroll) Increment() float64 {
	return self.incr
}
