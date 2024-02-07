package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 形状圆_创建, 创建圆形形状对象, 返回句柄.
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
func X形状圆_创建(X坐标 int, Y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) int {
	r, _, _ := xShapeEllipse_Create.Call(uintptr(X坐标), uintptr(Y坐标), uintptr(宽度), uintptr(高度), uintptr(父对象句柄))
	return int(r)
}

// 形状圆_置边框色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func X形状圆_置边框色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeEllipse_SetBorderColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状圆_置填充色.
//
// hShape: 形状对象句柄.
//
// color: ABGR 颜色值.
func X形状圆_置填充色(形状对象句柄 int, ABGR颜色值 int) int {
	r, _, _ := xShapeEllipse_SetFillColor.Call(uintptr(形状对象句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 形状圆_启用边框, 启用绘制圆边框.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func X形状圆_启用边框(形状对象句柄 int, 是否启用 bool) int {
	r, _, _ := xShapeEllipse_EnableBorder.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 形状圆_启用填充, 启用填充圆.
//
// hShape: 形状对象句柄.
//
// bEnable: 是否启用.
func X形状圆_启用填充(形状对象句柄 int, 是否启用 bool) int {
	r, _, _ := xShapeEllipse_EnableFill.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}
