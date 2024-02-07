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
	a := 炫彩App类.New(true)
	w := window.NewFrameWindow(0, 0, 1000, 800, "FrameWindow", 0, 炫彩常量类.Window_Style_Default)

	// 创建窗格
	hPane_left = 炫彩组件类.NewPane("left", 200, 280, w.Handle)
	hPane_right = 炫彩组件类.NewPane("right", 200, 280, w.Handle)
	hPane_bottom = 炫彩组件类.NewPane("bottom", 770, 170, w.Handle)

	// 添加窗格
	w.AddPane(0, hPane_left.Handle, 炫彩常量类.Pane_Align_Left)
	w.AddPane(0, hPane_right.Handle, 炫彩常量类.Pane_Align_Right)
	w.AddPane(0, hPane_bottom.Handle, 炫彩常量类.Pane_Align_Bottom)

	// 创建编辑框
	hEdit = 炫彩组件类.NewEdit(0, 0, 0, 0, w.Handle)
	hEdit.EnableMultiLine(true)
	// 设置主视图为编辑框
	w.SetView(hEdit.Handle)

	// 窗口_调整布局
	w.AdjustLayout()

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}
