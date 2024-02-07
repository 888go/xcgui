// 组合框
package main

import (
	"strconv"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 430, 300, "ComboBox", 0, 炫彩常量类.Window_Style_Default)

	// 创建组合框
	cbb := 炫彩组件类.X创建组合框(24, 50, 100, 30, w.Handle)
	// 创建数据适配器, 这个是必须创建的, 存储数据的
	cbb.X创建数据适配器()
	// 组合框加入项
	for i := 1; i <= 5; i++ {
		cbb.X添加项文本("item" + strconv.Itoa(i))
	}

	// 组合框选中项
	cbb.X置选择项(0)
	// 组合框禁止编辑项
	cbb.X启用编辑(false)

	// 创建编辑框
	edit := 炫彩组件类.X创建编辑框(138, 50, 100, 30, w.Handle)
	edit.X置文本("hello")

	// 注册组合框被选择事件
	cbb.X事件_下拉列表项选择完成(func(iItem int32, pbHandled *bool) int {
		edit.X置文本(cbb.X取项文本(iItem, 0))
		edit.X重绘(false)
		return 0
	})

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
