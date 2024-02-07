// 调用 wapi 打开/保存文件, 浏览文件夹
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/wapi/wutil"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

var (
	a *炫彩App类.App
	w *炫彩窗口基类.Window

	btn1 *炫彩组件类.Button
	btn2 *炫彩组件类.Button
	btn3 *炫彩组件类.Button
	btn4 *炫彩组件类.Button
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	btn1 = 炫彩组件类.X创建按钮(20, 40, 100, 30, "浏览文件夹", w.Handle)
	btn2 = 炫彩组件类.X创建按钮(20, 80, 100, 30, "单选打开文件", w.Handle)
	btn3 = 炫彩组件类.X创建按钮(130, 80, 100, 30, "多选打开文件", w.Handle)
	btn4 = 炫彩组件类.X创建按钮(20, 120, 100, 30, "保存文件", w.Handle)

	// 注册按钮事件
	btn1.X事件_被单击1(onBnClick)
	btn2.X事件_被单击1(onBnClick)
	btn3.X事件_被单击1(onBnClick)
	btn4.X事件_被单击1(onBnClick)

	a.X显示窗口并运行(w.Handle)
	a.X退出()
}

func onBnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn1.Handle:
		fmt.Println(炫彩WinApi工具类.X对话框打开文件夹(w.Handle))
	case btn2.Handle:
		fmt.Println(炫彩WinApi工具类.X对话框打开单个文件(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
	case btn3.Handle:
		fmt.Println(炫彩WinApi工具类.X对话框打开多个文件(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
	case btn4.Handle:
		fmt.Println(炫彩WinApi工具类.X对话框保存文件(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
	}
	return 0
}
