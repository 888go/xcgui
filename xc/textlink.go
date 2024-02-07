package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 文本链接_创建, 创建静态文本链接元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func X文本链接_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 文本内容 string, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xTextLink_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(文本内容), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 文本链接_启用离开状态下划线, 启用下划线, 鼠标离开状态.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X文本链接_启用离开状态下划线(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xTextLink_EnableUnderlineLeave.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 文本链接_停留状态下划线, 启用下划线, 鼠标停留状态.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X文本链接_停留状态下划线(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xTextLink_EnableUnderlineStay.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 文本链接_置停留状态文本颜色, 设置文本颜色, 鼠标停留状态.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X文本链接_置停留状态文本颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xTextLink_SetTextColorStay.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 文本链接_置离开状态下划线颜色, 设置下划线颜色, 鼠标离开状态.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X文本链接_置离开状态下划线颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xTextLink_SetUnderlineColorLeave.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 文本链接_置停留状态下划线颜色, 置下划线颜色, 鼠标停留状态.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X文本链接_置停留状态下划线颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xTextLink_SetUnderlineColorStay.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}
