// 事件拦截.
// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "xc", 0, 炫彩常量类.Window_Style_Default)

	// 创建一个按钮
	btn := 炫彩组件类.X创建按钮(50, 50, 70, 30, "button", w.Handle)

	// 一个事件可以注册多个处理函数，执行顺序为先执行最后注册的函数，最后执行第一个注册的函数.
	// 当你想拦截当前事件或不想向后传递，只需要将参数(*pbHnadled=true)即可.
	btn.X事件_被单击(event1)
	btn.X事件_被单击(event2)
	btn.X事件_被单击(event3)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

func event1(pbHandled *bool) int {
	fmt.Println("event1")
	return 0
}

func event2(pbHandled *bool) int {
	fmt.Println("event2")
	return 0
}

func event3(pbHandled *bool) int {
	fmt.Println("event3")
	*pbHandled = true
	return 0
}
