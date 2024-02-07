// 定时器
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/font"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
	"time"
)

var (
	a    *炫彩App类.App
	w    *炫彩窗口基类.Window
	text *炫彩组件类.ShapeText
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 500, 130, "定时器", 0, 炫彩常量类.Window_Style_Default)

	text = 炫彩组件类.X创建形状文本(50, 50, 400, 50, time.Now().Format("2006-01-02 15:04:05"), w.Handle)
	text.X置字体(炫彩字体类.X创建(30).Handle)

	// 定时器id是自己定的
	w.X置定时器(1, 1000)
	w.X线程_定时器消息(func(nIDEvent uint, pbHandled *bool) int {
		switch nIDEvent {
		case 1:
			text.X置文本(time.Now().Format("2006-01-02 15:04:05"))
			text.X重绘()
		}
		return 0
	})

	w.X显示(true)
	a.X运行()
	a.X退出()
}
