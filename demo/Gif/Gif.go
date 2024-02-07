// Gif.
package main

import (
	_ "embed"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/imagex"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

var (
	//go:embed res/bg1.gif
	bg1 []byte
)

var (
	a *炫彩App类.App
	w *window.Window
)

func main() {
	a = 炫彩App类.New(false)
	w = window.New(0, 0, 600, 400, "xcgui window", 0, 炫彩常量类.Window_Style_Default)

	shapeGif := 炫彩组件类.NewShapeGif(0, 30, 600, 400, w.Handle)
	shapeGif.SetImage(炫彩图片类.NewByMemAdaptive(bg1, 0, 0, 0, 0).Handle)

	w.Show(true)
	a.Run()
	a.Exit()
}
