package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// ShapeLine 直线(形状对象).
type ShapeLine struct {
	Shape
}

// 形状线_创建.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
//
// hParent: 父对象句柄.
func X创建形状线(坐标x1 int, 坐标y1 int, 坐标x2 int, 坐标y2 int, 父对象句柄 int) *ShapeLine {
	p := &ShapeLine{}
	p.X设置句柄(炫彩基类.X形状线_创建(坐标x1, 坐标y1, 坐标x2, 坐标y2, 父对象句柄))
	return p
}

// 从句柄创建对象.
func X创建形状线并按句柄(handle int) *ShapeLine {
	p := &ShapeLine{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建形状线并按名称(name string) *ShapeLine {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapeLine{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建形状线并按UID(nUID int) *ShapeLine {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapeLine{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建形状线并按UID名称(name string) *ShapeLine {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapeLine{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状线_置位置.
//
// x1: 坐标.
//
// y1: 坐标.
//
// x2: 坐标.
//
// y2: 坐标.
func (s *ShapeLine) X置位置(坐标x1 int, 坐标y1 int, 坐标x2 int, 坐标y2 int) int {
	return 炫彩基类.X形状线_置位置(s.Handle, 坐标x1, 坐标y1, 坐标x2, 坐标y2)
}

// 形状线_置颜色, 设置直线颜色.
//
// color: ABGR 颜色值.
func (s *ShapeLine) X置颜色(ABGR颜色值 int) int {
	return 炫彩基类.X形状线_置颜色(s.Handle, ABGR颜色值)
}
