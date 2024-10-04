package widgets

import (
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

type CommonButton struct {
	Clickable *widget.Clickable

	Content *CommonButtonContent

	Style *CommonButtonStyle
}

type CommonButtonContent struct {
	Text string

	IconName   string
	IconWidget *widget.Icon
	IconType   int
}

type CommonButtonStyle struct {
	MinWidth int

	TextSize  int
	TextColor color.NRGBA

	IconSize     int
	IconColor    color.NRGBA
	IconPosition int

	SpaceBetweenTextIcon int // 文本和图标之间的距离

	ContentPadding   int     // Content(Text&Icon)的整体与背景边界的距离
	ContentPosition  int     // Content在背景的位置(左中右)
	ContentSpaceBias float32 // Content在背景位置的偏移量(仅当Position不为居中时有效)[0~1]

	BgColor        color.NRGBA
	BgCornerRadius int
}

func NewCommonButton(clickable *widget.Clickable, content *CommonButtonContent, style *CommonButtonStyle) *CommonButton {
	b := &CommonButton{
		Clickable: clickable,
		Content:   content,
		Style:     style,
	}
	return b
}
func (b *CommonButton) Layout(gtx LC, theme THM) LD {
	var largerSizeTextIcon int // 用于计算文本和图标哪个更大，更大的会被用来作为文本和图标垂直对齐的标准

	if b.Content.IconName != "" || b.Content.IconWidget != nil {
		if b.Style.IconSize >= b.Style.TextSize {
			largerSizeTextIcon = b.Style.IconSize
		} else {
			largerSizeTextIcon = b.Style.TextSize
		}
	}
	// 文本布局
	textLabelLayout := layout.Rigid(func(gtx LC) LD {
		gtx.Constraints.Min.Y = gtx.Sp(unit.Sp(largerSizeTextIcon)) //用来垂直布局,要用Flex.Alignment必须写这个
		return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx LC) LD {
				l := material.Label(theme.GetMaterialTheme(), unit.Sp(b.Style.TextSize), b.Content.Text)
				l.Color = b.Style.TextColor
				return l.Layout(gtx)
			}),
		)

	})

	textIconLayout := []layout.FlexChild{textLabelLayout} // 构造文本布局，若无图标函数结尾会直接返回这个
	textIconAxis := layout.Horizontal                     // 默认为水平

	// 如果有图标，则处理图标布局
	if b.Content.IconName != "" || b.Content.IconWidget != nil {
		iconInsetTop := unit.Dp(0)
		iconInsetBottom := unit.Dp(0)
		iconInsetLeft := unit.Dp(0)
		iconInsetRight := unit.Dp(0)

		switch b.Style.IconPosition { // 根据图标位置确定间隙应该插在哪个方位
		case ButtonIconTop:
			iconInsetBottom = unit.Dp(b.Style.SpaceBetweenTextIcon)
		case ButtonIconDown:
			iconInsetTop = unit.Dp(b.Style.SpaceBetweenTextIcon)
		case ButtonIconStart:
			iconInsetRight = unit.Dp(b.Style.SpaceBetweenTextIcon)
		case ButtonIconEnd:
			iconInsetLeft = unit.Dp(b.Style.SpaceBetweenTextIcon)
		}

		iconLayoutWithInsetBetweenText := layout.Inset{Top: iconInsetTop, Bottom: iconInsetBottom, Left: iconInsetLeft, Right: iconInsetRight}

		iconLayout := layout.Rigid(func(gtx LC) LD {
			gtx.Constraints.Min.Y = gtx.Sp(unit.Sp(largerSizeTextIcon)) // 用Flex.Alignment必须写这个
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
				layout.Rigid(func(gtx LC) LD {
					// 插入文本和图标间隙
					iconColor := b.Style.TextColor

					return iconLayoutWithInsetBetweenText.Layout(gtx, func(gtx LC) LD {
						switch b.Content.IconType {
						case IconTypeFont:
							return FontIcons(b.Content.IconName, b.Style.IconSize, iconColor, theme).Layout(gtx)
						case IconTypeVg:
							gtx.Constraints.Min.X = gtx.Sp(unit.Sp(b.Style.IconSize))
							return b.Content.IconWidget.Layout(gtx, iconColor)
						}
						panic("Invalid Icon type")
					})
				}),
			)
		})

		// 确定文本和图标的布局方向
		if b.Style.IconPosition == ButtonIconTop || b.Style.IconPosition == ButtonIconDown {
			textIconAxis = layout.Vertical
		}
		switch b.Style.IconPosition {
		case ButtonIconStart, ButtonIconTop:
			textIconLayout = []layout.FlexChild{iconLayout, textLabelLayout}
		case ButtonIconEnd, ButtonIconDown:
			textIconLayout = []layout.FlexChild{textLabelLayout, iconLayout}
		}
	}

	// 处理Content布局
	btnContentLayout := func(gtx LC) LD {
		gtx.Constraints.Min.X = gtx.Dp(unit.Dp(b.Style.MinWidth)) //设置最小宽度
		regionMinX := gtx.Constraints.Min.X
		contentFlexLayout := layout.Flex{Axis: textIconAxis, Alignment: layout.Middle}

		if b.Style.ContentPosition != ButtonContentMiddle {
			var biasLayout layout.Inset
			switch b.Style.ContentPosition {
			case ButtonContentStart:
				contentFlexLayout = layout.Flex{Axis: textIconAxis, Alignment: layout.Start}
				biasLayout = layout.Inset{Right: unit.Dp(float32(regionMinX) * b.Style.ContentSpaceBias)}
			case ButtonContentEnd:
				contentFlexLayout = layout.Flex{Axis: textIconAxis, Alignment: layout.End}
				biasLayout = layout.Inset{Left: unit.Dp(float32(regionMinX) * b.Style.ContentSpaceBias)}
			}

			// 带偏置的ContentLayout
			return biasLayout.Layout(gtx, func(gtx LC) LD {
				return contentFlexLayout.Layout(gtx, textIconLayout...)
			})
		}

		return contentFlexLayout.Layout(gtx, textIconLayout...)
	}

	btnBGLayout := func(gtx LC) LD {
		gtx.Constraints.Min.X = gtx.Dp(unit.Dp(b.Style.MinWidth)) //设置最小宽度
		// 创建一个圆角矩形剪辑
		defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, b.Style.BgCornerRadius).Push(gtx.Ops).Pop()
		semantic.Button.Add(gtx.Ops)

		// 确定背景色
		background := b.Style.BgColor

		// 用调制好的背景色填充圆角矩形背景剪辑
		paint.Fill(gtx.Ops, background)
		return LD{Size: gtx.Constraints.Min}
	}

	return b.Clickable.Layout(gtx, func(gtx LC) LD {
		return layout.Background{}.Layout(gtx,
			btnBGLayout,
			func(gtx LC) LD {
				// 使用间隙布局，在Content(Icon&Text)周围添加间隙
				return layout.UniformInset(unit.Dp(b.Style.ContentPadding)).Layout(gtx, btnContentLayout)
			},
		)
	})
}
