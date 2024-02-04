package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeRect 矩形形状对象.
type ShapeRect struct {
	Shape
}

// 创建形状矩形
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
func NewShapeRect(x int, y int, cx int, cy int, hParent int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(xc.XShapeRect_Create(x, y, cx, cy, hParent))
	return p
}

// 创建形状矩形并按句柄.
func NewShapeRectByHandle(handle int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(handle)
	return p
}

// 创建形状矩形并按名称, 失败返回nil.
func NewShapeRectByName(name string) *ShapeRect {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建形状矩形并按UID
func NewShapeRectByUID(nUID int) *ShapeRect {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建形状矩形并按UID名称
func NewShapeRectByUIDName(name string) *ShapeRect {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 置边框色, 设置边框颜色.
//
// color: ABGR 颜色值.
func (s *ShapeRect) SetBorderColor(color int) int {
	return xc.XShapeRect_SetBorderColor(s.Handle, color)
}

// 置填充色, 设置填充颜色.
//
// color: ABGR 颜色值.
func (s *ShapeRect) SetFillColor(color int) int {
	return xc.XShapeRect_SetFillColor(s.Handle, color)
}

// 置圆角大小, 设置圆角大小.
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeRect) SetRoundAngle(nWidth int, nHeight int32) int {
	return xc.XShapeRect_SetRoundAngle(s.Handle, nWidth, nHeight)
}

// 取圆角大小, 获取圆角大小.
//
// pWidth: 圆角宽度.
//
// pHeight: 圆角高度.
func (s *ShapeRect) GetRoundAngle(pWidth *int, pHeight *int32) int {
	return xc.XShapeRect_GetRoundAngle(s.Handle, pWidth, pHeight)
}

// 启用边框, 启用绘制矩形边框.
//
// bEnable: 是否启用.
func (s *ShapeRect) EnableBorder(bEnable bool) int {
	return xc.XShapeRect_EnableBorder(s.Handle, bEnable)
}

// 启用填充, 启用填充矩形.
//
// bEnable: 是否启用.
func (s *ShapeRect) EnableFill(bEnable bool) int {
	return xc.XShapeRect_EnableFill(s.Handle, bEnable)
}

// 启用圆角.
//
// bEnable: 是否启用.
func (s *ShapeRect) EnableRoundAngle(bEnable bool) int {
	return xc.XShapeRect_EnableRoundAngle(s.Handle, bEnable)
}
