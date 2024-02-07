// 菜单
package main

import (
	"fmt"
	"github.com/888go/xcgui/wapi"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a   *炫彩App类.App
	w   *炫彩窗口基类.Window
	btn *炫彩组件类.Button

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
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 2.创建窗口
	w = 炫彩窗口基类.X创建窗口(0, 0, 400, 300, "Menu", 0, 炫彩常量类.Window_Style_Default)

	// 创建一个按钮
	btn = 炫彩组件类.X创建按钮(50, 50, 100, 30, "ShowMenu", w.Handle)
	// 注册按钮被单击事件
	btn.X事件_被单击(onBnClick)

	// 注册菜单被选择事件
	w.X线程_菜单选择(onMenuSelect)

	// 注册菜单弹出事件
	w.X线程_菜单弹出(onMenuPopup)

	// 注册菜单退出事件
	w.X线程_菜单退出(onMenuExit)

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}

// 按钮被单击事件
func onBnClick(pbHandled *bool) int {
	// 创建菜单
	menu := 炫彩组件类.X创建菜单()
	// 一级菜单
	menu.X添加项(item1, "item1", 0, 炫彩常量类.Menu_Item_Flag_Normal)
	menu.X添加项(item2, "item2", 0, 炫彩常量类.Menu_Item_Flag_Normal)
	if item_selected {
		menu.X添加项(item_select, "item_select", 0, 炫彩常量类.Menu_Item_Flag_Check)
	} else {
		menu.X添加项(item_select, "item_select", 0, 炫彩常量类.Menu_Item_Flag_Normal)
	}
	menu.X添加项(-1, "", 0, 炫彩常量类.Menu_Item_Flag_Separator) // 分隔栏
	menu.X添加项(item_disable, "item_disable", 0, 炫彩常量类.Menu_Item_Flag_Disable)

	// item1的二级菜单
	menu.X添加项(subitem1, "subitem1", item1, 炫彩常量类.Menu_Item_Flag_Normal)
	menu.X添加项(subitem2, "subitem2", item1, 炫彩常量类.Menu_Item_Flag_Normal)

	// 获取按钮坐标
	var rc 炫彩基类.RECT
	btn.X取窗口客户区坐标DPI(&rc)
	// 转换到屏幕坐标
	pt := 炫彩WinApi类.POINT{X: rc.Left, Y: rc.Bottom}
	炫彩WinApi类.X窗口取屏幕坐标(w.X取HWND(), &pt)
	// 弹出菜单
	menu.X弹出(w.X取HWND(), pt.X, pt.Y, 0, 炫彩常量类.Menu_Popup_Position_Left_Top)
	return 0
}

// 菜单被选择事件
func onMenuSelect(nID int32, pbHandled *bool) int {
	fmt.Println("菜单被选择:", nID)
	if nID == item_select {
		item_selected = !item_selected
	}
	return 0
}

// 菜单弹出事件
func onMenuPopup(HMENUX int, pbHandled *bool) int {
	fmt.Println("弹出菜单")
	return 0
}

// 菜单退出事件
func onMenuExit(pbHandled *bool) int {
	fmt.Println("菜单退出")
	return 0
}
