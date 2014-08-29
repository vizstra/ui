// +build !goci
package ui

// #cgo linux pkg-config: glfw3 xxf86vm x11 gl glu xrandr xi xcursor gthread-2.0
// #cgo linux LDFLAGS: -lm
// #define NANOVG_GL3_IMPLEMENTATION
// #include <stdlib.h>
// #include <math.h>
// #include <GL/gl.h>
// #include "nanovg.h"
// #include "nanovg_gl.h"
// #include "nanovg_gl_utils.h"
// #include "fontstash.h"
// #include "stb_image.h"
// #include "stb_truetype.h"
import "C"

import (
	"image/color"
	"unsafe"
)

type LineCap C.int

const (
	BUTT LineCap = iota
	ROUND
	SQUARE
	BEVEL
	MITER
)

type Winding C.int

const (
	CCW Winding = 1
	CW  Winding = 2
)

type PatternRepeat C.int

const (
	NOREPEAT = 0
	REPEATX  = 0x01 // Repeat image pattern in X direction
	REPEATY  = 0x02 // Repeat image pattern in Y direction
)

type Align C.int

const (
	// Horizontal align
	ALIGN_LEFT   Align = 1 << 0 // Default, align text horizontally to left.
	ALIGN_CENTER Align = 1 << 1 // Align text horizontally to center.
	ALIGN_RIGHT  Align = 1 << 2 // Align text horizontally to right.

	// Vertical align
	ALIGN_TOP      Align = 1 << 3 // Align text vertically to top.
	ALIGN_MIDDLE   Align = 1 << 4 // Align text vertically to middle.
	ALIGN_BOTTOM   Align = 1 << 5 // Align text vertically to bottom.
	ALIGN_BASELINE Align = 1 << 6 // Default, align text vertically to baseline.
)

type MipmapFlag C.int

const (
	IMAGE_GENERATE_MIPMAPS MipmapFlag = 1 << 0 // Generate mipmaps during creation of the image.
)

func toColor(c color.Color) C.NVGcolor {
	ri, gi, bi, ai := c.RGBA()
	var clr C.NVGcolor
	clr.r = C.float(ri) / 255.0
	clr.g = C.float(gi) / 255.0
	clr.b = C.float(bi) / 255.0
	clr.a = C.float(ai) / 255.0
	return clr
}

type Context struct {
	cbase *C.NVGcontext
}

// NewContext returns a New Context
func NewContext() Context {
	return Context{C.nvgCreateGL3(C.NVG_ANTIALIAS | C.NVG_STENCIL_STROKES)}
}

// BeginFrame begins drawing a new frame.
// BeginFrame defines the size of the window to render to in relation currently
// set viewport (i.e. glViewport on GL backends). Device pixel ration allows to
// control the rendering on Hi-DPI devices.
//
// For example, GLFW returns two dimensions for an opened window: window size and
// frame buffer size. In that case you would set the window Width/Height to the
// window size ratio to: frameBufferWidth / windowWidth.
func (self *Context) BeginFrame(w, h int, ratio float64) {
	C.nvgBeginFrame(self.cbase, C.int(w), C.int(h), C.float(ratio))
}

// EndFrame ends drawing, flushing remaining render state.
func (self *Context) EndFrame() {
	C.nvgEndFrame(self.cbase)
}

// Save pushes and saves the current render state into a state stack.
// A matching nvgRestore() must be used to restore the state.
func (self *Context) Save() {
	C.nvgSave(self.cbase)
}

// Restore pops and restores current render state.
func (self *Context) Restore() {
	C.nvgRestore(self.cbase)
}

// Reset will reset current render state to default
// values. Does not affect the render state stack.
func (self *Context) Reset() {
	C.nvgReset(self.cbase)
}

// StrokeColor sets current stroke style to a solid color.
func (self *Context) StrokeColor(color color.Color) {
	C.nvgStrokeColor(self.cbase, toColor(color))
}

// StrokePaint sets current stroke style to a paint, which can
// be a one of the gradients or a pattern.
func (self *Context) StrokePaint(paint Paint) {
	C.nvgStrokePaint(self.cbase, paint.cbase)
}

// FillColor sets current fill cstyle to a solid color.
func (self *Context) FillColor(color color.Color) {
	C.nvgFillColor(self.cbase, toColor(color))
}

// FillPaint sets current fill style to a paint, which can be
// a one of the gradients or a pattern.
func (self *Context) FillPaint(paint Paint) {
	C.nvgFillPaint(self.cbase, C.NVGpaint(paint.cbase))
}

// MiterLimit sets the miter limit of the stroke style.
// Miter limit controls when a sharp corner is beveled.
func (self *Context) MiterLimit(limit float64) {
	C.nvgMiterLimit(self.cbase, C.float(limit))
}

