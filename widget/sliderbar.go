package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// SliderBar 滑动条元素.
type SliderBar struct {
	Element
}

// 滑动条_创建, 创建滑动条元素.
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
func X创建滑动条(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *SliderBar {
	p := &SliderBar{}
	p.X设置句柄(炫彩基类.X滑动条_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建滑动条并按句柄(handle int) *SliderBar {
	p := &SliderBar{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建滑动条并按名称(name string) *SliderBar {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &SliderBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建滑动条并按UID(nUID int) *SliderBar {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &SliderBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建滑动条并按UID名称(name string) *SliderBar {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &SliderBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 滑动条_置范围, 设置滑动范围.
//
// range_: 范围.
func (s *SliderBar) X置范围(范围 int) int {
	return 炫彩基类.X滑动条_置范围(s.Handle, 范围)
}

// 滑动条_取范围, 获取滚动范围.
func (s *SliderBar) X取范围() int {
	return 炫彩基类.X滑动条_取范围(s.Handle)
}

// 滑动条_置进度图片, 设置进度贴图.
//
// hImage: 图片句柄.
func (s *SliderBar) X置进度图片(图片句柄 int) int {
	return 炫彩基类.X滑动条_置进度图片(s.Handle, 图片句柄)
}

// 滑动条_置滑块宽度, 设置滑块按钮宽度.
//
// width: 宽度.
func (s *SliderBar) X置滑块宽度(宽度 int) int {
	return 炫彩基类.X滑动条_置滑块宽度(s.Handle, 宽度)
}

// 滑动条_置滑块高度, 设置滑块按钮高度.
//
// height: 高度.
func (s *SliderBar) X置滑块高度(高度 int) int {
	return 炫彩基类.X滑动条_置滑块高度(s.Handle, 高度)
}

// 滑动条_置当前位置, 设置当前进度点.
//
// pos: 进度点.
func (s *SliderBar) X置当前位置(进度点 int) int {
	return 炫彩基类.X滑动条_置当前位置(s.Handle, 进度点)
}

// 滑动条_取当前位置, 获取当前进度点.
func (s *SliderBar) X取当前位置() int {
	return 炫彩基类.X滑动条_取当前位置(s.Handle)
}

// 滑动条_取滑块, 返回滑块按钮句柄.
func (s *SliderBar) X取滑块() int {
	return 炫彩基类.X滑动条_取滑块(s.Handle)
}

// 滑动条_置水平, 设置水平或垂直.
//
// bHorizon: 水平或垂直.
func (s *SliderBar) X置水平(水平或垂直 bool) int {
	return 炫彩基类.X滑动条_置水平(s.Handle, 水平或垂直)
}

/*
以下都是事件
*/

type XE_SLIDERBAR_CHANGE func(pos int32, pbHandled *bool) int            // 滑动条元素,滑块位置改变事件.
type XE_SLIDERBAR_CHANGE1 func(hEle int, pos int32, pbHandled *bool) int // 滑动条元素,滑块位置改变事件.

// 滑动条元素,滑块位置改变事件.
func (s *SliderBar) X事件_滑块位置改变(pFun XE_SLIDERBAR_CHANGE) bool {
	return 炫彩基类.X元素_注册事件C(s.Handle, 炫彩常量类.XE_SLIDERBAR_CHANGE, pFun)
}

// 滑动条元素,滑块位置改变事件.
func (s *SliderBar) X事件_滑块位置改变1(pFun XE_SLIDERBAR_CHANGE1) bool {
	return 炫彩基类.X元素_注册事件C1(s.Handle, 炫彩常量类.XE_SLIDERBAR_CHANGE, pFun)
}
