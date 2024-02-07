// 多窗口例子，比如你登录后销毁登录窗口载入主窗口
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

var (
	a    *炫彩App类.App
	w1   *炫彩窗口基类.Window
	w2   *炫彩窗口基类.Window
	btn1 *炫彩组件类.Button
	btn2 *炫彩组件类.Button
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)

	loadWindow1()

	a.X运行()
	a.X退出()
}

func loadWindow1() {
	w1 = 炫彩窗口基类.X创建窗口(0, 0, 200, 200, "窗口1", 0, 炫彩常量类.Window_Style_Default)
	btn1 = 炫彩组件类.X创建按钮(50, 50, 100, 30, "载入窗口2", w1.Handle)
	btn1.X事件_被单击1(onBnClick)
	w1.X显示(true)
}

func loadWindow2() {
	w2 = 炫彩窗口基类.X创建窗口(0, 0, 300, 300, "窗口2", 0, 炫彩常量类.Window_Style_Default)
	btn2 = 炫彩组件类.X创建按钮(100, 100, 100, 30, "载入窗口1", w2.Handle)
	btn2.X事件_被单击1(onBnClick)
	w2.X显示(true)
}

func onBnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn1.Handle:
		*pbHandled = true // 把单击事件拦截了, 载入新窗口时这是必要的
		w1.X关闭()
		loadWindow2()
	case btn2.Handle:
		*pbHandled = true // 把单击事件拦截了, 载入新窗口时这是必要的
		w2.X关闭()
		loadWindow1()
	}
	return 0
}
