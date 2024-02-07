package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 日期时间.
type DateTime struct {
	Element
}

// 日期_创建, 创建日期时间元素.
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
func X创建日期(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *DateTime {
	p := &DateTime{}
	p.X设置句柄(炫彩基类.X日期_创建(x坐标, y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建日期并按句柄(句柄 int) *DateTime {
	p := &DateTime{}
	p.X设置句柄(句柄)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建日期并按名称(名称 string) *DateTime {
	handle := 炫彩基类.X取对象从名称(名称)
	if handle > 0 {
		p := &DateTime{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建日期并按UID(nUID int) *DateTime {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &DateTime{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建日期并按UID名称(UID名称 string) *DateTime {
	handle := 炫彩基类.X取对象从UID名称(UID名称)
	if handle > 0 {
		p := &DateTime{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 日期_置样式, 设置样式.
//
// nStyle: 样式: 0为日期元素, 1为时间元素.
func (d *DateTime) X置样式(样式 int) int {
	return 炫彩基类.X日期_置样式(d.Handle, 样式)
}

// 日期_取样式, 返回元素样式.
func (d *DateTime) X取样式() int {
	return 炫彩基类.X日期_取样式(d.Handle)
}

// 日期_启用分割栏为斜线, 切换分割栏为: 斜线或横线.
//
// bSlash: TRUE: 斜线, FALSE: 横线.
func (d *DateTime) X启用分割栏为斜线(斜线 bool) int {
	return 炫彩基类.X日期_启用分割栏为斜线(d.Handle, 斜线)
}

// 日期_取内部按钮, 获取内部按钮元素.
//
// nType: 按钮类型.
func (d *DateTime) X取内部按钮(按钮类型 int) int {
	return 炫彩基类.X日期_取内部按钮(d.Handle, 按钮类型)
}

// 日期_取选择日期背景颜色, 获取被选择文字的背景颜色.
func (d *DateTime) X取选择日期背景颜色() int {
	return 炫彩基类.X日期_取选择日期背景颜色(d.Handle)
}

// 日期_置选择日期背景颜色, 设置被选择文字的背景颜色.
//
// crSelectBk: 文字被选中背景色, ABGR 颜色.
func (d *DateTime) X置选择日期背景颜色(背景色 int) int {
	return 炫彩基类.X日期_置选择日期背景颜色(d.Handle, 背景色)
}

// 日期_取当前日期.
//
// pnYear: 年.[OUT].
//
// pnMonth: 月.[OUT].
//
// pnDay: 日.[OUT].
func (d *DateTime) X取当前日期(年指针 *int32, 月指针 *int32, 日指针 *int32) int {
	return 炫彩基类.X日期_取当前日期(d.Handle, 年指针, 月指针, 日指针)
}

// 日期_置当前日期.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func (d *DateTime) X置当前日期(年 int32, 月 int32, 日 int32) int {
	return 炫彩基类.X日期_置当前日期(d.Handle, 年, 月, 日)
}

// 日期_取当前时间.
//
// pnHour: 时.[OUT].
//
// pnMinute: 分.[OUT].
//
// pnSecond: 秒.[OUT].
func (d *DateTime) X取当前时间(时指针 *int32, 分指针 *int32, 秒指针 *int32) int {
	return 炫彩基类.X日期_取当前时间(d.Handle, 时指针, 分指针, 秒指针)
}

// 日期_社区当前时间, 设置当前时分秒.
//
// nHour: 时.
//
// nMinute: 分.
//
// nSecond: 秒.
func (d *DateTime) X设置当前时间(时 int32, 分 int32, 秒 int32) int {
	return 炫彩基类.X日期_社区当前时间(d.Handle, 时, 分, 秒)
}

// 日期_弹出, 弹出月历卡片.
func (d *DateTime) X弹出月历卡片() int {
	return 炫彩基类.X日期_弹出(d.Handle)
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

// 日期时间元素,内容改变事件.
func (d *DateTime) X事件_内容改变(pFun XE_DATETIME_CHANGE) bool {
	return 炫彩基类.X元素_注册事件C(d.Handle, 炫彩常量类.XE_DATETIME_CHANGE, pFun)
}

// 日期时间元素,内容改变事件.
func (d *DateTime) X事件_内容改变1(pFun XE_DATETIME_CHANGE1) bool {
	return 炫彩基类.X元素_注册事件C1(d.Handle, 炫彩常量类.XE_DATETIME_CHANGE, pFun)
}

// 日期时间元素,弹出月历卡片事件.
func (d *DateTime) X事件_弹出月历卡片(pFun XE_DATETIME_POPUP_MONTHCAL) bool {
	return 炫彩基类.X元素_注册事件C(d.Handle, 炫彩常量类.XE_DATETIME_POPUP_MONTHCAL, pFun)
}

// 日期时间元素,弹出月历卡片事件.
func (d *DateTime) X事件_弹出月历卡片1(pFun XE_DATETIME_POPUP_MONTHCAL1) bool {
	return 炫彩基类.X元素_注册事件C1(d.Handle, 炫彩常量类.XE_DATETIME_POPUP_MONTHCAL, pFun)
}

// 日期时间元素,弹出的月历卡片退出事件.
func (d *DateTime) X事件_月历卡片退出(pFun XE_DATETIME_EXIT_MONTHCAL) bool {
	return 炫彩基类.X元素_注册事件C(d.Handle, 炫彩常量类.XE_DATETIME_EXIT_MONTHCAL, pFun)
}

// 日期时间元素,弹出的月历卡片退出事件.
func (d *DateTime) X事件_月历卡片退出1(pFun XE_DATETIME_EXIT_MONTHCAL1) bool {
	return 炫彩基类.X元素_注册事件C1(d.Handle, 炫彩常量类.XE_DATETIME_EXIT_MONTHCAL, pFun)
}
