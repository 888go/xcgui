// 窗口简单缓动
package main

import (
	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/ease"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
	"time"
)

var (
	a *app.App
	w *window.Window
)

func main() {
	a = app.I初始化(true)
	a.I置绘制频率(10)
	w = window.I窗口_创建(0, 0, 400, 300, "", 0, xcc.I常量_窗口样式_默认)
	// 窗口_调整布局
	w.I调整布局()
	w.I显示(xcc.I常量_窗口形式_显示并激活)

	time.AfterFunc(time.Millisecond*3, func() {
		// 获取窗口坐标
		var rect xc.RECT
		w.I取坐标(&rect)

		// 缓动
		for i := 1; i <= 30; i++ {
			v := ease.Bounce(float32(i)/30.0, xcc.I常量_缓动类型_快到慢)
			y := int(v * float32(rect.Top))

			w.I移动窗口(int(rect.Left), y)
			time.Sleep(time.Millisecond * 10)
		}
	})

	a.I运行()
	a.I退出()
}
