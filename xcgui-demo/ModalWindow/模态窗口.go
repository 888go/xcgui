// 模态窗口
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "xc", 0, xcc.I常量_窗口样式_默认)

	// 创建按钮_模态窗口
	btn := widget.I按钮_创建(30, 50, 100, 30, "ModalWindow", w.I句柄)
	// 给按钮绑定事件
	btn.I事件_按钮被单击(func(pbHandled *bool) int {
		// 创建模态窗口
		win_Modal := window.NewModalWindow(300, 200, "ModalWindow", w.I取HWND(), xcc.I常量_窗口样式_模态)
		// 显示模态窗口
		win_Modal.DoModal()
		return 0
	})

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
