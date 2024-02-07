// 列表: 添加行, 删除选中行, 清空行, 排序, 表头表项文本居中
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a    *炫彩App类.App
	w    *炫彩窗口基类.Window
	list *炫彩组件类.List

	btn_add   *炫彩组件类.Button
	btn_del   *炫彩组件类.Button
	btn_clear *炫彩组件类.Button
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 784, 400, "List", 0, 炫彩常量类.Window_Style_Default)

	// 创建List
	createList()
	// List添加行
	listAddItem()

	var startX int32 = 10
	btn_add = 炫彩组件类.X创建按钮(startX, 35, 100, 30, "添加20行", w.Handle)
	btn_add.X事件_被单击1(onBnClick)

	startX += 100 + 3
	btn_del = 炫彩组件类.X创建按钮(startX, 35, 100, 30, "删除选中行", w.Handle)
	btn_del.X事件_被单击1(onBnClick)

	startX += 100 + 3
	btn_clear = 炫彩组件类.X创建按钮(startX, 35, 100, 30, "删除所有行", w.Handle)
	btn_clear.X事件_被单击1(onBnClick)

	w.X显示(true)
	a.X运行()
	a.X退出()
}

// 按钮单击事件
func onBnClick(hEle int, pbHandled *bool) int {
	炫彩基类.X元素_启用(hEle, false) // 操作前禁用按钮

	switch hEle {
	case btn_add.Handle:
		listAddItem()
	case btn_del.Handle:
		listDelSelectItem()
	case btn_clear.Handle:
		list.X删除项全部()
		list.X重绘(true)
	}

	炫彩基类.X元素_启用(hEle, true) // 操作后解禁按钮
	return 0
}

// 创建List
func createList() {
	// 创建List
	list = 炫彩组件类.X创建列表(10, 70, 764, 315, w.Handle)
	// 创建表头数据适配器
	list.X创建列表头数据适配器()
	// 创建数据适配器: 5列
	list.X创建数据适配器(5)
	// 列表_置项默认高度和选中时高度
	list.X置项默认高度(24, 26)
	// 列表_绘制项分割线
	// list.SetDrawItemBkFlags(xcc.List_DrawItemBk_Flag_Line | xcc.List_DrawItemBk_Flag_LineV | xcc.List_DrawItemBk_Flag_Leave | xcc.List_DrawItemBk_Flag_Stay | xcc.List_DrawItemBk_Flag_Select)
	// 表头和表项居中
	listTextAlign()

	// 添加列
	// 如果想要更好看的多功能的List就需要到设计器里设计[列表项模板], 比如说可以在项里添加按钮, 编辑框, 选择框, 组合框等, 可以任意DIY. 可参照例子: List2
	list.X添加列文本(50, "name1", "序号")
	list.X添加列文本(147, "name2", "Column2")
	list.X添加列文本(147, "name3", "Column3")
	list.X添加列文本(147, "name4", "Column4")
	list.X添加列文本(147, "name5", "Column5")

	// 设置序号列可排序, 单击表头时排序
	list.X置排序(0, 0, true)
	// 这里我使用了置属性的方法是为了不新建多个变量, 因为考虑到组件可能会很多, 当然你也可以用变量来控制.
	// 这个置属性你可以理解为就是给元素绑定的map中赋值. 并不是在操作元素的属性.
	// 也是为了演示Set/GetProperty, 这个东西很有用, 比如说你的列表每1行都有隐藏的值, 就可以存在这里, 而不用自己新建一个map或slice, 看你需求了.
	list.X置属性("sortType", "1") // 1是正序, 0是倒序.
	list.X置属性("sortFlag", "0") // 只是我设定的标记

	// 列表头项被单击事件
	list.X事件_列表头项点击(func(iItem int32, pbHandled *bool) int {
		// 为了记录排序类型
		if iItem == 0 {
			// 下面这个sortFlag只是我设定的1个标记, 意义是让第1次单击表头排序时不设置sortType的值, 因为第1次默认就是正序
			if list.X取属性("sortFlag") != "1" {
				list.X置属性("sortFlag", "1")
			} else {
				if list.X取属性("sortType") == "1" {
					list.X置属性("sortType", "0")
					fmt.Println("列表当前排序: 倒序")
				} else {
					list.X置属性("sortType", "1")
					fmt.Println("列表当前排序: 正序")
				}
			}
		}
		return 0
	})
}

