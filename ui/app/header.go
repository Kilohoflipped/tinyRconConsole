package app

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/widgets"
	"image"
)

type Headbar struct {
	materialTheme *material.Theme
	cache         *op.Ops

	themeSwitchState *widget.Bool
	themeSwitcher    material.SwitchStyle

	iconDarkMode  material.LabelStyle
	iconLightMode material.LabelStyle
}

func NewHeadbar(theme *uitheme.Theme) *Headbar {
	h := &Headbar{
		materialTheme: theme.Material(),
		cache:         new(op.Ops),
	}
	h.iconDarkMode = widgets.MaterialIcons("dark_mode", theme)
	h.iconLightMode = widgets.MaterialIcons("light_mode", theme)

	h.themeSwitcher = material.Switch(theme.Material(), h.themeSwitchState, "")
	h.themeSwitcher.Color.Enabled = theme.SwitchBgColor
	h.themeSwitcher.Color.Disabled = theme.Palette.Fg
	return h
}

func (h *Headbar) Layout(gtx layout.Context, theme *uitheme.Theme) layout.Dimensions {
	inset := layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(4)}
	// 加上分割线的Flex的子组件
	headbarContent := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return layout.Background{}.Layout(gtx,
			// 绘制背景
			func(gtx layout.Context) layout.Dimensions {
				// 如果背景不是暗色模式
				if theme.IsThemeDark() {
					defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
					paint.Fill(gtx.Ops, theme.SideBarBgColor)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			// 绘制背景上的组件
			func(gtx layout.Context) layout.Dimensions {
				//间隙布局，
				return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					//所有的组件的实际框
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: layout.SpaceBetween}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return material.H6(h.materialTheme, "RconConsole").Layout(gtx)
									})
								}),
							)
						}),
					)
				})
			},
		)
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		headbarContent,
		widgets.DrawLineFlex(theme.SeparatorColor, unit.Dp(1), unit.Dp(gtx.Constraints.Max.X)),
	)
}
