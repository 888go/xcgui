// 月历卡片
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	a := app.I初始化(true)
	w := window.I窗口_创建(0, 0, 400, 300, "", 0, xcc.I常量_窗口样式_默认)

	// 创建MonthCal
	monthCal := widget.I月历_创建(30, 40, 290, 240, w.I句柄)
	// 注册月历元素日期改变事件
	monthCal.I事件_日期改变(func(pbHandled *bool) int {
		// 获取被选择的年月日
		var pnYear, pnMonth, pnDay int
		monthCal.I取选择日期(&pnYear, &pnMonth, &pnDay)

		fmt.Printf("%d年%d月%d日\n", pnYear, pnMonth, pnDay)
		return 0
	})

	w.I显示(xcc.I常量_窗口形式_显示并激活)
	a.I运行()
	a.I退出()
}
