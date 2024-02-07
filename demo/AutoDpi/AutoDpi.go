// 启用自动DPI的两种方法, 解决高分辨率屏幕界面模糊问题.
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := 炫彩App类.New(true)

	// 告诉系统本软件要自己控制DPI的两种方法:
	// 1.调用 a.EnableDPI(true)
	// 2.使用程序清单, 看这个: 程序清单方式启用DPI.7z, Windows文档里写的是推荐使用程序清单方式的, 应该和系统兼容性有关.
	a.EnableDPI(true)

	// 使用上面的两种方法之一, 然后调用这个函数启用自动DPI
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w := window.New(0, 0, 430, 300, "xcgui window", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	btn := 炫彩组件类.NewButton(165, 135, 100, 30, "Button", w.Handle)
	// 注册按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		a.MessageBox("提示", btn.GetText(), 炫彩常量类.MessageBox_Flag_Ok|炫彩常量类.MessageBox_Flag_Icon_Info, w.GetHWND(), 炫彩常量类.Window_Style_Modal)
		return 0
	})

	// 3.显示窗口
	w.Show(true)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
