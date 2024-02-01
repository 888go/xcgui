// 菜单
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	a   *app.App
	w   *window.Window
	btn *widget.Button

	item_selected = true // 控制item_select是否选中
)

const (
	item1 = iota + 10000
	subitem1
	subitem2

	item2
	item_select
	item_disable
)

func main() {
	// 1.初始化UI库
	a = app.I初始化(true)
	// 2.创建窗口
	w = window.I窗口_创建(0, 0, 366, 200, "xc", 0, xcc.I常量_窗口样式_默认)

	// 创建一个按钮
	btn = widget.I按钮_创建(50, 50, 100, 30, "ShowMenu", w.I句柄)
	// 注册按钮被单击事件
	btn.I事件_按钮被单击(onBnClick)

	// 注册菜单被选择事件
	btn.I事件_弹出菜单项被选择(onMenuSelect)

	// 注册菜单弹出事件
	btn.I事件_菜单弹出(onMenuPopup)

	// 注册菜单退出事件
	btn.I事件_菜单退出(onMenuExit)

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}

// 按钮被单击事件
func onBnClick(pbHandled *bool) int {
	// 创建菜单
	menu := widget.I菜单_创建()
	// 一级菜单
	menu.I添加项(item1, "item1", 0, xcc.I常量_菜单标识_正常)
	menu.I添加项(item2, "item2", 0, xcc.I常量_菜单标识_正常)
	if item_selected {
		menu.I添加项(item_select, "item_select", 0, xcc.I常量_菜单标识_勾选)
	} else {
		menu.I添加项(item_select, "item_select", 0, xcc.I常量_菜单标识_正常)
	}
	menu.I添加项(-1, "", 0, xcc.I常量_菜单标识_分隔栏) // 分隔栏
	menu.I添加项(item_disable, "item_disable", 0, xcc.I常量_菜单标识_禁用)

	// item1的二级菜单
	menu.I添加项(subitem1, "subitem1", item1, xcc.I常量_菜单标识_正常)
	menu.I添加项(subitem2, "subitem2", item1, xcc.I常量_菜单标识_正常)

	// 获取按钮坐标
	var r xc.RECT
	btn.I取坐标(&r)
	// 转换到屏幕坐标
	pt := xc.POINT{X: r.Left, Y: r.Bottom}
	xc.ClientToScreen(w.I句柄, &pt)
	// 弹出菜单
	menu.I弹出(w.I句柄, int(pt.X), int(pt.Y), btn.I句柄, xcc.I常量_菜单弹出方向_左上角)
	return 0
}

// 菜单被选择事件
func onMenuSelect(nID int, pbHandled *bool) int {
	fmt.Println("菜单被选择:", nID)
	if nID == item_select {
		item_selected = !item_selected
	}
	return 0
}

// 菜单弹出事件
func onMenuPopup(HMENUX int, pbHandled *bool) int {
	//fmt.Println("弹出菜单")
	return 0
}

// 菜单退出事件
func onMenuExit(pbHandled *bool) int {
	//fmt.Println("菜单退出")
	return 0
}
