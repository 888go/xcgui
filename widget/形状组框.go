package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
)

// 组框(形状对象).
type ShapeGroupBox struct {
	Shape
}

// I形状组框_创建, 创建组框形状对象.
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
func I形状组框_创建(x int, y int, cx int, cy int, pName string, hParent int) *ShapeGroupBox {
	p := &ShapeGroupBox{}
	p.SetHandle(xc.XShapeGroupBox_Create(x, y, cx, cy, pName, hParent))
	return p
}

// I形状组框_从句柄创建  对象.
func I形状组框_从句柄创建(handle int) *ShapeGroupBox {
	p := &ShapeGroupBox{}
	p.SetHandle(handle)
	return p
}

// I形状组框_从name创建  对象, 失败返回nil.
func I形状组框_从name创建(name string) *ShapeGroupBox {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状组框_从UID创建  对象, 失败返回nil.
func I形状组框_从UID创建(nUID int) *ShapeGroupBox {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状组框_从UID名称创建  对象, 失败返回nil.
func I形状组框_从UID名称创建(name string) *ShapeGroupBox {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeGroupBox{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状组框_  I置边框颜色 .
//
// color: ABGR 颜色值.
func (s *ShapeGroupBox) I置边框颜色(color int) int {
	return xc.XShapeGroupBox_SetBorderColor(s.I句柄, color)
}

// 形状组框_ I置文本颜色 .
//
// color: ABGR 颜色值.
func (s *ShapeGroupBox) I置文本颜色(color int) int {
	return xc.XShapeGroupBox_SetTextColor(s.I句柄, color)
}

// 形状组框_ I置字体 .
//
// hFontX: 炫彩字体.
func (s *ShapeGroupBox) I置字体(hFontX int) int {
	return xc.XShapeGroupBox_SetFontX(s.I句柄, hFontX)
}

// 形状组框_ I置文本偏移 , 设置文本偏移量.
//
// offsetX: 水平偏移.
//
// offsetY: 垂直偏移.
func (s *ShapeGroupBox) I置文本偏移(offsetX int, offsetY int) int {
	return xc.XShapeGroupBox_SetTextOffset(s.I句柄, offsetX, offsetY)
}

// 形状组框_ I置圆角大小 .
//
// nWidth: 圆角宽度.
//
// nHeight: 圆角高度.
func (s *ShapeGroupBox) I置圆角大小(nWidth int, nHeight int) int {
	return xc.XShapeGroupBox_SetRoundAngle(s.I句柄, nWidth, nHeight)
}

// 形状组框_ I置文本 .
//
// pText: 文本内容.
func (s *ShapeGroupBox) I置文本(pText string) int {
	return xc.XShapeGroupBox_SetText(s.I句柄, pText)
}

// 形状组框_ I取文本偏移 , 获取文本偏移量.
//
// pOffsetX: X坐标偏移量.
//
// pOffsetY: Y坐标偏移量.
func (s *ShapeGroupBox) I取文本偏移(pOffsetX *int, pOffsetY *int) int {
	return xc.XShapeGroupBox_GetTextOffset(s.I句柄, pOffsetX, pOffsetY)
}

// 形状组框_ I取圆角大小 .
//
// pWidth: 返回圆角宽度.
//
// pHeight: 返回圆角高度.
func (s *ShapeGroupBox) I取圆角大小(pWidth *int, pHeight *int) int {
	return xc.XShapeGroupBox_GetRoundAngle(s.I句柄, pWidth, pHeight)
}

// 形状组框_ I启用圆角 .
//
// bEnable: 是否启用.
func (s *ShapeGroupBox) I启用圆角(bEnable bool) int {
	return xc.XShapeGroupBox_EnableRoundAngle(s.I句柄, bEnable)
}