// StrokeWidth sets the stroke witdth of the stroke style.
func (self *Context) StrokeWidth(size float64) {
	C.nvgStrokeWidth(self.cbase, C.float(size))
}

// LineCap sets how the end of the line (cap) is drawn,
// Can be one of: BUTT (default), ROUND, SQUARE, etc.
func (self *Context) LineCap(cap LineCap) {
	C.nvgLineCap(self.cbase, C.int(cap))
}

// LineJoin sets how sharp path corners are drawn.
// Can be one of MITER (default), ROUND, BEVEL.
func (self *Context) LineJoin(join LineCap) {
	C.nvgLineJoin(self.cbase, C.int(join))
}

// Sets the transparency applied to all rendered shapes.
// Alreade transparent paths will get proportionally
// more transparent as well.
func (self *Context) GlobalAlpha(alpha float64) {
	C.nvgGlobalAlpha(self.cbase, C.float(alpha))
}

// ResetTransform resets current transform to a
// identity matrix.
func (self *Context) ResetTransform() {
	C.nvgResetTransform(self.cbase)
}

// Transform premultiplies current coordinate system by specified
// matrix. The parameters are interpreted as matrix as follows:
//   [a c e]
//   [b d f]
//   [0 0 1]
func (self *Context) Transform(a, b, c, d, e, f float64) {
	C.nvgTransform(self.cbase, C.float(a), C.float(b), C.float(c), C.float(d), C.float(e), C.float(f))
}

// Translate the current coordinate system.
func (self *Context) Translate(x, y float64) {
	C.nvgTranslate(self.cbase, C.float(x), C.float(y))
}

// Rotate the current coordinate system.
// The angle is specifid in radians.
func (self *Context) Rotate(angle float64) {
	C.nvgRotate(self.cbase, C.float(angle))
}

// Skew the current coordinate system along X axis.
// The angle is specifid in radians.
func (self *Context) SkewX(angle float64) {
	C.nvgSkewX(self.cbase, C.float(angle))
}

// Skew the current coordinate system along Y axis.
// The angle is specifid in radians.
func (self *Context) SkewY(angle float64) {
	C.nvgSkewY(self.cbase, C.float(angle))
}

// Scale the current coordinate system.
func (self *Context) Scale(x, y float64) {
	C.nvgScale(self.cbase, C.float(x), C.float(y))
}

func (self Context) BeginPath() {
	C.nvgBeginPath(self.cbase)
}

func (self Context) RoundedRect(x, y, w, h, r float64) {
	C.nvgRoundedRect(self.cbase, C.float(x), C.float(y), C.float(w), C.float(h), C.float(r))
}

func (self Context) Fill() {
	C.nvgFill(self.cbase)
}

// NewImage creates image by loading it from the disk from
// specified file name.  Returns the image handle.
func (self *Context) NewImage(filename string, flags int) Image {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	return Image{C.nvgCreateImage(self.cbase, cs, C.int(flags))}
}

// Creates image by loading it from the specified chunk of memory.
// Returns handle to the image.
// int nvgCreateImageMem(NVGcontext* ctx, int imageFlags, unsigned char* data, int ndata){

// }

// // Creates image from specified image data.
// // Returns handle to the image.
// int nvgCreateImageRGBA(NVGcontext* ctx, int w, int h, int imageFlags, const unsigned char* data);

// // Updates image data specified by image handle.
// void nvgUpdateImage(NVGcontext* ctx, int image, const unsigned char* data);

// // Returns the domensions of a created image.
// void nvgImageSize(NVGcontext* ctx, int image, int* w, int* h);

// // Deletes created image.
// void nvgDeleteImage(NVGcontext* ctx, int image);

type Image struct {
	handle C.int
}

type Paint struct {
	cbase C.NVGpaint
}

func (self Context) LinearGradient(x1, y1, x2, y2 float64, a, b color.Color) Paint {
	c, d := toColor(a), toColor(b)
	return Paint{C.nvgLinearGradient(self.cbase, C.float(x1), C.float(y1), C.float(x2), C.float(y2), c, d)}
}

func (self Context) RadialGradient(cx, cy, i, o float64, a, b color.Color) Paint {
	c, d := toColor(a), toColor(b)
	return Paint{C.nvgRadialGradient(self.cbase, C.float(cx), C.float(cy), C.float(i), C.float(o), c, d)}
}

func (self Context) BoxGradient(x, y, w, h, r, f float64, a, b color.Color) Paint {
	c, d := toColor(a), toColor(b)
	return Paint{C.nvgBoxGradient(self.cbase, C.float(x), C.float(y), C.float(w), C.float(h), C.float(r), C.float(f), c, d)}
}

type GlyphPosition struct {
	cbase *C.NVGglyphPosition
}

type TextRow struct {
	cbase *C.NVGtextRow
}
