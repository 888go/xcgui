// 月历卡片
// 要美化的话, 就得自绘, 看这个: http://www.xcgui.com/doc-ui/page_draw__month_cal.html
// 我以后有空会翻译几个好看的: http://mall.xcgui.com/1618.html
package main

import (
	"fmt"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w := window.New(0, 0, 400, 300, "月历卡片", 0, 炫彩常量类.Window_Style_Default)

	// 创建MonthCal
	monthCal := 炫彩组件类.NewMonthCal(30, 40, 290, 240, w.Handle)
	// 注册月历元素日期改变事件
	monthCal.Event_MONTHCAL_CHANGE(func(pbHandled *bool) int {
		// 获取被选择的年月日
		var pnYear, pnMonth, pnDay int32
		monthCal.GetSelDate(&pnYear, &pnMonth, &pnDay)

		fmt.Printf("%d年%d月%d日\n", pnYear, pnMonth, pnDay)
		return 0
	})

	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}
