// 炫彩_调用界面线程, 在主线程操作UI
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var w *window.Window

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)
	// 2.创建窗口
	w = window.I窗口_创建(0, 0, 466, 300, "xc", 0, xcc.I常量_窗口样式_简单|xcc.I常量_窗口样式_按钮_关闭)

	btn := widget.I按钮_创建(30, 40, 70, 30, "click", w.I句柄)
	btn.I事件_按钮被单击(func(pbHandled *bool) int {
		go func() {
			// 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
			a.I调用界面线程(test, 0)
		}()
		return 0
	})

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}

func test(data int) int {
	widget.I按钮_创建(30, 80, 80, 26, "new button", w.I句柄)
	w.I重绘(false)
	return 0
}
