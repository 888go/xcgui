// 事件拦截.
// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "xc", 0, xcc.I常量_窗口样式_默认)

	// 创建一个按钮
	btn := widget.I按钮_创建(50, 50, 70, 30, "button", w.I句柄)

	// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
	// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
	btn.I事件_按钮被单击(event1)
	btn.I事件_按钮被单击(event2)
	btn.I事件_按钮被单击(event3)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}

func event1(pbHandled *bool) int {
	println("event1")
	return 0
}

func event2(pbHandled *bool) int {
	println("event2")
	return 0
}

func event3(pbHandled *bool) int {
	println("event3")
	*pbHandled = true
	return 0
}
