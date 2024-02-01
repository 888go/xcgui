package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
)

// 形状对象GIF.
type ShapeGif struct {
	Shape
}

// I形状GIF_创建, 创建形状对象GIF.
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
func I形状GIF_创建(x int, y int, cx int, cy int, hParent int) *ShapeGif {
	p := &ShapeGif{}
	p.SetHandle(xc.XShapeGif_Create(x, y, cx, cy, hParent))
	return p
}

// I形状GIF_从句柄创建对象.
func I形状GIF_从句柄创建(handle int) *ShapeGif {
	p := &ShapeGif{}
	p.SetHandle(handle)
	return p
}

// I形状GIF_从name创建对象, 失败返回nil.
func I形状GIF_从name创建(name string) *ShapeGif {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeGif{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状GIF_从UID创建对象, 失败返回nil.
func I形状GIF_从UID创建(nUID int) *ShapeGif {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeGif{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状GIF_从UID名称创建对象, 失败返回nil.
func I形状GIF_从UID名称创建(name string) *ShapeGif {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeGif{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状GIF_置图片, 设置GIF图片.
//
// hImage: 图片句柄.
func (s *ShapeGif) I置图片(hImage int) int {
	return xc.XShapeGif_SetImage(s.I句柄, hImage)
}

// 形状GIF_取图片, 获取图片句柄.
func (s *ShapeGif) I取图片() int {
	return xc.XShapeGif_GetImage(s.I句柄)
}
