package widgets

import (
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
)

func MaterialIcons(name string, theme *uitheme.Theme) material.LabelStyle {
	l := material.Label(theme.Material(), unit.Sp(24), "")
	l.Font.Typeface = "MaterialIcons"
	l.Text = name
	return l
}
