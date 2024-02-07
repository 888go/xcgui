package 炫彩基类

import (
	"unsafe"
)

// 月历_创建, 创建日期时间元素, 返回元素句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func X月历_创建(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xMonthCal_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 月历_取内部按钮, 获取内部按钮元素.
//
// hEle: 元素句柄.
//
// nType: 按钮类型.
func X月历_取内部按钮(元素句柄 int, 按钮类型 int) int {
	r, _, _ := xMonthCal_GetButton.Call(uintptr(元素句柄), uintptr(按钮类型))
	return int(r)
}

// 月历_置当前日期, 设置月历选中的年月日.
//
// hEle: 元素句柄.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func X月历_置当前日期(元素句柄 int, 年 int32, 月 int32, 日 int32) int {
	r, _, _ := xMonthCal_SetToday.Call(uintptr(元素句柄), uintptr(年), uintptr(月), uintptr(日))
	return int(r)
}

// 月历_取当前日期, 获取月历当前年月日.
//
// hEle: 元素句柄.
//
// pnYear: 年.[INT.
//
// pnMonth: 月.[INT.
//
// pnDay: 日.[INT.
func X月历_取当前日期(元素句柄 int, 年指针 *int32, 月指针 *int32, 日指针 *int32) int {
	r, _, _ := xMonthCal_GetToday.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(年指针)), uintptr(unsafe.Pointer(月指针)), uintptr(unsafe.Pointer(日指针)))
	return int(r)
}

// 月历_取选择日期, 获取月历选中的年月日.
//
// hEle: 元素句柄.
//
// pnYear: 年.[INT.
//
// pnMonth: 月.[INT.
//
// pnDay: 日.[INT.
func X月历_取选择日期(元素句柄 int, 年指针 *int32, 月指针 *int32, 日指针 *int32) int {
	r, _, _ := xMonthCal_GetSelDate.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(年指针)), uintptr(unsafe.Pointer(月指针)), uintptr(unsafe.Pointer(日指针)))
	return int(r)
}

// 月历_置文本颜色.
//
// hEle: 元素句柄.
//
// nFlag: 1:周六, 周日文字颜色, 2:日期文字的颜色; 其它周文字颜色, 使用元素自身颜色.
//
// color: ABGR 颜色值.
func X月历_置文本颜色(元素句柄 int, 类型 int32, ABGR颜色值 int) int {
	r, _, _ := xMonthCal_SetTextColor.Call(uintptr(元素句柄), uintptr(类型), uintptr(ABGR颜色值))
	return int(r)
}
