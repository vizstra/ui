package color

import (
	"testing"
)

func expandClr(c Color) (R, G, B float64) {
	return c.R * 255.0, c.G * 255.0, c.B * 255.0
}

func testclr(t *testing.T, c Color, r, g, b float64) {
	R, G, B := expandClr(c)
	if R < r-1 || R > r+1 {
		t.Errorf("Bad Color Conversion: R=%v. want ~%v", R, r)
	}

	if G < g-1 || G > g+1 {
		t.Errorf("Bad Color Conversion: G=%v. want ~%v", G, g)
	}

	if B < b-1 || B > b+1 {
		t.Errorf("Bad Color Conversion: B=%v. want ~%v", B, b)
	}
}

func TestNewFromHSL(t *testing.T) {
	testclr(t, NewHSL(360.0, 1, .4), 204, 0, 0)
	testclr(t, NewHSL(0, 1, .4), 204, 0, 0)
	testclr(t, NewHSL(360.0, 0, .4), 102, 102, 102)
	testclr(t, NewHSL(360.0, 0, 0), 0, 0, 0)
	testclr(t, NewHSL(180.0, 1, .5), 0, 255, 255)
	testclr(t, NewHSL(360, 1, .4), 204, 0, 0)
	testclr(t, NewHSL(720, 1, .4), 204, 0, 0)
}

func testhsl(t *testing.T, c Color, h, s, l float64) {
	H, S, L := c.HSL()
	t.Log(c)
	if H < h-.01 || H > h+.01 {
		t.Errorf("Bad Color Conversion: H=%v S=%v L=%v. want H=~%v", H, S, L, h)
	}

	if S < s-.01 || S > s+.01 {
		t.Errorf("Bad Color Conversion: H=%v S=%v L=%v. want S=~%v", H, S, L, s)
	}

	if L < l-.01 || L > l+.01 {
		t.Errorf("Bad Color Conversion: H=%v S=%v L=%v. want L=~%v", H, S, L, l)
	}
}

func TestToHSL(t *testing.T) {
	testhsl(t, RGBA(204, 0, 0, 255), 0, 1, .4)
	testhsl(t, RGBA(102, 102, 102, 255), 0, 0, .4)
	testhsl(t, RGBA(0, 0, 0, 255), 0, 0, 0)
	testhsl(t, RGBA(0, 255, 255, 255), 180.0, 1, .5)
}
