// 树形框
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/font"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "Tree", 0, 炫彩常量类.Window_Style_Default)

	// 创建Tree
	tree := 炫彩组件类.X创建列表树(12, 33, 400, 260, w.Handle)
	// 创建数据适配器, 这个是必须的, 存储数据的
	tree.X创建数据适配器()

	// 循环添加数据
	for i := 0; i < 5; i++ {
		// 插入项
		index := tree.X插入项文本(fmt.Sprintf("item%d", i), 炫彩常量类.XC_ID_ROOT, 炫彩常量类.XC_ID_LAST)
		// 插入2个子项
		tree.X插入项文本("subitem-1", index, 炫彩常量类.XC_ID_LAST)
		subitemIndex := tree.X插入项文本("subitem-2", index, 炫彩常量类.XC_ID_LAST)
		// 给子项2插入2个子项
		tree.X插入项文本("subitem-2-1", subitemIndex, 炫彩常量类.XC_ID_LAST)
		tree.X插入项文本("subitem-2-2", subitemIndex, 炫彩常量类.XC_ID_LAST)
	}

	// 列表树元素-项模板创建完成.
	// 在此事件里可以对创建完成的项模板进行操作.
	// 你需要理解项模板的概念, Tree里面有一个默认的项模板, 它的每一项都是根据固定的项模板生成的.
	// 如果你想知道默认的项模板是什么结构的, 可以打开设计器, 在一个项目中新建文件, 选择树项模板.
	// 然后你就能理解项模板是什么样子了, 项模板就是各个基础元素组合而成的, 而你可以diy它达成你想要的样子.
	// 这就是项模板存在的意义, 然后可以通过tree.SetItemTemplate相关函数设置你自己的项模板.
	// 其它有项模板的元素还有: List, ListBox, ListView等.
	tree.X事件_项模板创建完成(func(pItem *炫彩基类.Tree_Item_, nFlag int32, pbHandled *bool) int {
		// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		if nFlag == 1 {
			// 获取项模板中(itemID=2)的形状文本句柄.
			// 在默认的项模板里, 文本是一个形状文本元素, 它的itemID就是2.
			hst := tree.X取模板对象(pItem.NID, 2)
			// 设置文本字体
			炫彩基类.X形状文本_置字体(hst, 炫彩字体类.X创建EX("Arial", 12, 炫彩常量类.FontStyle_Bold).Handle)
			// 设置文本颜色
			炫彩基类.X形状文本_置文本颜色(hst, 炫彩基类.ABGR(255, 34, 33, 255))
		}
		return 0
	})

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
