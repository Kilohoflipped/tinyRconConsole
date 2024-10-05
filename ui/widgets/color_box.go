package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
)

type ColorBox struct {
	Style *ColorBoxStyle
}

type ColorBoxStyle struct {
	size  image.Point
	color color.NRGBA
}

func NewColorBox(style *ColorBoxStyle) *ColorBox {
	b := &ColorBox{
		Style: style,
	}
	return b
}

func (b *ColorBox) Layout(gtx LC, theme THM) LD {
	defer clip.Rect{Max: b.Style.size}.Push(gtx.Ops).Pop()

	paint.ColorOp{Color: b.Style.color}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{Size: b.Style.size}
}
