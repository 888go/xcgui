// 窗口简单缓动
package main

import (
	"time"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/ease"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

var (
	a *炫彩App类.App
	w *window.Window
)

func main() {
	a = 炫彩App类.New(true)
	a.SetPaintFrequency(10)
	w = window.New(0, 0, 400, 300, "", 0, 炫彩常量类.Window_Style_Default)
	w.ShowWindow(炫彩常量类.SW_SHOW)

	time.AfterFunc(time.Millisecond*3, func() {
		// 获取窗口坐标
		var rect xc.RECT
		w.GetRect(&rect)

		// 缓动
		for i := 1; i <= 30; i++ {
			v := 炫彩缓动类.Bounce(float32(i)/30.0, 炫彩常量类.Ease_Type_Out)
			y := int32(v * float32(rect.Top))

			w.SetPosition(rect.Left, y)
			time.Sleep(time.Millisecond * 10)
		}
	})

	a.Run()
	a.Exit()
}
