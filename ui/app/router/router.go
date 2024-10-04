package router

import (
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/app/router/component"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/widgets"
	"image"
)

type Router struct {
	Theme THM

	NaviMenuBarList  *widget.List
	NaviMenuBarStyle *component.NaviMenuBarStyle
	NaviMenuBarComp  *component.NaviMenuBar

	cache *op.Ops

	// 当前选择的页面
	CurrSelectedPageIndex int
}

func NewRouter(theme THM) *Router {
	r := &Router{
		Theme:                 theme,
		CurrSelectedPageIndex: -1,
	}
	r.NaviMenuBarComp = r.makeNaviMenuBarComp(theme)
	return r
}

func (r *Router) makeNaviMenuBarComp(theme THM) *component.NaviMenuBar {
	r.NaviMenuBarList = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}
	r.NaviMenuBarStyle = &component.NaviMenuBarStyle{
		SpaceBetweenBtns: 0,
		ButtonPadding:    6,
	}
	return component.NewNaviMenuBar(r.NaviMenuBarList, r.NaviMenuBarStyle, theme)
}

func (r *Router) Layout(gtx LC, theme THM) LD {
	r.updateSelectedPageIndex(gtx)

	bgLayout := func(gtx LC) LD {
		defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, theme.WidgetsColorMap.ColorGeneralBg)
		return LD{Size: gtx.Constraints.Min}
	}

	naviMenuBarLayout := layout.Rigid(func(gtx LC) LD {
		return r.NaviMenuBarComp.Layout(gtx, r.CurrSelectedPageIndex, theme)
	})

	divideLineLayout := layout.Rigid(func(gtx LC) LD {
		if !theme.IsThemeDark {
			return LD{}
		}
		return widgets.DrawLine(gtx, theme.WidgetsColorMap.ColorThemeSeparator, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
	})

	return layout.Background{}.Layout(gtx,
		bgLayout,
		func(gtx LC) LD {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				naviMenuBarLayout,
				divideLineLayout,
			)
		},
	)
}
