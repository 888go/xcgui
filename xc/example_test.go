package 炫彩基类_test

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func ExampleXInitXCGUI() {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	炫彩基类.XLoadXCGUI()
	炫彩基类.X初始化(true)
	hWindow := 炫彩基类.X窗口_创建(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	炫彩基类.X窗口_显示方式(hWindow, 炫彩常量类.SW_SHOW)
	炫彩基类.X运行()
	炫彩基类.X退出()
}
