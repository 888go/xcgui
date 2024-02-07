// 设置默认字体, 获取字体信息
package main

import (
	"fmt"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/font"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := 炫彩App类.New(true)
	a.EnableDPI(true)
	a.EnableAutoDPI(true)

	// 创建字体
	f := 炫彩字体类.NewEX("Arial", 11, 炫彩常量类.FontStyle_Regular)
	// 设置程序默认字体
	a.SetDefaultFont(f.Handle)

	// 2.创建窗口
	w := window.New(0, 0, 466, 300, "xcgui", 0, 炫彩常量类.Window_Style_Default)

	// 创建一个按钮
	btn := 炫彩组件类.NewButton(30, 50, 150, 30, "GetFontInfo", w.Handle)
	btn.Event_BnClick(func(pbHandled *bool) int {
		// 获取字体信息
		var fontInfo xc.Font_Info_
		f.GetFontInfo(&fontInfo)
		w.MessageBox("Font Info", fmt.Sprintf("fontName=%s, fontSize=%d, fontStyle=%d", xc.Font_Info_Name(fontInfo.Name), fontInfo.NSize, fontInfo.NStyle), 炫彩常量类.MessageBox_Flag_Ok, 炫彩常量类.Window_Style_Pop)
		return 0
	})

	// 3.显示窗口
	w.ShowWindow(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.Run()
	// 5.释放UI库
	a.Exit()
}
