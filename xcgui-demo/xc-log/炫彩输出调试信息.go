// 炫彩输出调试信息
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	// 炫彩_启用debug文件, xcgui_debug.txt
	a.I启用debug文件(true)

	w := window.I窗口_创建(0, 0, 400, 300, "", 0, xcc.I常量_窗口样式_默认)

	btn := widget.I按钮_创建(20, 35, 70, 30, "Click", w.I句柄)
	btn.I事件_按钮被单击(func(pbHandled *bool) int {
		a.I输出调试信息到文件("hello")
		a.I输出调试信息到文件("word")
		return 0
	})

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
