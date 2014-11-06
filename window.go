package ui

import (
	glfw "github.com/go-gl/glfw3"
	gl "github.com/vizstra/opengl/gl43"
	"github.com/vizstra/vg"
	"log"
	"runtime"
)

var RR gl.Float = .0

type Window struct {
	name     string
	endchan  chan bool
	window   *glfw.Window
	renderer vg.Renderer
	child    Drawer
	KeyDispatch
	CharDispatch
	MouseDispatch
}

func NewWindow(name, title string, w, h, x, y int) *Window {
	glfw.SetErrorCallback(errorCallback)

	// TODO Monitor support needs to be added
	window, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		panic(err)
	}

	win := &Window{
		name,
		make(chan bool),
		window,
		nil,
		nil,
		NewKeyDispatch(),
		NewCharDispatch(),
		NewMouseDispatch(),
	}

	window.SetPosition(x, y)
	window.MakeContextCurrent()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		ke := KeyEvent{Key(key), scancode, Action(action), ModifierKey(mods)}
		win.KeyDispatch.Dispatch(ke)
	})

	window.SetCharacterCallback(func(w *glfw.Window, char uint) {
		win.CharDispatch.Dispatch(char)
	})

	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		for h, _ := range win.mouseClickHandlers {
			h.HandleMouseClick(MouseButtonState{MouseButton(button), Action(action), ModifierKey(mods)})
		}
	})

	window.SetCursorPositionCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		for h, _ := range win.mousePositionHandlers {
			h.HandleMousePosition(xpos, ypos)
		}
	})

	window.SetCursorEnterCallback(func(w *glfw.Window, entered bool) {
		for h, _ := range win.mouseEnterHandlers {
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

func (self *Window) Draw(x, y, w, h float64, ctx vg.Context) {
	fbw, fbh := self.FramebufferSize()

	// Calculate pixel ration for hi-dpi devices.
	gl.ClearColor(.87, .87, .87, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

	//Do OpenGL stuff
	// Iterate over scene, rendering what can be rendered
	if self.renderer != nil {
		self.renderer.Render()
	}

	////////////////////////////
	// gl.Viewport(0, 0, 50, 50)
	// gl.Begin(gl.TRIANGLES)
	// gl.Color3f(RR, 0.2, 0.3)
	// gl.Vertex3f(0, 0, 0)
	// gl.Vertex3f(1, 0, 0)
	// gl.Vertex3f(0, 1, 0)
	// gl.End()
	// gl.Rotatef(1, 0, 0, 1)
	///////////////////////////

	// 2D Display
	r := float64(fbw) / float64(w)
	gl.Viewport(0, 0, gl.Sizei(fbw), gl.Sizei(fbh))

	if self.child != nil {
		ctx.BeginFrame(int(w), int(h), r)
		self.child.Draw(0.0, 0.0, float64(w), float64(h), ctx)
		ctx.EndFrame()
	}
}

func (self *Window) Start() chan bool {
	e := gl.Init()
	if e != nil {
		panic(e)
	}

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
			self.Draw(0, 0, float64(w), float64(h), ctx)
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

func (self *Window) SetTitle(title string) {
	self.window.SetTitle(title)
}

func (self *Window) SetSize(w, h int) {
	self.window.SetSize(w, h)
}

func (self *Window) Size() (w, h int) {
	return self.window.GetSize()
}

func (self *Window) FramebufferSize() (w, h int) {
	return self.window.GetFramebufferSize()
}

func (self *Window) SetPosition(x, y int) {
	self.window.SetPosition(x, y)
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
}
