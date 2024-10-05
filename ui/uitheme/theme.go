package uitheme

import (
	"gio.tools/icons"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Theme struct {
	*material.Theme
	IsThemeDark bool

	BrandColorPalette   *BrandColorPaletteMap
	NeutralColorPalette *NeutralColorPaletteMap

	WidgetsColorMap *WidgetsColorMap
}

func New() *Theme {
	t := &Theme{
		Theme: &material.Theme{
			TextSize: unit.Sp(14),
		},
		IsThemeDark:         false,
		NeutralColorPalette: DefaultNeutralColorPalette,
		BrandColorPalette:   &BrandColorPaletteMap{},

		WidgetsColorMap: &WidgetsColorMap{},
	}

	t.Theme.Icon.RadioUnchecked = icons.ToggleRadioButtonUnchecked
	t.Theme.Icon.RadioChecked = icons.ToggleRadioButtonChecked
	t.Theme.Icon.CheckBoxUnchecked = icons.ToggleCheckBoxOutlineBlank
	t.Theme.Icon.CheckBoxChecked = icons.ToggleCheckBox

	t.BrandColorPalette = GenerateColorPaletteFromMainColor(GetNrgbaFromNumHex(0x0F6CBD))

	t.MapWidgetsColor()
	return t
}
func (t *Theme) GetMaterialTheme() *material.Theme {
	return t.Theme
}
