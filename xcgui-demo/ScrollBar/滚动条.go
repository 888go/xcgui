// 滚动条
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 430, 300, "ScrollBar", 0, xcc.I常量_窗口样式_默认)

	// 创建滚动条
	bar1 := widget.I滚动条_创建(12, 33, 300, 20, w.I句柄)
	bar2 := widget.I滚动条_创建(330, 33, 20, 240, w.I句柄)

	// 设置为垂直滚动条
	bar2.I置水平(false)

	// 注册滚动条元素滚动事件
	bar1.I事件_滚动条触发1(SBAR_SCROLL1)
	bar2.I事件_滚动条触发1(SBAR_SCROLL1)

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}

// 滚动条元素滚动事件
func SBAR_SCROLL1(hEle, pos int, pbHandled *bool) int {
	fmt.Println(pos)
	// 为了鼠标滚轮滚动和点击两端按钮实时显示效果而刷新
	xc.XEle_Redraw(hEle, true)
	return 0
}
