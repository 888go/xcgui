package 炫彩基类

import "github.com/888go/xcgui/common"

// 滚动条_创建, 创建滚动条元素, 返回元素句柄.
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

// ff:滚动条_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X滚动条_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xSBar_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 滚动条_置范围, 设置滚动范围.
//
// hEle: 元素句柄.
//
// range_: 范围.

// ff:滚动条_置范围
// range_:范围
// hEle:元素句柄
func X滚动条_置范围(元素句柄 int, 范围 int) int {
	r, _, _ := xSBar_SetRange.Call(uintptr(元素句柄), uintptr(范围))
	return int(r)
}

// 滚动条_取范围, 获取滚动范围.
//
// hEle: 元素句柄.

// ff:滚动条_取范围
// hEle:元素句柄
func X滚动条_取范围(元素句柄 int) int {
	r, _, _ := xSBar_GetRange.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动条_显示上下按钮, 显示隐藏滚动条上下按钮.
//
// hEle: 元素句柄.
//
// bShow: 是否显示.

// ff:滚动条_显示上下按钮
// bShow:是否显示
// hEle:元素句柄
func X滚动条_显示上下按钮(元素句柄 int, 是否显示 bool) int {
	r, _, _ := xSBar_ShowButton.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 滚动条_置滑块长度.
//
// hEle: 元素句柄.
//
// length: 长度.

// ff:滚动条_置滑块长度
// length:长度
// hEle:元素句柄
func X滚动条_置滑块长度(元素句柄 int, 长度 int) int {
	r, _, _ := xSBar_SetSliderLength.Call(uintptr(元素句柄), uintptr(长度))
	return int(r)
}

// 滚动条_置滑块最小长度.
//
// hEle: 元素句柄.
//
// minLength: 长度.

// ff:滚动条_置滑块最小长度
// minLength:长度
// hEle:元素句柄
func X滚动条_置滑块最小长度(元素句柄 int, 长度 int) int {
	r, _, _ := xSBar_SetSliderMinLength.Call(uintptr(元素句柄), uintptr(长度))
	return int(r)
}

// 滚动条_置滑块两边间隔, 设置滑块两边的间隔大小.
//
// hEle: 元素句柄.
//
// nPadding: 间隔大小.

// ff:滚动条_置滑块两边间隔
// nPadding:间隔大小
// hEle:元素句柄
func X滚动条_置滑块两边间隔(元素句柄 int, 间隔大小 int) int {
	r, _, _ := xSBar_SetSliderPadding.Call(uintptr(元素句柄), uintptr(间隔大小))
	return int(r)
}

// 滚动条_置水平, 设置水平或者垂直.
//
// hEle: 元素句柄.
//
// bHorizon: 水平或垂直.

// ff:滚动条_置水平
// bHorizon:水平或垂直
// hEle:元素句柄
func X滚动条_置水平(元素句柄 int, 水平或垂直 bool) bool {
	r, _, _ := xSBar_EnableHorizon.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(水平或垂直))
	return r != 0
}

// 滚动条_取滑块最大长度.
//
// hEle: 元素句柄.

// ff:滚动条_取滑块最大长度
// hEle:元素句柄
func X滚动条_取滑块最大长度(元素句柄 int) int {
	r, _, _ := xSBar_GetSliderMaxLength.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动条_向上滚动.
//
// hEle: 元素句柄.

// ff:滚动条_向上滚动
// hEle:元素句柄
func X滚动条_向上滚动(元素句柄 int) bool {
	r, _, _ := xSBar_ScrollUp.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动条_向下滚动.
//
// hEle: 元素句柄.

// ff:滚动条_向下滚动
// hEle:元素句柄
func X滚动条_向下滚动(元素句柄 int) bool {
	r, _, _ := xSBar_ScrollDown.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动条_滚动到顶部.
//
// hEle: 元素句柄.

// ff:滚动条_滚动到顶部
// hEle:元素句柄
func X滚动条_滚动到顶部(元素句柄 int) bool {
	r, _, _ := xSBar_ScrollTop.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动条_滚动到底部.
//
// hEle: 元素句柄.

// ff:滚动条_滚动到底部
// hEle:元素句柄
func X滚动条_滚动到底部(元素句柄 int) bool {
	r, _, _ := xSBar_ScrollBottom.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动条_滚动到指定位置, 滚动到指定位置点, 触发事件: XE_SBAR_SCROLL.
//
// hEle: 元素句柄.
//
// pos: 位置点.

// ff:滚动条_滚动到指定位置
// pos:位置点
// hEle:元素句柄
func X滚动条_滚动到指定位置(元素句柄 int, 位置点 int) bool {
	r, _, _ := xSBar_ScrollPos.Call(uintptr(元素句柄), uintptr(位置点))
	return r != 0
}

// 滚动条_取上按钮, 获取上按钮, 返回按钮句柄.
//
// hEle: 元素句柄.

// ff:滚动条_取上按钮
// hEle:元素句柄
func X滚动条_取上按钮(元素句柄 int) int {
	r, _, _ := xSBar_GetButtonUp.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动条_取下按钮, 获取下按钮, 返回按钮句柄.
//
// hEle: 元素句柄.

// ff:滚动条_取下按钮
// hEle:元素句柄
func X滚动条_取下按钮(元素句柄 int) int {
	r, _, _ := xSBar_GetButtonDown.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动条_取滑块, 获取滑动按钮, 返回按钮句柄.
//
// hEle: 元素句柄.

// ff:滚动条_取滑块
// hEle:元素句柄
func X滚动条_取滑块(元素句柄 int) int {
	r, _, _ := xSBar_GetButtonSlider.Call(uintptr(元素句柄))
	return int(r)
}
