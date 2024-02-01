// 列表框
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
	w := window.I窗口_创建(0, 0, 430, 500, "ListBox", 0, xcc.I常量_窗口样式_默认)

	// 创建ListBox
	lb := widget.I列表框_创建(12, 33, 400, 450, w.I句柄)

	// 创建数据适配器, 这个必须创建, 存储数据的
	lb.I创建数据适配器()

	for i := 0; i < 15; i++ {
		// 添加行
		lb.I添加项文本(fmt.Sprintf("item-%d", i))
	}

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
