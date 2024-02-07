package tf

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

// TFunc 测试用程序.
//
//	@Description: 测试时使用的函数.
//	@param f
func TFunc(f func(a *炫彩App类.App, w *炫彩窗口基类.Window)) {
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 600, 400, "Test", 0, 炫彩常量类.Window_Style_Default)
	f(a, w)
	w.X显示(true)
	a.X运行()
	a.X退出()
}
