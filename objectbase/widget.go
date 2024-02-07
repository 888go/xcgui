package 炫彩对象基类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 窗口组件.
type Widget struct {
	UI
}

// 窗口组件_判断显示, 判断UI对象是否显示.
func (w *Widget) X判断显示() bool {
	return 炫彩基类.X窗口组件_判断显示(w.Handle)
}

// 窗口组件_显示.
//
// bShow: 是否显示.
func (w *Widget) X显示(是否显示 bool) int {
	return 炫彩基类.X窗口组件_显示(w.Handle, 是否显示)
}

// 窗口组件_启用布局控制.
//
// bEnable:.
func (w *Widget) X启用布局控制(是否启用 bool) int {
	return 炫彩基类.X窗口组件_启用布局控制(w.Handle, 是否启用)
}

// 窗口组件_是否布局控制.
func (w *Widget) X是否布局控制() bool {
	return 炫彩基类.X窗口组件_是否布局控制(w.Handle)
}

// 窗口组件_取父元素.
func (w *Widget) X取父元素() int {
	return 炫彩基类.X窗口组件_取父元素(w.Handle)
}

// 窗口组件_取父对象, 获取父对象, 父可能是元素或窗口, 通过此函数可以检查是否有父.
func (w *Widget) X取父对象() int {
	return 炫彩基类.X窗口组件_取父对象(w.Handle)
}

// 窗口组件_取HWND.
func (w *Widget) X取HWND() uintptr {
	return 炫彩基类.X窗口组件_取HWND(w.Handle)
}

// 窗口组件_取HWINDOW.
func (w *Widget) X取HWINDOW() int {
	return 炫彩基类.X窗口组件_取HWINDOW(w.Handle)
}

// 窗口组件_布局项_启用换行, 强制换行.
//
// bWrap: 是否换行.
func (w *Widget) X布局项启用换行(是否换行 bool) int {
	return 炫彩基类.X窗口组件_布局项_启用换行(w.Handle, 是否换行)
}

// 窗口组件_布局项_启用交换, 根据水平垂直布局变换, 交换属性(宽度,高度,最小宽度,最小高度).
//
// bEnable: 是否启用.
func (w *Widget) X布局项启用交换(是否启用 bool) int {
	return 炫彩基类.X窗口组件_布局项_启用交换(w.Handle, 是否启用)
}

// 窗口组件_布局项_启用浮动, 向反方向对齐.
//
// bFloat: 是否浮动.
func (w *Widget) X布局项启用浮动(是否浮动 bool) int {
	return 炫彩基类.X窗口组件_布局项_启用浮动(w.Handle, 是否浮动)
}

// 窗口组件_布局项_置宽度.
//
// nType: 类型: xcc.Layout_Size_.
//
// nWidth: 宽度.
func (w *Widget) X布局项置宽度(类型 炫彩常量类.Layout_Size_, 宽度 int32) int {
	return 炫彩基类.X窗口组件_布局项_置宽度(w.Handle, 类型, 宽度)
}

// 窗口组件_布局项_置高度.
//
// nType: 类型: xcc.Layout_Size_.
//
// nHeight: 高度.
func (w *Widget) X布局项置高度(类型 炫彩常量类.Layout_Size_, 高度 int32) int {
	return 炫彩基类.X窗口组件_布局项_置高度(w.Handle, 类型, 高度)
}

// 窗口组件_布局项_取宽度.
//
// pType: 返回类型.
//
// pWidth: 返回宽度.
func (w *Widget) X布局项取宽度(返回类型 *炫彩常量类.Layout_Size_, 返回宽度 *int32) int {
	return 炫彩基类.X窗口组件_布局项_取宽度(w.Handle, 返回类型, 返回宽度)
}

// 窗口组件_布局项_取高度.
//
// pType: 返回类型.
//
// pHeight: 返回高度.
func (w *Widget) X布局项取高度(返回类型 *炫彩常量类.Layout_Size_, 返回高度 *int32) int {
	return 炫彩基类.X窗口组件_布局项_取高度(w.Handle, 返回类型, 返回高度)
}

// 窗口组件_布局项_置对齐, 根据水平垂直轴变化对齐.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func (w *Widget) X布局项置对齐(对齐方式 炫彩常量类.Layout_Align_Axis_) int {
	return 炫彩基类.X窗口组件_布局项_置对齐(w.Handle, 对齐方式)
}

// 窗口组件_布局项_置外间距.
func (w *Widget) X布局项置外间距(left, top, right, bottom int) int {
	return 炫彩基类.X窗口组件_布局项_置外间距(w.Handle, left, top, right, bottom)
}

// 窗口组件_布局项_置对齐.
//
// pMargin: 接收返回.
func (w *Widget) X布局项取外间距(接收返回 *炫彩基类.RECT) int {
	return 炫彩基类.X窗口组件_布局项_取外间距(w.Handle, 接收返回)
}

// 窗口组件_布局项_置最小大小, 限制大小仅针对缩放有效(自动, 填充父, 比例, 百分比).
//
// width: 最小宽度.
//
// height: 最小高度.
func (w *Widget) X布局项置最小大小(最小宽度, 最小高度 int) int {
	return 炫彩基类.X窗口组件_布局项_置最小大小(w.Handle, 最小宽度, 最小高度)
}

// 窗口组件_布局项_置位置, 相对位置, 值大于等于0有效.
//
// left: 左边距离.
//
// top: 上边距离.
//
// right: 右边距离.
//
// bottom: 下边距离.
func (w *Widget) X布局项置位置(左边距离, 上边距离, 右边距离, 下边距离 int) int {
	return 炫彩基类.X窗口组件_布局项_置位置(w.Handle, 左边距离, 上边距离, 右边距离, 下边距离)
}

// 窗口组件_置ID, 设置元素ID.
//
// nID: ID值.
func (w *Widget) X置ID(ID值 int32) int {
	return 炫彩基类.X窗口组件_置ID(w.Handle, ID值)
}

// 窗口组件_置UID, 设置元素UID, 全局唯一标识符.
//
// nUID: UID值.
func (w *Widget) X置UID(UID值 int32) int {
	return 炫彩基类.X窗口组件_置UID(w.Handle, UID值)
}

// 窗口组件_置名称 设置元素name.
//
// pName: name值.
func (w *Widget) X置名称(name值 string) int {
	return 炫彩基类.X窗口组件_置名称(w.Handle, name值)
}

// 窗口组件_取ID, 获取元素ID.
func (w *Widget) X取ID() int32 {
	return 炫彩基类.X窗口组件_取ID(w.Handle)
}

// 窗口组件_取UID, 获取元素UID, 全局唯一标识符.
func (w *Widget) X取UID() int32 {
	return 炫彩基类.X窗口组件_取UID(w.Handle)
}

// 窗口组件_取名称 获取元素name.
func (w *Widget) X取名称() string {
	return 炫彩基类.X窗口组件_取名称(w.Handle)
}
