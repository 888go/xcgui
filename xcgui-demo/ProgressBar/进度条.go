// 进度条
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	bar     *widget.ProgressBar
	btn_Add *widget.Button
	btn_Sub *widget.Button
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 436, 104, "xc", 0, xcc.I常量_窗口样式_默认)

	// 创建一个进度条
	bar = widget.I进度条_创建(24, 60, 200, 10, w.I句柄)
	// 设置进度条边框大小
	bar.I置边框大小(1, 1, 1, 1)
	// 设置进度条不显示进度文字
	bar.I启用进度文本(false)
	// 设置进度条最大值
	bar.I置范围(100)
	// 设置进度条进度为0
	bar.I置进度(0)

	// 创建按钮_进度加
	btn_Add = widget.I按钮_创建(238, 50, 70, 30, "+", w.I句柄)
	btn_Add.I事件_按钮被单击1(onBtnClick)
	// 创建按钮_进度减
	btn_Sub = widget.I按钮_创建(318, 50, 70, 30, "-", w.I句柄)
	btn_Sub.I事件_按钮被单击1(onBtnClick)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}

// 事件_按钮被单击
func onBtnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn_Add.I句柄:
		bar.I置进度(bar.I取进度() + 10)
	case btn_Sub.I句柄:
		bar.I置进度(bar.I取进度() - 10)
	}
	bar.I重绘(true)
	return 0
}
