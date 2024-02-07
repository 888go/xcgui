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
func TFunc(f func(a *炫彩App类.App, w *window.Window)) {
	a := 炫彩App类.New(true)
	w := window.New(0, 0, 600, 400, "Test", 0, 炫彩常量类.Window_Style_Default)
	f(a, w)
	w.Show(true)
	a.Run()
	a.Exit()
}
