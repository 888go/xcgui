// 模仿选择夹: 仅为演示ScrollView, TabBar的使用, 真正的选择夹不是以这种方式实现.
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

var (
	sv  *widget.ScrollView
	tab *widget.TabBar
)

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)
	// 2.创建窗口
	w := window.I窗口_创建(0, 0, 400, 270, "xc", 0, xcc.I常量_窗口样式_简单|xcc.I常量_窗口样式_按钮_关闭)

	// 创建选择夹顶部Tab条
	tab = widget.TAB条_创建(10, 33, 380, 28, w.I句柄)
	tab.I添加标签("Page 1")
	tab.I添加标签("Page 2")
	tab.I置选择(0)

	// 注册选择夹顶部按钮事件
	tab.I事件_标签按钮选择改变(tabBarSelect)

	// 创建滚动视图, 即ScrollView
	sv = widget.I滚动视图_创建(10, 60, 380, 200, w.I句柄)
	// 隐藏滚动视图的纵/横滚动条
	sv.I显示水平滚动条(false)
	sv.I显示垂直滚动条(false)
	// 设置视图内容大小, 视图内容总高度400, 可分成两页, 每页高200
	sv.I置视图大小(380, 400)
	// 禁用接收鼠标滚轮事件, 防止用户手动滚动视图
	sv.I启用事件_XE_MOUSEWHEEL(false)

	// 第一页, 从0开始
	widget.I按钮_创建(10, 10, 100, 20, "Button1", sv.I句柄)
	widget.I按钮_创建(10, 40, 100, 20, "Button2", sv.I句柄)
	widget.I按钮_创建(10, 70, 100, 20, "Button3", sv.I句柄)
	// 第二页, 从200开始
	widget.I按钮_创建(10, 200+10, 100, 20, "Btn4", sv.I句柄)
	widget.I按钮_创建(10, 200+40, 100, 20, "Btn5", sv.I句柄)
	widget.I按钮_创建(10, 200+70, 100, 20, "Btn6", sv.I句柄)

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}

func tabBarSelect(iItem int, pbHandled *bool) int {
	switch iItem {
	case 0:
		sv.I垂直滚动到Y(0)
	case 1:
		sv.I垂直滚动到Y(200)
	}
	return 0
}
