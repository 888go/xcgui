package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 日期时间.
type DateTime struct {
	Element
}

// I日期_创建, 创建日期时间元素.
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
func I日期_创建(x int, y int, cx int, cy int, hParent int) *DateTime {
	p := &DateTime{}
	p.SetHandle(xc.XDateTime_Create(x, y, cx, cy, hParent))
	return p
}

// I日期_从句柄创建对象.
func I日期_从句柄创建(handle int) *DateTime {
	p := &DateTime{}
	p.SetHandle(handle)
	return p
}

// I日期_从name创建对象, 失败返回nil.
func I日期_从name创建(name string) *DateTime {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &DateTime{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I日期_从UID创建对象, 失败返回nil.
func I日期_从UID创建(nUID int) *DateTime {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &DateTime{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I日期_从UID名称创建对象, 失败返回nil.
func I日期_从UID名称创建(name string) *DateTime {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &DateTime{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 日期_置样式, 设置样式.
//
// nStyle: 样式: 0为日期元素, 1为时间元素.
func (d *DateTime) I置样式(nStyle int) int {
	return xc.XDateTime_SetStyle(d.I句柄, nStyle)
}

// 日期_取样式, 返回元素样式.
func (d *DateTime) I取样式() int {
	return xc.XDateTime_GetStyle(d.I句柄)
}

// 日期_启用分割栏为斜线, 切换分割栏为: 斜线或横线.
//
// bSlash: TRUE: 斜线, FALSE: 横线.
func (d *DateTime) I启用分割栏为斜线(bSlash bool) int {
	return xc.XDateTime_EnableSplitSlash(d.I句柄, bSlash)
}

// 日期_取内部按钮, 获取内部按钮元素.
//
// nType: 按钮类型.
func (d *DateTime) I取内部按钮(nType int) int {
	return xc.XDateTime_GetButton(d.I句柄, nType)
}

// 日期_取选择日期背景颜色, 获取被选择文字的背景颜色.
func (d *DateTime) I取选择日期背景颜色() int {
	return xc.XDateTime_GetSelBkColor(d.I句柄)
}

// 日期_置选择日期背景颜色, 设置被选择文字的背景颜色.
//
// crSelectBk: 文字被选中背景色, ABGR 颜色.
func (d *DateTime) I置选择日期背景颜色(crSelectBk int) int {
	return xc.XDateTime_SetSelBkColor(d.I句柄, crSelectBk)
}

// 日期_取当前日期.
//
// pnYear: 年.[OUT].
//
// pnMonth: 月.[OUT].
//
// pnDay: 日.[OUT].
func (d *DateTime) I取当前日期(pnYear *int, pnMonth *int, pnDay *int) int {
	return xc.XDateTime_GetDate(d.I句柄, pnYear, pnMonth, pnDay)
}

// 日期_置当前日期.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func (d *DateTime) I置当前日期(nYear int, nMonth int, nDay int) int {
	return xc.XDateTime_SetDate(d.I句柄, nYear, nMonth, nDay)
}

// 日期_取当前时间.
//
// pnHour: 时.[OUT].
//
// pnMinute: 分.[OUT].
//
// pnSecond: 秒.[OUT].
func (d *DateTime) I取当前时间(pnHour *int, pnMinute *int, pnSecond *int) int {
	return xc.XDateTime_GetTime(d.I句柄, pnHour, pnMinute, pnSecond)
}

// 日期_置当前时间, 设置当前时分秒.
//
// nHour: 时.
//
// nMinute: 分.
//
// nSecond: 秒.
func (d *DateTime) I置当前时间(nHour int, nMinute int, nSecond int) int {
	return xc.XDateTime_SetTime(d.I句柄, nHour, nMinute, nSecond)
}

// 日期_弹出, 弹出月历卡片.
func (d *DateTime) I弹出() int {
	return xc.XDateTime_Popup(d.I句柄)
}

/*
以下都是事件
*/

type XE_DATETIME_CHANGE func(pbHandled *bool) int                                                     // 日期时间元素,内容改变事件.
type XE_DATETIME_CHANGE1 func(hEle int, pbHandled *bool) int                                          // 日期时间元素,内容改变事件.
type XE_DATETIME_POPUP_MONTHCAL func(hMonthCalWnd int, hMonthCal int, pbHandled *bool) int            // 日期时间元素,弹出月历卡片事件.
type XE_DATETIME_POPUP_MONTHCAL1 func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int // 日期时间元素,弹出月历卡片事件.
type XE_DATETIME_EXIT_MONTHCAL func(hMonthCalWnd int, hMonthCal int, pbHandled *bool) int             // 日期时间元素,弹出的月历卡片退出事件.
type XE_DATETIME_EXIT_MONTHCAL1 func(hEle int, hMonthCalWnd int, hMonthCal int, pbHandled *bool) int  // 日期时间元素,弹出的月历卡片退出事件.

// 日期时间元素,I事件_内容改变 事件.
func (d *DateTime) I事件_内容改变(pFun XE_DATETIME_CHANGE) bool {
	return xc.XEle_RegEventC(d.I句柄, xcc.XE_DATETIME_CHANGE, pFun)
}

// 日期时间元素,I事件_内容改变 事件.
func (d *DateTime) I事件_内容改变1(pFun XE_DATETIME_CHANGE1) bool {
	return xc.XEle_RegEventC1(d.I句柄, xcc.XE_DATETIME_CHANGE, pFun)
}

// 日期时间元素,I事件_弹出月历 卡片事件.
func (d *DateTime) I事件_弹出月历(pFun XE_DATETIME_POPUP_MONTHCAL) bool {
	return xc.XEle_RegEventC(d.I句柄, xcc.XE_DATETIME_POPUP_MONTHCAL, pFun)
}

// 日期时间元素,I事件_弹出月历 卡片事件.
func (d *DateTime) I事件_弹出月历1(pFun XE_DATETIME_POPUP_MONTHCAL1) bool {
	return xc.XEle_RegEventC1(d.I句柄, xcc.XE_DATETIME_POPUP_MONTHCAL, pFun)
}

// 日期时间元素,I事件_弹出的月历退出 事件.
func (d *DateTime) I事件_弹出的月历退出(pFun XE_DATETIME_EXIT_MONTHCAL) bool {
	return xc.XEle_RegEventC(d.I句柄, xcc.XE_DATETIME_EXIT_MONTHCAL, pFun)
}

// 日期时间元素,I事件_弹出的月历退出 事件.
func (d *DateTime) I事件_弹出的月历退出1(pFun XE_DATETIME_EXIT_MONTHCAL1) bool {
	return xc.XEle_RegEventC1(d.I句柄, xcc.XE_DATETIME_EXIT_MONTHCAL, pFun)
}
