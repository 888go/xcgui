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
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)

	// 创建字体
	f := 炫彩字体类.X创建EX("Arial", 11, 炫彩常量类.FontStyle_Regular)
	// 设置程序默认字体
	a.X置默认字体(f.Handle)

	// 2.创建窗口
	w := 炫彩窗口基类.X创建窗口(0, 0, 466, 300, "xcgui", 0, 炫彩常量类.Window_Style_Default)

	// 创建一个按钮
	btn := 炫彩组件类.X创建按钮(30, 50, 150, 30, "GetFontInfo", w.Handle)
	btn.X事件_被单击(func(pbHandled *bool) int {
		// 获取字体信息
		var fontInfo 炫彩基类.Font_Info_
		f.X取信息(&fontInfo)
		w.X消息框("Font Info", fmt.Sprintf("fontName=%s, fontSize=%d, fontStyle=%d", 炫彩基类.Font_Info_Name(fontInfo.Name), fontInfo.NSize, fontInfo.NStyle), 炫彩常量类.MessageBox_Flag_Ok, 炫彩常量类.Window_Style_Pop)
		return 0
	})

	// 3.显示窗口
	w.X显示方式(炫彩常量类.SW_SHOW)
	// 4.运行程序
	a.X运行()
	// 5.释放UI库
	a.X退出()
}
