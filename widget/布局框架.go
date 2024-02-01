package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// LayoutFrame 布局框架.
type LayoutFrame struct {
	ScrollView
}

// I布局框架_创建 I布局框架_创建.
//	@param x 元素x坐标.
//	@param y 元素y坐标.
//	@param cx 宽度.
//	@param cy 高度.
//	@param hParent 父为窗口句柄或元素句柄.
//	@return *LayoutFrame
//
func I布局框架_创建(x int, y int, cx int, cy int, hParent int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(xc.XLayoutFrame_Create(x, y, cx, cy, hParent))
	return p
}

// I布局框架_从句柄创建 I布局框架_从句柄创建对象.
//	@param handle
//	@return *LayoutFrame
//
func I布局框架_从句柄创建(handle int) *LayoutFrame {
	p := &LayoutFrame{}
	p.SetHandle(handle)
	return p
}

// I布局框架_从name创建 I布局框架_从name创建对象, 失败返回nil.
//	@param name
//	@return *LayoutFrame
//
func I布局框架_从name创建(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局框架_从UID创建 I布局框架_从UID创建对象, 失败返回nil.
//	@param nUID
//	@return *LayoutFrame
//
func I布局框架_从UID创建(nUID int) *LayoutFrame {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I布局框架_从UID名称创建 I布局框架_从UID名称创建对象, 失败返回nil.
//	@param name
//	@return *LayoutFrame
//
func I布局框架_从UID名称创建(name string) *LayoutFrame {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &LayoutFrame{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I显示布局边界 布局框架_显示布局边界.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutFrame) I显示布局边界(bEnable bool) int {
	return xc.XLayoutFrame_ShowLayoutFrame(l.I句柄, bEnable)
}

/*
LayoutBox-布局盒子
*/

// I启用水平 布局盒子_启用水平.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutFrame) I启用水平(bEnable bool) int {
	return xc.XLayoutBox_EnableHorizon(l.I句柄, bEnable)
}

// I启用自动换行 布局盒子_启用自动换行.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutFrame) I启用自动换行(bEnable bool) int {
	return xc.XLayoutBox_EnableAutoWrap(l.I句柄, bEnable)
}

// I启用溢出隐藏 布局盒子_启用溢出隐藏.
//	@param bEnable 是否启用.
//	@return int
//
func (l *LayoutFrame) I启用溢出隐藏(bEnable bool) int {
	return xc.XLayoutBox_EnableOverflowHide(l.I句柄, bEnable)
}

// I置水平对齐 布局盒子_置水平对齐.
//	@param nAlign 对齐方式: xcc.I常量_布局对齐_.
//	@return int
//
func (l *LayoutFrame) I置水平对齐(nAlign xcc.I常量_布局对齐_) int {
	return xc.XLayoutBox_SetAlignH(l.I句柄, nAlign)
}

// I置垂直对齐 布局盒子_置垂直对齐.
//	@param nAlign 对齐方式: xcc.I常量_布局对齐_.
//	@return int
//
func (l *LayoutFrame) I置垂直对齐(nAlign xcc.I常量_布局对齐_) int {
	return xc.XLayoutBox_SetAlignV(l.I句柄, nAlign)
}

// I置对齐基线 布局盒子_置对齐基线.
//	@param nAlign 对齐方式: xcc.I常量_布局轴对齐_.
//	@return int
//
func (l *LayoutFrame) I置对齐基线(nAlign xcc.I常量_布局轴对齐_) int {
	return xc.XLayoutBox_SetAlignBaseline(l.I句柄, nAlign)
}

// I置间距 布局盒子_置间距.
//	@param nSpace 项间距大小.
//	@return int
//
func (l *LayoutFrame) I置间距(nSpace int) int {
	return xc.XLayoutBox_SetSpace(l.I句柄, nSpace)
}

// I置行距 布局盒子_置行距.
//	@param nSpace 行间距大小.
//	@return int
//
func (l *LayoutFrame) I置行距(nSpace int) int {
	return xc.XLayoutBox_SetSpaceRow(l.I句柄, nSpace)
}
