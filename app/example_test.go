package 炫彩App类_test

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func ExampleNew() {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	w.X显示方式(炫彩常量类.SW_SHOW)
	a.X运行()
	a.X退出()
}

func ExampleApp_ShowAndRun() {
	a := 炫彩App类.X创建(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	a.X显示窗口并运行(w.Handle)
	a.X退出()
}
