package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapePicture 形状对象图片.
type ShapePicture struct {
	Shape
}

// NewShapePicture 形状图片_创建, 创建形状对象-图片.
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

// ff:创建形状图片
// hParent:父对象句柄
// cy:高度
// cx:宽度
// y:y坐标
// x:x坐标
func NewShapePicture(x int, y int, cx int, cy int, hParent int) *ShapePicture {
	p := &ShapePicture{}
	p.SetHandle(xc.XShapePic_Create(x, y, cx, cy, hParent))
	return p
}

// 从句柄创建对象.

// ff:创建形状图片并按句柄
// handle:
func NewShapePictureByHandle(handle int) *ShapePicture {
	p := &ShapePicture{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建形状图片并按名称
// name:
func NewShapePictureByName(name string) *ShapePicture {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapePicture{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建形状图片并按UID
// nUID:
func NewShapePictureByUID(nUID int) *ShapePicture {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapePicture{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建形状图片并按UID名称
// name:
func NewShapePictureByUIDName(name string) *ShapePicture {
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

// ff:置图片
// hImage:图片句柄
func (s *ShapePicture) SetImage(hImage int) int {
	return xc.XShapePic_SetImage(s.Handle, hImage)
}

// 形状图片_取图片, 获取图片句柄.

// ff:取图片
func (s *ShapePicture) GetImage() int {
	return xc.XShapePic_GetImage(s.Handle)
}
