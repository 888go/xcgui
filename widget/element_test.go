package 炫彩组件类_test

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/tf"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"testing"
)

func TestElement_SetFocus(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		layout := 炫彩组件类.NewLayoutEle(50, 50, 150, 50, w.Handle)
		edit := 炫彩组件类.NewEdit(0, 0, 100, 30, layout.Handle)
		炫彩组件类.NewButton(50, 100, 100, 30, "setfocus", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			edit.SetFocus()
			return 0
		})
	})
}

func TestElement_SetLeft(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		btn := 炫彩组件类.NewButton(10, 40, 100, 30, "setleft+5", w.Handle)
		btn.Event_BnClick(func(pbHandled *bool) int {
			btn.SetLeft(btn.GetLeft()+5, true)
			return 0
		})
	})
}

func TestElement_SetTop(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		btn2 := 炫彩组件类.NewButton(10, 80, 100, 30, "settop+5", w.Handle)
		btn2.Event_BnClick(func(pbHandled *bool) int {
			btn2.SetTop(btn2.GetTop()+5, true)
			return 0
		})
	})
}
