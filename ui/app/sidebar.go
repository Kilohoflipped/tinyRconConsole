package app

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/widgets"
	"image"
)

var selectedPageIndexBuffer int = -1

type Sidebar struct {
	Theme *uitheme.Theme
	// 按钮组件
	naviButtonsInfo      []*naviButtonInfo
	naviButtonsComp      []*widgets.FlatButton
	naviButtonsClickable []*widget.Clickable

	list *widget.List

	cache *op.Ops

	// 当前选择的页面
	selectedPageIndex int
}

func NewSidebar(theme *uitheme.Theme) *Sidebar {
	s := &Sidebar{
		Theme: theme,
		cache: new(op.Ops),

		naviButtonsInfo: []*naviButtonInfo{
			{Text: "Home"},
			{Text: "ServerList"},
			{Text: "Console"},
			{Text: "Settings"},
			{Text: "Help"},
			{Text: "About"},
		},

		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}

	s.naviButtonsClickable = make([]*widget.Clickable, 0)
	for range s.naviButtonsInfo {
		s.naviButtonsClickable = append(s.naviButtonsClickable, &widget.Clickable{})
	}
	s.makeButtons(theme)

	return s
}

func (s *Sidebar) makeButtons(theme *uitheme.Theme) {
	s.naviButtonsComp = make([]*widgets.FlatButton, 0)
	for i, b := range s.naviButtonsInfo {
		s.naviButtonsComp = append(s.naviButtonsComp, &widgets.FlatButton{
			Icon: b.Icon,
			Text: b.Text,
			//IconPosition:      widgets.FlatButtonIconTop,
			Clickable: s.naviButtonsClickable[i],
			//SpaceBetween:      unit.Dp(4),
			BgTopPadding:    unit.Dp(4),
			BgBottomPadding: unit.Dp(0),
			BgLeftPadding:   unit.Dp(4),
			BgRightPadding:  unit.Dp(4),
			CornerRadius:    10,
			MinWidth:        unit.Dp(240),
			BackgroundColor: theme.SideBarBgColor,
			TextColor:       theme.SideBarTextColor,
			TextSize:        20,
			ContentPadding:  unit.Dp(13),
		})
	}
}

// 获得被点击的导航按钮的索引
func (s *Sidebar) getClickedNaviBtnIndex(gtx layout.Context) int {
	for i, c := range s.naviButtonsClickable {
		if c.Clicked(gtx) {
			return i
		}
	}
	return -1 // 没有按钮被点击则返回-1
}

// 当有导航按钮被点击后,更新当前正浏览的页面索引
func (s *Sidebar) updateSelectedPageIndex(gtx layout.Context) {
	clickedIndex := s.getClickedNaviBtnIndex(gtx)
	if clickedIndex != -1 {
		selectedPageIndexBuffer = clickedIndex
	}
	s.selectedPageIndex = selectedPageIndexBuffer
}

func (s *Sidebar) Layout(gtx LC, theme THM) LD {
	s.updateSelectedPageIndex(gtx)

	bgLayout := func(gtx LC) LD {
		if theme.IsThemeDark() {
			defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, theme.SideBarBgColor)
		}
		return LD{Size: gtx.Constraints.Min}
	}

	naviBtnsLayout := layout.Rigid(func(gtx LC) LD {
		return s.list.Layout(gtx, len(s.naviButtonsInfo), func(gtx LC, i int) LD {
			btn := s.naviButtonsComp[i]
			if s.selectedPageIndex == i {
				btn.TextColor = theme.SideBarTextColor
			} else {
				btn.TextColor = widgets.Disabled(theme.SideBarTextColor)
			}
			return btn.Layout(gtx, theme)
		})
	})

	divideLineLayout := layout.Rigid(func(gtx LC) LD {
		if !theme.IsThemeDark() {
			return LD{}
		}
		return widgets.DrawLine(gtx, theme.SeparatorColor, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
	})

	// 背景布局
	return layout.Background{}.Layout(gtx,
		bgLayout,
		func(gtx LC) LD {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				naviBtnsLayout,
				divideLineLayout,
			)
		},
	)
}

type naviButtonInfo struct {
	Icon *widget.Icon
	Text string
}
