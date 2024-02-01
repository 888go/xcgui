// 注册元素事件
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)
	// 2.创建窗口
	w := window.I窗口_创建(0, 0, 430, 300, "xc", 0, xcc.I常量_窗口样式_简单|xcc.I常量_窗口样式_按钮_关闭)

	// 创建一个按钮
	btn := widget.I按钮_创建(50, 50, 120, 40, "button", w.I句柄)

	// 所有的封装好的组件事件都是以Event_开头, 你跳转到事件定义的源码之后, 可以看到这个组件所有的事件
	// 注册按钮被单击事件
	btn.I事件_按钮被单击(func(pbHandled *bool) int {
		println("I事件_按钮被单击")
		return 0
	})

	// 注册鼠标进入事件
	btn.I事件_鼠标进入(func(pbHandled *bool) int {
		println("I事件_鼠标进入")
		return 0
	})

	// 注册鼠标离开事件
	btn.I事件_鼠标离开(func(hEleStay int, pbHandled *bool) int {
		println("I事件_窗口鼠标离开")
		return 0
	})

	// 注册鼠标滚轮滚动事件
	btn.I启用事件_XE_MOUSEWHEEL(true)
	btn.I事件_鼠标滚轮滚动(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		println("I事件_窗口鼠标滚轮滚动", nFlags, pPt.X, pPt.Y)
		return 0
	})

	//注册鼠标移动事件
	/* 	btn.I事件_窗口鼠标移动(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		println("I事件_窗口鼠标移动", nFlags, pPt.X, pPt.Y)
		return 0
	}) */

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}
