package 炫彩组件类_test

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/tf"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"testing"
)

func TestElement_SetFocus(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		layout := 炫彩组件类.X创建布局(50, 50, 150, 50, w.Handle)
		edit := 炫彩组件类.X创建编辑框(0, 0, 100, 30, layout.Handle)
		炫彩组件类.X创建按钮(50, 100, 100, 30, "setfocus", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			edit.X置焦点()
			return 0
		})
	})
}

func TestElement_SetLeft(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		btn := 炫彩组件类.X创建按钮(10, 40, 100, 30, "setleft+5", w.Handle)
		btn.X事件_被单击(func(pbHandled *bool) int {
			btn.X置左边(btn.X取左边()+5, true)
			return 0
		})
	})
}

func TestElement_SetTop(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		btn2 := 炫彩组件类.X创建按钮(10, 80, 100, 30, "settop+5", w.Handle)
		btn2.X事件_被单击(func(pbHandled *bool) int {
			btn2.X置顶边(btn2.X取顶边()+5, true)
			return 0
		})
	})
}
