package app

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/app/router"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/fonts"
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

	router  *router.Router
	headbar *Headbar

	homePage *home.Page
}

func NewUI(w *app.Window) (*UI, error) {
	// 创建mainUI对象
	u := &UI{
		window: w,
	}

	fontCollection := fonts.AllFontFaces

	u.Theme = uitheme.New()
	u.Theme.Shaper = text.NewShaper(text.WithCollection(fontCollection))
	u.Theme.SetThemeState(uitheme.StateThemeDark)

	u.router = router.NewRouter(u.Theme)
	u.headbar = NewHeadbar(u.Theme)

	u.homePage = home.NewPage(u.Theme)

	return u, nil
}

func (u *UI) Run() error {
	// Run 主窗体
	var ops op.Ops
	// 主循环
	for {
		switch e := u.window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// 帧事件:正常运行
			gtx := app.NewContext(&ops, e) // 每个帧都会创建全局布局上下文
			u.Layout(gtx)                  // 在全局上下文里渲染UI
			e.Frame(gtx.Ops)               // 根据当前帧的操作树绘制帧
		}
	}
}

func (u *UI) Layout(gtx LC) LD {
	headbarLayout := layout.Rigid(func(gtx LC) LD {
		return u.headbar.Layout(gtx, u.Theme)
	})

	routerLayout := layout.Rigid(func(gtx LC) LD {
		return u.router.Layout(gtx, u.Theme)
	})

	homePageLayout := layout.Flexed(1, func(gtx LC) LD {
		return u.homePage.Layout(gtx, u.Theme)
	})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		headbarLayout,
		// 侧边栏和multiView
		layout.Flexed(1, func(gtx LC) LD {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				routerLayout,
				homePageLayout,
			)
		}),
	)
}
