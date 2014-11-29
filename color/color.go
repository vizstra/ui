package color

import (
	"math"
)

type Color struct {
	R, G, B, A float64
}

func (self Color) RGBA() (r, g, b, a uint32) {
	return uint32(self.R * 255), uint32(self.G * 255), uint32(self.B * 255), uint32(self.A * 255)
}

func RGBA(r, g, b, a uint8) Color {
	return Color{
		float64(r) / 255.0,
		float64(g) / 255.0,
		float64(b) / 255.0,
		float64(a) / 255.0,
	}
}

func clritof(c uint8) float64 {
	return float64(c) / 255
}

func HexRGBA(hex uint32) Color {
	return Color{clritof(uint8(hex >> 24 & 0xFF)),
		clritof(uint8(hex >> 16 & 0xFF)),
		clritof(uint8(hex >> 8 & 0xFF)),
		clritof(uint8(hex & 0xFF))}
}

func CloneColor(a Color) Color {
	return Color{a.R, a.G, a.B, a.A}
}

// HSL returns the color components of this color as HSL.
func (self Color) HSL() (h, s, l float64) {
	r, g, b := self.R, self.G, self.B
	max := math.Max(b, math.Max(g, r))
	min := math.Min(b, math.Min(g, r))
	l = (max + min) / 2.0

	if max == min {
		h = 0.0
		s = 0.0
	} else {
		var d = max - min

		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			v := 0.0
			if g < b {
				v = 6.0
			}
			h = (g-b)/d + (v)
		case g:
			h = (b-r)/d + 2.0
		case b:
			h = (r-g)/d + 4.0
		}
		h /= 6.0
	}

	return h * 360, s, l
}

// NewHSL creates a Color from the given HSL parameters.
// Hue        (h) [0.0 .. 360.0]
// Saturation (s) [0.0 .. 1.0]
// Lightness  (l) [0.0 .. 1.0]
func NewHSL(h, s, l float64) Color {

	if h > 360 {
		h -= (360 * math.Mod(h, 360))
	} else if h < 0 {
		h += (360 * math.Mod(h, 360))
	}

	if s > 1 {
		s = 1
	} else if s < 0 {
		s = 0
	}

	if l > 1 {
		l = 1
	} else if l < 0 {
		l = 0
	}

	h /= 360.0
	var c Color

	if s == 0 {
		c.R, c.G, c.B = l, l, l

	} else {
		hue2rgb := func(p, q, t float64) float64 {
			if t < 0 {
				t += 1
			}

			if t > 1 {
				t -= 1
			}

			if t < 1.0/6 {
				return p + (q-p)*6*t
			}

			if t < 0.5 {
				return q
			}

			if t < 2.0/3 {
				return p + (q-p)*(2.0/3-t)*6.0
			}

			return p
		}

		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}

		p := 2.0*l - q
		c.R = hue2rgb(p, q, h+1.0/3.0)
		c.G = hue2rgb(p, q, h)
		c.B = hue2rgb(p, q, h-1.0/3.0)
	}

	c.A = 1
	return c
}

// Lighten by the given percentage (p) [0..1].
func (self Color) Lighten(p float64) Color {
	h, s, l := self.HSL()
	l += (l * p)
	return NewHSL(h, s, l)
}

// func (self Color) Darken(p float64) Color {
// 	return self
// }

// func (self Color) Saturate(p float64) Color {
// 	return self
// }

// func (self Color) Desaturate(p float64) Color {
// 	return self
// }

// func (self Color) Mix(c Color, weight float64) Color {
// 	return self
// }

// Multiply multiplies this color by the given color
// and returns the result; the resultant color adopts
// is assigned the alpha channel value from this color.
func (self Color) Multiply(c Color) Color {
	return Color{
		self.R * c.R,
		self.G * c.G,
		self.B * c.B,
		self.A,
	}
}

// Difference subtracts the given color and returns
// the result; this new color adopts is assigned the
// alpha channel value from the first color.
func (self Color) Difference(c Color) Color {
	return Color{
		math.Abs(self.R - c.R),
		math.Abs(self.G - c.G),
		math.Abs(self.B - c.B),
		self.A,
	}
}

// Average returns a new computed color of the two
// colors' RGA channels.  The returned color adopts
// the same alpha channel value as the first color.
func (self Color) Average(c Color) Color {
	return Color{
		(self.R + c.R) / 2,
		(self.G + c.G) / 2,
		(self.B + c.B) / 2,
		self.A,
	}
}
