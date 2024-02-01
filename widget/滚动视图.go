package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 滚动视图.
type ScrollView struct {
	Element
}

// I滚动视图_创建, 创建滚动视图元素, 返回元素句柄.
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
func I滚动视图_创建(x int, y int, cx int, cy int, hParent int) *ScrollView {
	p := &ScrollView{}
	p.SetHandle(xc.XSView_Create(x, y, cx, cy, hParent))
	return p
}

// I滚动视图_从句柄创建对象.
func I滚动视图_从句柄创建(handle int) *ScrollView {
	p := &ScrollView{}
	p.SetHandle(handle)
	return p
}

// I滚动视图_从name创建对象, 失败返回nil.
func I滚动视图_从name创建(name string) *ScrollView {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ScrollView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I滚动视图_从UID创建对象, 失败返回nil.
func I滚动视图_从UID创建(nUID int) *ScrollView {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ScrollView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I滚动视图_从UID名称创建对象, 失败返回nil.
func I滚动视图_从UID名称创建(name string) *ScrollView {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ScrollView{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 滚动视_置视图大小, 设置内容大小, 如果内容改变返回TRUE否则返回FALSE.
//
// cx: 宽度.
//
// cy: 高度.
func (s *ScrollView) I置视图大小(cx int, cy int) bool {
	return xc.XSView_SetTotalSize(s.I句柄, cx, cy)
}

// 滚动视_取视图大小, 获取内容总大小.
//
// pSize: 大小.
func (s *ScrollView) I取视图大小(pSize *xc.SIZE) int {
	return xc.XSView_GetTotalSize(s.I句柄, pSize)
}

// 滚动视_置滚动单位大小, 设置滚动单位大小, 如果内容改变返回TRUE否则返回FALSE.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *ScrollView) I置滚动单位大小(nWidth int, nHeight int) bool {
	return xc.XSView_SetLineSize(s.I句柄, nWidth, nHeight)
}

// 滚动视_取滚动单位大小, 获取滚动单位大小.
//
// pSize: 返回大小.
func (s *ScrollView) I取滚动单位大小(pSize *xc.SIZE) int {
	return xc.XSView_GetLineSize(s.I句柄, pSize)
}

// 滚动视_置滚动条大小.
//
// size: 滚动条大小.
func (s *ScrollView) I置滚动条大小(size int) int {
	return xc.XSView_SetScrollBarSize(s.I句柄, size)
}

// 滚动视_取视口原点X, 获取视口原点X坐标.
func (s *ScrollView) I取视口原点X() int {
	return xc.XSView_GetViewPosH(s.I句柄)
}

// 滚动视_取视口原点Y, 获取视口原点Y坐标.
func (s *ScrollView) I取视口原点Y() int {
	return xc.XSView_GetViewPosV(s.I句柄)
}

// 滚动视_取视口宽度.
func (s *ScrollView) I取视口宽度() int {
	return xc.XSView_GetViewWidth(s.I句柄)
}

// 滚动视_取视口高度.
func (s *ScrollView) I取视口高度() int {
	return xc.XSView_GetViewHeight(s.I句柄)
}

// 滚动视_取视口坐标.
//
// pRect: 坐标.
func (s *ScrollView) I取视口坐标(pRect *xc.RECT) int {
	return xc.XSView_GetViewRect(s.I句柄, pRect)
}

// 滚动视_取水平滚动条, 返回滚动条句柄.
func (s *ScrollView) I取水平滚动条() int {
	return xc.XSView_GetScrollBarH(s.I句柄)
}

// 滚动视_取垂直滚动条, 返回滚动条句柄.
func (s *ScrollView) I取垂直滚动条() int {
	return xc.XSView_GetScrollBarV(s.I句柄)
}

// 滚动视_水平滚动, 水平滚动条, 滚动到指定位置点.
//
// pos: 位置点.
func (s *ScrollView) I水平滚动(pos int) bool {
	return xc.XSView_ScrollPosH(s.I句柄, pos)
}

// 滚动视_垂直滚动, 垂直滚动条, 滚动到指定位置点.
//
// pos: 位置点.
func (s *ScrollView) I垂直滚动(pos int) bool {
	return xc.XSView_ScrollPosV(s.I句柄, pos)
}

// 滚动视_水平滚动到X, 水平滚动条, 滚动到指定坐标.
//
// posX: X坐标.
func (s *ScrollView) I水平滚动到X(posX int) bool {
	return xc.XSView_ScrollPosXH(s.I句柄, posX)
}

// 滚动视_垂直滚动到Y, 垂直滚动条, 滚动到指定坐标.
//
// posY: Y坐标.
func (s *ScrollView) I垂直滚动到Y(posY int) bool {
	return xc.XSView_ScrollPosYV(s.I句柄, posY)
}

// 滚动视_显示水平滚动条.
//
// bShow: 是否显示.
func (s *ScrollView) I显示水平滚动条(bShow bool) int {
	return xc.XSView_ShowSBarH(s.I句柄, bShow)
}

// 滚动视_显示垂直滚动条.
//
// bShow: 是否显示.
func (s *ScrollView) I显示垂直滚动条(bShow bool) int {
	return xc.XSView_ShowSBarV(s.I句柄, bShow)
}

// 滚动视_启用自动显示滚动条.
//
// bEnable: 是否启用.
func (s *ScrollView) I启用自动显示滚动条(bEnable bool) int {
	return xc.XSView_EnableAutoShowScrollBar(s.I句柄, bEnable)
}

// 滚动视_向左滚动.
func (s *ScrollView) I向左滚动() bool {
	return xc.XSView_ScrollLeftLine(s.I句柄)
}

// 滚动视_向右滚动.
func (s *ScrollView) I向右滚动() bool {
	return xc.XSView_ScrollRightLine(s.I句柄)
}

// 滚动视_向上滚动.
func (s *ScrollView) I向上滚动() bool {
	return xc.XSView_ScrollTopLine(s.I句柄)
}

// 滚动视_向下滚动.
func (s *ScrollView) I向下滚动() bool {
	return xc.XSView_ScrollBottomLine(s.I句柄)
}

// 滚动视_滚动到左侧, 水平滚动到左侧.
func (s *ScrollView) I滚动到左侧() bool {
	return xc.XSView_ScrollLeft(s.I句柄)
}

// 滚动视_滚动到右侧, 水平滚动到右侧.
func (s *ScrollView) I滚动到右侧() bool {
	return xc.XSView_ScrollRight(s.I句柄)
}

// 滚动视_滚动到顶部, 垂直滚动到顶部.
func (s *ScrollView) I滚动到顶部() bool {
	return xc.XSView_ScrollTop(s.I句柄)
}

// 滚动视_滚动到底部, 垂直滚动到底部.
func (s *ScrollView) I滚动到底部() bool {
	return xc.XSView_ScrollBottom(s.I句柄)
}

/*
以下都是事件
*/

type XE_SCROLLVIEW_SCROLL_H func(pos int, pbHandled *bool) int            // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_H1 func(hEle int, pos int, pbHandled *bool) int // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V func(pos int, pbHandled *bool) int            // 滚动视图元素垂直滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V1 func(hEle int, pos int, pbHandled *bool) int // 滚动视图元素垂直滚动事件,滚动视图触发.

// 滚动视图元素  I事件_水平滚动 事件,滚动视图触发.
func (s *ScrollView) I事件_水平滚动(pFun XE_SCROLLVIEW_SCROLL_H) bool {
	return xc.XEle_RegEventC(s.I句柄, xcc.XE_SCROLLVIEW_SCROLL_H, pFun)
}

// 滚动视图元素 I事件_水平滚动 事件,滚动视图触发.
func (s *ScrollView) I事件_水平滚动1(pFun XE_SCROLLVIEW_SCROLL_H1) bool {
	return xc.XEle_RegEventC1(s.I句柄, xcc.XE_SCROLLVIEW_SCROLL_H, pFun)
}

// 滚动视图元素 I事件_垂直滚动 事件,滚动视图触发.
func (s *ScrollView) I事件_垂直滚动(pFun XE_SCROLLVIEW_SCROLL_V) bool {
	return xc.XEle_RegEventC(s.I句柄, xcc.XE_SCROLLVIEW_SCROLL_V, pFun)
}

// 滚动视图元素 I事件_垂直滚动 事件,滚动视图触发.
func (s *ScrollView) I事件_垂直滚动1(pFun XE_SCROLLVIEW_SCROLL_V1) bool {
	return xc.XEle_RegEventC1(s.I句柄, xcc.XE_SCROLLVIEW_SCROLL_V, pFun)
}
