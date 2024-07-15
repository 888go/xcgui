package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// ShapeGif 形状对象GIF.
type ShapeGif struct {
	Shape
}

// 形状GIF_创建, 创建形状对象GIF.
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

// ff:创建形状GIF
// hParent:父对象句柄
// cy:高度
// cx:宽度
// y:Y坐标
// x:坐标
func X创建形状GIF(坐标 int, Y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) *ShapeGif {
	p := &ShapeGif{}
	p.X设置句柄(炫彩基类.X形状GIF_创建(坐标, Y坐标, 宽度, 高度, 父对象句柄))
	return p
}

// 从句柄创建对象.

// ff:创建形状GIF并按句柄
// handle:
func X创建形状GIF并按句柄(handle int) *ShapeGif {
	p := &ShapeGif{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建形状GIF并按名称
// name:
func X创建形状GIF并按名称(name string) *ShapeGif {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapeGif{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建形状GIF并按UID
// nUID:
func X创建形状GIF并按UID(nUID int) *ShapeGif {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapeGif{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建形状GIF并按UID名称
// name:
func X创建形状GIF并按UID名称(name string) *ShapeGif {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapeGif{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状GIF_置图片, 设置GIF图片.
//
// hImage: 图片句柄.

// ff:置图片
// hImage:图片句柄
func (s *ShapeGif) X置图片(图片句柄 int) int {
	return 炫彩基类.X形状GIF_置图片(s.Handle, 图片句柄)
}

// 形状GIF_取图片, 获取图片句柄.

// ff:取图片
func (s *ShapeGif) X取图片() int {
	return 炫彩基类.X形状GIF_取图片(s.Handle)
}
