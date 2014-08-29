package ui

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	// "image/color"
	"log"
	// "time"
)

var RR float32 = .0

type Window struct {
	name                  string
	endchan               chan bool
	window                *glfw.Window
	renderer              Renderer
	child                 View
	keyHandlers           map[KeyHandler]bool
	charHandlers          map[CharHandler]bool
	mouseButtonHandlers   map[MouseButtonHandler]bool
	mousePositionHandlers map[MousePositionHandler]bool
	mouseEnterHandlers    map[MouseEnterHandler]bool
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
		make(map[KeyHandler]bool, 0),
		make(map[CharHandler]bool, 0),
		make(map[MouseButtonHandler]bool, 0),
		make(map[MousePositionHandler]bool, 0),
		make(map[MouseEnterHandler]bool, 0),
	}

	window.SetPosition(x, y)
	window.MakeContextCurrent()

	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		for h, _ := range win.keyHandlers {
			h.HandleKey(Key(key), scancode, Action(action), ModifierKey(mods))
		}
	})

	window.SetCharacterCallback(func(w *glfw.Window, char uint) {
		for h, _ := range win.charHandlers {
			h.HandleChar(char)
		}
	})

	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		for h, _ := range win.mouseButtonHandlers {
			h.HandleMouseButton(MouseButton(button), Action(action), ModifierKey(mods))
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

func (self *Window) Start() chan bool {
	// This is the ever important draw goroutine
	go func() {
		vg := NewContext()
		for !self.window.ShouldClose() {
			w, h := self.Size()
			fbw, fbh := self.FramebufferSize()
			// Calculate pixel ration for hi-dpi devices.
			gl.ClearColor(.88, .9, .88, 1)
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

			//Do OpenGL stuff
			// Iterate over scene, rendering what can be rendered
			if self.renderer != nil {
				self.renderer.Render()
			}

			// gl.Viewport(0, 0, 50, 50)
			gl.Begin(gl.TRIANGLES)
			gl.Color3f(RR, 0.2, 0.3)
			gl.Vertex3f(0, 0, 0)
			gl.Vertex3f(1, 0, 0)
			gl.Vertex3f(0, 1, 0)
			gl.End()

			gl.Rotatef(1, 0, 0, 1)

			// 2D Display
			r := float64(fbw) / float64(w)
			gl.Viewport(0, 0, fbw, fbh)

			if self.child != nil {
				vg.BeginFrame(w, h, r)

				self.child.Draw(0.0, 0.0, float64(w), float64(h), vg)
				vg.EndFrame()
			}

			self.window.SwapBuffers()

			glfw.PollEvents()
		}
		glfw.Terminate()
		self.endchan <- true
	}()
	return self.endchan
}

func (self *Window) SetChild(child View) {
	self.child = child
	if h, ok := child.(KeyHandler); ok {
		self.AddKeyHandler(h)
	}

	if h, ok := child.(CharHandler); ok {
		self.AddCharHandler(h)
	}

	if h, ok := child.(MouseButtonHandler); ok {
		self.AddMouseButtonHandler(h)
	}

	if h, ok := child.(MousePositionHandler); ok {
		self.AddMousePositionHandler(h)
	}

	if h, ok := child.(MouseEnterHandler); ok {
		self.AddMouseEnterHandler(h)
	}
}

func (self *Window) Child() View {
	return self.child
}

func (self *Window) SetRenderer(renderer Renderer) {
	self.renderer = renderer
}

func (self *Window) Renderer() Renderer {
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

func (self *Window) AddKeyHandler(h KeyHandler) {
	if _, ok := self.keyHandlers[h]; !ok {
		self.keyHandlers[h] = true
	}
}

func (self *Window) AddCharHandler(h CharHandler) {
	if _, ok := self.charHandlers[h]; !ok {
		self.charHandlers[h] = true
	}
}

func (self *Window) AddMouseButtonHandler(h MouseButtonHandler) {
	if _, ok := self.mouseButtonHandlers[h]; !ok {
		self.mouseButtonHandlers[h] = true
	}
}

func (self *Window) AddMousePositionHandler(h MousePositionHandler) {
	if _, ok := self.mousePositionHandlers[h]; !ok {
		self.mousePositionHandlers[h] = true
	}
}

func (self *Window) AddMouseEnterHandler(h MouseEnterHandler) {
	if _, ok := self.mouseEnterHandlers[h]; !ok {
		self.mouseEnterHandlers[h] = true
	}
}

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Printf("%v: %v\n", err, desc)
}

func init() {
	if !glfw.Init() {
		panic("Cannot initialize windowing library.")
	}
}
