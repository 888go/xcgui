package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// 组框(形状对象).
type ShapeGroupBox struct {
	Shape
}

// 形状组框_创建, 创建组框形状对象.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 名称.
//
// hParent: 父对象句柄.
func X创建形状组框(X坐标 int, Y坐标 int, 宽度 int, 高度 int, 名称 string, 父对象句柄 int) *ShapeGroupBox {
	p := &ShapeGroupBox{}
	p.X设置句柄(炫彩基类.X形状组框_创建(X坐标, Y坐标, 宽度, 高度, 名称, 父对象句柄))
	return p
}

// 从句柄创建对象.
func X创建形状组框并按句柄(handle int) *ShapeGroupBox {
	p := &ShapeGroupBox{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建形状组框并按名称(name string) *ShapeGroupBox {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建形状组框并按UID(nUID int) *ShapeGroupBox {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建形状组框并按UID名称(name string) *ShapeGroupBox {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状组框_置边框颜色.
//
// color: ABGR 颜色值.
func (s *ShapeGroupBox) X置边框颜色(ABGR颜色值 int) int {
	return 炫彩基类.X形状组框_置边框颜色(s.Handle, ABGR颜色值)
}

// 形状组框_置文本颜色.
//
// color: ABGR 颜色值.
func (s *ShapeGroupBox) X置文本颜色(ABGR颜色值 int) int {
	return 炫彩基类.X形状组框_置文本颜色(s.Handle, ABGR颜色值)
}

// 形状组框_置字体.
//
// hFontX: 炫彩字体.
func (s *ShapeGroupBox) X置字体(炫彩字体 int) int {
	return 炫彩基类.X形状组框_置字体(s.Handle, 炫彩字体)
}

// 形状组框_置文本偏移, 设置文本偏移量.
//
// offsetX: 水平偏移.
//
// offsetY: 垂直偏移.
func (s *ShapeGroupBox) X置文本偏移(水平偏移 int32, 垂直偏移 int32) int {
	return 炫彩基类.X形状组框_置文本偏移(s.Handle, 水平偏移, 垂直偏移)
}

// 形状组框_置圆角大小.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeGroupBox) X置圆角大小(圆角宽度 int32, 圆角高度 int32) int {
	return 炫彩基类.X形状组框_置圆角大小(s.Handle, 圆角宽度, 圆角高度)
}

// 形状组框_置文本.
//
// pText: 文本内容.
func (s *ShapeGroupBox) X置文本(文本内容 string) int {
	return 炫彩基类.X形状组框_置文本(s.Handle, 文本内容)
}

// 形状组框_取文本偏移, 获取文本偏移量.
//
// pOffsetX: X坐标偏移量.
//
// pOffsetY: Y坐标偏移量.
func (s *ShapeGroupBox) X取文本偏移(X坐标偏移量 *int32, Y坐标偏移量 *int32) int {
	return 炫彩基类.X形状组框_取文本偏移(s.Handle, X坐标偏移量, Y坐标偏移量)
}

// 形状组框_取圆角大小.
//
// pWidth: 返回圆角宽度.
//
// pHeight: 返回圆角高度.
func (s *ShapeGroupBox) X取圆角大小(返回圆角宽度 *int32, 返回圆角高度 *int32) int {
	return 炫彩基类.X形状组框_取圆角大小(s.Handle, 返回圆角宽度, 返回圆角高度)
}

// 形状组框_启用圆角.
//
// bEnable: 是否启用.
func (s *ShapeGroupBox) X启用圆角(是否启用 bool) int {
	return 炫彩基类.X形状组框_启用圆角(s.Handle, 是否启用)
}
