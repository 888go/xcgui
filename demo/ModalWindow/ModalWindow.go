// 模态窗口
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "模态窗口", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮_模态窗口
	btn := 炫彩组件类.X创建按钮(30, 50, 100, 30, "ModalWindow", w.Handle)
	// 给按钮绑定事件
	btn.X事件_被单击(func(pbHandled *bool) int {
		// 创建模态窗口
		mw := 炫彩窗口基类.X创建模态窗口(300, 200, "ModalWindow", w.X取HWND(), 炫彩常量类.Window_Style_Modal)
		// 显示模态窗口
		mw.X启动()
		return 0
	})

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
