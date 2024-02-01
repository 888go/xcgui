package widget

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/xc"
)

// 形状对象基类.
type Shape struct {
	objectbase.Widget
}

// 从句柄创建对象.
func NewShapeByHandle(handle int) *Shape {
	p := &Shape{}
	p.SetHandle(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func NewShapeByName(name string) *Shape {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Shape{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func NewShapeByUID(nUID int) *Shape {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Shape{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func NewShapeByUIDName(name string) *Shape {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Shape{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 形状_移除, 从父UI元素或窗口,和父布局对象中移除.
func (s *Shape) I移除() int {
	return xc.XShape_RemoveShape(s.I句柄)
}

// 形状_取Z序, 获取形状对象Z序, 成功返回索引值, 否则返回 XC_ID_ERROR.
func (s *Shape) I取Z序() int {
	return xc.XShape_GetZOrder(s.I句柄)
}

// 形状_重绘, 重绘形状对象.
func (s *Shape) I重绘() int {
	return xc.XShape_Redraw(s.I句柄)
}

// 形状_取宽度, 获取内容宽度.
func (s *Shape) I取宽度() int {
	return xc.XShape_GetWidth(s.I句柄)
}

// 形状_取高度, 获取内容高度.
func (s *Shape) I取高度() int {
	return xc.XShape_GetHeight(s.I句柄)
}

// 形状_移动位置.
//
// x: x坐标.
//
// y: y坐标.
func (s *Shape) I移动位置(x int, y int) int {
	return xc.XShape_SetPosition(s.I句柄, x, y)
}

// 形状_取坐标.
//
// pRect: 接收返回坐标.
func (s *Shape) I取坐标(pRect *xc.RECT) int {
	return xc.XShape_GetRect(s.I句柄, pRect)
}

// 形状_置坐标.
//
// pRect: 坐标.
func (s *Shape) I置坐标(pRect *xc.RECT) int {
	return xc.XShape_SetRect(s.I句柄, pRect)
}

// 形状_置逻辑坐标, 设置元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
//
// bRedraw: 是否重绘.
func (s *Shape) I置逻辑坐标(pRect *xc.RECT, bRedraw bool) bool {
	return xc.XShape_SetRectLogic(s.I句柄, pRect, bRedraw)
}

// 形状_取逻辑坐标, 获取元素坐标, 逻辑坐标, 包含滚动视图偏移.
//
// pRect: 坐标.
func (s *Shape) I取逻辑坐标(pRect *xc.RECT) int {
	return xc.XShape_GetRectLogic(s.I句柄, pRect)
}

// 形状_取基于窗口客户区坐标, 基于窗口客户区坐标.
//
// pRect: 坐标.
func (s *Shape) I取基于窗口客户区坐标(pRect *xc.RECT) int {
	return xc.XShape_GetWndClientRect(s.I句柄, pRect)
}

// 形状_取内容大小 ,仅计算有效内容, 填充父, 权重依赖父级所以无法计算.
//
// pSize: 接收返回内容大小值.
func (s *Shape) I取内容大小(pSize *xc.SIZE) int {
	return xc.XShape_GetContentSize(s.I句柄, pSize)
}

// 形状_显示布局边界, 是否显示布局边界.
//
// bShow: 是否显示.
func (s *Shape) I显示布局边界(bShow bool) int {
	return xc.XShape_ShowLayout(s.I句柄, bShow)
}

// 形状_调整布局.
func (s *Shape) I调整布局() int {
	return xc.XShape_AdjustLayout(s.I句柄)
}

// 形状_销毁, 销毁形状对象.
func (s *Shape) I销毁() int {
	return xc.XShape_Destroy(s.I句柄)
}

// 形状_取位置.
//
// pOutX: 返回X坐标.
//
// pOutY: 返回Y坐标.
func (s *Shape) I取位置(pOutX *int, pOutY *int) int {
	return xc.XShape_GetPosition(s.I句柄, pOutX, pOutY)
}

// 形状_置大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (s *Shape) I置大小(nWidth int, nHeight int) int {
	return xc.XShape_SetSize(s.I句柄, nWidth, nHeight)
}

// 形状_取大小.
//
// pOutWidth: 返回宽度.
//
// pOutHeight: 返回高度.
func (s *Shape) I取大小(pOutWidth *int, pOutHeight *int) int {
	return xc.XShape_GetSize(s.I句柄, pOutWidth, pOutHeight)
}

// 形状_置透明度.
//
// alpha: 透明度.
func (s *Shape) I置透明度(alpha uint8) int {
	return xc.XShape_SetAlpha(s.I句柄, alpha)
}

// 形状_取透明度, 返回透明度.
func (s *Shape) I取透明度() int {
	return xc.XShape_GetAlpha(s.I句柄)
}
