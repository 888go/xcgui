package 炫彩窗口基类_test

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/tf"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"math/rand"
	"testing"
	"time"
)

func Test_windowBase_SetSize(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		rand.Seed(time.Now().Unix())
		炫彩组件类.X创建按钮(50, 100, 100, 30, "SetSize", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			width := rand.Int31n(400) + 200
			height := rand.Int31n(400) + 200
			w.X置大小(width, height)
			return 0
		})
	})
}

func Test_windowBase_SetWidth(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		rand.Seed(time.Now().Unix())
		炫彩组件类.X创建按钮(50, 100, 100, 30, "SetWidth", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			width := rand.Int31n(400) + 200
			w.X置宽度(width)
			return 0
		})
	})
}

func Test_windowBase_SetHeight(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		rand.Seed(time.Now().Unix())
		炫彩组件类.X创建按钮(50, 100, 100, 30, "SetHeight", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			height := rand.Int31n(400) + 200
			w.X置高度(height)
			return 0
		})
	})
}

func Test_windowBase_SetLeft(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		炫彩组件类.X创建按钮(50, 100, 100, 30, "SetLeft+5", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			w.X置左边(w.X取左边() + 5)
			return 0
		})
	})
}

func Test_windowBase_SetTopEdge(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		炫彩组件类.X创建按钮(50, 100, 100, 30, "SetTopEdge+5", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			w.X置顶边(w.X取顶边() + 5)
			return 0
		})
	})
}
