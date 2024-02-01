// 内存加载图片
package main

import (
	_ "embed"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

//go:embed 1.png
var img1 []byte

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 415, 296, "", 0, xcc.I常量_窗口样式_默认)

	// 加载图片从内存
	hImg := xc.XImage_LoadMemory(img1)

	//创建形状对象-图片
	shapePic := widget.I形状图片_创建(8, 30, 400, 260, w.I句柄)
	shapePic.I置图片(hImg)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
