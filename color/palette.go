package color

import ()

var (
	White     Color = HexRGBA(0xFFFFFFFF)
	Gray1     Color = HexRGBA(0x8d8d8dFF)
	Gray2     Color = HexRGBA(0x858585FF)
	Gray3     Color = HexRGBA(0x7d7d7dFF)
	Gray4     Color = HexRGBA(0x757575FF)
	Gray5     Color = HexRGBA(0x6d6d6dFF)
	Gray6     Color = HexRGBA(0x656565FF)
	Gray7     Color = HexRGBA(0x5d5d5dFF)
	Gray8     Color = HexRGBA(0x555555FF)
	Gray9     Color = HexRGBA(0x4d4d4dFF)
	Gray10    Color = HexRGBA(0x454545FF)
	Gray11    Color = HexRGBA(0x3d3d3dFF)
	Gray12    Color = HexRGBA(0x353535FF)
	Gray13    Color = HexRGBA(0x2d2d2dFF)
	Black     Color = HexRGBA(0x050505FF)
	Blue1     Color = HexRGBA(0xCDDBECFF)
	Blue2     Color = HexRGBA(0xA3C3ECFF)
	Blue3     Color = HexRGBA(0x5D9CECFF)
	Blue4     Color = HexRGBA(0x4A89DCFF)
	Blue5     Color = HexRGBA(0x2B486CFF)
	Cyan1     Color = HexRGBA(0x4FC1E9FF)
	Cyan2     Color = HexRGBA(0x3BAFDAFF)
	Teal1     Color = HexRGBA(0x48CFADFF)
	Teal2     Color = HexRGBA(0x37BC9BFF)
	Green1    Color = HexRGBA(0xA0D468FF)
	Green2    Color = HexRGBA(0x8CC152FF)
	Yellow1   Color = HexRGBA(0xFFCE54FF)
	Yellow2   Color = HexRGBA(0xF6BB42FF)
	Orange1   Color = HexRGBA(0xFC6E51FF)
	Orange2   Color = HexRGBA(0xE9573FFF)
	Red1      Color = HexRGBA(0xED5565FF)
	Red2      Color = HexRGBA(0xDA4453FF)
	Purple1   Color = HexRGBA(0xAC92ECFF)
	Purple2   Color = HexRGBA(0x967ADCFF)
	Pink1     Color = HexRGBA(0xAC92ECFF)
	Pink2     Color = HexRGBA(0xD770ADFF)
	Selection Color = HexRGBA(0xFF0000FF)
)

var Palette map[int]Color

func init() {
	Palette = make(map[int]Color)
}
