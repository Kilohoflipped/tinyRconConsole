package app

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/pages/home"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
)

type (
	LC  = layout.Context
	LD  = layout.Dimensions
	THM = *uitheme.Theme
)

type UI struct {
	window *app.Window
	Theme  *uitheme.Theme

	sidebar *Sidebar
	headbar *Headbar

	homeView *home.View
}

func NewUi(w *app.Window) (*UI, error) {
	// 创建mainUI对象
	u := &UI{
		window: w,
	}

	gioTheme := material.NewTheme()
	//fontCollection := fonts.AllFontFaces
	//theme.Shaper = text.NewShaper(text.WithCollection(fontCollection))
	u.Theme = uitheme.New(gioTheme, true)

	u.sidebar = NewSidebar(u.Theme)
	u.headbar = NewHeadbar(u.Theme)

	u.homeView = home.NewView(u.Theme)

	return u, nil
}

// Run 主窗体
func (u *UI) Run() error {
	var ops op.Ops
	for {
		switch e := u.window.Event().(type) {

		case app.DestroyEvent:
			// 用户关闭窗体
			return e.Err

		case app.FrameEvent:
			// 帧事件:正常运行
			gtx := app.NewContext(&ops, e)
			// 渲染UI
			u.Layout(gtx)
			// 最终在 GPU 上启动绘制。
			e.Frame(gtx.Ops)
		}
	}
}

func (u *UI) Layout(gtx LC) LD {

	headbarLayout := layout.Rigid(func(gtx LC) LD {
		return u.headbar.Layout(gtx, u.Theme)
	})

	sidebarLayout := layout.Rigid(func(gtx LC) LD {
		return u.sidebar.Layout(gtx, u.Theme)
	})

	multiViewLayout := layout.Flexed(1, func(gtx LC) LD {
		switch selectedPageIndexBuffer {
		case 0:
			return u.homeView.Layout(gtx, u.Theme)
		}
		return LD{}
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		headbarLayout,
		// 侧边栏和multiView
		layout.Flexed(1, func(gtx LC) LD {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				sidebarLayout,
				multiViewLayout,
			)
		}),
	)
}