// 表头和表项居中, 纯代码实现需要记一些api, 需要有清晰的思维, 还是用设计器来的简单, 真要写大程序不可能离开设计器的
func listTextAlign() {
	list.X事件_列表头项模板创建完成(func(pItem *炫彩基类.List_Header_Item_, pbHandled *bool) int {
		for i := 0; i < list.X取列数量(); i++ {
			hEle := list.X取列表头模板对象(i, 1)
			if a.X判断句柄包含类型(hEle, 炫彩常量类.XC_SHAPE_TEXT) { // 是形状文本
				炫彩基类.X形状文本_置文本对齐(hEle, 炫彩常量类.TextAlignFlag_Center|炫彩常量类.TextAlignFlag_Vcenter)
			}
		}
		return 0
	})

	list.X事件_项模板创建完成(func(pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int {
		// nFlag  0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		if nFlag == 1 {
			for i := 0; i < list.X取列数量(); i++ {
				hEle := list.X取项模板对象(int(pItem.Index), i, 1)
				if a.X判断句柄包含类型(hEle, 炫彩常量类.XC_SHAPE_TEXT) { // 是形状文本
					炫彩基类.X形状文本_置文本对齐(hEle, 炫彩常量类.TextAlignFlag_Center|炫彩常量类.TextAlignFlag_Vcenter)
				}
			}
		}
		return 0
	})
}

// List添加20行
func listAddItem() {
	// 循环添加数据
	for i := 0; i < 20; i++ {
		num := list.X取项数量AD() + 1

		// 添加行
		var index int
		if list.X取属性("sortType") == "1" { // 正序
			index = list.X添加项文本EX("name2", fmt.Sprintf("item%d-Column2", num))
		} else { // 倒序
			index = list.X插入项文本EX(0, "name2", fmt.Sprintf("item%d-Column2", num))
		}
		fmt.Printf("添加行索引: %d\n", index)

		// 置行数据
		// 序号列设置int型的数据才能按数字大小排序
		list.X置项指数值(index, 0, num)
		list.X置项文本(index, 2, fmt.Sprintf("item%d-Column3", num))
		list.X置项文本(index, 3, fmt.Sprintf("item%d-Column4", num))
		list.X置项文本(index, 4, fmt.Sprintf("item%d-Column5", num))
	}

	list.X重绘(true)
}

// List删除选中行
func listDelSelectItem() {
	count := list.X取选择项数量()
	if count == 0 {
		w.X消息框("提示", "你没有选中列表任何行!", 炫彩常量类.MessageBox_Flag_Ok|炫彩常量类.MessageBox_Flag_Icon_Info, 炫彩常量类.Window_Style_Modal)
		return
	}

	var indexArr []int32
	// 取选中行索引数组
	list.X取全部选择(&indexArr, count)
	// 根据选中行索引数组倒着删, 正着删的话你每删1行下面的行索引就变了
	for i := count - 1; i > -1; i-- {
		list.X删除项(int(indexArr[i]))
		fmt.Printf("删除行索引: %d\n", indexArr[i])
	}

	// 重排剩余行序号
	count = list.X取项数量AD()
	if list.X取属性("sortType") == "1" { // 正序
		for i := 0; i < count; i++ {
			list.X置项指数值(i, 0, i+1)
		}
	} else { // 倒序
		for i, num := 0, count; i < count; i, num = i+1, num-1 {
			list.X置项指数值(i, 0, num)
		}
	}

	// 刷新列表项数据
	list.X刷新项数据()
	// 列表重绘
	list.X重绘(true)
}
