package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// 滚动视_创建, 创建滚动视图元素, 返回元素句柄.
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

// ff:滚动视_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X滚动视_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xSView_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 滚动视_置视图大小, 设置内容大小, 如果内容改变返回TRUE否则返回FALSE.
//
// hEle: 元素句柄.
//
// cx: 宽度.
//
// cy: 高度.

// ff:滚动视_置视图大小
// cy:高度
// cx:宽度
// hEle:元素句柄
func X滚动视_置视图大小(元素句柄 int, 宽度 int, 高度 int) bool {
	r, _, _ := xSView_SetTotalSize.Call(uintptr(元素句柄), uintptr(宽度), uintptr(高度))
	return r != 0
}

// 滚动视_取视图大小, 获取内容总大小.
//
// hEle: 元素句柄.
//
// pSize: 大小.

// ff:滚动视_取视图大小
// pSize:大小
// hEle:元素句柄
func X滚动视_取视图大小(元素句柄 int, 大小 *SIZE) int {
	r, _, _ := xSView_GetTotalSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(大小)))
	return int(r)
}

// 滚动视_置滚动单位大小, 设置滚动单位大小, 如果内容改变返回TRUE否则返回FALSE.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
//
// nHeight: 高度.

// ff:滚动视_置滚动单位大小
// nHeight:高度
// nWidth:宽度
// hEle:元素句柄
func X滚动视_置滚动单位大小(元素句柄 int, 宽度 int, 高度 int) bool {
	r, _, _ := xSView_SetLineSize.Call(uintptr(元素句柄), uintptr(宽度), uintptr(高度))
	return r != 0
}

// 滚动视_取滚动单位大小, 获取滚动单位大小.
//
// hEle: 元素句柄.
//
// pSize: 返回大小.

// ff:滚动视_取滚动单位大小
// pSize:返回大小
// hEle:元素句柄
func X滚动视_取滚动单位大小(元素句柄 int, 返回大小 *SIZE) int {
	r, _, _ := xSView_GetLineSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(返回大小)))
	return int(r)
}

// 滚动视_置滚动条大小.
//
// hEle: 元素句柄.
//
// size: 滚动条大小.

// ff:滚动视_置滚动条大小
// size:滚动条大小
// hEle:元素句柄
func X滚动视_置滚动条大小(元素句柄 int, 滚动条大小 int) int {
	r, _, _ := xSView_SetScrollBarSize.Call(uintptr(元素句柄), uintptr(滚动条大小))
	return int(r)
}

// 滚动视_取视口原点X, 获取视口原点X坐标.
//
// hEle: 元素句柄.

