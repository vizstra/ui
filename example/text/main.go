package main

import (
	"github.com/vizstra/ui"
	"github.com/vizstra/ui/color"
	"github.com/vizstra/ui/layout"
	"github.com/vizstra/ui/text"
)

const TEXT = `Pea horseradish azuki bean lettuce avocado asparagus okra. Kohlrabi radish okra azuki bean corn fava bean mustard tigernut jÃ­cama green bean celtuce collard greens avocado quandong fennel gumbo black-eyed pea. Grape silver beet watercress potato tigernut corn groundnut. Chickweed okra pea winter purslane coriander yarrow sweet pepper radish garlic brussels sprout groundnut summer purslane earthnut pea tomato spring onion azuki bean gourd. Gumbo kakadu plum komatsuna black-eyed pea green bean zucchini gourd winter purslane silver beet rock melon radish asparagus spinach.`

func main() {
	window := ui.NewWindow("", "Text Example", 2650, 25, 1237, 580)
	table := layout.NewTable(window)
	table.SetCellMargin(ui.Margin{3, 3, 10, 3})
	table.SetRowHeight(0, 40)
	table.SetRowHeight(1, 400)
	table.SetRowHeight(2, 40)
	table.SetRowHeight(3, 80)

	addText := func(a text.Alignment, atxt string, c color.Color, col, row, w, h int) {
		table.SetColWidth(col, 300)
		txt := text.New(table, "", atxt)
		txt.Background = color.Gray10
		txt.Foreground = color.White
		txt.Alignment = text.CENTER
		table.AddMultiCell(txt, col, row-1, w, h)

		txt = text.New(table, "", TEXT)
		txt.Alignment = a
		txt.Foreground = c
		table.AddMultiCell(txt, col, row, w, h)

	}

	addText(text.LEFT, "Left", color.Gray13, 0, 1, 1, 1)
	addText(text.JUSTIFY, "Justify", color.Purple2, 1, 1, 1, 1)
	addText(text.CENTER, "Center", color.Green2, 2, 1, 1, 1)
	addText(text.RIGHT, "Right", color.Blue4, 3, 1, 1, 1)
	addText(text.NOWRAP, "No Wrapping", color.Red2, 0, 3, 4, 1)

	window.SetChild(table)
	end := window.Start()
	<-end
}
