package uitheme

import (
	"gio.tools/icons"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Theme struct {
	*material.Theme
	IsThemeDark bool

	brandColorPalette   *BrandColorPaletteMap
	neutralColorPalette *NeutralColorPaletteMap

	WidgetsColorMap *WidgetsColorMap
}

func New() *Theme {
	t := &Theme{
		Theme: &material.Theme{
			TextSize: unit.Sp(14),
		},
		IsThemeDark:         false,
		neutralColorPalette: DefaultNeutralColorPalette,
		brandColorPalette:   &BrandColorPaletteMap{},

		WidgetsColorMap: &WidgetsColorMap{},
	}

	t.Theme.Icon.RadioUnchecked = icons.ToggleRadioButtonUnchecked
	t.Theme.Icon.RadioChecked = icons.ToggleRadioButtonChecked
	t.Theme.Icon.CheckBoxUnchecked = icons.ToggleCheckBoxOutlineBlank
	t.Theme.Icon.CheckBoxChecked = icons.ToggleCheckBox

	t.brandColorPalette.GenerateColorPaletteFromMainColor(GetNrgbaFromNumHex(0x13F694))

	t.MapWidgetsColor()
	return t
}
func (t *Theme) GetMaterialTheme() *material.Theme {
	return t.Theme
}
