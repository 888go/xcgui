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
	w  *炫彩窗口基类.Window
	mb *炫彩组件类.MenuBar
)

func main() {
	// 1.初始化UI库
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	// 2.创建窗口
	w = 炫彩窗口基类.X创建窗口(0, 0, 570, 400, "MenuBar", 0, 炫彩常量类.Window_Style_Default)
	w.X置边大小(1, 30, 1, 1)

	// 创建菜单条
	mb = 炫彩组件类.X创建菜单条(5, 32, w.X取宽度()-10, 30, w.Handle)
	// 菜单条禁用按钮自动宽度
	mb.X启用自动宽度(false)
	// 菜单条禁用绘制边框
	mb.X启用绘制边框(false)
	// 菜单条禁用绘制焦点
	mb.X启用绘制焦点(false)

	// 添加菜单条按钮
	mb.X添加按钮("文件(F)")
	mb.X添加按钮("编辑(E)")
	mb.X添加按钮("搜索(S)")
	mb.X添加按钮("视图(V)")
	mb.X添加按钮("编码(N)")
	mb.X添加按钮("语言(L)")
	mb.X添加按钮("设置(T)")
	mb.X添加按钮("工具(O)")

	// 取消绘制菜单条按钮边框
	for i := 0; i < mb.X取子对象数量(); i++ {
		hele := mb.X取子对象并按索引(i)
		炫彩基类.X元素_启用绘制边框(hele, false)
		炫彩基类.X元素_注册事件C1(hele, 炫彩常量类.XE_BNCLICK, onMenuBarBnClick)
		炫彩基类.X元素_注册事件C1(hele, 炫彩常量类.XE_MENU_SELECT, onMenuSelect)
	}

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}

func onMenuBarBnClick(hEle int, pbHandled *bool) int {
	fmt.Println(炫彩基类.X按钮_取文本(hEle) + "被单击了")
	// 创建菜单
	menu := 炫彩组件类.X创建菜单()
	// 一级菜单
	menu.X添加项(10001, "item1", 0, 炫彩常量类.Menu_Item_Flag_Normal)
	menu.X添加项(10002, "item2", 0, 炫彩常量类.Menu_Item_Flag_Normal)

	// 获取按钮坐标
	var rc 炫彩基类.RECT
	炫彩基类.X元素_取窗口客户区坐标DPI(hEle, &rc)
	// 转换到屏幕坐标
	pt := 炫彩WinApi类.POINT{X: rc.Left, Y: rc.Bottom}
	炫彩WinApi类.X窗口取屏幕坐标(w.X取HWND(), &pt)
	// 弹出菜单
	menu.X弹出(w.X取HWND(), pt.X, pt.Y, hEle, 炫彩常量类.Menu_Popup_Position_Left_Top)
	return 0
}

// 菜单被选择事件
func onMenuSelect(hEle int, nID int32, pbHandled *bool) int {
	fmt.Println(炫彩基类.X按钮_取文本(hEle)+"下的菜单被选择:", nID)
	return 0
}
