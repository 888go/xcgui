package 炫彩基类_test

import (
	"fmt"
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/tf"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"testing"
)

func TestSetBnClicks(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		a.X显示布局边界(true)
		炫彩组件类.X创建按钮(100, 50, 100, 30, "窗口中按钮", w.Handle)

		lay1 := 炫彩组件类.X创建布局(50, 80, 400, 300, w.Handle)
		炫彩组件类.X创建按钮(0, 0, 100, 30, "布局元素1中按钮", lay1.Handle)

		lay2 := 炫彩组件类.X创建布局(0, 0, 300, 200, lay1.Handle)
		炫彩组件类.X创建按钮(0, 0, 100, 30, "布局元素2中按钮", lay2.Handle)

		炫彩基类.SetBnClicks(w.Handle, func(hEle int, pbHandled *bool) int {
			fmt.Println("被单击按钮:", 炫彩基类.X按钮_取文本(hEle))
			return 0
		})
	})
}
