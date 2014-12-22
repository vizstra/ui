package chart

import (
	. "github.com/vizstra/ui/color"
)

const (
	CHART_BACKGROUND = 0x042f0173 + iota
)

func init() {
	Palette[CHART_BACKGROUND] = White
}
