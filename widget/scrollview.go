package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ScrollView 滚动视图.
type ScrollView struct {
	Element
}

// 滚动视_创建, 创建滚动视图元素, 返回元素句柄.
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
func X创建滚动视(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *ScrollView {
	p := &ScrollView{}
	p.X设置句柄(炫彩基类.X滚动视_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建滚动视并按句柄(handle int) *ScrollView {
	p := &ScrollView{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建滚动视并按名称(name string) *ScrollView {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ScrollView{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建滚动视并按UID(nUID int) *ScrollView {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ScrollView{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建滚动视并按UID名称(name string) *ScrollView {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ScrollView{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 滚动视_置视图大小, 设置内容大小, 如果内容改变返回TRUE否则返回FALSE.
//
// cx: 宽度.
//
// cy: 高度.
func (s *ScrollView) X置视图大小(宽度 int, 高度 int) bool {
	return 炫彩基类.X滚动视_置视图大小(s.Handle, 宽度, 高度)
}

// 滚动视_取视图大小, 获取内容总大小.
//
// pSize: 大小.
func (s *ScrollView) X取视图大小(大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X滚动视_取视图大小(s.Handle, 大小)
}

// 滚动视_置滚动单位大小, 设置滚动单位大小, 如果内容改变返回TRUE否则返回FALSE.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *ScrollView) X置滚动单位大小(宽度 int, 高度 int) bool {
	return 炫彩基类.X滚动视_置滚动单位大小(s.Handle, 宽度, 高度)
}

// 滚动视_取滚动单位大小, 获取滚动单位大小.
//
// pSize: 返回大小.
func (s *ScrollView) X取滚动单位大小(返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X滚动视_取滚动单位大小(s.Handle, 返回大小)
}

// 滚动视_置滚动条大小.
//
// size: 滚动条大小.
func (s *ScrollView) X置滚动条大小(滚动条大小 int) int {
	return 炫彩基类.X滚动视_置滚动条大小(s.Handle, 滚动条大小)
}

// 滚动视_取视口原点X, 获取视口原点X坐标.
func (s *ScrollView) X取视口原点X() int {
	return 炫彩基类.X滚动视_取视口原点X(s.Handle)
}

// 滚动视_取视口原点Y, 获取视口原点Y坐标.
func (s *ScrollView) X取视口原点Y() int {
	return 炫彩基类.X滚动视_取视口原点Y(s.Handle)
}

// 滚动视_取视口宽度.
func (s *ScrollView) X取视口宽度() int {
	return 炫彩基类.X滚动视_取视口宽度(s.Handle)
}

// 滚动视_取视口高度.
func (s *ScrollView) X取视口高度() int {
	return 炫彩基类.X滚动视_取视口高度(s.Handle)
}

// 滚动视_取视口坐标.
//
// pRect: 坐标.
func (s *ScrollView) X取视口坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X滚动视_取视口坐标(s.Handle, 坐标)
}

// 滚动视_取水平滚动条, 返回滚动条句柄.
func (s *ScrollView) X取水平滚动条() int {
	return 炫彩基类.X滚动视_取水平滚动条(s.Handle)
}

// 滚动视_取垂直滚动条, 返回滚动条句柄.
func (s *ScrollView) X取垂直滚动条() int {
	return 炫彩基类.X滚动视_取垂直滚动条(s.Handle)
}

// 滚动视_水平滚动, 水平滚动条, 滚动到指定位置点.
//
// pos: 位置点.
func (s *ScrollView) X水平滚动(位置点 int) bool {
	return 炫彩基类.X滚动视_水平滚动(s.Handle, 位置点)
}

// 滚动视_垂直滚动, 垂直滚动条, 滚动到指定位置点.
//
// pos: 位置点.
func (s *ScrollView) X垂直滚动(位置点 int) bool {
	return 炫彩基类.X滚动视_垂直滚动(s.Handle, 位置点)
}

// 滚动视_水平滚动到X, 水平滚动条, 滚动到指定坐标.
//
// posX: X坐标.
func (s *ScrollView) X水平滚动到X(X坐标 int) bool {
	return 炫彩基类.X滚动视_水平滚动到X(s.Handle, X坐标)
}

// 滚动视_垂直滚动到Y, 垂直滚动条, 滚动到指定坐标.
//
// posY: Y坐标.
func (s *ScrollView) X垂直滚动到Y(Y坐标 int) bool {
	return 炫彩基类.X滚动视_垂直滚动到Y(s.Handle, Y坐标)
}

// 滚动视_显示水平滚动条.
//
// bShow: 是否显示.
func (s *ScrollView) X显示水平滚动条(是否显示 bool) int {
	return 炫彩基类.X滚动视_显示水平滚动条(s.Handle, 是否显示)
}

// 滚动视_显示垂直滚动条.
//
// bShow: 是否显示.
func (s *ScrollView) X显示垂直滚动条(是否显示 bool) int {
	return 炫彩基类.X滚动视_显示垂直滚动条(s.Handle, 是否显示)
}

// 滚动视_启用自动显示滚动条.
//
// bEnable: 是否启用.
func (s *ScrollView) X启用自动显示滚动条(是否启用 bool) int {
	return 炫彩基类.X滚动视_启用自动显示滚动条(s.Handle, 是否启用)
}

// 滚动视_向左滚动.
func (s *ScrollView) X向左滚动() bool {
	return 炫彩基类.X滚动视_向左滚动(s.Handle)
}

// 滚动视_向右滚动.
func (s *ScrollView) X向右滚动() bool {
	return 炫彩基类.X滚动视_向右滚动(s.Handle)
}

// 滚动视_向上滚动.
func (s *ScrollView) X向上滚动() bool {
	return 炫彩基类.X滚动视_向上滚动(s.Handle)
}

// 滚动视_向下滚动.
func (s *ScrollView) X向下滚动() bool {
	return 炫彩基类.X滚动视_向下滚动(s.Handle)
}

// 滚动视_滚动到左侧, 水平滚动到左侧.
func (s *ScrollView) X滚动到左侧() bool {
	return 炫彩基类.X滚动视_滚动到左侧(s.Handle)
}

// 滚动视_滚动到右侧, 水平滚动到右侧.
func (s *ScrollView) X滚动到右侧() bool {
	return 炫彩基类.X滚动视_滚动到右侧(s.Handle)
}

// 滚动视_滚动到顶部, 垂直滚动到顶部.
func (s *ScrollView) X滚动到顶部() bool {
	return 炫彩基类.X滚动视_滚动到顶部(s.Handle)
}

// 滚动视_滚动到底部, 垂直滚动到底部.
func (s *ScrollView) X滚动到底部() bool {
	return 炫彩基类.X滚动视_滚动到底部(s.Handle)
}

/*
以下都是事件
*/

type XE_SCROLLVIEW_SCROLL_H func(pos int32, pbHandled *bool) int            // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_H1 func(hEle int, pos int32, pbHandled *bool) int // 滚动视图元素水平滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V func(pos int32, pbHandled *bool) int            // 滚动视图元素垂直滚动事件,滚动视图触发.
type XE_SCROLLVIEW_SCROLL_V1 func(hEle int, pos int32, pbHandled *bool) int // 滚动视图元素垂直滚动事件,滚动视图触发.

// 滚动视图元素水平滚动事件,滚动视图触发.
func (s *ScrollView) X事件_水平滚动(pFun XE_SCROLLVIEW_SCROLL_H) bool {
	return 炫彩基类.X元素_注册事件C(s.Handle, 炫彩常量类.XE_SCROLLVIEW_SCROLL_H, pFun)
}

// 滚动视图元素水平滚动事件,滚动视图触发.
func (s *ScrollView) X事件_水平滚动1(pFun XE_SCROLLVIEW_SCROLL_H1) bool {
	return 炫彩基类.X元素_注册事件C1(s.Handle, 炫彩常量类.XE_SCROLLVIEW_SCROLL_H, pFun)
}

// 滚动视图元素垂直滚动事件,滚动视图触发.
func (s *ScrollView) X事件_垂直滚动(pFun XE_SCROLLVIEW_SCROLL_V) bool {
	return 炫彩基类.X元素_注册事件C(s.Handle, 炫彩常量类.XE_SCROLLVIEW_SCROLL_V, pFun)
}

// 滚动视图元素垂直滚动事件,滚动视图触发.
func (s *ScrollView) X事件_垂直滚动1(pFun XE_SCROLLVIEW_SCROLL_V1) bool {
	return 炫彩基类.X元素_注册事件C1(s.Handle, 炫彩常量类.XE_SCROLLVIEW_SCROLL_V, pFun)
}
