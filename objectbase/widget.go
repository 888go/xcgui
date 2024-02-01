package objectbase

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 窗口组件.
type Widget struct {
	UI
}

// 窗口组件_ I判断显示, 判断UI对象是否显示.
func (w *Widget) I判断显示() bool {
	return xc.XWidget_IsShow(w.I句柄)
}

// 窗口组件_ I显示.
//
// bShow: 是否显示.
func (w *Widget) I显示(bShow bool) int {
	return xc.XWidget_Show(w.I句柄, bShow)
}

// 窗口组件_ I启用布局控制.
//
// bEnable:.
func (w *Widget) I启用布局控制(bEnable bool) int {
	return xc.XWidget_EnableLayoutControl(w.I句柄, bEnable)
}

// 窗口组件_ I是否布局控制.
func (w *Widget) I是否布局控制() bool {
	return xc.XWidget_IsLayoutControl(w.I句柄)
}

// 窗口组件_ I取父元素.
func (w *Widget) I取父元素() int {
	return xc.XWidget_GetParentEle(w.I句柄)
}

// 窗口组件_ I取父对象, 获取父对象, 父可能是元素或窗口, 通过此函数可以检查是否有父.
func (w *Widget) I取父对象() int {
	return xc.XWidget_GetParent(w.I句柄)
}

// 窗口组件_ I取HWND.
func (w *Widget) I取HWND() int {
	return xc.XWidget_GetHWND(w.I句柄)
}

// 窗口组件_ I取HWINDOW.
func (w *Widget) I取HWINDOW() int {
	return xc.XWidget_GetHWINDOW(w.I句柄)
}

// 窗口组件_ I布局项_启用换行, 强制换行.
//
// bWrap: 是否换行.
func (w *Widget) I布局项_启用换行(bWrap bool) int {
	return xc.XWidget_LayoutItem_EnableWrap(w.I句柄, bWrap)
}

// 窗口组件_ I布局项_启用交换, 根据水平垂直布局变换, 交换属性(宽度,高度,最小宽度,最小高度).
//
// bEnable: 是否启用.
func (w *Widget) I布局项_启用交换(bEnable bool) int {
	return xc.XWidget_LayoutItem_EnableSwap(w.I句柄, bEnable)
}

// 窗口组件_ I布局项_启用浮动, 向反方向对齐.
//
// bFloat: 是否浮动.
func (w *Widget) I布局项_启用浮动(bFloat bool) int {
	return xc.XWidget_LayoutItem_EnableFloat(w.I句柄, bFloat)
}

// 窗口组件_ I布局项_置宽度.
//
// nType: 类型: xcc.I常量_布局大小_.
//
// nWidth: 宽度.
func (w *Widget) I布局项_置宽度(nType xcc.I常量_布局大小_, nWidth int) int {
	return xc.XWidget_LayoutItem_SetWidth(w.I句柄, nType, nWidth)
}

// 窗口组件_ I布局项_置高度.
//
// nType: 类型: xcc.I常量_布局大小_.
//
// nHeight: 高度.
func (w *Widget) I布局项_置高度(nType xcc.I常量_布局大小_, nHeight int) int {
	return xc.XWidget_LayoutItem_SetHeight(w.I句柄, nType, nHeight)
}

// 窗口组件_ I布局项_取宽度.
//
// pType: 返回类型.
//
// pWidth: 返回宽度.
func (w *Widget) I布局项_取宽度(pType, pWidth *int) int {
	return xc.XWidget_LayoutItem_GetWidth(w.I句柄, pType, pWidth)
}

// 窗口组件_ I布局项_取高度.
//
// pType: 返回类型.
//
// pHeight: 返回高度.
func (w *Widget) I布局项_取高度(pType, pHeight *int) int {
	return xc.XWidget_LayoutItem_GetHeight(w.I句柄, pType, pHeight)
}

// 窗口组件_ I布局项_置对齐, 根据水平垂直轴变化对齐.
//
// nAlign: 对齐方式: xcc.I常量_布局轴对齐_.
func (w *Widget) I布局项_置对齐(nAlign xcc.I常量_布局轴对齐_) int {
	return xc.XWidget_LayoutItem_SetAlign(w.I句柄, nAlign)
}

// 窗口组件_ I布局项_置外间距.
func (w *Widget) I布局项_置外间距(left, top, right, bottom int) int {
	return xc.XWidget_LayoutItem_SetMargin(w.I句柄, left, top, right, bottom)
}

// 窗口组件_ I布局项_置对齐.
//
// pMargin: 接收返回.
func (w *Widget) I布局项_置对齐1(pMargin *xc.RECT) int {
	return xc.XWidget_LayoutItem_GetMargin(w.I句柄, pMargin)
}

// 窗口组件_ I布局项_置最小大小, 限制大小仅针对缩放有效(自动, 填充父, 比例, 百分比).
//
// width: 最小宽度.
//
// height: 最小高度.
func (w *Widget) I布局项_置最小大小(width, height int) int {
	return xc.XWidget_LayoutItem_SetMinSize(w.I句柄, width, height)
}

// 窗口组件_ I布局项_置位置, 相对位置, 值大于等于0有效.
//
// left: 左边距离.
//
// top: 上边距离.
//
// right: 右边距离.
//
// bottom: 下边距离.
func (w *Widget) I布局项_置位置(left, top, right, bottom int) int {
	return xc.XWidget_LayoutItem_SetPosition(w.I句柄, left, top, right, bottom)
}

// 窗口组件_ I置ID, 设置元素ID.
//
// nID: ID值.
func (w *Widget) I置ID(nID int) int {
	return xc.XWidget_SetID(w.I句柄, nID)
}

// 窗口组件_ I置UID, 设置元素UID, 全局唯一标识符.
//
// nUID: UID值.
func (w *Widget) I置UID(nUID int) int {
	return xc.XWidget_SetUID(w.I句柄, nUID)
}

// 窗口组件_ I置名称 设置元素name.
//
// pName: name值.
func (w *Widget) I置名称(pName string) int {
	return xc.XWidget_SetName(w.I句柄, pName)
}

// 窗口组件_ I取ID, 获取元素ID.
func (w *Widget) I取ID() int {
	return xc.XWidget_GetID(w.I句柄)
}

// 窗口组件_ I取UID, 获取元素UID, 全局唯一标识符.
func (w *Widget) I取UID() int {
	return xc.XWidget_GetUID(w.I句柄)
}

// 窗口组件_ I取名称 获取元素name.
func (w *Widget) I取名称() string {
	return xc.XWidget_GetName(w.I句柄)
}
