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
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 400, 300, "日期时间框", 0, 炫彩常量类.Window_Style_Default)

	dt := 炫彩组件类.NewDateTime(20, 50, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt.SetStyle(0)

	dt2 := 炫彩组件类.NewDateTime(200, 50, 120, 26, w.Handle)
	// 0为日期元素, 1为时间元素.
	dt2.SetStyle(1)

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}
