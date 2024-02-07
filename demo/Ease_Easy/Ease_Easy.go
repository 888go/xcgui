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
	a *app.App
	w *window.Window
)

func main() {
	a = app.New(true)
	a.SetPaintFrequency(10)
	w = window.New(0, 0, 400, 300, "", 0, xcc.Window_Style_Default)
	w.ShowWindow(xcc.SW_SHOW)

	time.AfterFunc(time.Millisecond*3, func() {
		// 获取窗口坐标
		var rect xc.RECT
		w.GetRect(&rect)

		// 缓动
		for i := 1; i <= 30; i++ {
			v := ease.Bounce(float32(i)/30.0, xcc.Ease_Type_Out)
			y := int32(v * float32(rect.Top))

			w.SetPosition(rect.Left, y)
			time.Sleep(time.Millisecond * 10)
		}
	})

	a.Run()
	a.Exit()
}
