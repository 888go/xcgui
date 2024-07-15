// 日期时间框
// 要美化的话, 就得自绘, 看这个: http://www.xcgui.com/doc-ui/page_draw__month_cal.html
// 我以后有空会翻译几个好看的: http://mall.xcgui.com/1618.html
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 400, 300, "日期时间框", 0, 炫彩常量类.Window_Style_Default)

	dt := 炫彩组件类.X创建日期(20, 50, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt.X置样式(0)

	dt2 := 炫彩组件类.X创建日期(200, 50, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt2.X置样式(1)

	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}
