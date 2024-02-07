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
	w *window.Window

	btn1 *炫彩组件类.Button
	btn2 *炫彩组件类.Button
	btn3 *炫彩组件类.Button
	btn4 *炫彩组件类.Button
)

func main() {
	a = 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 430, 300, "", 0, 炫彩常量类.Window_Style_Default)

	// 创建按钮
	btn1 = 炫彩组件类.NewButton(20, 40, 100, 30, "浏览文件夹", w.Handle)
	btn2 = 炫彩组件类.NewButton(20, 80, 100, 30, "单选打开文件", w.Handle)
	btn3 = 炫彩组件类.NewButton(130, 80, 100, 30, "多选打开文件", w.Handle)
	btn4 = 炫彩组件类.NewButton(20, 120, 100, 30, "保存文件", w.Handle)

	// 注册按钮事件
	btn1.Event_BnClick1(onBnClick)
	btn2.Event_BnClick1(onBnClick)
	btn3.Event_BnClick1(onBnClick)
	btn4.Event_BnClick1(onBnClick)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

func onBnClick(hEle int, pbHandled *bool) int {
	switch hEle {
	case btn1.Handle:
		fmt.Println(炫彩WinApi工具类.OpenDir(w.Handle))
	case btn2.Handle:
		fmt.Println(炫彩WinApi工具类.OpenFile(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
	case btn3.Handle:
		fmt.Println(炫彩WinApi工具类.OpenFiles(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
	case btn4.Handle:
		fmt.Println(炫彩WinApi工具类.SaveFile(w.Handle, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
	}
	return 0
}
