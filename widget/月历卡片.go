package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 月历卡片.
type MonthCal struct {
	Element
}

// I月历_创建, 创建日期时间元素.
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
func I月历_创建(x int, y int, cx int, cy int, hParent int) *MonthCal {
	p := &MonthCal{}
	p.SetHandle(xc.XMonthCal_Create(x, y, cx, cy, hParent))
	return p
}

// I月历_从句柄创建对象.
func I月历_从句柄创建(handle int) *MonthCal {
	p := &MonthCal{}
	p.SetHandle(handle)
	return p
}

// I月历_从name创建对象, 失败返回nil.
func I月历_从name创建(name string) *MonthCal {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &MonthCal{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I月历_从UID创建对象, 失败返回nil.
func I月历_从UID创建(nUID int) *MonthCal {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &MonthCal{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I月历_从UID名称创建对象, 失败返回nil.
func I月历_从UID名称创建(name string) *MonthCal {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &MonthCal{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 月历_取内部按钮, 获取内部按钮元素.
//
// nType: 按钮类型.
func (m *MonthCal) I取内部按钮(nType int) int {
	return xc.XMonthCal_GetButton(m.I句柄, nType)
}

// 月历_置当前日期, 设置月历选中的年月日.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.
func (m *MonthCal) I置当前日期(nYear int, nMonth int, nDay int) int {
	return xc.XMonthCal_SetToday(m.I句柄, nYear, nMonth, nDay)
}

// 月历_取当前日期, 获取月历当前年月日.
//
// pnYear: 年.[INT.
//
// pnMonth: 月.[INT.
//
// pnDay: 日.[INT.
func (m *MonthCal) I取当前日期(pnYear *int, pnMonth *int, pnDay *int) int {
	return xc.XMonthCal_GetToday(m.I句柄, pnYear, pnMonth, pnDay)
}

// 月历_取选择日期, 获取月历选中的年月日.
//
// pnYear: 年.[INT.
//
// pnMonth: 月.[INT.
//
// pnDay: 日.[INT.
func (m *MonthCal) I取选择日期(pnYear *int, pnMonth *int, pnDay *int) int {
	return xc.XMonthCal_GetSelDate(m.I句柄, pnYear, pnMonth, pnDay)
}

// 月历_置文本颜色.
//
// nFlag: 1:周六, 周日文字颜色, 2:日期文字的颜色; 其它周文字颜色, 使用元素自身颜色.
//
// color: ABGR 颜色值.
func (m *MonthCal) I置文本颜色(nFlag int, color int) int {
	return xc.XMonthCal_SetTextColor(m.I句柄, nFlag, color)
}

/*
以下都是事件
*/

type XE_MONTHCAL_CHANGE func(pbHandled *bool) int            // 月历元素日期改变事件.
type XE_MONTHCAL_CHANGE1 func(hEle int, pbHandled *bool) int // 月历元素日期改变事件.

// 月历元素  I事件_日期改变 事件.
func (m *MonthCal) I事件_日期改变(pFun XE_MONTHCAL_CHANGE) bool {
	return xc.XEle_RegEventC(m.I句柄, xcc.XE_MONTHCAL_CHANGE, pFun)
}

// 月历元素 I事件_日期改变 事件.
func (m *MonthCal) I事件_日期改变1(pFun XE_MONTHCAL_CHANGE1) bool {
	return xc.XEle_RegEventC1(m.I句柄, xcc.XE_MONTHCAL_CHANGE, pFun)
}
