package fonts

import (
	"embed"
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"log"
)

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

	cascadiaCodeTTF, err := getFont("CascadiaCode.ttf")
	if err != nil {
		return nil, err
	}
	cascadiaCode, err := opentype.Parse(cascadiaCodeTTF)
	if err != nil {
		return nil, err
	}

	fontFaces = append(fontFaces,
		font.FontFace{Font: font.Font{}, Face: cascadiaCode},
	)
	return fontFaces, nil
}

func getFont(path string) ([]byte, error) {
	data, err := fonts.ReadFile(fmt.Sprintf("fonts/%s", path))
	if err != nil {
		return nil, err
	}
	return data, err
}
