// 形状文本自动根据内容改变大小, 设置字体、字体大小、字体样式
package main

import (
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/font"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func main() {
	a := 炫彩App类.X创建(true)
	a.X启用DPI(true)
	a.X启用自动DPI(true)
	w := 炫彩窗口基类.X创建窗口(0, 0, 500, 500, "ShapeText", 0, 炫彩常量类.Window_Style_Default)

	st := 炫彩组件类.X创建形状文本(15, 35, 100, 30, "测试字体大小", w.Handle)
	// 自动根据内容改变大小
	st.X布局项置宽度(炫彩常量类.Layout_Size_Auto, -1)
	st.X布局项置高度(炫彩常量类.Layout_Size_Auto, -1)

	// 设置字体大小
	st.X置字体(炫彩字体类.X创建(50).Handle)
	// 设置个新字体， 粗体
	// st.SetFont(font.NewEX("幼圆", 50, xcc.FontStyle_Bold).Handle)

	shapeText := 炫彩组件类.X创建形状文本(15, 235, 150, 30, "测试文字自动换行测试文字自动换行测试文字自动换行测试文字自动换行测试文字自动换行测试文字自动换行测试文字自动换行测试文字自动换行", w.Handle)
	shapeText.X置文本对齐(炫彩常量类.TextAlignFlag_Left | 炫彩常量类.TextAlignFlag_Top) // 置文本对齐方式
	shapeText.X布局项置高度(炫彩常量类.Layout_Size_Auto, -1)               // 高度自动

	w.X显示(true)
	a.X运行()
	a.X退出()
}
