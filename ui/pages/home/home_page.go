package home

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"image"
)

type Page struct {
}

func NewPage(theme *uitheme.Theme) *Page {
	v := &Page{}
	return v
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
			return LD{}
		},
	)
}
