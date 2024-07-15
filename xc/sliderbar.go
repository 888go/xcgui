package 炫彩基类

import "github.com/888go/xcgui/common"

// 滑动条_创建, 创建滑动条元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.

// ff:滑动条_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X滑动条_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xSliderBar_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 滑动条_置范围, 设置滑动范围.
//
// hEle: 元素句柄.
//
// range_: 范围.

// ff:滑动条_置范围
// range_:范围
// hEle:元素句柄
func X滑动条_置范围(元素句柄 int, 范围 int) int {
	r, _, _ := xSliderBar_SetRange.Call(uintptr(元素句柄), uintptr(范围))
	return int(r)
}

// 滑动条_取范围, 获取滚动范围.
//
// hEle: 元素句柄.

// ff:滑动条_取范围
// hEle:元素句柄
func X滑动条_取范围(元素句柄 int) int {
	r, _, _ := xSliderBar_GetRange.Call(uintptr(元素句柄))
	return int(r)
}

// 滑动条_置进度图片, 设置进度贴图.
//
// hEle: 元素句柄.
//
// hImage: 图片句柄.

// ff:滑动条_置进度图片
// hImage:图片句柄
// hEle:元素句柄
func X滑动条_置进度图片(元素句柄 int, 图片句柄 int) int {
	r, _, _ := xSliderBar_SetImageLoad.Call(uintptr(元素句柄), uintptr(图片句柄))
	return int(r)
}

// 滑动条_置滑块宽度, 设置滑块按钮宽度.
//
// hEle: 元素句柄.
//
// width: 宽度.

// ff:滑动条_置滑块宽度
// width:宽度
// hEle:元素句柄
func X滑动条_置滑块宽度(元素句柄 int, 宽度 int) int {
	r, _, _ := xSliderBar_SetButtonWidth.Call(uintptr(元素句柄), uintptr(宽度))
	return int(r)
}

// 滑动条_置滑块高度, 设置滑块按钮高度.
//
// hEle: 元素句柄.
//
// height: 高度.

// ff:滑动条_置滑块高度
// height:高度
// hEle:元素句柄
func X滑动条_置滑块高度(元素句柄 int, 高度 int) int {
	r, _, _ := xSliderBar_SetButtonHeight.Call(uintptr(元素句柄), uintptr(高度))
	return int(r)
}

// 滑动条_置当前位置, 设置当前进度点.
//
// hEle: 元素句柄.
//
// pos: 进度点.

// ff:滑动条_置当前位置
// pos:进度点
// hEle:元素句柄
func X滑动条_置当前位置(元素句柄 int, 进度点 int) int {
	r, _, _ := xSliderBar_SetPos.Call(uintptr(元素句柄), uintptr(进度点))
	return int(r)
}

// 滑动条_取当前位置, 获取当前进度点.
//
// hEle: 元素句柄.

// ff:滑动条_取当前位置
// hEle:元素句柄
func X滑动条_取当前位置(元素句柄 int) int {
	r, _, _ := xSliderBar_GetPos.Call(uintptr(元素句柄))
	return int(r)
}

// 滑动条_取滑块, 返回滑块按钮句柄.
//
// hEle: 元素句柄.

// ff:滑动条_取滑块
// hEle:元素句柄
func X滑动条_取滑块(元素句柄 int) int {
	r, _, _ := xSliderBar_GetButton.Call(uintptr(元素句柄))
	return int(r)
}

// 滑动条_置水平, 设置水平或垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.

// ff:滑动条_置水平
// bHorizon:水平或垂直
// hEle:元素句柄
func X滑动条_置水平(元素句柄 int, 水平或垂直 bool) int {
	r, _, _ := xSliderBar_EnableHorizon.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(水平或垂直))
	return int(r)
}
