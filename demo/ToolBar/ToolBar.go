// 工具条.
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
	"strconv"
)

var (
	a  *炫彩App类.App
	w  *window.Window
	tb *炫彩组件类.ToolBar
)

func main() {
	// 1.初始化UI库
	a = 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w = window.New(0, 0, 570, 400, "ToolBar", 0, 炫彩常量类.Window_Style_Default)
	w.SetBorderSize(1, 30, 1, 1)

	// 创建工具条
	tb = 炫彩组件类.NewToolBar(5, 32, w.GetWidth()-10, 30, w.Handle)

	// 插入元素
	for i := 1; i < 10; i++ {
		btn := 炫彩组件类.NewButton(0, 0, 100, 30, "按钮"+strconv.Itoa(i), tb.Handle)
		btn.EnableDrawBorder(false)          // 不绘制边框
		btn.EnableDrawFocus(false)           // 不绘制焦点
		btn.Event_BnClick1(onToolBarBnClick) // 注册按钮单击事件
		tb.InsertEle(btn.Handle, -1)
	}

	// 3.显示窗口
	w.ShowWindow(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func onToolBarBnClick(hEle int, pbHandled *bool) int {
	fmt.Println(xc.XBtn_GetText(hEle) + "被单击了")
	return 0
}
