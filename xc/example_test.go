package xc_test

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
	xc.LoadXCGUI()
	xc.XInitXCGUI(true)
	hWindow := xc.XWnd_Create(0, 0, 500, 500, "", 0, 炫彩常量类.Window_Style_Default)
	xc.XWnd_ShowWindow(hWindow, 炫彩常量类.SW_SHOW)
	xc.XRunXCGUI()
	xc.XExitXCGUI()
}
