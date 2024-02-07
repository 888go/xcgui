// 注册元素事件
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := 炫彩App类.X创建(true)
	// 2.创建窗口
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "xc", 0, 炫彩常量类.Window_Style_Simple|炫彩常量类.Window_Style_Btn_Close)

	// 创建一个按钮
	btn := 炫彩组件类.X创建按钮(50, 50, 120, 40, "button", w.Handle)

	// 所有的封装好的组件事件都是以Event_开头, 你跳转到事件定义的源码之后, 可以看到这个组件所有的事件
	// 注册按钮被单击事件
	btn.X事件_被单击(func(pbHandled *bool) int {
		println("Event_BnClick")
		return 0
	})

	// 注册鼠标进入事件
	btn.X事件_鼠标进入(func(pbHandled *bool) int {
		println("Event_MOUSESTAY")
		return 0
	})

	// 注册鼠标离开事件
	btn.X事件_鼠标离开(func(hEleStay int, pbHandled *bool) int {
		println("Event_MOUSELEAVE")
		return 0
	})

	// 注册鼠标滚轮滚动事件
	btn.X启用事件_XE_MOUSEWHEEL(true)
	btn.X事件_鼠标滚轮滚动(func(nFlags int, pPt *炫彩基类.POINT, pbHandled *bool) int {
		println("Event_MOUSEWHEEL", nFlags, pPt.X, pPt.Y)
		return 0
	})

	// 注册鼠标移动事件
	/* 	btn.Event_MOUSEMOVE(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		println("Event_MOUSEMOVE", nFlags, pPt.X, pPt.Y)
		return 0
	}) */

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}
