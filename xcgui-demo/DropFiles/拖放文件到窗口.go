// 拖放文件到窗口or元素.
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/wapi"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	a *app.App
	w *window.Window

	edit *widget.Edit
)

func main() {
	a = app.I初始化(true)
	defer a.I退出()

	w = window.I窗口_创建(0, 0, 600, 600, "拖放文件到窗口or元素", 0, xcc.I常量_窗口样式_默认)
	// 窗口_启用拖放文件.
	w.I启用拖放文件(true)
	// 注册窗口文件拖放事件.
	w.I事件_拖动文件到窗口1(事件_窗口文件拖放)

	// 创建编辑框.
	edit = widget.I编辑框_创建(15, 40, 570, 300, w.I句柄)
	// 编辑框允许多行.
	edit.I启用多行(true)
	// 注册元素文件拖放事件, 可注册也可不注册, 因为在窗口拖放事件里也可以处理元素的, 根据自己的需求来吧, 很灵活.
	// edit.I事件_拖动文件到窗口1(onEleDropFiles)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
}

// 事件_窗口文件拖放.
func 事件_窗口文件拖放(HXCGUI, hDropInfo int, pbHandled *bool) int {
	// 在窗口拖放事件这里可以实现对其他元素的拖放事件进行处理, 所以即使不注册元素拖放事件也行, 自己灵活使用..
	// 窗口_取鼠标停留元素.
	hEle := w.I取鼠标停留元素()
	fmt.Println("鼠标停留元素句柄:", hEle)
	if hEle == edit.I句柄 {
		return onEleDropFiles(hEle, hDropInfo, pbHandled)
	}

	fmt.Println("***************************************拖放文件到窗口***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt xc.POINT
	wapi.DragQueryPoint(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	// 循环获取拖放进窗口的所有文件.
	i := 0
	for {
		filePath := ""
		length := wapi.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}

		fmt.Println("文件路径:", filePath)
		i++ // 索引+1检索下一个文件
	}

	wapi.DragFinish(hDropInfo)
	return 0
}

// 事件_元素文件拖放.
func onEleDropFiles(HXCGUI, hDropInfo int, pbHandled *bool) int {
	fmt.Println("***************************************拖放文件到元素***************************************")
	// 获取拖放文件到窗口时鼠标的坐标.
	var pt xc.POINT
	wapi.DragQueryPoint(hDropInfo, &pt)
	fmt.Println("鼠标坐标:", pt)

	// 循环获取拖放进元素的所有文件.
	i := 0
	for {
		filePath := ""
		length := wapi.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}

		edit.I添加文本(filePath + "\r\n")
		fmt.Println("文件路径:", filePath)
		i++ // 索引+1检索下一个文件
	}

	wapi.DragFinish(hDropInfo)
	return 0
}
