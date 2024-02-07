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
	w    *window.Window
	text *炫彩组件类.ShapeText
)

func main() {
	a = 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)
	w = window.New(0, 0, 500, 130, "定时器", 0, 炫彩常量类.Window_Style_Default)

	text = 炫彩组件类.NewShapeText(50, 50, 400, 50, time.Now().Format("2006-01-02 15:04:05"), w.Handle)
	text.SetFont(炫彩字体类.New(30).Handle)

	// 定时器id是自己定的
	w.SetTimer(1, 1000)
	w.Event_TIMER(func(nIDEvent uint, pbHandled *bool) int {
		switch nIDEvent {
		case 1:
			text.SetText(time.Now().Format("2006-01-02 15:04:05"))
			text.Redraw()
		}
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}
