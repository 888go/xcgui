package xc_test

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func ExampleXInitXCGUI() {
	// 可自定义xcgui.dll的路径, 如无需自定义, 则删掉这句代码.
	/*	err := xc.SetXcguiPath(`C:\Users\Administrator\Desktop\XCGUI.dll`)
		if err != nil {
			panic(err)
		}*/
	xc.LoadXCGUI()
	xc.XInitXCGUI(true)
	hWindow := xc.XWnd_Create(0, 0, 500, 500, "", 0, xcc.I常量_窗口样式_默认)
	xc.XWnd_ShowWindow(hWindow, xcc.I常量_窗口形式_显示并激活)
	xc.XRunXCGUI()
	xc.XExitXCGUI()
}
