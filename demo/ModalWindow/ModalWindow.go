// 模态窗口
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 430, 300, "模态窗口", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮_模态窗口
	btn := 炫彩组件类.NewButton(30, 50, 100, 30, "ModalWindow", w.Handle)
	// 给按钮绑定事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 创建模态窗口
		mw := window.NewModalWindow(300, 200, "ModalWindow", w.GetHWND(), 炫彩常量类.Window_Style_Modal)
		// 显示模态窗口
		mw.DoModal()
		return 0
	})

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}
