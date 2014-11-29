package color

import ()

const (
	White = iota
	Gray1
	Gray2
	Gray3
	Gray4
	Gray5
	Gray6
	Gray7
	Gray8
	Gray9
	Gray10
	Gray11
	Gray12
	Gray13
	Black
	Blue1
	Blue2
	Blue3
	Blue4
	Blue5
	Cyan1
	Cyan2
	Teal1
	Teal2
	Green1
	Green2
	Yellow1
	Yellow2
	Orange1
	Orange2
	Red1
	Red2
	Purple1
	Purple2
	Pink1
	Pink2
	Selection
)

var DefaultPalette map[int]Color
var Palette map[int]Color

func init() {
	DefaultPalette = make(map[int]Color)
	Palette = DefaultPalette

	DefaultPalette[Selection] = HexRGBA(0xFF0000FF)
	DefaultPalette[White] = HexRGBA(0xFFFFFFFF)
	DefaultPalette[Gray1] = HexRGBA(0x8d8d8dFF)
	DefaultPalette[Gray2] = HexRGBA(0x858585FF)
	DefaultPalette[Gray3] = HexRGBA(0x7d7d7dFF)
	DefaultPalette[Gray4] = HexRGBA(0x757575FF)
	DefaultPalette[Gray5] = HexRGBA(0x6d6d6dFF)
	DefaultPalette[Gray6] = HexRGBA(0x656565FF)
	DefaultPalette[Gray7] = HexRGBA(0x5d5d5dFF)
	DefaultPalette[Gray8] = HexRGBA(0x555555FF)
	DefaultPalette[Gray9] = HexRGBA(0x4d4d4dFF)
	DefaultPalette[Gray10] = HexRGBA(0x454545FF)
	DefaultPalette[Gray11] = HexRGBA(0x3d3d3dFF)
	DefaultPalette[Gray12] = HexRGBA(0x353535FF)
	DefaultPalette[Gray13] = HexRGBA(0x2d2d2dFF)
	DefaultPalette[Black] = HexRGBA(0x050505FF)
	DefaultPalette[Blue1] = HexRGBA(0xCDDBECFF)
	DefaultPalette[Blue2] = HexRGBA(0xA3C3ECFF)
	DefaultPalette[Blue3] = HexRGBA(0x5D9CECFF)
	DefaultPalette[Blue4] = HexRGBA(0x4A89DCFF)
	DefaultPalette[Blue5] = HexRGBA(0x2B486CFF)
	DefaultPalette[Cyan1] = HexRGBA(0x4FC1E9FF)
	DefaultPalette[Cyan2] = HexRGBA(0x3BAFDAFF)
	DefaultPalette[Teal1] = HexRGBA(0x48CFADFF)
	DefaultPalette[Teal2] = HexRGBA(0x37BC9BFF)
	DefaultPalette[Green1] = HexRGBA(0xA0D468FF)
	DefaultPalette[Green2] = HexRGBA(0x8CC152FF)
	DefaultPalette[Yellow1] = HexRGBA(0xFFCE54FF)
	DefaultPalette[Yellow2] = HexRGBA(0xF6BB42FF)
	DefaultPalette[Orange1] = HexRGBA(0xFC6E51FF)
	DefaultPalette[Orange2] = HexRGBA(0xE9573FFF)
	DefaultPalette[Red1] = HexRGBA(0xED5565FF)
	DefaultPalette[Red2] = HexRGBA(0xDA4453FF)
	DefaultPalette[Purple1] = HexRGBA(0xAC92ECFF)
	DefaultPalette[Purple2] = HexRGBA(0x967ADCFF)
	DefaultPalette[Pink1] = HexRGBA(0xEC87C0FF)
	DefaultPalette[Pink2] = HexRGBA(0xD770ADFF)
}
