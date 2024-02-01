package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 滚动条.
type ScrollBar struct {
	Element
}

// I滚动条_创建, 创建滚动条元素, 返回元素句柄.
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
func I滚动条_创建(x int, y int, cx int, cy int, hParent int) *ScrollBar {
	p := &ScrollBar{}
	p.SetHandle(xc.XSBar_Create(x, y, cx, cy, hParent))
	return p
}

// I滚动条_从句柄创建对象.
func I滚动条_从句柄创建(handle int) *ScrollBar {
	p := &ScrollBar{}
	p.SetHandle(handle)
	return p
}

// I滚动条_从name创建对象, 失败返回nil.
func I滚动条_从name创建(name string) *ScrollBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ScrollBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I滚动条_从UID创建对象, 失败返回nil.
func I滚动条_从UID创建(nUID int) *ScrollBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ScrollBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I滚动条_从UID名称创建对象, 失败返回nil.
func I滚动条_从UID名称创建(name string) *ScrollBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ScrollBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 滚动条_置范围, 设置滚动范围.
//
// range_: 范围.
func (s *ScrollBar) I置范围(range_ int) int {
	return xc.XSBar_SetRange(s.I句柄, range_)
}

// 滚动条_取范围, 获取滚动范围.
func (s *ScrollBar) I取范围() int {
	return xc.XSBar_GetRange(s.I句柄)
}

// 滚动条_显示上下按钮, 显示隐藏滚动条上下按钮.
//
// bShow: 是否显示.
func (s *ScrollBar) I显示上下按钮(bShow bool) int {
	return xc.XSBar_ShowButton(s.I句柄, bShow)
}

// 滚动条_置滑块长度.
//
// length: 长度.
func (s *ScrollBar) I置滑块长度(length int) int {
	return xc.XSBar_SetSliderLength(s.I句柄, length)
}

// 滚动条_置滑块最小长度.
//
// minLength: 长度.
func (s *ScrollBar) I置滑块最小长度(minLength int) int {
	return xc.XSBar_SetSliderMinLength(s.I句柄, minLength)
}

// 滚动条_置滑块两边间隔, 设置滑块两边的间隔大小.
//
// nPadding: 间隔大小.
func (s *ScrollBar) I置滑块两边间隔(nPadding int) int {
	return xc.XSBar_SetSliderPadding(s.I句柄, nPadding)
}

// 滚动条_置水平, 设置水平或者垂直.
//
// bHorizon: 水平或垂直.
func (s *ScrollBar) I置水平(bHorizon bool) bool {
	return xc.XSBar_EnableHorizon(s.I句柄, bHorizon)
}

// 滚动条_取滑块最大长度.
func (s *ScrollBar) I取滑块最大长度() int {
	return xc.XSBar_GetSliderMaxLength(s.I句柄)
}

// 滚动条_向上滚动.
func (s *ScrollBar) I向上滚动() bool {
	return xc.XSBar_ScrollUp(s.I句柄)
}

// 滚动条_向下滚动.
func (s *ScrollBar) I向下滚动() bool {
	return xc.XSBar_ScrollDown(s.I句柄)
}

// 滚动条_滚动到顶部.
func (s *ScrollBar) I滚动到顶部() bool {
	return xc.XSBar_ScrollTop(s.I句柄)
}

// 滚动条_滚动到底部.
func (s *ScrollBar) I滚动到底部() bool {
	return xc.XSBar_ScrollBottom(s.I句柄)
}

// 滚动条_滚动到指定位置, 滚动到指定位置点, 触发事件: XE_SBAR_SCROLL.
//
// pos: 位置点.
func (s *ScrollBar) I滚动到指定位置(pos int) bool {
	return xc.XSBar_ScrollPos(s.I句柄, pos)
}

// 滚动条_取上按钮, 获取上按钮, 返回按钮句柄.
func (s *ScrollBar) GetButtonUp() int {
	return xc.XSBar_GetButtonUp(s.I句柄)
}

// 滚动条_取下按钮, 获取下按钮, 返回按钮句柄.
func (s *ScrollBar) I取下按钮() int {
	return xc.XSBar_GetButtonDown(s.I句柄)
}

// 滚动条_取滑块, 获取滑动按钮, 返回按钮句柄.
func (s *ScrollBar) I取滑块() int {
	return xc.XSBar_GetButtonSlider(s.I句柄)
}

/*
以下都是事件
*/

type XE_SBAR_SCROLL func(pos int, pbHandled *bool) int            // 滚动条元素滚动事件,滚动条触发.
type XE_SBAR_SCROLL1 func(hEle int, pos int, pbHandled *bool) int // 滚动条元素滚动事件,滚动条触发.

// 滚动条元素滚动事件,I事件_滚动条触发.
func (s *ScrollBar) I事件_滚动条触发(pFun XE_SBAR_SCROLL) bool {
	return xc.XEle_RegEventC(s.I句柄, xcc.XE_SBAR_SCROLL, pFun)
}

// 滚动条元素滚动事件,滚动条触发.
func (s *ScrollBar) I事件_滚动条触发1(pFun XE_SBAR_SCROLL1) bool {
	return xc.XEle_RegEventC1(s.I句柄, xcc.XE_SBAR_SCROLL, pFun)
}
