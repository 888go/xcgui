package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// LayoutFrame 布局框架.
type LayoutFrame struct {
	ScrollView
}

// NewLayoutFrame 布局框架_创建.
//
//	@param x 元素x坐标.
//	@param y 元素y坐标.
//	@param cx 宽度.
//	@param cy 高度.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutFrame
func X创建布局框架(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *LayoutFrame {
	p := &LayoutFrame{}
	p.X设置句柄(炫彩基类.X布局框架_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// NewLayoutFrameByHandle 从句柄创建对象.
//
//	@param handle
//	@return *LayoutFrame
func X创建布局框架并按句柄(句柄 int) *LayoutFrame {
	p := &LayoutFrame{}
	p.X设置句柄(句柄)
	return p
}

// NewLayoutFrameByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutFrame
func X创建布局框架并按名称(名称 string) *LayoutFrame {
	handle := 炫彩基类.X取对象从名称(名称)
	if handle > 0 {
		p := &LayoutFrame{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutFrameByUID 从UID创建对象, 失败返回nil.
//
//	@param nUID
//	@return *LayoutFrame
func X创建布局框架并按UID(nUID int) *LayoutFrame {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &LayoutFrame{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// NewLayoutFrameByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutFrame
func X创建布局框架并按UID名称(名称 string) *LayoutFrame {
	handle := 炫彩基类.X取对象从UID名称(名称)
	if handle > 0 {
		p := &LayoutFrame{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// ShowLayoutFrame 布局框架_显示布局边界.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) X显示布局边界(是否启用 bool) int {
	return 炫彩基类.X布局框架_显示布局边界(l.Handle, 是否启用)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) X启用水平(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用水平(l.Handle, 是否启用)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) X启用自动换行(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用自动换行(l.Handle, 是否启用)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int
func (l *LayoutFrame) X启用溢出隐藏(是否启用 bool) int {
	return 炫彩基类.X布局盒子_启用溢出隐藏(l.Handle, 是否启用)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (l *LayoutFrame) X置水平对齐(对齐方式 炫彩常量类.Layout_Align_) int {
	return 炫彩基类.X布局盒子_置水平对齐(l.Handle, 对齐方式)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func (l *LayoutFrame) X置垂直对齐(对齐方式 炫彩常量类.Layout_Align_) int {
	return 炫彩基类.X布局盒子_置垂直对齐(l.Handle, 对齐方式)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
func (l *LayoutFrame) X置对齐基线(对齐方式 炫彩常量类.Layout_Align_Axis_) int {
	return 炫彩基类.X布局盒子_置对齐基线(l.Handle, 对齐方式)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int
func (l *LayoutFrame) X置间距(项间距大小 int) int {
	return 炫彩基类.X布局盒子_置间距(l.Handle, 项间距大小)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int
func (l *LayoutFrame) X置行距(行间距大小 int) int {
	return 炫彩基类.X布局盒子_置行距(l.Handle, 行间距大小)
}
