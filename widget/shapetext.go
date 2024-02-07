package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// ShapeText 形状对象文本.
type ShapeText struct {
	Shape
}

// 形状文本_创建, 创建形状对象文本.
//
// x: X坐标.
//
// y: Y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// pName: 文本内容.
//
// hParent: 父对象句柄.
func X创建形状文本(X坐标, Y坐标, 宽度, 高度 int32, 文本内容 string, 父对象句柄 int) *ShapeText {
	p := &ShapeText{}
	p.X设置句柄(炫彩基类.X形状文本_创建(X坐标, Y坐标, 宽度, 高度, 文本内容, 父对象句柄))
	return p
}

// 从句柄创建对象.
func X创建形状文本并按句柄(handle int) *ShapeText {
	p := &ShapeText{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建形状文本并按名称(name string) *ShapeText {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapeText{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建形状文本并按UID(nUID int) *ShapeText {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapeText{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建形状文本并按UID名称(name string) *ShapeText {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapeText{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状文本_置文本, 设置文本内容.
//
// pName: 文本内容.
func (s *ShapeText) X置文本(文本内容 string) int {
	return 炫彩基类.X形状文本_置文本(s.Handle, 文本内容)
}

// 形状文本_取文本, 获取文本内容.
func (s *ShapeText) X取文本() string {
	return 炫彩基类.X形状文本_取文本(s.Handle)
}

// 形状文本_取文本长度, 获取文本长度.
func (s *ShapeText) X取文本长度() int {
	return 炫彩基类.X形状文本_取文本长度(s.Handle)
}

// 形状文本_置字体.
//
// hFontx: 字体句柄.
func (s *ShapeText) X置字体(字体句柄 int) int {
	return 炫彩基类.X形状文本_置字体(s.Handle, 字体句柄)
}

// 形状文本_取字体, 返回字体句柄.
func (s *ShapeText) X取字体() int {
	return 炫彩基类.X形状文本_取字体(s.Handle)
}

// 形状文本_置文本颜色, 设置文本颜色.
//
// color: ABGR 颜色值.
func (s *ShapeText) X置文本颜色(ABGR颜色值 int) int {
	return 炫彩基类.X形状文本_置文本颜色(s.Handle, ABGR颜色值)
}

// 形状文本_取文本颜色.
func (s *ShapeText) X取文本颜色() int {
	return 炫彩基类.X形状文本_取文本颜色(s.Handle)
}

// 形状文本_置文本对齐.
//
// align: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (s *ShapeText) X置文本对齐(文本对齐方式 炫彩常量类.TextFormatFlag_) int {
	return 炫彩基类.X形状文本_置文本对齐(s.Handle, 文本对齐方式)
}

// 形状文本_置偏移, 设置内容偏移.
//
// x: X坐标.
//
// y: Y坐标.
func (s *ShapeText) X置偏移(X坐标 int, Y坐标 int) int {
	return 炫彩基类.X形状文本_置偏移(s.Handle, X坐标, Y坐标)
}
