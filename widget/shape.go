package 炫彩组件类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
)

// Shape 形状对象基类.
type Shape struct {
	炫彩对象基类.Widget
}

// 从句柄创建对象.

// ff:创建形状对象基类并按句柄
// handle:
func X创建形状对象基类并按句柄(handle int) *Shape {
	p := &Shape{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建形状对象基类并按名称
// name:
func X创建形状对象基类并按名称(name string) *Shape {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Shape{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.

// ff:创建形状对象基类并按UID
// nUID:
func X创建形状对象基类并按UID(nUID int) *Shape {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Shape{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.

// ff:创建形状对象基类并按UID名称
// name:
func X创建形状对象基类并按UID名称(name string) *Shape {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Shape{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 形状_移除, 从父UI元素或窗口,和父布局对象中移除.

// ff:移除
func (s *Shape) X移除() int {
	return 炫彩基类.X形状_移除(s.Handle)
}

// 形状_取Z序, 获取形状对象Z序, 成功返回索引值, 否则返回 XC_ID_ERROR.

// ff:取Z序
func (s *Shape) X取Z序() int {
	return 炫彩基类.X形状_取Z序(s.Handle)
}

// 形状_重绘, 重绘形状对象.

// ff:重绘
func (s *Shape) X重绘() int {
	return 炫彩基类.X形状_重绘(s.Handle)
}

// 形状_取宽度, 获取内容宽度.

// ff:取宽度
func (s *Shape) X取宽度() int32 {
	return 炫彩基类.X形状_取宽度(s.Handle)
}

// 形状_取高度, 获取内容高度.

// ff:取高度
func (s *Shape) X取高度() int32 {
	return 炫彩基类.X形状_取高度(s.Handle)
}

// 形状_移动位置.
//
// x: x坐标.
//
// y: y坐标.

// ff:移动位置
// y:y坐标
// x:x坐标
func (s *Shape) X移动位置(x坐标, y坐标 int32) int {
	return 炫彩基类.X形状_移动位置(s.Handle, x坐标, y坐标)
}

// 形状_取坐标.
//
// pRect: 接收返回坐标.

// ff:取坐标
// pRect:接收返回坐标
func (s *Shape) X取坐标(接收返回坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X形状_取坐标(s.Handle, 接收返回坐标)
}

// 形状_置坐标.
//
// pRect: 坐标.

// ff:置坐标
// pRect:坐标
func (s *Shape) X置坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X形状_置坐标(s.Handle, 坐标)
}

// 形状_置逻辑坐标, 设置元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.

// ff:置逻辑坐标
// bRedraw:
// pRect:坐标
func (s *Shape) X置逻辑坐标(坐标 *炫彩基类.RECT, bRedraw bool) bool {
	return 炫彩基类.X形状_置逻辑坐标(s.Handle, 坐标, bRedraw)
}

// 形状_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.

// ff:取逻辑坐标
// pRect:坐标
func (s *Shape) X取逻辑坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X形状_取逻辑坐标(s.Handle, 坐标)
}

// 形状_取基于窗口客户区坐标, 基于窗口客户区坐标.
//
// pRect: 坐标.

// ff:取基于窗口客户区坐标
// pRect:坐标
func (s *Shape) X取基于窗口客户区坐标(坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X形状_取基于窗口客户区坐标(s.Handle, 坐标)
}

// 形状_取内容大小 ,仅计算有效内容, 填充父, 权重依赖父级所以无法计算.
//
// pSize: 接收返回内容大小值.

// ff:取内容大小
// pSize:接收返回大小值
func (s *Shape) X取内容大小(接收返回大小值 *炫彩基类.SIZE) int {
	return 炫彩基类.X形状_取内容大小(s.Handle, 接收返回大小值)
}

// 形状_显示布局边界, 是否显示布局边界.
//
// bShow: 是否显示.

// ff:显示布局边界
// bShow:是否显示
func (s *Shape) X显示布局边界(是否显示 bool) int {
	return 炫彩基类.X形状_显示布局边界(s.Handle, 是否显示)
}

// 形状_调整布局.

// ff:调整布局
func (s *Shape) X调整布局() int {
	return 炫彩基类.X形状_调整布局(s.Handle)
}

// 形状_销毁, 销毁形状对象.

// ff:销毁
func (s *Shape) X销毁() int {
	return 炫彩基类.X形状_销毁(s.Handle)
}

// 形状_取位置.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.

// ff:取位置
// pOutY:返回Y坐标
// pOutX:返回X坐标
func (s *Shape) X取位置(返回X坐标, 返回Y坐标 *int32) int {
	return 炫彩基类.X形状_取位置(s.Handle, 返回X坐标, 返回Y坐标)
}

// 形状_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.

// ff:置大小
// nHeight:高度
// nWidth:宽度
func (s *Shape) X置大小(宽度, 高度 int32) int {
	return 炫彩基类.X形状_置大小(s.Handle, 宽度, 高度)
}

// 形状_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.

// ff:取大小
// pOutHeight:返回高度
// pOutWidth:返回宽度
func (s *Shape) X取大小(返回宽度, 返回高度 *int32) int {
	return 炫彩基类.X形状_取大小(s.Handle, 返回宽度, 返回高度)
}

// 形状_置透明度.
//
// alpha: 透明度.

// ff:置透明度
// alpha:透明度
func (s *Shape) X置透明度(透明度 uint8) int {
	return 炫彩基类.X形状_置透明度(s.Handle, 透明度)
}

// 形状_取透明度, 返回透明度.

// ff:取透明度
func (s *Shape) X取透明度() int {
	return 炫彩基类.X形状_取透明度(s.Handle)
}
