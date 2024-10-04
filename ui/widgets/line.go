package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
	"image/color"
)

// DrawLineFlex 一个线条Flex子组件
func DrawLineFlex(background color.NRGBA, height, width unit.Dp) layout.FlexChild {
	return layout.Rigid(func(gtx LC) LD {
		return DrawLine(gtx, background, height, width)
	})
}

// DrawLine 一个线条组件
func DrawLine(gtx LC, background color.NRGBA, height, width unit.Dp) LD {
	// 将大小从Dp转换为像素
	w, h := gtx.Dp(width), gtx.Dp(height)
	// 创建一个以左上角为起点的矩形
	tabRect := image.Rect(0, 0, w, h)
	paint.FillShape(gtx.Ops, background, clip.Rect(tabRect).Op())
	return layout.Dimensions{Size: image.Pt(w, h)}
}
