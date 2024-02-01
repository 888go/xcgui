// 列表
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
	w := window.I窗口_创建(0, 0, 784, 308, "List", 0, xcc.I常量_窗口样式_默认)

	// 创建List
	list := widget.I列表_创建(10, 33, 764, 263, w.I句柄)
	// 创建表头数据适配器
	list.I列表头创建数据适配器()
	// 创建数据适配器: 5列
	list.I创建数据适配器(5)
	// 列表_置项默认高度
	list.I置项默认高度(24, 24)

	// 添加列
	// 如果想要更好看的多功能的List就需要到设计器里设计[列表项模板], 比如说可以在项里添加按钮, 编辑框, 选择框, 组合框等, 可以任意DIY.
	list.I添加列文本(147, "name1", "Column1")
	list.I添加列文本(147, "name2", "Column2")
	list.I添加列文本(147, "name3", "Column3")
	list.I添加列文本(147, "name4", "Column4")
	list.I添加列文本(147, "name5", "Column5")

	// 循环添加数据
	for i := 0; i < 20; i++ {
		// 添加行
		index := list.I添加项文本(fmt.Sprintf("Column1-item%d", i))
		fmt.Printf("index: %v\n", index)
		// 置行数据
		list.I置项文本(index, 1, fmt.Sprintf("Column2-item%d", i))
		list.I置项文本(index, 2, fmt.Sprintf("Column3-item%d", i))
		list.I置项文本(index, 3, fmt.Sprintf("Column4-item%d", i))
		list.I置项文本(index, 4, fmt.Sprintf("Column5-item%d", i))
	}

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
