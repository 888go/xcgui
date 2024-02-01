// 滑块条.
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "SliderBar", 0, xcc.I常量_窗口样式_默认)

	// 创建SliderBar
	sb := widget.I滑动条_创建(12, 33, 300, 60, w.I句柄)
	// 设置滑动范围
	sb.I置范围(10)

	// 设置滑块按钮高度和宽度
	sb.I置滑块高度(27)
	sb.I置滑块宽度(27)

	// 启用背景透明
	sb.I启用背景透明(true)

	// 注册滑块位置改变事件
	sb.I事件_滑块位置改变(func(pos int, pbHandled *bool) int {
		fmt.Println(pos)
		return 0
	})

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
