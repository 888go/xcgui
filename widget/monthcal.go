package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 月历卡片.
type MonthCal struct {
	Element
}

// 月历_创建, 创建日期时间元素.
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

// ff:创建月历
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:y坐标
// x:x坐标
func X创建月历(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *MonthCal {
	p := &MonthCal{}
	p.X设置句柄(炫彩基类.X月历_创建(x坐标, y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.

// ff:创建月历并按句柄
// handle:
func X创建月历并按句柄(handle int) *MonthCal {
	p := &MonthCal{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建月历并按名称
// name:
func X创建月历并按名称(name string) *MonthCal {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &MonthCal{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建月历并按UID
// nUID:
func X创建月历并按UID(nUID int) *MonthCal {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &MonthCal{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建月历并按UID名称
// name:
func X创建月历并按UID名称(name string) *MonthCal {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &MonthCal{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 月历_取内部按钮, 获取内部按钮元素.
//
// nType: 按钮类型.

// ff:取内部按钮
// nType:按钮类型
func (m *MonthCal) X取内部按钮(按钮类型 int) int {
	return 炫彩基类.X月历_取内部按钮(m.Handle, 按钮类型)
}

// 月历_置当前日期, 设置月历选中的年月日.
//
// nYear: 年.
//
// nMonth: 月.
//
// nDay: 日.

// ff:置当前日期
// nDay:日
// nMonth:月
// nYear:年
func (m *MonthCal) X置当前日期(年 int32, 月 int32, 日 int32) int {
	return 炫彩基类.X月历_置当前日期(m.Handle, 年, 月, 日)
}

// 月历_取当前日期, 获取月历当前年月日.
//
// pnYear: 年.[INT.
//
// pnMonth: 月.[INT.
//
// pnDay: 日.[INT.

// ff:取当前日期
// pnDay:日指针
// pnMonth:月指针
// pnYear:年指针
func (m *MonthCal) X取当前日期(年指针 *int32, 月指针 *int32, 日指针 *int32) int {
	return 炫彩基类.X月历_取当前日期(m.Handle, 年指针, 月指针, 日指针)
}

// 月历_取选择日期, 获取月历选中的年月日.
//
// pnYear: 年.[INT.
//
// pnMonth: 月.[INT.
//
// pnDay: 日.[INT.

// ff:取选择日期
// pnDay:日指针
// pnMonth:月指针
// pnYear:年指针
func (m *MonthCal) X取选择日期(年指针 *int32, 月指针 *int32, 日指针 *int32) int {
	return 炫彩基类.X月历_取选择日期(m.Handle, 年指针, 月指针, 日指针)
}

// 月历_置文本颜色.
//
// nFlag: 1:周六, 周日文字颜色, 2:日期文字的颜色; 其它周文字颜色, 使用元素自身颜色.
//
// color: ABGR 颜色值.

// ff:置文本颜色
// color:ABGR颜色值
// nFlag:类型
func (m *MonthCal) X置文本颜色(类型 int32, ABGR颜色值 int) int {
	return 炫彩基类.X月历_置文本颜色(m.Handle, 类型, ABGR颜色值)
}

/*
以下都是事件
*/

type XE_MONTHCAL_CHANGE func(pbHandled *bool) int            // 月历元素日期改变事件.
type XE_MONTHCAL_CHANGE1 func(hEle int, pbHandled *bool) int // 月历元素日期改变事件.

// 月历元素日期改变事件.

// ff:事件_日期改变
// pFun:
func (m *MonthCal) X事件_日期改变(pFun XE_MONTHCAL_CHANGE) bool {
	return 炫彩基类.X元素_注册事件C(m.Handle, 炫彩常量类.XE_MONTHCAL_CHANGE, pFun)
}

// 月历元素日期改变事件.

// ff:事件_日期改变1
// pFun:
func (m *MonthCal) X事件_日期改变1(pFun XE_MONTHCAL_CHANGE1) bool {
	return 炫彩基类.X元素_注册事件C1(m.Handle, 炫彩常量类.XE_MONTHCAL_CHANGE, pFun)
}
