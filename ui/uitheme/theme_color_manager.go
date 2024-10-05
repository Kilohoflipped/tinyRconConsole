package uitheme

import "image/color"

const (
	StateThemeDark  = 0
	StateThemeLight = 1
)

type WidgetsColorMap struct {
	ColorGeneralFg         color.NRGBA
	ColorGeneralBg         color.NRGBA
	ColorGeneralContrastFg color.NRGBA
	ColorGeneralContrastBg color.NRGBA

	ColorThemeSeparator color.NRGBA
}

type NeutralColorPaletteMap struct {
	gray1  color.NRGBA
	gray2  color.NRGBA
	gray3  color.NRGBA
	gray4  color.NRGBA
	gray5  color.NRGBA
	gray6  color.NRGBA
	gray7  color.NRGBA
	gray8  color.NRGBA
	gray9  color.NRGBA
	gray10 color.NRGBA
	gray11 color.NRGBA
	gray12 color.NRGBA
	gray13 color.NRGBA
	gray14 color.NRGBA
	gray15 color.NRGBA
	gray16 color.NRGBA
}

type BrandColorPaletteMap struct {
	BrandMainColor color.NRGBA
	BrandColors    [16]color.NRGBA
}

var DefaultNeutralColorPalette = &NeutralColorPaletteMap{
	gray1:  GetNrgbaFromNumHex(0x000000),
	gray2:  GetNrgbaFromNumHex(0x050505),
	gray3:  GetNrgbaFromNumHex(0x1a1a1a),
	gray4:  GetNrgbaFromNumHex(0x242424),
	gray5:  GetNrgbaFromNumHex(0x2e2e2e),
	gray6:  GetNrgbaFromNumHex(0x424242),
	gray7:  GetNrgbaFromNumHex(0x575757),
	gray8:  GetNrgbaFromNumHex(0x6b6b6b),
	gray9:  GetNrgbaFromNumHex(0x808080),
	gray10: GetNrgbaFromNumHex(0x949494),
	gray11: GetNrgbaFromNumHex(0xa8a8a8),
	gray12: GetNrgbaFromNumHex(0xbdbdbd),
	gray13: GetNrgbaFromNumHex(0xd1d1d1),
	gray14: GetNrgbaFromNumHex(0xe6e6e6),
	gray15: GetNrgbaFromNumHex(0xfafafa),
	gray16: GetNrgbaFromNumHex(0xffffff),
}

func (t *Theme) SetThemeState(themeState int) {
	switch themeState {
	case StateThemeDark:
		t.IsThemeDark = true
	case StateThemeLight:
		t.IsThemeDark = false
	}
	t.MapWidgetsColor()
}

func (t *Theme) MapWidgetsColor() {
	if t.IsThemeDark {
		t.MapWidgetsColorDark()
	} else {
		t.MapWidgetsColorLight()
	}
}

func (t *Theme) MapWidgetsColorLight() {
	tncp := t.NeutralColorPalette
	tbcp := t.BrandColorPalette
	t.Theme.Palette.Fg = tncp.gray4
	t.Theme.Palette.Bg = tncp.gray16
	t.Theme.Palette.ContrastFg = tbcp.BrandMainColor
	t.Theme.Palette.ContrastBg = tbcp.BrandColors[16-1]

	t.WidgetsColorMap.ColorGeneralFg = tncp.gray4
	t.WidgetsColorMap.ColorGeneralBg = tncp.gray16
	t.WidgetsColorMap.ColorGeneralContrastFg = tbcp.BrandMainColor
	t.WidgetsColorMap.ColorGeneralContrastBg = tbcp.BrandColors[16-1]

	t.WidgetsColorMap.ColorThemeSeparator = tncp.gray16
}

func (t *Theme) MapWidgetsColorDark() {
	tncp := t.NeutralColorPalette
	tbcp := t.BrandColorPalette

	t.Theme.Palette.Fg = tncp.gray16
	t.Theme.Palette.Bg = tncp.gray4
	t.Theme.Palette.ContrastFg = tbcp.BrandColors[10-1]
	t.Theme.Palette.ContrastBg = tbcp.BrandColors[3-1]

	t.WidgetsColorMap.ColorGeneralFg = tncp.gray16
	t.WidgetsColorMap.ColorGeneralBg = tncp.gray4
	t.WidgetsColorMap.ColorGeneralContrastFg = tbcp.BrandColors[10-1]
	t.WidgetsColorMap.ColorGeneralContrastBg = tbcp.BrandColors[3-1]

	t.WidgetsColorMap.ColorThemeSeparator = tncp.gray7
}
