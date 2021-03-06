package utils

import (
	color "github.com/fatih/color"
)

// 警告信息采用红色显示
var colors []*color.Color

type ColorFunc func(a ...interface{}) string

// 新增服务等采用绿色显示
var Red, Green, Magenta, Cyan, Blue ColorFunc

func init() {
	redFg := color.New(color.FgRed)
	Red = redFg.SprintFunc()

	greenFg := color.New(color.FgGreen)
	Green = greenFg.SprintFunc()

	magentaFg := color.New(color.FgMagenta)
	Magenta = magentaFg.SprintFunc()

	blueFg := color.New(color.FgBlue)
	Blue = blueFg.SprintFunc()

	cyanFg := color.New(color.FgCyan)
	Cyan = cyanFg.SprintFunc()

	colors = append(colors, redFg, greenFg, magentaFg, blueFg, cyanFg)
}

func DisableColorOutput() {
	for i := 0; i < len(colors); i++ {
		colors[i].DisableColor()
	}
}
