Vistra UI is an experimental OpenGL based GUI library for the Go programming language. It a personal project and should be viewed as such.  Although it is in the earliest stages of development, feel free to use it and submit pull requests as bugs are found or features created.  Areas I would like help with are:

- testing
- documentation
- architecture
- charts and data visualizations
- font fallbacks for unicode
- optimization
- nanovg porting?
- automated testing framework
- other ui elements (see [Semantic UI](http://semantic-ui.com/)) for ideas

I chose OpenGL (NanoVG) for rendering to have a less event driven UI experience and a simple frame based architecture.  There is a 3D render pass, followed by a vector graphics drawing pass.  Each frame renders itself as often as possible in a separate goroutine which is kicked off by the Window.Start() function.  You can currently set mouse event callbacks or request a mouse event channel.  Keyboard input will follow after the text rendering system is built.

To get the library for use:
```
go get github.com/vizstra/ui
```

Here is an example of a simple window with a single button:
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

#####Developer Protip
Use rerun to automatically rebuild and test UI changes; I do.
```
go install https://github.com/skelterjohn/rerun
```

#####Screenshots
A couple notes about these screenshots, I am working on a hi resolution 4K monitor and didn't bother to look at these images on other platforms or modify them in anyway.  Theming is not yet supported, and I have not begun to explore that aspect of the architecture yet, so I wouldn't say things are in their final form:

###### Text Rendering
<img src=https://raw.githubusercontent.com/vizstra/ui/master/res/img/screenshots/text.png>

###### Buttons
<img src=https://raw.githubusercontent.com/vizstra/ui/master/res/img/screenshots/calculator.png>

###### Charts
<img src=https://raw.githubusercontent.com/vizstra/ui/master/res/img/screenshots/chart1.png>