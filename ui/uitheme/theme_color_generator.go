package uitheme

import (
	"embed"
	"encoding/json"
	"fmt"
	"image/color"
	"math"
)

//go:embed brand_color_generator_parameters/*
var brandColorGeneratorParameters embed.FS

type ColorRegressionEquationParams struct {
	k, kb [9]float32
	kc    [3]float32
	m, mb [9]float32
}

var StdColorRegressionEquationParams *ColorRegressionEquationParams

func init() {
	var err error
	StdColorRegressionEquationParams, err = ReadJSON2GainParams("std_brand_color_generator_parameters.json")
	if err != nil {
		panic("load standard color regression equation params fatal")
	}
}

func GenerateColorPaletteFromMainColor(brandMainColor color.NRGBA) *BrandColorPaletteMap {
	p := &BrandColorPaletteMap{}
	p.BrandMainColor = brandMainColor
	for i := 1; i <= 16; i++ {
		brandColor, err := CalculateBrandColorByChannel(brandMainColor, i)
		if err != nil {
			panic("calculate brand color fatal")
		}
		p.BrandColors[i-1] = brandColor
	}
	return p
}

func CalculateBrandColorByChannel(brandMainColor color.NRGBA, channel int) (color.NRGBA, error) {
	// 把NRGBA颜色转换到非标准化的LAB颜色
	colorLabx, err := GetUnStdLabFromNrgba(brandMainColor)
	if err != nil {
		return color.NRGBA{}, err
	}

	// 根据channel计算拟合的LAB色板颜色
	colorLaby := StdColorRegressionEquationParams.CalculateColorLaby(channel, colorLabx)

	// 返回NRGBA颜色
	colorNrgbay := GetNrgbaFromUnStdLab(colorLaby)
	return colorNrgbay, nil
}

func CalculateColorRegressionEquation(channel int, labValue float32, k, kb [3]float32, kc float32, m, mb [3]float32) float32 {

	fittedLabValue := float64(k[0]*(labValue-kb[0])) + float64(k[1])*math.Pow(float64(labValue-kb[1]), 2.0) + float64(k[2])*math.Pow(float64(labValue-kb[2]), -1.0) +
		float64(kc) +
		float64(m[0]*(float32(channel)-mb[0])) + float64(m[1])*math.Pow(float64(float32(channel)-mb[1]), 2.0) + float64(m[2])*math.Pow(float64(float32(channel)-mb[2]), 3.0)

	return float32(fittedLabValue)
}

func ReadJSON2GainParams(paramsFilePath string) (*ColorRegressionEquationParams, error) {
	// 读取json文件
	stdJSONFile, err := brandColorGeneratorParameters.ReadFile(
		fmt.Sprintf("brand_color_generator_parameters/%s",
			paramsFilePath))
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据
	var stdJSONData map[string]float32
	err = json.Unmarshal(stdJSONFile, &stdJSONData)
	if err != nil {
		return nil, err
	}

	p := &ColorRegressionEquationParams{}

	// 提取值
	for key, value := range stdJSONData {
		switch key {
		case "k1":
			p.k[0] = value
		case "k2":
			p.k[1] = value
		case "k3":
			p.k[2] = value
		case "k4":
			p.k[3] = value
		case "k5":
			p.k[4] = value
		case "k6":
			p.k[5] = value
		case "k7":
			p.k[6] = value
		case "k8":
			p.k[7] = value
		case "k9":
			p.k[8] = value

		case "kb1":
			p.kb[0] = value
		case "kb2":
			p.kb[1] = value
		case "kb3":
			p.kb[2] = value
		case "kb4":
			p.kb[3] = value
		case "kb5":
			p.kb[4] = value
		case "kb6":
			p.kb[5] = value
		case "kb7":
			p.kb[6] = value
		case "kb8":
			p.kb[7] = value
		case "kb9":
			p.kb[8] = value

		case "kc1":
			p.kc[0] = value
		case "kc2":
			p.kc[1] = value
		case "kc3":
			p.kc[2] = value

		case "m1":
			p.m[0] = value
		case "m2":
			p.m[1] = value
		case "m3":
			p.m[2] = value
		case "m4":
			p.m[3] = value
		case "m5":
			p.m[4] = value
		case "m6":
			p.m[5] = value
		case "m7":
			p.m[6] = value
		case "m8":
			p.m[7] = value
		case "m9":
			p.m[8] = value

		case "mb1":
			p.mb[0] = value
		case "mb2":
			p.mb[1] = value
		case "mb3":
			p.mb[2] = value
		case "mb4":
			p.mb[3] = value
		case "mb5":
			p.mb[4] = value
		case "mb6":
			p.mb[5] = value
		case "mb7":
			p.mb[6] = value
		case "mb8":
			p.mb[7] = value
		case "mb9":
			p.mb[8] = value
		}
	}
	return p, nil
}

func (e *ColorRegressionEquationParams) CalculateColorLy(channel int, colorLx float32) float32 {
	kL := [3]float32{e.k[0], e.k[3], e.k[6]}
	kbL := [3]float32{e.kb[0], e.kb[3], e.kb[6]}

	mL := [3]float32{e.m[0], e.m[3], e.m[6]}
	mbL := [3]float32{e.mb[0], e.mb[3], e.mb[6]}

	return CalculateColorRegressionEquation(channel, colorLx, kL, kbL, e.kc[0], mL, mbL)
}

func (e *ColorRegressionEquationParams) CalculateColorBy(channel int, colorBx float32) float32 {
	kB := [3]float32{e.k[1], e.k[4], e.k[7]}
	kbB := [3]float32{e.kb[1], e.kb[4], e.kb[7]}

	mB := [3]float32{e.m[1], e.m[4], e.m[7]}
	mbB := [3]float32{e.mb[1], e.mb[4], e.mb[7]}

	return CalculateColorRegressionEquation(channel, colorBx, kB, kbB, e.kc[1], mB, mbB)
}

func (e *ColorRegressionEquationParams) CalculateColorAy(channel int, colorAx float32) float32 {
	kA := [3]float32{e.k[2], e.k[5], e.k[8]}
	kbA := [3]float32{e.kb[2], e.kb[5], e.kb[8]}

	mA := [3]float32{e.m[2], e.m[5], e.m[8]}
	mbA := [3]float32{e.mb[2], e.mb[5], e.mb[8]}

	return CalculateColorRegressionEquation(channel, colorAx, kA, kbA, e.kc[2], mA, mbA)
}

func (e *ColorRegressionEquationParams) CalculateColorLaby(channel int, colorLabx [3]float32) [3]float32 {
	ly := e.CalculateColorLy(channel, colorLabx[0])
	ay := e.CalculateColorAy(channel, colorLabx[1])
	by := e.CalculateColorBy(channel, colorLabx[2])

	if ly < 0 {
		ly = 0
	} else if ly > 100 {
		ly = 100
	}

	if ay < -100 {
		ay = -100
	} else if ay > 100 {
		ay = 100
	}

	if by < -100 {
		by = -100
	} else if by > 100 {
		by = 100
	}

	return [3]float32{ly, ay, by}
}
