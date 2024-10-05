package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/Mr-Ao-Dragon/tinyRconConsole/ui/uitheme"
	"image"
	"image/color"
)

type ThemeSwatch struct {
	List *widget.List

	ColorBoxComps []*ColorBox

	Style *ThemeSwatchStyle
}

type ThemeSwatchStyle struct {
	BrandMainColor color.NRGBA
}

func NewThemeSwatch(list *widget.List, style *ThemeSwatchStyle, theme THM) *ThemeSwatch {
	s := &ThemeSwatch{
		List:  list,
		Style: style,
	}
	s.ColorBoxComps = s.makeColorBoxComps(theme)
	return s
}

func (s *ThemeSwatch) makeColorBoxComps(theme THM) []*ColorBox {
	colorsPalette := uitheme.GenerateColorPaletteFromMainColor(s.Style.BrandMainColor)

	colorBoxStyles := make([]*ColorBoxStyle, 0)
	for i := range colorsPalette.BrandColors {
		colorBoxStyles = append(colorBoxStyles, &ColorBoxStyle{
			color: colorsPalette.BrandColors[i],
			size:  image.Point{X: 50, Y: 500},
		})
	}

	colorBoxComps := make([]*ColorBox, 0)
	for i := range colorBoxStyles {
		colorBoxComps = append(colorBoxComps, NewColorBox(colorBoxStyles[i]))
	}
	return colorBoxComps
}

func (s *ThemeSwatch) Layout(gtx LC, theme THM) LD {
	bgLayout := func(gtx LC) LD {
		defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, 0).Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, theme.WidgetsColorMap.ColorGeneralBg)
		return LD{Size: gtx.Constraints.Min}
	}

	colorBoxesLayout := func(gtx LC) LD {
		return s.List.Layout(gtx, len(s.ColorBoxComps), func(gtx LC, index int) LD {
			return layout.Flex{}.Layout(gtx,
				layout.Rigid(func(gtx LC) LD {
					return s.ColorBoxComps[index].Layout(gtx, theme)
				}))
		})
	}

	return layout.Background{}.Layout(gtx,
		bgLayout,
		colorBoxesLayout,
	)
}
