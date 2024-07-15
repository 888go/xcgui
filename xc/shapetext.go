package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 形状文本_创建, 创建形状对象文本.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父对象句柄.

// ff:形状文本_创建
// hParent:父对象句柄
// pName:文本内容
// cy:高度
// cx:宽度
// y:Y坐标
// x:坐标
func X形状文本_创建(坐标, Y坐标, 宽度, 高度 int32, 文本内容 string, 父对象句柄 int) int {
	r, _, _ := xShapeText_Create.Call(uintptr(坐标), uintptr(Y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(文本内容), uintptr(父对象句柄))
	return int(r)
}

// 形状文本_置文本, 设置文本内容.
//
// hTextBlock: 形状对象文本句柄.
//
// pName: 文本内容.

// ff:形状文本_置文本
// pName:文本内容
// hTextBlock:形状对象文本句柄
func X形状文本_置文本(形状对象文本句柄 int, 文本内容 string) int {
	r, _, _ := xShapeText_SetText.Call(uintptr(形状对象文本句柄), 炫彩工具类.StrPtr(文本内容))
	return int(r)
}

// 形状文本_取文本, 获取文本内容.
//
// hTextBlock: 形状对象文本句柄.

// ff:形状文本_取文本
// hTextBlock:形状对象文本句柄
func X形状文本_取文本(形状对象文本句柄 int) string {
	r, _, _ := xShapeText_GetText.Call(uintptr(形状对象文本句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 形状文本_取文本长度, 获取文本长度.
//
// hTextBlock: 形状对象文本句柄.

// ff:形状文本_取文本长度
// hTextBlock:形状对象文本句柄
func X形状文本_取文本长度(形状对象文本句柄 int) int {
	r, _, _ := xShapeText_GetTextLength.Call(uintptr(形状对象文本句柄))
	return int(r)
}

// 形状文本_置字体.
//
// hTextBlock: 形状对象文本句柄.
//
// hFontx: 字体句柄.

// ff:形状文本_置字体
// hFontx:字体句柄
// hTextBlock:形状对象文本句柄
func X形状文本_置字体(形状对象文本句柄 int, 字体句柄 int) int {
	r, _, _ := xShapeText_SetFont.Call(uintptr(形状对象文本句柄), uintptr(字体句柄))
	return int(r)
}

// 形状文本_取字体, 返回字体句柄.
//
// hTextBlock: 形状对象文本句柄.

// ff:形状文本_取字体
// hTextBlock:形状对象文本句柄
func X形状文本_取字体(形状对象文本句柄 int) int {
	r, _, _ := xShapeText_GetFont.Call(uintptr(形状对象文本句柄))
	return int(r)
}

// 形状文本_置文本颜色, 设置文本颜色.
//
// hTextBlock: 形状对象文本句柄.
//
// color: ABGR 颜色值.

// ff:形状文本_置文本颜色
// color:ABGR颜色值
// hTextBlock:形状对象文本句柄
func X形状文本_置文本颜色(形状对象文本句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeText_SetTextColor.Call(uintptr(形状对象文本句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状文本_取文本颜色.
//
// hTextBlock: 形状对象文本句柄.

// ff:形状文本_取文本颜色
// hTextBlock:形状对象文本句柄
func X形状文本_取文本颜色(形状对象文本句柄 int) int {
	r, _, _ := xShapeText_GetTextColor.Call(uintptr(形状对象文本句柄))
	return int(r)
}

// 形状文本_置文本对齐.
//
// hTextBlock: 形状对象文本句柄.
//
// align: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.

// ff:形状文本_置文本对齐
// align:文本对齐方式
// hTextBlock:形状对象文本句柄
func X形状文本_置文本对齐(形状对象文本句柄 int, 文本对齐方式 炫彩常量类.TextFormatFlag_) int {
	r, _, _ := xShapeText_SetTextAlign.Call(uintptr(形状对象文本句柄), uintptr(文本对齐方式))
	return int(r)
}

// 形状文本_置偏移, 设置内容偏移.
//
// hTextBlock: 形状对象文本句柄.
//
// x: X坐标.
//
// y: Y坐标.

// ff:形状文本_置偏移
// y:Y坐标
// x:坐标
// hTextBlock:形状对象文本句柄
func X形状文本_置偏移(形状对象文本句柄 int, 坐标 int, Y坐标 int) int {
	r, _, _ := xShapeText_SetOffset.Call(uintptr(形状对象文本句柄), uintptr(坐标), uintptr(Y坐标))
	return int(r)
}
