package component

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/widgets"
)

type (
	LC  = layout.Context
	LD  = layout.Dimensions
	THM = *uitheme.Theme
)

type NaviMenuBar struct {
	List *widget.List

	MenuButtonsContent   []*widgets.MenuButtonContent
	MenuButtonStyle      *widgets.MenuButtonStyle
	MenuButtonsClickable []*widget.Clickable

	MenuButtonsComp []*widgets.MenuButton

	Style *NaviMenuBarStyle
}

type NaviMenuBarStyle struct {
	SpaceBetweenBtns int
	ButtonPadding    int
}

func NewNaviMenuBar(list *widget.List, style *NaviMenuBarStyle, theme THM) *NaviMenuBar {
	n := &NaviMenuBar{
		List:  list,
		Style: style,
	}
	n.MenuButtonsComp = n.makeMenuButtonsComp(theme)
	return n
}

func (n *NaviMenuBar) makeMenuButtonsComp(theme THM) []*widgets.MenuButton {
	n.MenuButtonsContent = []*widgets.MenuButtonContent{
		{IconName: "house", IconType: widgets.IconTypeFont, Text: "Home"},
		{IconName: "server", IconType: widgets.IconTypeFont, Text: "ServerList"},
		{IconName: "terminal", IconType: widgets.IconTypeFont, Text: "Console"},
		{IconName: "gear", IconType: widgets.IconTypeFont, Text: "Settings"},
		{IconName: "book", IconType: widgets.IconTypeFont, Text: "Help"},
		{IconName: "circle-info", IconType: widgets.IconTypeFont, Text: "About"},
	}
	n.MenuButtonStyle = &widgets.MenuButtonStyle{
		MinWidth: 190,

		TextSize:  20,
		TextColor: theme.WidgetsColorMap.ColorGeneralContrastFg,

		IconSize:             24,
		IconPosition:         widgets.ButtonIconStart,
		IconColor:            theme.WidgetsColorMap.ColorGeneralContrastFg,
		SpaceBetweenTextIcon: 8,

		ContentPadding:   10,
		ContentPosition:  widgets.ButtonContentEnd,
		ContentSpaceBias: 0.15,

		BgCornerRadius: 10,
		BgColor:        theme.WidgetsColorMap.ColorGeneralContrastBg,
	}
	n.MenuButtonsClickable = make([]*widget.Clickable, 0)
	for range n.MenuButtonsContent {
		n.MenuButtonsClickable = append(n.MenuButtonsClickable, &widget.Clickable{})
	}
	menuButtonComp := make([]*widgets.MenuButton, 0)
	for i := range n.MenuButtonsContent {
		menuButtonComp = append(menuButtonComp, widgets.NewMenuButtons(
			n.MenuButtonsClickable[i],
			n.MenuButtonsContent[i],
			n.MenuButtonStyle))
	}
	return menuButtonComp
}

func (n *NaviMenuBar) Layout(gtx LC, currSelectedPageIndex int, theme THM) LD {
	// 构造菜单栏按钮组件
	insetLayout := layout.Inset{Left: unit.Dp(n.Style.ButtonPadding), Right: unit.Dp(n.Style.ButtonPadding), Top: unit.Dp(n.Style.ButtonPadding)}
	return n.List.Layout(gtx, len(n.MenuButtonsComp), func(gtx LC, i int) LD {
		// 这里面是每个编号为i的item的layout, i从0开始到len()结束
		if i == currSelectedPageIndex {
			n.MenuButtonsComp[i].IsNavigated = true
		} else {
			n.MenuButtonsComp[i].IsNavigated = false
		}
		return layout.Inset{Bottom: unit.Dp(n.Style.SpaceBetweenBtns)}.Layout(gtx, func(gtx LC) LD {
			return insetLayout.Layout(gtx, func(gtx LC) LD {
				return n.MenuButtonsComp[i].Layout(gtx, theme)
			})
		})
	})
}