// ff:滚动视_取视口原点X
// hEle:元素句柄
func X滚动视_取视口原点X(元素句柄 int) int {
	r, _, _ := xSView_GetViewPosH.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动视_取视口原点Y, 获取视口原点Y坐标.
//
// hEle: 元素句柄.

// ff:滚动视_取视口原点Y
// hEle:元素句柄
func X滚动视_取视口原点Y(元素句柄 int) int {
	r, _, _ := xSView_GetViewPosV.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动视_取视口宽度.
//
// hEle: 元素句柄.

// ff:滚动视_取视口宽度
// hEle:元素句柄
func X滚动视_取视口宽度(元素句柄 int) int {
	r, _, _ := xSView_GetViewWidth.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动视_取视口高度.
//
// hEle: 元素句柄.

// ff:滚动视_取视口高度
// hEle:元素句柄
func X滚动视_取视口高度(元素句柄 int) int {
	r, _, _ := xSView_GetViewHeight.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动视_取视口坐标.
//
// hEle: 元素句柄.
//
// pRect: 坐标.

// ff:滚动视_取视口坐标
// pRect:坐标
// hEle:元素句柄
func X滚动视_取视口坐标(元素句柄 int, 坐标 *RECT) int {
	r, _, _ := xSView_GetViewRect.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(坐标)))
	return int(r)
}

// 滚动视_取水平滚动条, 返回滚动条句柄.
//
// hEle: 元素句柄.

// ff:滚动视_取水平滚动条
// hEle:元素句柄
func X滚动视_取水平滚动条(元素句柄 int) int {
	r, _, _ := xSView_GetScrollBarH.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动视_取垂直滚动条, 返回滚动条句柄.
//
// hEle: 元素句柄.

// ff:滚动视_取垂直滚动条
// hEle:元素句柄
func X滚动视_取垂直滚动条(元素句柄 int) int {
	r, _, _ := xSView_GetScrollBarV.Call(uintptr(元素句柄))
	return int(r)
}

// 滚动视_水平滚动, 水平滚动条, 滚动到指定位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.

// ff:滚动视_水平滚动
// pos:位置点
// hEle:元素句柄
func X滚动视_水平滚动(元素句柄 int, 位置点 int) bool {
	r, _, _ := xSView_ScrollPosH.Call(uintptr(元素句柄), uintptr(位置点))
	return r != 0
}

// 滚动视_垂直滚动, 垂直滚动条, 滚动到指定位置点.
//
// hEle: 元素句柄.
//
// pos: 位置点.

// ff:滚动视_垂直滚动
// pos:位置点
// hEle:元素句柄
func X滚动视_垂直滚动(元素句柄 int, 位置点 int) bool {
	r, _, _ := xSView_ScrollPosV.Call(uintptr(元素句柄), uintptr(位置点))
	return r != 0
}

// 滚动视_水平滚动到X, 水平滚动条, 滚动到指定坐标.
//
// hEle: 元素句柄.
//
// posX: X坐标.

// ff:滚动视_水平滚动到X
// posX:坐标
// hEle:元素句柄
func X滚动视_水平滚动到X(元素句柄 int, 坐标 int) bool {
	r, _, _ := xSView_ScrollPosXH.Call(uintptr(元素句柄), uintptr(坐标))
	return r != 0
}

// 滚动视_垂直滚动到Y, 垂直滚动条, 滚动到指定坐标.
//
// hEle: 元素句柄.
//
// posY: Y坐标.

// ff:滚动视_垂直滚动到Y
// posY:Y坐标
// hEle:元素句柄
func X滚动视_垂直滚动到Y(元素句柄 int, Y坐标 int) bool {
	r, _, _ := xSView_ScrollPosYV.Call(uintptr(元素句柄), uintptr(Y坐标))
	return r != 0
}

// 滚动视_显示水平滚动条.
//
// hEle: 元素句柄.
//
// bShow: 是否显示.

// ff:滚动视_显示水平滚动条
// bShow:是否显示
// hEle:元素句柄
func X滚动视_显示水平滚动条(元素句柄 int, 是否显示 bool) int {
	r, _, _ := xSView_ShowSBarH.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 滚动视_显示垂直滚动条.
//
// hEle: 元素句柄.
//
// bShow: 是否显示.

// ff:滚动视_显示垂直滚动条
// bShow:是否显示
// hEle:元素句柄
func X滚动视_显示垂直滚动条(元素句柄 int, 是否显示 bool) int {
	r, _, _ := xSView_ShowSBarV.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 滚动视_启用自动显示滚动条.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:滚动视_启用自动显示滚动条
// bEnable:是否启用
// hEle:元素句柄
func X滚动视_启用自动显示滚动条(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xSView_EnableAutoShowScrollBar.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 滚动视_向左滚动.
//
// hEle: 元素句柄.

// ff:滚动视_向左滚动
// hEle:元素句柄
func X滚动视_向左滚动(元素句柄 int) bool {
	r, _, _ := xSView_ScrollLeftLine.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_向右滚动.
//
// hEle: 元素句柄.

// ff:滚动视_向右滚动
// hEle:元素句柄
func X滚动视_向右滚动(元素句柄 int) bool {
	r, _, _ := xSView_ScrollRightLine.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_向上滚动.
//
// hEle: 元素句柄.

// ff:滚动视_向上滚动
// hEle:元素句柄
func X滚动视_向上滚动(元素句柄 int) bool {
	r, _, _ := xSView_ScrollTopLine.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_向下滚动.
//
// hEle: 元素句柄.

// ff:滚动视_向下滚动
// hEle:元素句柄
func X滚动视_向下滚动(元素句柄 int) bool {
	r, _, _ := xSView_ScrollBottomLine.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_滚动到左侧, 水平滚动到左侧.
//
// hEle: 元素句柄.

// ff:滚动视_滚动到左侧
// hEle:元素句柄
func X滚动视_滚动到左侧(元素句柄 int) bool {
	r, _, _ := xSView_ScrollLeft.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_滚动到右侧, 水平滚动到右侧.
//
// hEle: 元素句柄.

// ff:滚动视_滚动到右侧
// hEle:元素句柄
func X滚动视_滚动到右侧(元素句柄 int) bool {
	r, _, _ := xSView_ScrollRight.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_滚动到顶部, 垂直滚动到顶部.
//
// hEle: 元素句柄.

// ff:滚动视_滚动到顶部
// hEle:元素句柄
func X滚动视_滚动到顶部(元素句柄 int) bool {
	r, _, _ := xSView_ScrollTop.Call(uintptr(元素句柄))
	return r != 0
}

// 滚动视_滚动到底部, 垂直滚动到底部.
//
// hEle: 元素句柄.

// ff:滚动视_滚动到底部
// hEle:元素句柄
func X滚动视_滚动到底部(元素句柄 int) bool {
	r, _, _ := xSView_ScrollBottom.Call(uintptr(元素句柄))
	return r != 0
}
