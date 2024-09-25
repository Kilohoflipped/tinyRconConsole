package home

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/widgets"
	"image"
)

type View struct {
	actionButtons          []*actionButton
	actionButtonsComp      []*widgets.FlatButton
	actionButtonsClickable []*widget.Clickable

	list *layout.List
}

func NewView(theme *uitheme.Theme) *View {
	v := &View{

		actionButtons: []*actionButton{
			{Text: "NewServer"},
		},
	}
	v.list = &layout.List{}
	v.actionButtonsClickable = make([]*widget.Clickable, 0)
	for range v.actionButtons {
		v.actionButtonsClickable = append(v.actionButtonsClickable, &widget.Clickable{})
	}
	v.makeActionButtons(theme)
	return v
}

func (v *View) makeActionButtons(theme *uitheme.Theme) {
	v.actionButtonsComp = make([]*widgets.FlatButton, 0)
	for i, b := range v.actionButtons {
		v.actionButtonsComp = append(v.actionButtonsComp, &widgets.FlatButton{
			Icon:            b.Icon,
			Text:            b.Text,
			Clickable:       v.actionButtonsClickable[i],
			BgTopPadding:    unit.Dp(4),
			BgBottomPadding: unit.Dp(0),
			BgLeftPadding:   unit.Dp(4),
			BgRightPadding:  unit.Dp(4),
			CornerRadius:    10,
			MinWidth:        unit.Dp(120),
			BackgroundColor: theme.SideBarBgColor,
			TextColor:       theme.SideBarTextColor,
			TextSize:        18,
			ContentPadding:  unit.Dp(9),
		})
	}
}

func (v *View) Layout(gtx layout.Context, theme *uitheme.Theme) layout.Dimensions {
	bgLayout := func(gtx layout.Context) layout.Dimensions {
		if theme.IsThemeDark() {
			defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, theme.SideBarBgColor)
		}
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}

	return layout.Background{}.Layout(gtx,

		// 绘制背景
		bgLayout,
		// 绘制背景上的组件
		func(gtx layout.Context) layout.Dimensions {
			// 页面四周留空
			return layout.UniformInset(unit.Dp(6)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				//功能按钮和区域说明
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return v.list.Layout(gtx, len(v.actionButtons), func(gtx layout.Context, i int) layout.Dimensions {
							btn := v.actionButtonsComp[i]
							return btn.Layout(gtx, theme)
						})
					}))
			})
		},
	)
}

type actionButton struct {
	Icon *widget.Icon
	Text string
}
