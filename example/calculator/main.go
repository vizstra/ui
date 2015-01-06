package main

import (
	"fmt"
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/button"
	"github.com/vizstra/ui/color"
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
	window := ui.NewWindow("", "Calculator", 1940, 0, 228, 335)
	fill := layout.NewFill(window)
	fill.SetMargin(ui.Margin{10, 10, 10, 10})

	table := layout.NewTable(fill)
	table.SetDefaultCellDimensions(50, 50)
	table.SetCellMargin(ui.Margin{3, 3, 3, 3})

	fill.SetChild(table)
	window.SetChild(fill)

	buffer := 0.0
	action := NONE
	bartext := NewBartext(table, "", "")

	makebutton := func(text string, c, r, w, h int, cb func(ui.MouseButtonState)) *button.Button {
		b := button.NewButton(table, text, text)
		b.AddMouseClickCB(cb)
		table.AddMultiCell(b, c, r, w, h)
		return b
	}

	makeDefaultButton := func(text string, c, r, w, h int) *button.Button {
		return makebutton(text, c, r, w, h, func(state ui.MouseButtonState) {
			if state.Action == ui.RELEASE {
				bartext.Text += text
			}
		})
	}

	table.AddMultiCell(bartext, 0, 0, 4, 1)

	makeDefaultButton("7", 0, 2, 1, 1)
	makeDefaultButton("4", 0, 3, 1, 1)
	makeDefaultButton("1", 0, 4, 1, 1)
	makeDefaultButton("0", 0, 5, 2, 1)
	makeDefaultButton("8", 1, 2, 1, 1)
	makeDefaultButton("5", 1, 3, 1, 1)
	makeDefaultButton("2", 1, 4, 1, 1)
	makeDefaultButton("9", 2, 2, 1, 1)
	makeDefaultButton("6", 2, 3, 1, 1)
	makeDefaultButton("3", 2, 4, 1, 1)
	makeDefaultButton(".", 2, 5, 1, 1)

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

	self.SetCommonBackground(color.Purple2)

	self.DrawCB = func(ctx vg.Context) {
		x, y, w, h := self.Bounds()
		ctx.Scissor(x, y, w, h)
		ctx.FillColor(self.Foreground)
		ctx.SetFontSize(40)
		ctx.FindFont(vg.FONT_DEFAULT)
		ctx.TextAlign(vg.ALIGN_CENTER | vg.ALIGN_RIGHT)
		ctx.Text(x+w, y+h/1.5, self.Text)
		ctx.ResetScissor()
	}

	return self
}
