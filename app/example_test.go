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
	a := 炫彩App类.New(true)
	w := window.New(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	w.ShowWindow(炫彩常量类.SW_SHOW)
	a.Run()
	a.Exit()
}

func ExampleApp_ShowAndRun() {
	a := 炫彩App类.New(true)
	w := window.New(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	a.ShowAndRun(w.Handle)
	a.Exit()
}
