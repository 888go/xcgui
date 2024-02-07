package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// ShapeRect 矩形形状对象.
type ShapeRect struct {
	Shape
}

// 形状矩形_创建, 创建矩形形状对象.
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
func X创建形状矩形(X坐标 int, Y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) *ShapeRect {
	p := &ShapeRect{}
	p.X设置句柄(炫彩基类.X形状矩形_创建(X坐标, Y坐标, 宽度, 高度, 父对象句柄))
	return p
}

// 从句柄创建对象.
func X创建形状矩形并按句柄(handle int) *ShapeRect {
	p := &ShapeRect{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建形状矩形并按名称(name string) *ShapeRect {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建形状矩形并按UID(nUID int) *ShapeRect {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapeRect{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建形状矩形并按UID名称(name string) *ShapeRect {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状矩形_置边框色, 设置边框颜色.
//
// color: ABGR 颜色值.
func (s *ShapeRect) X置边框色(ABGR颜色值 int) int {
	return 炫彩基类.X形状矩形_置边框色(s.Handle, ABGR颜色值)
}

// 形状矩形_置填充色, 设置填充颜色.
//
// color: ABGR 颜色值.
func (s *ShapeRect) X置填充色(ABGR颜色值 int) int {
	return 炫彩基类.X形状矩形_置填充色(s.Handle, ABGR颜色值)
}

// 形状矩形_置圆角大小, 设置圆角大小.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeRect) X置圆角大小(圆角宽度 int, 圆角高度 int32) int {
	return 炫彩基类.X形状矩形_置圆角大小(s.Handle, 圆角宽度, 圆角高度)
}

// 形状矩形_取圆角大小, 获取圆角大小.
//
// pWidth: 圆角宽度.
//
// pHeight: 圆角高度.
func (s *ShapeRect) X取圆角大小(圆角宽度 *int, 圆角高度 *int32) int {
	return 炫彩基类.X形状矩形_取圆角大小(s.Handle, 圆角宽度, 圆角高度)
}

// 形状矩形_启用边框, 启用绘制矩形边框.
//
// bEnable: 是否启用.
func (s *ShapeRect) X启用边框(是否启用 bool) int {
	return 炫彩基类.X形状矩形_启用边框(s.Handle, 是否启用)
}

// 形状矩形_启用填充, 启用填充矩形.
//
// bEnable: 是否启用.
func (s *ShapeRect) X启用填充(是否启用 bool) int {
	return 炫彩基类.X形状矩形_启用填充(s.Handle, 是否启用)
}

// 形状矩形_启用圆角.
//
// bEnable: 是否启用.
func (s *ShapeRect) X启用圆角(是否启用 bool) int {
	return 炫彩基类.X形状矩形_启用圆角(s.Handle, 是否启用)
}
