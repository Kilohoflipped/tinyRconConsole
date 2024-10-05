package uitheme

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
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
	a *= 100
	b *= 100
	return [3]float32{float32(l), float32(a), float32(b)}, nil
}

func GetNrgbaFromUnStdLab(lab [3]float32) color.NRGBA {
	c := colorful.Lab(float64(lab[0]/100), float64(lab[1]/100), float64(lab[2]/100))
	if c.R < 0 {
		c.R = 0
	} else if c.R > 1 {
		c.R = 1
	}

	if c.G < 0 {
		c.G = 0
	} else if c.G > 1 {
		c.G = 1
	}

	if c.B < 0 {
		c.B = 0
	} else if c.B > 1 {
		c.B = 1
	}

	return color.NRGBA{R: uint8(c.R * 255), G: uint8(c.G * 255), B: uint8(c.B * 255), A: 255}
}
