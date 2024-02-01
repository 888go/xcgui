package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
)

// ShapePicture 形状对象图片.
type ShapePicture struct {
	Shape
}

// I形状图片_创建, 创建形状对象-图片.
//
// x: x坐标.
//
// y: y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父对象句柄.
func I形状图片_创建(x int, y int, cx int, cy int, hParent int) *ShapePicture {
	p := &ShapePicture{}
	p.SetHandle(xc.XShapePic_Create(x, y, cx, cy, hParent))
	return p
}

// I形状图片_从句柄创建对象.
func I形状图片_从句柄创建(handle int) *ShapePicture {
	p := &ShapePicture{}
	p.SetHandle(handle)
	return p
}

// I形状图片_从name创建对象, 失败返回nil.
func I形状图片_从name创建(name string) *ShapePicture {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapePicture{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状图片_从UID创建对象, 失败返回nil.
func I形状图片_从UID创建(nUID int) *ShapePicture {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapePicture{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I形状图片_从UID名称创建对象, 失败返回nil.
func I形状图片_从UID名称创建(name string) *ShapePicture {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapePicture{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状图片_置图片, 设置图片.
//
// hImage: 图片句柄.
func (s *ShapePicture) I置图片(hImage int) int {
	return xc.XShapePic_SetImage(s.I句柄, hImage)
}

// 形状图片_取图片, 获取图片句柄.
func (s *ShapePicture) I取图片() int {
	return xc.XShapePic_GetImage(s.I句柄)
}
