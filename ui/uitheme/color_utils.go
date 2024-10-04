package uitheme

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
)

func GetNrgbaFromNumHex(c uint32) color.NRGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)} //nolint:gosec
}

func GetStrHexFromNrgba(color color.NRGBA) string {
	return fmt.Sprintf("#%02x%02x%02x", color.R, color.G, color.B)
}

func GetNrgbaFromStrHex(hexColor string) (color.NRGBA, error) {
	var r, g, b uint8
	_, err := fmt.Sscanf(hexColor[1:], "%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return color.NRGBA{}, err
	}
	return color.NRGBA{R: r, G: g, B: b, A: 255}, nil
}

func GetUnStdLabFromNrgba(color color.NRGBA) ([3]float32, error) {
	c, err := colorful.Hex(GetStrHexFromNrgba(color))
	if err != nil {
		return [3]float32{}, err
	}
	l, a, b := c.Lab()
	l *= 100
	a *= 128
	b *= 128
	return [3]float32{float32(l), float32(a), float32(b)}, nil
}

func GetNrgbaFromUnStdLab(lab [3]float32) color.NRGBA {
	c := colorful.Lab(float64(lab[0]/100), float64(lab[1]/128), float64(lab[2]/128))
	return color.NRGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}
}

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	// 创建矩形剪辑区域，右下角坐标为size,使用Push开始剪辑绘制，使用Pop在返回前结束绘制剪辑限制
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	// 设置当前paint画笔颜色
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	// 使用当前paint画笔填充当前剪辑区域
	paint.PaintOp{}.Add(gtx.Ops)
	// 返回布局实例
	return layout.Dimensions{Size: size}
}
