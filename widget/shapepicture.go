package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
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
func X创建形状图片(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父对象句柄 int) *ShapePicture {
	p := &ShapePicture{}
	p.X设置句柄(炫彩基类.X形状图片_创建(x坐标, y坐标, 宽度, 高度, 父对象句柄))
	return p
}

// 从句柄创建对象.

// ff:创建形状图片并按句柄
// handle:
func X创建形状图片并按句柄(handle int) *ShapePicture {
	p := &ShapePicture{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建形状图片并按名称
// name:
func X创建形状图片并按名称(name string) *ShapePicture {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &ShapePicture{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建形状图片并按UID
// nUID:
func X创建形状图片并按UID(nUID int) *ShapePicture {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &ShapePicture{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建形状图片并按UID名称
// name:
func X创建形状图片并按UID名称(name string) *ShapePicture {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &ShapePicture{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状图片_置图片, 设置图片.
//
// hImage: 图片句柄.

// ff:置图片
// hImage:图片句柄
func (s *ShapePicture) X置图片(图片句柄 int) int {
	return 炫彩基类.X形状图片_置图片(s.Handle, 图片句柄)
}

// 形状图片_取图片, 获取图片句柄.

// ff:取图片
func (s *ShapePicture) X取图片() int {
	return 炫彩基类.X形状图片_取图片(s.Handle)
}
