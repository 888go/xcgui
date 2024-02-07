package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// 日期_创建, 创建日期时间元素, 返回元素句柄.
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
func X日期_创建(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xDateTime_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 日期_置样式, 设置样式.
//
// hEle: 元素句柄.
//
// nStyle: 样式: 0为日期元素, 1为时间元素.
func X日期_置样式(元素句柄 int, 样式 int) int {
	r, _, _ := xDateTime_SetStyle.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 日期_取样式, 返回元素样式.
//
// hEle: 元素句柄.
func X日期_取样式(元素句柄 int) int {
	r, _, _ := xDateTime_GetStyle.Call(uintptr(元素句柄))
	return int(r)
}

// 日期_启用分割栏为斜线, 切换分割栏为: 斜线或横线.
//
// hEle: 元素句柄.
//
// bSlash: TRUE: 斜线, FALSE: 横线.
func X日期_启用分割栏为斜线(元素句柄 int, TRUE bool) int {
	r, _, _ := xDateTime_EnableSplitSlash.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(TRUE))
	return int(r)
}

// 日期_取内部按钮, 获取内部按钮元素.
//
// hEle: 元素句柄.
//
// nType: 按钮类型.
func X日期_取内部按钮(元素句柄 int, 按钮类型 int) int {
	r, _, _ := xDateTime_GetButton.Call(uintptr(元素句柄), uintptr(按钮类型))
	return int(r)
}

// 日期_取选择日期背景颜色, 获取被选择文字的背景颜色.
//
// hEle: 元素句柄.
func X日期_取选择日期背景颜色(元素句柄 int) int {
	r, _, _ := xDateTime_GetSelBkColor.Call(uintptr(元素句柄))
	return int(r)
}

// 日期_置选择日期背景颜色, 设置被选择文字的背景颜色.
//
// hEle: 元素句柄.
//
// crSelectBk: 文字被选中背景色, ABGR 颜色.
func X日期_置选择日期背景颜色(元素句柄 int, 文字被选中背景色 int) int {
	r, _, _ := xDateTime_SetSelBkColor.Call(uintptr(元素句柄), uintptr(文字被选中背景色))
	return int(r)
}

// 日期_取当前日期.
//
// hEle: 元素句柄.
//
// pnYear: 年.[OUT].
//
// pnMonth: 月.[OUT].
//
// pnDay: 日.[OUT].
func X日期_取当前日期(元素句柄 int, 年指针 *int32, 月指针 *int32, 日指针 *int32) int {
	r, _, _ := xDateTime_GetDate.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(年指针)), uintptr(unsafe.Pointer(月指针)), uintptr(unsafe.Pointer(日指针)))
	return int(r)
}

// 日期_置当前日期.
//
// hEle: 元素句柄.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func X日期_置当前日期(元素句柄 int, 年 int32, 月 int32, 日 int32) int {
	r, _, _ := xDateTime_SetDate.Call(uintptr(元素句柄), uintptr(年), uintptr(月), uintptr(日))
	return int(r)
}

// 日期_取当前时间.
//
// hEle: 元素句柄.
//
// pnHour: 时.[OUT].
//
// pnMinute: 分.[OUT].
//
// pnSecond: 秒.[OUT].
func X日期_取当前时间(元素句柄 int, 时指针 *int32, 分指针 *int32, 秒指针 *int32) int {
	r, _, _ := xDateTime_GetTime.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(时指针)), uintptr(unsafe.Pointer(分指针)), uintptr(unsafe.Pointer(秒指针)))
	return int(r)
}

// 日期_社区当前时间, 设置当前时分秒.
//
// hEle: 元素句柄.
//
// nHour: 时.
//
// nMinute: 分.
//
// nSecond: 秒.
func X日期_社区当前时间(元素句柄 int, 时 int32, 分 int32, 秒 int32) int {
	r, _, _ := xDateTime_SetTime.Call(uintptr(元素句柄), uintptr(时), uintptr(分), uintptr(秒))
	return int(r)
}

// 日期_弹出, 弹出月历卡片.
//
// hEle: 元素句柄.
func X日期_弹出(元素句柄 int) int {
	r, _, _ := xDateTime_Popup.Call(uintptr(元素句柄))
	return int(r)
}
