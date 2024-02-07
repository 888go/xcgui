// 列表, 模板进阶操作.
package main

import (
	_ "embed"
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	//go:embed list2/list2.zip
	zipData []byte

	a  *炫彩App类.App
	w  *炫彩窗口基类.Window
	ls *炫彩组件类.List
)

func main() {
	a = 炫彩App类.X创建(true)
	// 从内存zip加载资源文件
	a.X加载资源文件内存ZIP(zipData, "resource.res", "")
	w = 炫彩窗口基类.X创建窗口(0, 0, 302, 308, "列表, 模板进阶操作", 0, 炫彩常量类.Window_Style_Default)

	// 创建List
	ls = 炫彩组件类.X创建列表(10, 33, 282, 263, w.Handle)

	// List的模板得设置两遍, 一遍是列表头, 一遍是列表项
	var hTemp int
	if hTemp = 炫彩基类.X模板_加载从内存ZIP(炫彩常量类.ListItemTemp_Type_List_Head, zipData, "tmpl_list.xml", ""); hTemp == 0 {
		panic("ListItemTemp_Type_List_Head: hTemp==0")
	}
	ls.X置项模板(hTemp)

	if hTemp = 炫彩基类.X模板_加载从内存ZIP(炫彩常量类.ListItemTemp_Type_List_Item, zipData, "tmpl_list.xml", ""); hTemp == 0 {
		panic("ListItemTemp_Type_List_Item: hTemp==0")
	}
	ls.X置项模板(hTemp)

	// 创建表头数据适配器
	ls.X创建列表头数据适配器()
	// 创建数据适配器: 2列
	ls.X创建数据适配器(2)
	// 列表_置项默认高度
	ls.X置项默认高度(28, 28)

	// 添加列
	ls.X添加列文本(68, "name1", "状态")
	ls.X添加列文本(192, "name2", "操作")

	// 循环添加数据
	for i := 0; i < 8; i++ {
		// 添加行
		index := ls.X添加项文本("")
		// 置行数据
		ls.X置项文本(index, 1, "")
	}

	// 注册项模板创建完成事件
	ls.X事件_项模板创建完成(onLIST_TEMP_CREATE_END)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

// 项模板创建完成事件. 在此事件中获取按钮并注册事件
func onLIST_TEMP_CREATE_END(pItem *炫彩基类.List_Item_, nFlag int32, pbHandled *bool) int {
	// 只在创建新模板实例的时候, 给按钮注册事件, 这样是为了避免重复注册事件
	if nFlag == 1 { // 0:状态改变(复用); 1:新模板实例; 2:旧模板复用
		index := int(pItem.Index)
		hBtn := ls.X取项模板对象(index, 0, 2) // 前两个参数是项索引和列索引, 第三个参数是项模板里按钮的itemID, 在设计器里是可以自己填的, 必须填了, 这里才能获取
		fmt.Println(炫彩基类.X按钮_取文本(hBtn))
		// 注册按钮事件
		炫彩基类.X元素_注册事件C1(hBtn, 炫彩常量类.XE_BNCLICK, onBnClick)

		// 项模板里按钮的itemID是2,3,4
		for i := 2; i < 5; i++ {
			hBtn = ls.X取项模板对象(index, 1, i)
			fmt.Println(炫彩基类.X按钮_取文本(hBtn))
			炫彩基类.X元素_注册事件C1(hBtn, 炫彩常量类.XE_BNCLICK, onBnClick)
		}
	}
	return 0
}

// 按钮事件
func onBnClick(hEle int, pbHandled *bool) int {
	row := ls.X取对象所在行(hEle) // 获取项索引
	btnText := 炫彩基类.X按钮_取文本(hEle)
	var col int // 列索引
	switch btnText {
	case "运行中":
		col = 0
	case "日志", "复制", "删除":
		col = 1
	}

	炫彩基类.X消息框("提示", fmt.Sprintf("你点击了按钮: %s, 行: %d, 列: %d", btnText, row, col), 炫彩常量类.MessageBox_Flag_Ok, w.X取HWND(), 炫彩常量类.Window_Style_Default)
	return 0
}
