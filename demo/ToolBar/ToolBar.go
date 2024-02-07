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
	w  *炫彩窗口基类.Window
	tb *炫彩组件类.ToolBar
)

func main() {
	// 1.初始化UI库
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 2.创建窗口
	w = 炫彩窗口基类.X创建窗口(0, 0, 570, 400, "ToolBar", 0, 炫彩常量类.Window_Style_Default)
	w.X置边大小(1, 30, 1, 1)

	// 创建工具条
	tb = 炫彩组件类.X创建工具条(5, 32, w.X取宽度()-10, 30, w.Handle)

	// 插入元素
	for i := 1; i < 10; i++ {
		btn := 炫彩组件类.X创建按钮(0, 0, 100, 30, "按钮"+strconv.Itoa(i), tb.Handle)
		btn.X启用绘制边框(false)          // 不绘制边框
		btn.X启用绘制焦点(false)           // 不绘制焦点
		btn.X事件_被单击1(onToolBarBnClick) // 注册按钮单击事件
		tb.X插入元素(btn.Handle, -1)
	}

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}

func onToolBarBnClick(hEle int, pbHandled *bool) int {
	fmt.Println(炫彩基类.X按钮_取文本(hEle) + "被单击了")
	return 0
}
