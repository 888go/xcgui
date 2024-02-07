package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// XUI_SetStyle 可视对象_置样式, 设置UI对象样式.
//
//	@param hXCGUI 对象句柄.
//	@param nStyle xcc.XC_OBJECT_STYLE, 样式值: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
//	@return int
func X可视对象_置样式(对象句柄 int, 样式 炫彩常量类.XC_OBJECT_STYLE) int {
	r, _, _ := xUI_SetStyle.Call(uintptr(对象句柄), uintptr(样式))
	return int(r)
}

// XUI_GetStyle 可视对象_取样式, 获取UI对象样式
//
//	@param hXCGUI 对象句柄.
//	@return xcc.XC_OBJECT_STYLE 返回: xcc.Button_Style_ , xcc.Element_Style_ , xcc.ListBox_Style_.
func X可视对象_取样式(对象句柄 int) 炫彩常量类.XC_OBJECT_STYLE {
	r, _, _ := xUI_GetStyle.Call(uintptr(对象句柄))
	return 炫彩常量类.XC_OBJECT_STYLE(r)
}

// XUI_EnableCSS 可视对象_启用CSS, 启用或禁用样式.
//
//	@param hXCGUI 对象句柄.
//	@param bEnable 是否启用.
//	@return int
func X可视对象_启用CSS(对象句柄 int, 是否启用 bool) int {
	r, _, _ := xUI_EnableCSS.Call(uintptr(对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// XUI_SetCssName 可视对象_置CSS名称, 设置CSS[套用样式]名称.
//
//	@param hXCGUI 对象句柄.
//	@param pName 套用样式名称.
//	@return int
func X可视对象_置CSS名称(对象句柄 int, 套用样式名称 string) int {
	r, _, _ := xUI_SetCssName.Call(uintptr(对象句柄), 炫彩工具类.StrPtr(套用样式名称))
	return int(r)
}

// XUI_GetCssName 可视对象_取CSS名称, 获取CSS样式名称.
//
//	@param hXCGUI 对象句柄.
//	@return string
func X可视对象_取CSS名称(对象句柄 int) string {
	r, _, _ := xUI_GetCssName.Call(uintptr(对象句柄))
	return 炫彩工具类.UintPtrToString(r)
}
