package color

import (
	// "fmt"
	"time"
)

type tweenState int

const (
	TWEEN_STOPPED tweenState = iota
	TWEEN_STARTED
	TWEEN_PAUSED
)

type Tween struct {
	a         Color
	b         Color
	length    time.Duration
	startTime time.Time

	// Used to calculate an amended start time when restarted
	pausedTime time.Time
	state      tweenState
}

func NewTween(a, b Color, length time.Duration) *Tween {
	return &Tween{
		a:      a,
		b:      b,
		length: length,
	}
}

func (self *Tween) Start() {
	switch self.state {
	case TWEEN_STOPPED:
		self.startTime = time.Now()
	case TWEEN_PAUSED:
		d := self.pausedTime.Sub(self.startTime)
		// subtracting the delta for new start time
		self.startTime = time.Now().Add(-d)
	default:
		//no-op
	}
	self.state = TWEEN_STARTED
}

func (self *Tween) Pause() {
	self.state = TWEEN_PAUSED
	self.pausedTime = time.Now()
}

func (self *Tween) Stop() {
	self.state = TWEEN_STOPPED
}

func (self *Tween) Color() Color {
	d := time.Since(self.startTime)

	switch self.state {
	case TWEEN_STOPPED:
		// fmt.Println("STOP")
		return self.a
	case TWEEN_STARTED:
		return Color{
			self.lin(self.a.R, self.b.R, d),
			self.lin(self.a.G, self.b.G, d),
			self.lin(self.a.B, self.b.B, d),
			self.lin(self.a.A, self.b.A, d),
		}

	case TWEEN_PAUSED:
		d := self.pausedTime.Sub(self.startTime)
		return Color{
			self.lin(self.a.R, self.b.R, d),
			self.lin(self.a.G, self.b.G, d),
			self.lin(self.a.B, self.b.B, d),
			self.lin(self.a.A, self.b.A, d),
		}
	}

	return self.b
}

func (self *Tween) lin(a, b float64, duration time.Duration) float64 {
	v := (float64(duration) / float64(self.length))
	if v > 1 {
		// fmt.Println(duration, self.length, v)
		return b
	}
	d := 0.0
	if a >= b {
		d = a - b
	} else {
		d = b - a
	}
	// fmt.Println(d, self.length, duration)
	r := d * v
	if a >= b {
		return a - (d * r)
	} else {
		return a + (d * r)
	}
}
