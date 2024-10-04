package fonts

import (
	"embed"
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"log"
)

//go:embed fonts/*
var fonts embed.FS
var AllFontFaces []font.FontFace

func init() {
	var err error
	AllFontFaces, err = Prepare()
	if err != nil {
		log.Fatal(err)
	}
}

func Prepare() ([]font.FontFace, error) {
	var fontFaces []font.FontFace

	sourceHanSansRegular, err := registerFonts("SourceHanSansSC-Regular.otf")
	if err != nil {
		return nil, err
	}

	sourceHanSansMedium, err := registerFonts("SourceHanSansSC-Medium.otf")
	if err != nil {
		return nil, err
	}

	cascadiaCode, err := registerFonts("CascadiaCode.ttf")
	if err != nil {
		return nil, err
	}

	fontAwesome6Regular, err := registerFonts("Font Awesome 6 Free-Regular-400.otf")
	if err != nil {
		return nil, err
	}

	fontAwesome6Solid, err := registerFonts("Font Awesome 6 Free-Solid-900.otf")
	if err != nil {
		return nil, err
	}

	fontFaces = append(fontFaces,
		font.FontFace{Font: font.Font{Weight: font.Medium}, Face: sourceHanSansMedium},
		font.FontFace{Font: font.Font{Typeface: "FontAwes6Regular"}, Face: fontAwesome6Regular},
		font.FontFace{Font: font.Font{Weight: font.Light, Typeface: "FontAwes6Solid"}, Face: fontAwesome6Solid},
		font.FontFace{Font: font.Font{}, Face: cascadiaCode},
		font.FontFace{Font: font.Font{Weight: font.Normal}, Face: sourceHanSansRegular},
	)
	return fontFaces, nil
}
func getFontFromPath(path string) ([]byte, error) {
	data, err := fonts.ReadFile(fmt.Sprintf("fonts/%s", path))
	if err != nil {
		return nil, err
	}
	return data, err
}

func registerFonts(name string) (opentype.Face, error) {
	fontNameTF, err := getFontFromPath(name)
	if err != nil {
		return opentype.Face{}, err
	}
	fontName, err := opentype.Parse(fontNameTF)
	if err != nil {
		return opentype.Face{}, err
	}
	return fontName, nil
}
