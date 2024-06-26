package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeRect 矩形形状对象.
type ShapeRect struct {
	Shape
}

// 形状矩形_创建, 创建矩形形状对象.
//.
//.
//.
//.
//.
// ff:创建形状矩形
// x:坐标
// y:Y坐标
// cx:宽度
// cy:高度
// hParent:父对象句柄
func NewShapeRect(x int, y int, cx int, cy int, hParent int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(xc.XShapeRect_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.
// ff:创建形状矩形并按句柄
// handle:
func NewShapeRectByHandle(handle int) *ShapeRect {
	p := &ShapeRect{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
// ff:创建形状矩形并按名称
// name:
func NewShapeRectByName(name string) *ShapeRect {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
// ff:创建形状矩形并按UID
// nUID:
func NewShapeRectByUID(nUID int) *ShapeRect {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
// ff:创建形状矩形并按UID名称
// name:
func NewShapeRectByUIDName(name string) *ShapeRect {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeRect{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状矩形_置边框色, 设置边框颜色.
//.
// ff:置边框色
// color:ABGR颜色值
func (s *ShapeRect) SetBorderColor(color int) int {
	return xc.XShapeRect_SetBorderColor(s.Handle, color)
}

// 形状矩形_置填充色, 设置填充颜色.
//.
// ff:置填充色
// color:ABGR颜色值
func (s *ShapeRect) SetFillColor(color int) int {
	return xc.XShapeRect_SetFillColor(s.Handle, color)
}

// 形状矩形_置圆角大小, 设置圆角大小.
//.
//.
// ff:置圆角大小
// nWidth:圆角宽度
// nHeight:圆角高度
func (s *ShapeRect) SetRoundAngle(nWidth int, nHeight int32) int {
	return xc.XShapeRect_SetRoundAngle(s.Handle, nWidth, nHeight)
}

// 形状矩形_取圆角大小, 获取圆角大小.
//.
//.
// ff:取圆角大小
// pWidth:圆角宽度
// pHeight:圆角高度
func (s *ShapeRect) GetRoundAngle(pWidth *int, pHeight *int32) int {
	return xc.XShapeRect_GetRoundAngle(s.Handle, pWidth, pHeight)
}

// 形状矩形_启用边框, 启用绘制矩形边框.
//.
// ff:启用边框
// bEnable:是否启用
func (s *ShapeRect) EnableBorder(bEnable bool) int {
	return xc.XShapeRect_EnableBorder(s.Handle, bEnable)
}

// 形状矩形_启用填充, 启用填充矩形.
//.
// ff:启用填充
// bEnable:是否启用
func (s *ShapeRect) EnableFill(bEnable bool) int {
	return xc.XShapeRect_EnableFill(s.Handle, bEnable)
}

// 形状矩形_启用圆角.
//.
// ff:启用圆角
// bEnable:是否启用
func (s *ShapeRect) EnableRoundAngle(bEnable bool) int {
	return xc.XShapeRect_EnableRoundAngle(s.Handle, bEnable)
}
