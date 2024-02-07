package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 进度条_创建, 创建进度条元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄.如果是窗口资源句柄将被添加到窗口.
func X进度条_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xProgBar_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 进度条_置范围, 设置范围.
//
// hEle: 元素句柄.
//
// range_: 范围.
func X进度条_置范围(元素句柄 int, 范围 int) int {
	r, _, _ := xProgBar_SetRange.Call(uintptr(元素句柄), uintptr(范围))
	return int(r)
}

// 进度条_取范围.
//
// hEle: 元素句柄.
func X进度条_取范围(元素句柄 int) int {
	r, _, _ := xProgBar_GetRange.Call(uintptr(元素句柄))
	return int(r)
}

// 进度条_置进度图片.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.
func X进度条_置进度图片(元素句柄 int, 图片句柄 int) int {
	r, _, _ := xProgBar_SetImageLoad.Call(uintptr(元素句柄), uintptr(图片句柄))
	return int(r)
}

// 进度条_置进度, 设置位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.
func X进度条_置进度(元素句柄 int, 位置点 int) int {
	r, _, _ := xProgBar_SetPos.Call(uintptr(元素句柄), uintptr(位置点))
	return int(r)
}

// 进度条_取进度, 获取当前位置点.
//
// hEle: 元素句柄.
func X进度条_取进度(元素句柄 int) int {
	r, _, _ := xProgBar_GetPos.Call(uintptr(元素句柄))
	return int(r)
}

// 进度条_置水平, 设置水平或垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.
func X进度条_置水平(元素句柄 int, 水平或垂直 bool) int {
	r, _, _ := xProgBar_EnableHorizon.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(水平或垂直))
	return int(r)
}

// 进度条_启用缩放, 缩放进度贴图为当前进度区域(当前进度所显示区域), 否则为整体100进度区域.
//
// hEle: 元素句柄.
//
// bStretch: 缩放.
func X进度条_启用缩放(元素句柄 int, 缩放 bool) bool {
	r, _, _ := xProgBar_EnableStretch.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(缩放))
	return r != 0
}

// 进度条_启用进度文本 显示进度值文本.
//
// hEle: 元素句柄.
//
// bShow: 是否启用.
func X进度条_启用进度文本(元素句柄 int, 是否启用 bool) bool {
	r, _, _ := xProgBar_EnableShowText.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return r != 0
}

// 进度条_置进度颜色. 设置进度颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.
func X进度条_置进度颜色(元素句柄 int, ABGR颜色 int) bool {
	r, _, _ := xProgBar_SetColorLoad.Call(uintptr(元素句柄), uintptr(ABGR颜色))
	return r != 0
}
