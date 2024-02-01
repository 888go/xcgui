package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
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
func I形状线_创建(x1 int, y1 int, x2 int, y2 int, hParent int) *ShapeLine {
	p := &ShapeLine{}
	p.SetHandle(xc.XShapeLine_Create(x1, y1, x2, y2, hParent))
	return p
}

// I形状线_从句柄创建 对象.
func I形状线_从句柄创建(handle int) *ShapeLine {
	p := &ShapeLine{}
	p.SetHandle(handle)
	return p
}

// I形状线_从name创建 对象, 失败返回nil.
func I形状线_从name创建(name string) *ShapeLine {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeLine{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状线_从UID创建 对象, 失败返回nil.
func I形状线_从UID创建(nUID int) *ShapeLine {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeLine{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状线_从UID名称创建 对象, 失败返回nil.
func I形状线_从UID名称创建(name string) *ShapeLine {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeLine{}
		p.SetHandle(handle)
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
func (s *ShapeLine) I置位置(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XShapeLine_SetPosition(s.I句柄, x1, y1, x2, y2)
}

// 形状线_置颜色, 设置直线颜色.
//
// color: ABGR 颜色值.
func (s *ShapeLine) I置颜色(color int) int {
	return xc.XShapeLine_SetColor(s.I句柄, color)
}
