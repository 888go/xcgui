// 框架窗口
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	hPane_left   *widget.Pane
	hPane_right  *widget.Pane
	hPane_bottom *widget.Pane

	hEdit *widget.Edit
)

func main() {
	a := app.I初始化(true)
	w := window.NewFrameWindow(0, 0, 1000, 800, "FrameWindow", 0, xcc.I常量_窗口样式_默认)

	// 创建窗格
	hPane_left = widget.I窗格_创建("left", 200, 280, w.I句柄)
	hPane_right = widget.I窗格_创建("right", 200, 280, w.I句柄)
	hPane_bottom = widget.I窗格_创建("bottom", 770, 170, w.I句柄)

	// 添加窗格
	w.AddPane(0, hPane_left.I句柄, xcc.I常量_框架窗口对齐_左)
	w.AddPane(0, hPane_right.I句柄, xcc.I常量_框架窗口对齐_右)
	w.AddPane(0, hPane_bottom.I句柄, xcc.I常量_框架窗口对齐_底)

	// 创建编辑框
	hEdit = widget.I编辑框_创建(0, 0, 0, 0, w.I句柄)
	hEdit.I启用多行(true)
	// 设置主视图为编辑框
	w.SetView(hEdit.I句柄)

	// 窗口_调整布局
	w.I调整布局()

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
