// 日期时间框
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 400, 300, "", 0, xcc.I常量_窗口样式_默认)

	dt := widget.I日期_创建(20, 50, 120, 26, w.I句柄)
	// 0为日期元素, 1为时间元素.
	dt.I置样式(0)

	dt2 := widget.I日期_创建(20, 150, 120, 26, w.I句柄)
	// 0为日期元素, 1为时间元素.
	dt2.I置样式(1)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
