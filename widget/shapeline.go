package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeLine 直线(形状对象).
type ShapeLine struct {
	Shape
}

// 形状线_创建.
//.
//.
//.
//.
//.
// ff:创建形状线
// x1:坐标x1
// y1:坐标y1
// x2:坐标x2
// y2:坐标y2
// hParent:父对象句柄
func NewShapeLine(x1 int, y1 int, x2 int, y2 int, hParent int) *ShapeLine {
	p := &ShapeLine{}
	p.SetHandle(xc.XShapeLine_Create(x1, y1, x2, y2, hParent))
	return p
}

// 从句柄创建对象.
// ff:创建形状线并按句柄
// handle:
func NewShapeLineByHandle(handle int) *ShapeLine {
	p := &ShapeLine{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
// ff:创建形状线并按名称
// name:
func NewShapeLineByName(name string) *ShapeLine {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeLine{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
// ff:创建形状线并按UID
// nUID:
func NewShapeLineByUID(nUID int) *ShapeLine {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeLine{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
// ff:创建形状线并按UID名称
// name:
func NewShapeLineByUIDName(name string) *ShapeLine {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeLine{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状线_置位置.
//.
//.
//.
//.
// ff:置位置
// x1:坐标x1
// y1:坐标y1
// x2:坐标x2
// y2:坐标y2
func (s *ShapeLine) SetPosition(x1 int, y1 int, x2 int, y2 int) int {
	return xc.XShapeLine_SetPosition(s.Handle, x1, y1, x2, y2)
}

// 形状线_置颜色, 设置直线颜色.
//.
// ff:置颜色
// color:ABGR颜色值
func (s *ShapeLine) SetColor(color int) int {
	return xc.XShapeLine_SetColor(s.Handle, color)
}
