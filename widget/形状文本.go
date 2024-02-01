package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// ShapeText 形状对象文本.
type ShapeText struct {
	Shape
}

// I形状文本_创建, 创建形状对象文本.
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
func I形状文本_创建(x int, y int, cx int, cy int, pName string, hParent int) *ShapeText {
	p := &ShapeText{}
	p.SetHandle(xc.XShapeText_Create(x, y, cx, cy, pName, hParent))
	return p
}

// I形状文本_从句柄创建 对象.
func I形状文本_从句柄创建(handle int) *ShapeText {
	p := &ShapeText{}
	p.SetHandle(handle)
	return p
}

// I形状文本_从name创建 对象, 失败返回nil.
func I形状文本_从name创建(name string) *ShapeText {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeText{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状文本_从UID创建  对象, 失败返回nil.
func I形状文本_从UID创建(nUID int) *ShapeText {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeText{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状文本_从UID名称创建  对象, 失败返回nil.
func I形状文本_从UID名称创建(name string) *ShapeText {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeText{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状文本_置文本, 设置文本内容.
//
// pName: 文本内容.
func (s *ShapeText) I置文本(pName string) int {
	return xc.XShapeText_SetText(s.I句柄, pName)
}

// 形状文本_取文本, 获取文本内容.
func (s *ShapeText) I取文本() string {
	return xc.XShapeText_GetText(s.I句柄)
}

// 形状文本_取文本长度, 获取文本长度.
func (s *ShapeText) I取文本长度() int {
	return xc.XShapeText_GetTextLength(s.I句柄)
}

// 形状文本_置字体.
//
// hFontx: 字体句柄.
func (s *ShapeText) I置字体(hFontx int) int {
	return xc.XShapeText_SetFont(s.I句柄, hFontx)
}

// 形状文本_取字体, 返回字体句柄.
func (s *ShapeText) I取字体() int {
	return xc.XShapeText_GetFont(s.I句柄)
}

// 形状文本_置文本颜色, 设置文本颜色.
//
// color: ABGR 颜色值.
func (s *ShapeText) I置文本颜色(color int) int {
	return xc.XShapeText_SetTextColor(s.I句柄, color)
}

// 形状文本_取文本颜色.
func (s *ShapeText) I取文本颜色() int {
	return xc.XShapeText_GetTextColor(s.I句柄)
}

// 形状文本_置文本对齐.
//
// align: 文本对齐方式, I常量_文本对齐_, TextAlignFlag_, TextTrimming_.
func (s *ShapeText) I置文本对齐(align xcc.I常量_文本对齐_) int {
	return xc.XShapeText_SetTextAlign(s.I句柄, align)
}

// 形状文本_置偏移, 设置内容偏移.
//
// x: X坐标.
//
// y: Y坐标.
func (s *ShapeText) I置偏移(x int, y int) int {
	return xc.XShapeText_SetOffset(s.I句柄, x, y)
}
