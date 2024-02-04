package widget

import (
	"github.com/twgh/xcgui/xc"
)

// ShapeGif 形状对象GIF.
type ShapeGif struct {
	Shape
}

// 创建形状GIF, 创建形状对象GIF.
// x: X坐标.
// y: Y坐标.
// cx: 宽度.
// cy: 高度.
// hParent: 父对象句柄.
func NewShapeGif(x int, y int, cx int, cy int, hParent int) *ShapeGif {
	p := &ShapeGif{}
	p.SetHandle(xc.XShapeGif_Create(x, y, cx, cy, hParent))
	return p
}

// 创建形状GIF并按句柄.
func NewShapeGifByHandle(handle int) *ShapeGif {
	p := &ShapeGif{}
	p.SetHandle(handle)
	return p
}

// 创建形状GIF并按名称 失败返回nil.
func NewShapeGifByName(name string) *ShapeGif {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &ShapeGif{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建形状GIF并按UID
func NewShapeGifByUID(nUID int) *ShapeGif {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &ShapeGif{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建形状GIF并按UID名称
func NewShapeGifByUIDName(name string) *ShapeGif {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &ShapeGif{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 置图片, 设置GIF图片.
// hImage: 图片句柄.
func (s *ShapeGif) SetImage(hImage int) int {
	return xc.XShapeGif_SetImage(s.Handle, hImage)
}

// 取图片, 获取图片句柄.
func (s *ShapeGif) GetImage() int {
	return xc.XShapeGif_GetImage(s.Handle)
}
