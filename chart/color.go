package chart

import (
	. "github.com/vizstra/ui/color"
)

const (
	CHART_BACKGROUND = 0x042f0173 + iota
)

func init() {
	DefaultPalette[CHART_BACKGROUND] = DefaultPalette[White]
}
