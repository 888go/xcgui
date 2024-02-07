package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// 圆形(形状对象).
type ShapeEllipse struct {
	Shape
}

// 形状圆_创建, 创建圆形形状对象.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func X创建形状圆(X坐标 int, Y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) *ShapeEllipse {
	p := &ShapeEllipse{}
	p.X设置句柄(炫彩基类.X形状圆_创建(X坐标, Y坐标, 宽度, 高度, 父对象句柄))
	return p
}

// 从句柄创建对象.
func X创建形状圆并按句柄(handle int) *ShapeEllipse {
	p := &ShapeEllipse{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建形状圆并按名称(name string) *ShapeEllipse {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapeEllipse{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建形状圆并按UID(nUID int) *ShapeEllipse {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapeEllipse{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建形状圆并按UID名称(name string) *ShapeEllipse {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapeEllipse{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状圆_置边框色.
//
// color: ABGR 颜色值.
func (s *ShapeEllipse) X置边框色(ABGR颜色值 int) int {
	return 炫彩基类.X形状圆_置边框色(s.Handle, ABGR颜色值)
}

// 形状圆_置填充色.
//
// color: ABGR 颜色值.
func (s *ShapeEllipse) X置填充色(ABGR颜色值 int) int {
	return 炫彩基类.X形状圆_置填充色(s.Handle, ABGR颜色值)
}

// 形状圆_启用边框, 启用绘制圆边框.
//
// bEnable: 是否启用.
func (s *ShapeEllipse) X启用边框(是否启用 bool) int {
	return 炫彩基类.X形状圆_启用边框(s.Handle, 是否启用)
}

// 形状圆_启用填充, 启用填充圆.
//
// bEnable: 是否启用.
func (s *ShapeEllipse) X启用填充(是否启用 bool) int {
	return 炫彩基类.X形状圆_启用填充(s.Handle, 是否启用)
}
