// 设置默认字体, 获取字体信息
package main

import (
	"fmt"

	"e.coding.net/gogit/go/xcgui/app"
	"e.coding.net/gogit/go/xcgui/font"
	"e.coding.net/gogit/go/xcgui/widget"
	"e.coding.net/gogit/go/xcgui/window"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

func main() {
	// 1.初始化UI库
	a := app.I初始化(true)

	// 创建字体
	f := font.NewFontEX("Arial", 11, xcc.I常量_字体样式_正常)
	// 设置程序默认字体
	a.I置默认字体(f.I句柄)

	// 2.创建窗口
	w := window.I窗口_创建(0, 0, 466, 300, "xcgui", 0, xcc.I常量_窗口样式_默认)

	// 创建一个按钮
	btn := widget.I按钮_创建(30, 50, 150, 30, "GetFontInfo", w.I句柄)
	btn.I事件_按钮被单击(func(pbHandled *bool) int {
		// 获取字体信息
		var fontInfo xc.Font_Info_
		f.GetFontInfo(&fontInfo)
		w.I消息框("Font Info", fmt.Sprintf("fontName=%s, fontSize=%d, fontStyle=%d", xc.Font_Info_Name(fontInfo.Name), fontInfo.NSize, fontInfo.NStyle), xcc.I常量_弹出消息框_确定按钮, xcc.I常量_窗口样式_弹出)
		return 0
	})

	// 3.显示窗口
	w.I显示(xcc.I常量_窗口形式_显示并激活)
	// 4.运行程序
	a.I运行()
	// 5.释放UI库
	a.I退出()
}
