// 树形框
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
	w := window.I窗口_创建(0, 0, 430, 300, "", 0, xcc.I常量_窗口样式_默认)

	// 创建Tree
	tree := widget.I树形框_创建(12, 33, 400, 260, w.I句柄)
	// 创建数据适配器, 这个是必须的, 存储数据的
	tree.I创建数据适配器()

	// 循环添加数据
	for i := 0; i < 5; i++ {
		// 插入项
		索引 := tree.I插入项文本(fmt.Sprintf("item%d", i), xcc.I常量_根节点, xcc.I常量_插入末尾位置)
		// 插入2个子项
		tree.I插入项文本("subitem-1", 索引, xcc.I常量_插入末尾位置)
		subitemIndex := tree.I插入项文本("subitem-2", 索引, xcc.I常量_插入末尾位置)
		// 给子项2插入2个子项
		tree.I插入项文本("subitem-2-1", subitemIndex, xcc.I常量_插入末尾位置)
		tree.I插入项文本("subitem-2-2", subitemIndex, xcc.I常量_插入末尾位置)
	}

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
