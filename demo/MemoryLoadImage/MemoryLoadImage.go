// 内存加载图片
package main

import (
	_ "embed"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

//go:embed 1.png
var img1 []byte

func main() {
	a := app.New(true)
	w := window.New(0, 0, 415, 296, "", 0, xcc.Window_Style_Default)

	// 加载图片从内存
	hImg := xc.XImage_LoadMemory(img1)

	// 创建形状对象-图片
	shapePic := widget.NewShapePicture(8, 30, 400, 260, w.Handle)
	shapePic.SetImage(hImg)

	w.ShowWindow(xcc.SW_SHOW)
	a.Run()
	a.Exit()
}
