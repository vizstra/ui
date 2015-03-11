package ui

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
	gl "github.com/vizstra/opengl/gl42"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/vg"
	"log"
	"runtime"
	"time"
)

var debug bool = false

var RR gl.Float = 0.0

type Window struct {
	Element
	endchan  chan bool
	window   *glfw.Window
	renderer vg.Renderer
	child    Drawer
	KeyDispatch
	CharDispatch
	MouseDispatch
	ScrollDispatch
}

// NewWindow returns a new Window. The returned Window
// will init opengl when its Start function is called.
func NewWindow(name, title string, x, y, w, h int) *Window {
	glfw.SetErrorCallback(errorCallback)

	// TODO Monitor support needs to be added
	window, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		panic(err)
	}

	win := &Window{
		NewElement(nil, title),

		make(chan bool),
		window,
		nil,
		nil,
		NewKeyDispatch(),
		NewCharDispatch(),
		NewMouseDispatch(),
		NewScrollDispatch(),
	}

	window.SetPosition(x, y)
	window.MakeContextCurrent()

	var lastScroll time.Time
	var yoffTicks, xoffTicks float64
	var maxScrollTick = 25.0
	window.SetScrollCallback(func(w *glfw.Window, xoff float64, yoff float64) {
		if yoff != 0 && time.Since(lastScroll) < 50*time.Millisecond {
			if yoff > -maxScrollTick || yoff < maxScrollTick {
				yoffTicks = yoffTicks + yoff
			}
		} else {
			yoffTicks = yoff
		}

		if xoff != 0 && time.Since(lastScroll) < 100*time.Millisecond {
			if xoff > -maxScrollTick || xoff < maxScrollTick {
				xoffTicks = xoffTicks + xoff
			}
		} else {
			xoffTicks = xoff
		}

		lastScroll = time.Now()

		for h, _ := range win.ScrollHandlers {
			h.HandleScroll(xoffTicks, yoffTicks)
		}
	})

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		ke := KeyEvent{Key(key), scancode, Action(action), ModifierKey(mods)}
		win.KeyDispatch.Dispatch(ke)
	})

	window.SetCharacterCallback(func(w *glfw.Window, char uint) {
		win.CharDispatch.Dispatch(char)
	})

	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		for h, _ := range win.MouseClickHandlers {
			h.HandleMouseClick(MouseButtonState{MouseButton(button), Action(action), ModifierKey(mods)})
		}
	})

	window.SetCursorPositionCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		for h, _ := range win.MousePositionHandlers {
			h.HandleMousePosition(xpos, ypos)
		}
	})

	window.SetCursorEnterCallback(func(w *glfw.Window, entered bool) {
		for h, _ := range win.MouseEnterHandlers {
			h.HandleMouseEnter(entered)
		}
	})

	win.AddMouseEnterHandler(win)
	return win
}

func (self *Window) HandleMouseEnter(b bool) {
	if b {
		RR = 1
	} else {
		RR = 0
	}
}

func (self *Window) Draw(ctx vg.Context) {
	now := time.Now()
	_, _, w, h := self.Bounds()
	fbw, fbh := self.FramebufferSize()

	// Calculate pixel ration for hi-dpi devices.
	bg := color.Gray13

	gl.ClearColor(gl.Float(bg.R), gl.Float(bg.G), gl.Float(bg.B), gl.Float(bg.A))
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

	//Do OpenGL stuff
	// Iterate over scene, rendering what can be rendered
	if self.renderer != nil {
		self.renderer.Render()
	}

	// 2D Display
	// r := float64(fbw) / float64(w)
	gl.Viewport(0, 0, gl.Sizei(fbw), gl.Sizei(fbh))

	if self.child != nil {
		ctx.BeginFrame(int(w), int(h), 1.0)
		self.child.SetBounds(0.0, 0.0, float64(w), float64(h))
		self.child.Draw(ctx)
		ctx.EndFrame()
	}
	if debug {
		fmt.Println(time.Since(now))
	}
}

func (self *Window) Start() chan bool {
	now := time.Now()
	// This is the ever important draw goroutine
	go func() {

		// super important
		runtime.LockOSThread()
		ctx := vg.NewContext()
		for !self.window.ShouldClose() {
			// gl.Enable(gl.BLEND)
			// gl.Enable(gl.LINE_SMOOTH)
			// gl.Hint(gl.LINE_SMOOTH_HINT, gl.NICEST)
			// gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

			w, h := self.Size()
			self.SetBounds(0, 0, float64(w), float64(h))

			if time.Since(now) > 2*time.Second {
				debug = true

				self.Draw(ctx)
				debug = false
				now = time.Now()
			} else {
				self.Draw(ctx)
			}
			self.window.SwapBuffers()
			glfw.PollEvents()

			// var total gl.Int
			// gl.GetIntegerv(0x9048, &total)

			// var available gl.Int
			// gl.GetIntegerv(0x9049, &available)

			// fmt.Println(total, available)
		}
		glfw.Terminate()
		self.endchan <- true
	}()
	return self.endchan
}

func (self *Window) SetChild(child Drawer) {
	self.child = child
	self.AddDrawerKeyHandler(child)
	self.AddDrawerCharHandler(child)
	self.AddDrawerMouseClickHandler(child)
	self.AddDrawerMousePositionHandler(child)
	self.AddDrawerMouseEnterHandler(child)
}

func (self *Window) Child() Drawer {
	return self.child
}

func (self *Window) SetRenderer(renderer vg.Renderer) {
	self.renderer = renderer
}

func (self *Window) Renderer() vg.Renderer {
	return self.renderer
}

// SetRenderFunc wraps the function in Renderer for convienence in some
// instances.  This overrides previously set instances.
func (self *Window) SetRenderFunc(f func()) {
	self.renderer = &vg.RenderFunc{f}
}

func (self *Window) SetTitle(title string) {
	self.window.SetTitle(title)
}

func (self *Window) SetSize(size Size) {
	self.window.SetSize(int(size.W), int(size.H))
}

func (self *Window) Size() (w, h int) {
	return self.window.GetSize()
}

func (self *Window) FramebufferSize() (w, h int) {
	return self.window.GetFramebufferSize()
}

func (self *Window) SetPosition(pos Position) {
	self.window.SetPosition(int(pos.X), int(pos.Y))
}

func (self *Window) Position() (x, y int) {
	return self.window.GetPosition()
}

func (self *Window) Hide() {
	self.window.Hide()
}

func (self *Window) Iconify() {
	self.window.Iconify()
}

func (self *Window) Show() {
	self.window.Show()
}

func (self *Window) Restore() {
	self.window.Restore()
}

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Printf("%v: %v\n", err, desc)
}

func init() {
	if !glfw.Init() {
		panic("Cannot initialize windowing library.")
	}

	e := gl.Init()
	if e != nil {
		panic(e)
	}
}
