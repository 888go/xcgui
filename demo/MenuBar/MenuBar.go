// 菜单条.
package main

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a  *炫彩App类.App
	w  *window.Window
	mb *炫彩组件类.MenuBar
)

func main() {
	// 1.初始化UI库
	a = 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	// 2.创建窗口
	w = window.New(0, 0, 570, 400, "MenuBar", 0, 炫彩常量类.Window_Style_Default)
	w.SetBorderSize(1, 30, 1, 1)

	// 创建菜单条
	mb = 炫彩组件类.NewMenuBar(5, 32, w.GetWidth()-10, 30, w.Handle)
	// 菜单条禁用按钮自动宽度
	mb.EnableAutoWidth(false)
	// 菜单条禁用绘制边框
	mb.EnableDrawBorder(false)
	// 菜单条禁用绘制焦点
	mb.EnableDrawFocus(false)

	// 添加菜单条按钮
	mb.AddButton("文件(F)")
	mb.AddButton("编辑(E)")
	mb.AddButton("搜索(S)")
	mb.AddButton("视图(V)")
	mb.AddButton("编码(N)")
	mb.AddButton("语言(L)")
	mb.AddButton("设置(T)")
	mb.AddButton("工具(O)")

	// 取消绘制菜单条按钮边框
	for i := 0; i < mb.GetChildCount(); i++ {
		hele := mb.GetChildByIndex(i)
		xc.XEle_EnableDrawBorder(hele, false)
		xc.XEle_RegEventC1(hele, 炫彩常量类.XE_BNCLICK, onMenuBarBnClick)
		xc.XEle_RegEventC1(hele, 炫彩常量类.XE_MENU_SELECT, onMenuSelect)
	}

	// 3.显示窗口
	w.ShowWindow(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}

func onMenuBarBnClick(hEle int, pbHandled *bool) int {
	fmt.Println(xc.XBtn_GetText(hEle) + "被单击了")
	// 创建菜单
	menu := 炫彩组件类.NewMenu()
	// 一级菜单
	menu.AddItem(10001, "item1", 0, 炫彩常量类.Menu_Item_Flag_Normal)
	menu.AddItem(10002, "item2", 0, 炫彩常量类.Menu_Item_Flag_Normal)

	// 获取按钮坐标
	var rc xc.RECT
	xc.XEle_GetWndClientRectDPI(hEle, &rc)
	// 转换到屏幕坐标
	pt := 炫彩WinApi类.POINT{X: rc.Left, Y: rc.Bottom}
	炫彩WinApi类.ClientToScreen(w.GetHWND(), &pt)
	// 弹出菜单
	menu.Popup(w.GetHWND(), pt.X, pt.Y, hEle, 炫彩常量类.Menu_Popup_Position_Left_Top)
	return 0
}

// 菜单被选择事件
func onMenuSelect(hEle int, nID int32, pbHandled *bool) int {
	fmt.Println(xc.XBtn_GetText(hEle)+"下的菜单被选择:", nID)
	return 0
}
