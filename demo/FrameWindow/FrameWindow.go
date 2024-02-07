// 框架窗口
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

var (
	hPane_left   *炫彩组件类.Pane
	hPane_right  *炫彩组件类.Pane
	hPane_bottom *炫彩组件类.Pane

	hEdit *炫彩组件类.Edit
)

func main() {
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建框架窗口(0, 0, 1000, 800, "FrameWindow", 0, 炫彩常量类.Window_Style_Default)

	// 创建窗格
	hPane_left = 炫彩组件类.X创建窗格("left", 200, 280, w.Handle)
	hPane_right = 炫彩组件类.X创建窗格("right", 200, 280, w.Handle)
	hPane_bottom = 炫彩组件类.X创建窗格("bottom", 770, 170, w.Handle)

	// 添加窗格
	w.X添加窗格(0, hPane_left.Handle, 炫彩常量类.Pane_Align_Left)
	w.X添加窗格(0, hPane_right.Handle, 炫彩常量类.Pane_Align_Right)
	w.X添加窗格(0, hPane_bottom.Handle, 炫彩常量类.Pane_Align_Bottom)

	// 创建编辑框
	hEdit = 炫彩组件类.X创建编辑框(0, 0, 0, 0, w.Handle)
	hEdit.X启用多行(true)
	// 设置主视图为编辑框
	w.X置视图(hEdit.Handle)

	// 窗口_调整布局
	w.X调整布局()

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
