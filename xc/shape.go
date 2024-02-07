package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// 形状_移除, 从父UI元素或窗口,和父布局对象中移除.
//
// hShape: 形状对象句柄.
func X形状_移除(形状对象句柄 int) int {
	r, _, _ := xShape_RemoveShape.Call(uintptr(形状对象句柄))
	return int(r)
}

// 形状_取Z序, 获取形状对象Z序, 成功返回索引值, 否则返回 XC_ID_ERROR.
//
// hShape: 形状对象句柄.
func X形状_取Z序(形状对象句柄 int) int {
	r, _, _ := xShape_GetZOrder.Call(uintptr(形状对象句柄))
	return int(r)
}

// 形状_重绘, 重绘形状对象.
//
// hShape: 形状对象句柄.
func X形状_重绘(形状对象句柄 int) int {
	r, _, _ := xShape_Redraw.Call(uintptr(形状对象句柄))
	return int(r)
}

// 形状_取宽度, 获取内容宽度.
//
// hShape: 形状对象句柄.
func X形状_取宽度(形状对象句柄 int) int32 {
	r, _, _ := xShape_GetWidth.Call(uintptr(形状对象句柄))
	return int32(r)
}

// 形状_取高度, 获取内容高度.
//
// hShape: 形状对象句柄.
func X形状_取高度(形状对象句柄 int) int32 {
	r, _, _ := xShape_GetHeight.Call(uintptr(形状对象句柄))
	return int32(r)
}

// 形状_移动位置.
//
// hShape: 形状对象句柄.
//
// x: x坐标.
//
// y: y坐标.
func X形状_移动位置(形状对象句柄 int, x坐标, y坐标 int32) int {
	r, _, _ := xShape_SetPosition.Call(uintptr(形状对象句柄), uintptr(x坐标), uintptr(y坐标))
	return int(r)
}

// 形状_取坐标.
//
// hShape: 形状对象句柄.
//
// pRect: 接收返回坐标.
func X形状_取坐标(形状对象句柄 int, 接收返回坐标 *RECT) int {
	r, _, _ := xShape_GetRect.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(接收返回坐标)))
	return int(r)
}

// 形状_置坐标.
//
// hShape: 形状对象句柄.
//
// pRect: 坐标.
func X形状_置坐标(形状对象句柄 int, 坐标 *RECT) int {
	r, _, _ := xShape_SetRect.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 形状_置逻辑坐标, 设置元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// hShape:.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
func X形状_置逻辑坐标(形状对象句柄 int, 坐标 *RECT, 是否重绘 bool) bool {
	r, _, _ := xShape_SetRectLogic.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(坐标)), 炫彩工具类.BoolPtr(是否重绘))
	return r != 0
}

// 形状_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// hShape: 形状对象句柄.
//
// pRect: 坐标.
func X形状_取逻辑坐标(形状对象句柄 int, 坐标 *RECT) int {
	r, _, _ := xShape_GetRectLogic.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 形状_取基于窗口客户区坐标, 基于窗口客户区坐标.
//
// hShape: 形状对象句柄.
//
// pRect: 坐标.
func X形状_取基于窗口客户区坐标(形状对象句柄 int, 坐标 *RECT) int {
	r, _, _ := xShape_GetWndClientRect.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 形状_取内容大小 ,仅计算有效内容, 填充父, 权重依赖父级所以无法计算.
//
// hShape: 形状对象句柄.
//
// pSize: 接收返回内容大小值.
func X形状_取内容大小(形状对象句柄 int, 接收返回内容大小值 *SIZE) int {
	r, _, _ := xShape_GetContentSize.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(接收返回内容大小值)))
	return int(r)
}

// 形状_显示布局边界, 是否显示布局边界.
//
// hShape: 形状对象句柄.
//
// bShow: 是否显示.
func X形状_显示布局边界(形状对象句柄 int, 是否显示 bool) int {
	r, _, _ := xShape_ShowLayout.Call(uintptr(形状对象句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 形状_调整布局.
//
// hShape: 形状对象句柄.
func X形状_调整布局(形状对象句柄 int) int {
	r, _, _ := xShape_AdjustLayout.Call(uintptr(形状对象句柄))
	return int(r)
}

// 形状_销毁, 销毁形状对象.
//
// hShape: 形状对象句柄.
func X形状_销毁(形状对象句柄 int) int {
	r, _, _ := xShape_Destroy.Call(uintptr(形状对象句柄))
	return int(r)
}

// 形状_取位置.
//
// hShape: 形状对象句柄.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func X形状_取位置(形状对象句柄 int, 返回X坐标, 返回Y坐标 *int32) int {
	r, _, _ := xShape_GetPosition.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(返回X坐标)), uintptr(unsafe.Pointer(返回Y坐标)))
	return int(r)
}

// 形状_置大小.
//
// hShape: 形状对象句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func X形状_置大小(形状对象句柄 int, 宽度, 高度 int32) int {
	r, _, _ := xShape_SetSize.Call(uintptr(形状对象句柄), uintptr(宽度), uintptr(高度))
	return int(r)
}

// 形状_取大小.
//
// hShape: 形状对象句柄.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func X形状_取大小(形状对象句柄 int, 返回宽度, 返回高度 *int32) int {
	r, _, _ := xShape_GetSize.Call(uintptr(形状对象句柄), uintptr(unsafe.Pointer(返回宽度)), uintptr(unsafe.Pointer(返回高度)))
	return int(r)
}

// 形状_置透明度.
//
// hShape: 形状对象句柄.
//
// alpha: 透明度.
func X形状_置透明度(形状对象句柄 int, 透明度 uint8) int {
	r, _, _ := xShape_SetAlpha.Call(uintptr(形状对象句柄), uintptr(透明度))
	return int(r)
}

// 形状_取透明度, 返回透明度.
//
// hShape: 形状对象句柄.
func X形状_取透明度(形状对象句柄 int) int {
	r, _, _ := xShape_GetAlpha.Call(uintptr(形状对象句柄))
	return int(r)
}
