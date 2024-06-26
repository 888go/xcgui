package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
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
// ff:创建布局框架
// x:元素x坐标
// y:元素y坐标
// cx:宽度
// cy:高度
// hParent:父窗口句柄或元素句柄
func NewLayoutFrame(x int, y int, cx int, cy int, hParent int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(xc.XLayoutFrame_Create(x, y, cx, cy, hParent))
	return p
}

// NewLayoutFrameByHandle 从句柄创建对象.
//
//	@param handle
//	@return *LayoutFrame
// ff:创建布局框架并按句柄
// handle:句柄
func NewLayoutFrameByHandle(handle int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(handle)
	return p
}

// NewLayoutFrameByName 从name创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutFrame
// ff:创建布局框架并按名称
// name:名称
func NewLayoutFrameByName(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutFrameByUID 从UID创建对象, 失败返回nil.
//
//	@param nUID
//	@return *LayoutFrame
// ff:创建布局框架并按UID
// nUID:
func NewLayoutFrameByUID(nUID int) *LayoutFrame {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// NewLayoutFrameByUIDName 从UID名称创建对象, 失败返回nil.
//
//	@param name
//	@return *LayoutFrame
// ff:创建布局框架并按UID名称
// name:名称
func NewLayoutFrameByUIDName(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// ShowLayoutFrame 布局框架_显示布局边界.
//
//	@param bEnable 是否启用.
//	@return int
// ff:显示布局边界
// bEnable:是否启用
func (l *LayoutFrame) ShowLayoutFrame(bEnable bool) int {
	return xc.XLayoutFrame_ShowLayoutFrame(l.Handle, bEnable)
}

/*
LayoutBox-布局盒子
*/

// EnableHorizon 布局盒子_启用水平.
//
//	@param bEnable 是否启用.
//	@return int
// ff:启用水平
// bEnable:是否启用
func (l *LayoutFrame) EnableHorizon(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.Handle, bEnable)
}

// EnableAutoWrap 布局盒子_启用自动换行.
//
//	@param bEnable 是否启用.
//	@return int
// ff:启用自动换行
// bEnable:是否启用
func (l *LayoutFrame) EnableAutoWrap(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.Handle, bEnable)
}

// EnableOverflowHide 布局盒子_启用溢出隐藏.
//
//	@param bEnable 是否启用.
//	@return int
// ff:启用溢出隐藏
// bEnable:是否启用
func (l *LayoutFrame) EnableOverflowHide(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.Handle, bEnable)
}

// SetAlignH 布局盒子_置水平对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
// ff:置水平对齐
// nAlign:对齐方式
func (l *LayoutFrame) SetAlignH(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignH(l.Handle, nAlign)
}

// SetAlignV 布局盒子_置垂直对齐.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
// ff:置垂直对齐
// nAlign:对齐方式
func (l *LayoutFrame) SetAlignV(nAlign xcc.Layout_Align_) int {
	return xc.XLayoutBox_SetAlignV(l.Handle, nAlign)
}

// SetAlignBaseline 布局盒子_置对齐基线.
//
//	@param nAlign 对齐方式: xcc.Layout_Align_Axis_.
//	@return int
// ff:置对齐基线
// nAlign:对齐方式
func (l *LayoutFrame) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {
	return xc.XLayoutBox_SetAlignBaseline(l.Handle, nAlign)
}

// SetSpace 布局盒子_置间距.
//
//	@param nSpace 项间距大小.
//	@return int
// ff:置间距
// nSpace:项间距大小
func (l *LayoutFrame) SetSpace(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.Handle, nSpace)
}

// SetSpaceRow 布局盒子_置行距.
//
//	@param nSpace 行间距大小.
//	@return int
// ff:置行距
// nSpace:行间距大小
func (l *LayoutFrame) SetSpaceRow(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.Handle, nSpace)
}
