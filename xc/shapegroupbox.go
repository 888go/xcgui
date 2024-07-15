package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// 形状组框_创建, 创建组框形状对象, 返回句柄.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 名称.
//
// hParent: 父对象句柄.

// ff:形状组框_创建
// hParent:父对象句柄
// pName:名称
// cy:高度
// cx:宽度
// y:Y坐标
// x:坐标
func X形状组框_创建(坐标 int, Y坐标 int, 宽度 int, 高度 int, 名称 string, 父对象句柄 int) int {
	r, _, _ := xShapeGroupBox_Create.Call(uintptr(坐标), uintptr(Y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(名称), uintptr(父对象句柄))
	return int(r)
}

// 形状组框_置边框颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.

// ff:形状组框_置边框颜色
// color:ABGR颜色值
// hShape:形状对象句柄
func X形状组框_置边框颜色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeGroupBox_SetBorderColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状组框_置文本颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.

// ff:形状组框_置文本颜色
// color:ABGR颜色值
// hShape:形状对象句柄
func X形状组框_置文本颜色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeGroupBox_SetTextColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状组框_置字体.
//
// hShape: 形状对象句柄.
//
// hFontX: 炫彩字体.

// ff:形状组框_置字体
// hFontX:炫彩字体
// hShape:形状对象句柄
func X形状组框_置字体(形状对象句柄 int, 炫彩字体 int) int {
	r, _, _ := xShapeGroupBox_SetFontX.Call(uintptr(形状对象句柄), uintptr(炫彩字体))
	return int(r)
}

// 形状组框_置文本偏移, 设置文本偏移量.
//
// hShape: 形状对象句柄.
//
// offsetX: 水平偏移.
//
// offsetY: 垂直偏移.

// ff:形状组框_置文本偏移
// offsetY:垂直偏移
// offsetX:水平偏移
// hShape:形状对象句柄
func X形状组框_置文本偏移(形状对象句柄 int, 水平偏移 int32, 垂直偏移 int32) int {
	r, _, _ := xShapeGroupBox_SetTextOffset.Call(uintptr(形状对象句柄), uintptr(水平偏移), uintptr(垂直偏移))
	return int(r)
}

// 形状组框_置圆角大小.
//
// hShape: 形状对象句柄.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.

// ff:形状组框_置圆角大小
// nHeight:圆角高度
// nWidth:圆角宽度
// hShape:形状对象句柄
func X形状组框_置圆角大小(形状对象句柄 int, 圆角宽度 int32, 圆角高度 int32) int {
	r, _, _ := xShapeGroupBox_SetRoundAngle.Call(uintptr(形状对象句柄), uintptr(圆角宽度), uintptr(圆角高度))
	return int(r)
}

// 形状组框_置文本.
//
// hShape: 形状对象句柄.
//
// pText: 文本内容.

// ff:形状组框_置文本
// pText:文本内容
// hShape:形状对象句柄
func X形状组框_置文本(形状对象句柄 int, 文本内容 string) int {
	r, _, _ := xShapeGroupBox_SetText.Call(uintptr(形状对象句柄), 炫彩工具类.StrPtr(文本内容))
	return int(r)
}

// 形状组框_取文本偏移, 获取文本偏移量.
//
// hShape: 形状对象句柄.
//
// pOffsetX: X坐标偏移量.
//
// pOffsetY: Y坐标偏移量.

// ff:形状组框_取文本偏移
// pOffsetY:Y坐标偏移量
// pOffsetX:坐标偏移量
// hShape:形状对象句柄
func X形状组框_取文本偏移(形状对象句柄 int, 坐标偏移量 *int32, Y坐标偏移量 *int32) int {
	r, _, _ := xShapeGroupBox_GetTextOffset.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(坐标偏移量)), uintptr(unsafe.Pointer(Y坐标偏移量)))
	return int(r)
}

// 形状组框_取圆角大小.
//
// hShape: 形状对象句柄.
//
// pWidth: 返回圆角宽度.
//
// pHeight: 返回圆角高度.

// ff:形状组框_取圆角大小
// pHeight:返回圆角高度
// pWidth:返回圆角宽度
// hShape:形状对象句柄
func X形状组框_取圆角大小(形状对象句柄 int, 返回圆角宽度 *int32, 返回圆角高度 *int32) int {
	r, _, _ := xShapeGroupBox_GetRoundAngle.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(返回圆角宽度)), uintptr(unsafe.Pointer(返回圆角高度)))
	return int(r)
}

// 形状组框_启用圆角.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.

// ff:形状组框_启用圆角
// bEnable:是否启用
// hShape:形状对象句柄
func X形状组框_启用圆角(形状对象句柄 int, 是否启用 bool) int {
	r, _, _ := xShapeGroupBox_EnableRoundAngle.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}
