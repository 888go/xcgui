package 炫彩缓动类_test//bm:炫彩缓动类_test

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/ease"
	"github.com/888go/xcgui/tf"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
	"testing"
	"time"
)

func TestEx(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		time.AfterFunc(time.Millisecond*3, func() {
			// 获取窗口坐标
			var rect 炫彩基类.RECT
			w.X取坐标(&rect)

			for i := 0; i < 30; i++ {
				v := 炫彩缓动类.X缓动EX(float32(i)/30.0, 炫彩常量类.Ease_Flag_Back|炫彩常量类.Ease_Flag_Out)
				y := int32(v * float32(rect.Top))
				w.X移动窗口(rect.Left, y)
				time.Sleep(time.Millisecond * 10)
			}
		})
	})
}
