package widgets

import (
	"gio.tools/icons"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image/color"
)

func FontIcons(iconName string, iconSize int, iconColor color.NRGBA, theme THM) material.LabelStyle {
	if iconSize <= 0 {
		iconSize = 24
	}
	i := material.Label(theme.GetMaterialTheme(), unit.Sp(iconSize), "")
	i.Font.Typeface = "FontAwes6Solid"
	i.Text = iconName
	i.Color = iconColor
	return i
}
func GoVgIcons() *widget.Icon {
	return icons.AVAVTimer
}
