// 组合框
package main

import (
	"strconv"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "ComboBox", 0, xcc.I常量_窗口样式_默认)

	// 创建组合框
	cbb := widget.I组合框_创建(24, 50, 100, 30, w.I句柄)
	// 创建数据适配器, 这个是必须创建的, 存储数据的
	cbb.I创建数据适配器()
	// 组合框加入项
	for i := 1; i <= 5; i++ {
		cbb.I添加项文本("item" + strconv.Itoa(i))
	}

	// 组合框选中项
	cbb.I置选择项(0)
	// 组合框禁止编辑项
	cbb.I启用编辑(false)

	// 创建编辑框
	edit := widget.I编辑框_创建(138, 50, 100, 30, w.I句柄)
	edit.I置文本("hello")

	// 注册组合框被选择事件
	cbb.I事件_选择完成(func(iItem int, pbHandled *bool) int {
		edit.I置文本(cbb.I取项文本(iItem, 0))
		edit.I重绘(true)
		return 0
	})

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
