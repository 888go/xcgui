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
	w *炫彩窗口基类.Window
)

func main() {
	a = 炫彩App类.X创建(true)
	a.X置绘制频率(10)
	w = 炫彩窗口基类.X创建窗口(0, 0, 400, 300, "", 0, 炫彩常量类.Window_Style_Default)
	w.X显示方式(炫彩常量类.SW_SHOW)

	time.AfterFunc(time.Millisecond*3, func() {
		// 获取窗口坐标
		var rect 炫彩基类.RECT
		w.X取坐标(&rect)

		// 缓动
		for i := 1; i <= 30; i++ {
			v := 炫彩缓动类.X缓动弹跳(float32(i)/30.0, 炫彩常量类.Ease_Type_Out)
			y := int32(v * float32(rect.Top))

			w.X移动窗口(rect.Left, y)
			time.Sleep(time.Millisecond * 10)
		}
	})

	a.X运行()
	a.X退出()
}
