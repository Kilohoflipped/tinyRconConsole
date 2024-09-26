package widgets

import (
	"gioui.org/io/input"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"image"
	"image/color"
)

type FlatButton struct {
	Icon         *widget.Icon
	IconPosition int
	SpaceBetween unit.Dp

	Clickable *widget.Clickable

	MinWidth        unit.Dp
	BackgroundColor color.NRGBA
	TextColor       color.NRGBA
	TextSize        int
	Text            string

	CornerRadius    int
	BgTopPadding    unit.Dp
	BgBottomPadding unit.Dp
	BgLeftPadding   unit.Dp
	BgRightPadding  unit.Dp
	ContentPadding  unit.Dp
}

func (f *FlatButton) Layout(gtx layout.Context, theme *uitheme.Theme) layout.Dimensions {

	textLabelLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		l := material.Label(theme.Material(), unit.Sp(f.TextSize), f.Text)

		l.Color = f.TextColor
		return l.Layout(gtx)
	})

	//创建可点击区域
	return f.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// 使用插入间隙布局，使得按钮整体边缘与外框有一定的间隙
		return layout.Inset{Left: f.BgLeftPadding, Right: f.BgRightPadding, Top: f.BgTopPadding, Bottom: f.BgBottomPadding}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				// 所有Insert的间隙内的部分都可以触发Button
				semantic.Button.Add(gtx.Ops)

				// 使用背景布局，主要是实现Hover交互和选中交互时按钮背景的变换
				return layout.Background{}.Layout(gtx,
					// 创建圆角矩形背景
					func(gtx layout.Context) layout.Dimensions {
						// 绘制边界
						gtx.Constraints.Min.X = gtx.Dp(f.MinWidth)
						// 创建一个圆角矩形剪辑
						defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, f.CornerRadius).Push(gtx.Ops).Pop()
						// 根据按钮状态调整按钮没被选中/被选中/被悬停时的颜色
						background := f.BackgroundColor
						if gtx.Source == (input.Source{}) {
							background = Disabled(f.BackgroundColor)
						} else if f.Clickable.Hovered() || gtx.Focused(f.Clickable) {
							background = Hovered(f.BackgroundColor)
						}
						// 用调制好的背景色填充圆角矩形背景剪辑
						paint.Fill(gtx.Ops, background)
						return layout.Dimensions{Size: gtx.Constraints.Min}
					},

					// 绘制文本
					func(gtx layout.Context) layout.Dimensions {
						// 使用间隙布局，在文本周围添加间隙
						return layout.UniformInset(f.ContentPadding).Layout(gtx,
							func(gtx layout.Context) layout.Dimensions {
								// 添加文本
								return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: layout.SpaceBetween}.Layout(gtx, textLabelLayout)
							})
					},
				)
			},
		)
	})
}
