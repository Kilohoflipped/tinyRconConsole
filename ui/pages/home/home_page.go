package home

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/widgets"
	"image"
)

type Page struct {
	themeSwatchComp *widgets.ThemeSwatch
}

func NewPage(theme *uitheme.Theme) *Page {
	p := &Page{}
	p.themeSwatchComp = p.makeThemeSwatch(theme)
	return p
}

func (p *Page) makeThemeSwatch(theme THM) *widgets.ThemeSwatch {
	themeSwatchStyle := &widgets.ThemeSwatchStyle{
		BrandMainColor: theme.BrandColorPalette.BrandMainColor,
	}

	themeSwatchList := &widget.List{
		List: layout.List{
			Axis: layout.Horizontal,
		},
	}

	themeSwatchComp := widgets.NewThemeSwatch(themeSwatchList, themeSwatchStyle, theme)

	return themeSwatchComp
}

func (p *Page) Layout(gtx LC, theme THM) LD {
	bgLayout := func(gtx LC) LD {
		defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, theme.WidgetsColorMap.ColorGeneralBg)
		return LD{Size: gtx.Constraints.Min}
	}

	return layout.Background{}.Layout(gtx,
		bgLayout,
		func(gtx LC) LD {
			return layout.Center.Layout(gtx, func(gtx LC) LD {
				return p.themeSwatchComp.Layout(gtx, theme)
			})
		},
	)
}
