package main

import (
	"fmt"
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/button"
	"github.com/vizstra/ui/layout"
	"github.com/vizstra/vg"
	"strconv"
)

const (
	NONE = iota
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
)

func main() {
	window := ui.NewWindow("", "Calculator", 1440, 0, 350, 375)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{30, 30, 30, 30})

	table := layout.NewTable(fill)
	table.SetDefaultCellDimensions(50, 50)
	table.SetCellMargin(ui.Margin{2, 2, 2, 2})

	fill.SetChild(table)
	window.SetChild(fill)

	makebutton := func(text string, c, r, w, h int, cb func(ui.MouseButtonState)) *button.Button {
		b := button.NewButton(table, text, text)
		b.AddMouseClickCB(cb)
		table.AddMultiCell(b, c, r, w, h)
		return b
	}
	buffer := 0.0
	action := NONE
	bartext := NewBartext(table, "", "")
	table.AddMultiCell(bartext, 0, 0, 4, 1)

	makebutton("7", 0, 2, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "7"
		}
	})
	makebutton("4", 0, 3, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "4"
		}
	})
	makebutton("1", 0, 4, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "1"
		}
	})

	makebutton("8", 1, 2, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "8"
		}
	})
	makebutton("5", 1, 3, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "5"
		}
	})
	makebutton("2", 1, 4, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "2"
		}
	})

	makebutton("9", 2, 2, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "9"
		}
	})

	makebutton("6", 2, 3, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "6"
		}
	})

	makebutton("3", 2, 4, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "3"
		}
	})

	makebutton("Ce", 3, 2, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text = ""
			buffer = 0.0
		}
	})
	makebutton("÷", 3, 1, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "0"
		}
	})
	makebutton("x", 2, 1, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			buffer, _ = strconv.ParseFloat(bartext.Text, 64)
			bartext.Text = ""
		}
	})

	makebutton("+", 0, 1, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			b, err := strconv.ParseFloat(bartext.Text, 64)
			fmt.Println(b, buffer)
			text := ""
			if err != nil {
				text = "Malformed Number"
			} else if action == NONE {
				text = ""
			} else if action == ADD {

				text = fmt.Sprintf("%f", b+buffer)
			} else {
				buffer = b
			}
			bartext.Text = text
			action = ADD
		}
	})
	makebutton("-", 1, 1, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			buffer, _ = strconv.ParseFloat(bartext.Text, 64)
			bartext.Text = ""
		}
	})
	makebutton("=", 3, 3, 1, 3, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			buffer, _ = strconv.ParseFloat(bartext.Text, 64)
			bartext.Text = ""
		}
	})

	makebutton("0", 0, 5, 2, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			bartext.Text += "0"
		}
	})

	makebutton(".", 2, 5, 1, 1, func(state ui.MouseButtonState) {
		if state.Action == ui.RELEASE {
			// TODO: Makes invalid #s. Fix.
			bartext.Text += "."
		}
	})

	end := window.Start()
	<-end
}

type bartext struct {
	ui.Element
	Text string
}

func NewBartext(parent ui.Drawer, name, text string) *bartext {
	self := &bartext{
		ui.NewElement(parent, name),
		text,
	}

	self.DrawCB = func(x, y, w, h float64, ctx vg.Context) {
		ctx.Scissor(x, y, w, h)
		ctx.FillColor(self.Foreground)
		ctx.TextAlign(vg.ALIGN_CENTER | vg.ALIGN_RIGHT)

		ctx.SetFontSize(25)
		ctx.FindFont(vg.FONT_DEFAULT)
		ctx.Text(x+w, y+h/2, self.Text)
		ctx.ResetScissor()
	}

	return self
}
