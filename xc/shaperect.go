package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// 形状矩形_创建, 创建矩形形状对象, 返回句柄.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func X形状矩形_创建(X坐标 int, Y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) int {
	r, _, _ := xShapeRect_Create.Call(uintptr(X坐标), uintptr(Y坐标), uintptr(宽度), uintptr(高度), uintptr(父对象句柄))
	return int(r)
}

// 形状矩形_置边框色, 设置边框颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func X形状矩形_置边框色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeRect_SetBorderColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状矩形_置填充色, 设置填充颜色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func X形状矩形_置填充色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeRect_SetFillColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状矩形_置圆角大小, 设置圆角大小.
//
// hShape: 形状对象句柄.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func X形状矩形_置圆角大小(形状对象句柄 int, 圆角宽度 int, 圆角高度 int32) int {
	r, _, _ := xShapeRect_SetRoundAngle.Call(uintptr(形状对象句柄), uintptr(圆角宽度), uintptr(圆角高度))
	return int(r)
}

// 形状矩形_取圆角大小, 获取圆角大小.
//
// hShape: 形状对象句柄.
//
// pWidth: 圆角宽度.
//
// pHeight: 圆角高度.
func X形状矩形_取圆角大小(形状对象句柄 int, 圆角宽度 *int, 圆角高度 *int32) int {
	r, _, _ := xShapeRect_GetRoundAngle.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(圆角宽度)), uintptr(unsafe.Pointer(圆角高度)))
	return int(r)
}

// 形状矩形_启用边框, 启用绘制矩形边框.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func X形状矩形_启用边框(形状对象句柄 int, 是否启用 bool) int {
	r, _, _ := xShapeRect_EnableBorder.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 形状矩形_启用填充, 启用填充矩形.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func X形状矩形_启用填充(形状对象句柄 int, 是否启用 bool) int {
	r, _, _ := xShapeRect_EnableFill.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 形状矩形_启用圆角.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func X形状矩形_启用圆角(形状对象句柄 int, 是否启用 bool) int {
	r, _, _ := xShapeRect_EnableRoundAngle.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}
