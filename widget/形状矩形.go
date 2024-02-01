package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
)

// ShapeRect 矩形形状对象.
type ShapeRect struct {
	Shape
}

// I形状矩形_创建, 创建矩形形状对象.
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
func I形状矩形_创建(x int, y int, cx int, cy int, hParent int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(xc.XShapeRect_Create(x, y, cx, cy, hParent))
	return p
}

// I形状矩形_从句柄创建  对象.
func I形状矩形_从句柄创建(handle int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(handle)
	return p
}

// I形状矩形_从name创建  对象, 失败返回nil.
func I形状矩形_从name创建(name string) *ShapeRect {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状矩形_从UID创建  对象, 失败返回nil.
func NewShapeRectByUID(nUID int) *ShapeRect {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状矩形_从UID名称创建  对象, 失败返回nil.
func I形状矩形_从UID名称创建(name string) *ShapeRect {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状矩形_置边框色, 设置边框颜色.
//
// color: ABGR 颜色值.
func (s *ShapeRect) I置边框色(color int) int {
	return xc.XShapeRect_SetBorderColor(s.I句柄, color)
}

// 形状矩形_置填充色, 设置填充颜色.
//
// color: ABGR 颜色值.
func (s *ShapeRect) I置填充色(color int) int {
	return xc.XShapeRect_SetFillColor(s.I句柄, color)
}

// 形状矩形_置圆角大小, 设置圆角大小.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeRect) I置圆角大小(nWidth int, nHeight int) int {
	return xc.XShapeRect_SetRoundAngle(s.I句柄, nWidth, nHeight)
}

// 形状矩形_取圆角大小, 获取圆角大小.
//
// pWidth: 圆角宽度.
//
// pHeight: 圆角高度.
func (s *ShapeRect) I取圆角大小(pWidth *int, pHeight *int) int {
	return xc.XShapeRect_GetRoundAngle(s.I句柄, pWidth, pHeight)
}

// 形状矩形_启用边框, 启用绘制矩形边框.
//
// bEnable: 是否启用.
func (s *ShapeRect) I启用边框(bEnable bool) int {
	return xc.XShapeRect_EnableBorder(s.I句柄, bEnable)
}

// 形状矩形_启用填充, 启用填充矩形.
//
// bEnable: 是否启用.
func (s *ShapeRect) I启用填充(bEnable bool) int {
	return xc.XShapeRect_EnableFill(s.I句柄, bEnable)
}

// 形状矩形_启用圆角.
//
// bEnable: 是否启用.
func (s *ShapeRect) I启用圆角(bEnable bool) int {
	return xc.XShapeRect_EnableRoundAngle(s.I句柄, bEnable)
}
