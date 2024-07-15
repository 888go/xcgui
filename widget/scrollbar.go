package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ScrollBar 滚动条.
type ScrollBar struct {
	Element
}

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

// ff:创建滚动条
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X创建滚动条(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *ScrollBar {
	p := &ScrollBar{}
	p.X设置句柄(炫彩基类.X滚动条_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.

// ff:创建滚动条并按句柄
// handle:
func X创建滚动条并按句柄(handle int) *ScrollBar {
	p := &ScrollBar{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建滚动条并按名称
// name:
func X创建滚动条并按名称(name string) *ScrollBar {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ScrollBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建滚动条并按UID
// nUID:
func X创建滚动条并按UID(nUID int) *ScrollBar {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ScrollBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建滚动条并按UID名称
// name:
func X创建滚动条并按UID名称(name string) *ScrollBar {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ScrollBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 滚动条_置范围, 设置滚动范围.
//
// range_: 范围.

// ff:置范围
// range_:范围
func (s *ScrollBar) X置范围(范围 int) int {
	return 炫彩基类.X滚动条_置范围(s.Handle, 范围)
}

// 滚动条_取范围, 获取滚动范围.

// ff:取范围
func (s *ScrollBar) X取范围() int {
	return 炫彩基类.X滚动条_取范围(s.Handle)
}

// 滚动条_显示上下按钮, 显示隐藏滚动条上下按钮.
//
// bShow: 是否显示.

// ff:显示上下按钮
// bShow:是否显示
func (s *ScrollBar) X显示上下按钮(是否显示 bool) int {
	return 炫彩基类.X滚动条_显示上下按钮(s.Handle, 是否显示)
}

// 滚动条_置滑块长度.
//
// length: 长度.

// ff:置滑块长度
// length:长度
func (s *ScrollBar) X置滑块长度(长度 int) int {
	return 炫彩基类.X滚动条_置滑块长度(s.Handle, 长度)
}

// 滚动条_置滑块最小长度.
//
// minLength: 长度.

// ff:置滑块最小长度
// minLength:长度
func (s *ScrollBar) X置滑块最小长度(长度 int) int {
	return 炫彩基类.X滚动条_置滑块最小长度(s.Handle, 长度)
}

// 滚动条_置滑块两边间隔, 设置滑块两边的间隔大小.
//
// nPadding: 间隔大小.

// ff:置滑块两边间隔
// nPadding:间隔大小
func (s *ScrollBar) X置滑块两边间隔(间隔大小 int) int {
	return 炫彩基类.X滚动条_置滑块两边间隔(s.Handle, 间隔大小)
}

// 滚动条_置水平, 设置水平或者垂直.
//
// bHorizon: 水平或垂直.

// ff:置水平
// bHorizon:水平或垂直
func (s *ScrollBar) X置水平(水平或垂直 bool) bool {
	return 炫彩基类.X滚动条_置水平(s.Handle, 水平或垂直)
}

// 滚动条_取滑块最大长度.

// ff:取滑块最大长度
func (s *ScrollBar) X取滑块最大长度() int {
	return 炫彩基类.X滚动条_取滑块最大长度(s.Handle)
}

// 滚动条_向上滚动.

// ff:向上滚动
func (s *ScrollBar) X向上滚动() bool {
	return 炫彩基类.X滚动条_向上滚动(s.Handle)
}

// 滚动条_向下滚动.

// ff:向下滚动
func (s *ScrollBar) X向下滚动() bool {
	return 炫彩基类.X滚动条_向下滚动(s.Handle)
}

// 滚动条_滚动到顶部.

// ff:滚动到顶部
func (s *ScrollBar) X滚动到顶部() bool {
	return 炫彩基类.X滚动条_滚动到顶部(s.Handle)
}

// 滚动条_滚动到底部.

// ff:滚动到底部
func (s *ScrollBar) X滚动到底部() bool {
	return 炫彩基类.X滚动条_滚动到底部(s.Handle)
}

// 滚动条_滚动到指定位置, 滚动到指定位置点, 触发事件: XE_SBAR_SCROLL.
//
// pos: 位置点.

// ff:滚动到指定位置
// pos:位置点
func (s *ScrollBar) X滚动到指定位置(位置点 int) bool {
	return 炫彩基类.X滚动条_滚动到指定位置(s.Handle, 位置点)
}

// 滚动条_取上按钮, 获取上按钮, 返回按钮句柄.

// ff:取上按钮
func (s *ScrollBar) X取上按钮() int {
	return 炫彩基类.X滚动条_取上按钮(s.Handle)
}

// 滚动条_取下按钮, 获取下按钮, 返回按钮句柄.

// ff:取下按钮
func (s *ScrollBar) X取下按钮() int {
	return 炫彩基类.X滚动条_取下按钮(s.Handle)
}

// 滚动条_取滑块, 获取滑动按钮, 返回按钮句柄.

// ff:取滑块
func (s *ScrollBar) X取滑块() int {
	return 炫彩基类.X滚动条_取滑块(s.Handle)
}

/*
以下都是事件
*/

type XE_SBAR_SCROLL func(pos int32, pbHandled *bool) int            // 滚动条元素滚动事件,滚动条触发.
type XE_SBAR_SCROLL1 func(hEle int, pos int32, pbHandled *bool) int // 滚动条元素滚动事件,滚动条触发.

// 滚动条元素滚动事件,滚动条触发.

// ff:事件_滚动事件
// pFun:
func (s *ScrollBar) X事件_滚动事件(pFun XE_SBAR_SCROLL) bool {
	return 炫彩基类.X元素_注册事件C(s.Handle, 炫彩常量类.XE_SBAR_SCROLL, pFun)
}

// 滚动条元素滚动事件,滚动条触发.

// ff:事件_滚动事件1
// pFun:
func (s *ScrollBar) X事件_滚动事件1(pFun XE_SBAR_SCROLL1) bool {
	return 炫彩基类.X元素_注册事件C1(s.Handle, 炫彩常量类.XE_SBAR_SCROLL, pFun)
}
