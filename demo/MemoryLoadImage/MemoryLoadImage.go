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
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 415, 296, "", 0, 炫彩常量类.Window_Style_Default)

	// 加载图片从内存
	hImg := 炫彩基类.X图片_加载从内存(img1)

	// 创建形状对象-图片
	shapePic := 炫彩组件类.X创建形状图片(8, 30, 400, 260, w.Handle)
	shapePic.X置图片(hImg)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
