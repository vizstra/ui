Vistra UI is an experimental OpenGL based GUI library for the Go programming language. It a personal project in the and should be viewed as such.  Although it is in the earliest stages of development, feel free to use it and submit pull requests as bugs are found or features created.  Areas I would like help with are a charting library, font fallback, glyph caching (maybe), and various other widgets (see [Semantic UI](http://semantic-ui.com/)). 

I chose OpenGL for rendering to have a less event driven UI experience and a simple frame based architecture.  There is a 3D render pass, followed by a vector graphics drawing pass.  Each frame renders itself as often as possible in a separate goroutine which is kicked off by the Window.Start() function.  You can currently set mouse event callbacks or request a mouse event channel.  Keyboard input will follow after the text rendering system is built.

To get the library for use:
```
go get github.com/vizstra/ui
```

Here is a simple example of using 
```Go
package main
import (
  "github.com/vizstra/ui"
  "github.com/vizstra/ui/button"
  "github.com/vizstra/ui/layout"
)

func main() {
  window := ui.NewWindow("", "Button Example", 240, 60, 1570, 60)
  fill := layout.NewFill(window)
  fill.SetMargin(ui.Margin{10, 10, 10, 10})
  b := button.NewButton(fill, "", "Click Here!")
  b.AddMousePositionCB(func(x, y float64) { })
  b.AddMouseClickCB(func(m ui.MouseButtonState) { })
  fill.SetChild(b)
  window.SetChild(fill)
  end := window.Start()
  <-end
}
```

######Developer Protip
Use rerun to automatically rebuild and test UI changes; I do.
```
go install https://github.com/skelterjohn/rerun
```

#### Current UI Goals
- [x] Button
- [ ] Text (This is a very large area currently being worked)
- [ ] Label
- [ ] Scrollbar
- [ ] List
- [ ] Radio
- [ ] Check
- [x] Progress Bar
- [ ] Image Button
- [ ] Tree
- [ ] Table
- [x] Line Chart
- [ ] Pie Chart

###### Layouts
- [x] Fill
- [x] Table
- [ ] Grid


