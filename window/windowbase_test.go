package window_test

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
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		rand.Seed(time.Now().Unix())
		炫彩组件类.NewButton(50, 100, 100, 30, "SetSize", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			width := rand.Int31n(400) + 200
			height := rand.Int31n(400) + 200
			w.SetSize(width, height)
			return 0
		})
	})
}

func Test_windowBase_SetWidth(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		rand.Seed(time.Now().Unix())
		炫彩组件类.NewButton(50, 100, 100, 30, "SetWidth", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			width := rand.Int31n(400) + 200
			w.SetWidth(width)
			return 0
		})
	})
}

func Test_windowBase_SetHeight(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		rand.Seed(time.Now().Unix())
		炫彩组件类.NewButton(50, 100, 100, 30, "SetHeight", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			height := rand.Int31n(400) + 200
			w.SetHeight(height)
			return 0
		})
	})
}

func Test_windowBase_SetLeft(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		炫彩组件类.NewButton(50, 100, 100, 30, "SetLeft+5", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			w.SetLeft(w.GetLeft() + 5)
			return 0
		})
	})
}

func Test_windowBase_SetTopEdge(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *window.Window) {
		炫彩组件类.NewButton(50, 100, 100, 30, "SetTopEdge+5", w.Handle).Event_BnClick(func(pbHandled *bool) int {
			w.SetTopEdge(w.GetTop() + 5)
			return 0
		})
	})
}
