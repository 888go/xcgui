// 拖放文件到窗口or元素.
package main

import (
	"fmt"
	"github.com/888go/xcgui/wapi/wutil"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a    *炫彩App类.App
	w    *炫彩窗口基类.Window
	edit *炫彩组件类.Edit
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)

	w = 炫彩窗口基类.X创建窗口(0, 0, 600, 600, "拖放文件到窗口or元素", 0, 炫彩常量类.Window_Style_Default|炫彩常量类.Window_Style_Drag_Window)
	// 创建编辑框.
	edit = 炫彩组件类.X创建编辑框(15, 40, 570, 300, w.Handle)
	// 编辑框允许多行.
	edit.X启用多行(true)

	// 窗口_启用拖放文件.
	w.X启用拖放文件(true)

	// 注册元素文件拖放事件
	edit.X事件_文件拖放事件1(onEleDropFiles)

	// 注册窗口文件拖放事件.
	// w.Event_DROPFILES1(onWndDropFiles)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

// 事件_元素文件拖放.
func onEleDropFiles(hEle int, hDropInfo uintptr, pbHandled *bool) int {
	fmt.Println("***************************************拖放文件到元素***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt 炫彩基类.POINT
	炫彩WinApi类.X拖放文件取鼠标位置(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	files := 炫彩WinApi工具类.X拖放文件取路径(hDropInfo)
	for _, v := range files {
		edit.X添加文本(v + "\n")
		fmt.Println("文件路径:", v)
	}
	return 0
}

// 事件_窗口文件拖放.
func onWndDropFiles(HXCGUI int, hDropInfo uintptr, pbHandled *bool) int {
	// win7在窗口拖放事件这里利用 [窗口_取鼠标停留元素] 可以实现对元素的拖放事件进行处理, 所以即使不注册元素拖放事件也行, 自己灵活使用..
	hEle := w.X取鼠标停留元素() // win10 好像获取不到, 那就还是去注册元素拖放事件, 不注册窗口拖放事件
	fmt.Println("鼠标停留元素句柄:", hEle)
	if hEle == edit.Handle {
		return onEleDropFiles(hEle, hDropInfo, pbHandled)
	}

	fmt.Println("***************************************拖放文件到窗口***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt 炫彩基类.POINT
	炫彩WinApi类.X拖放文件取鼠标位置(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	files := 炫彩WinApi工具类.X拖放文件取路径(hDropInfo)
	for _, v := range files {
		fmt.Println("文件路径:", v)
	}
	return 0
}
