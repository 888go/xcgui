package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
	
	"github.com/888go/xcgui/xcc"
)

// 窗口组件_判断显示, 判断UI对象是否显示.
//
// hXCGUI: 对象句柄.
func X窗口组件_判断显示(对象句柄 int) bool {
	r, _, _ := xWidget_IsShow.Call(uintptr(对象句柄))
	return r != 0
}

// 窗口组件_显示.
//
// hXCGUI: 对象句柄.
//
// bShow: 是否显示.
func X窗口组件_显示(对象句柄 int, 是否显示 bool) int {
	r, _, _ := xWidget_Show.Call(uintptr(对象句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 窗口组件_启用布局控制.
//
// hXCGUI:.
//
// bEnable:.
func X窗口组件_启用布局控制(对象句柄 int, 是否启用 bool) int {
	r, _, _ := xWidget_EnableLayoutControl.Call(uintptr(对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口组件_是否布局控制.
//
// hXCGUI:.
func X窗口组件_是否布局控制(对象句柄 int) bool {
	r, _, _ := xWidget_IsLayoutControl.Call(uintptr(对象句柄))
	return r != 0
}

// 窗口组件_取父元素.
//
// hXCGUI: 对象句柄.
func X窗口组件_取父元素(对象句柄 int) int {
	r, _, _ := xWidget_GetParentEle.Call(uintptr(对象句柄))
	return int(r)
}

// 窗口组件_取父对象, 获取父对象, 父可能是元素或窗口, 通过此函数可以检查是否有父.
//
// hXCGUI: 对象句柄.
func X窗口组件_取父对象(对象句柄 int) int {
	r, _, _ := xWidget_GetParent.Call(uintptr(对象句柄))
	return int(r)
}

// 窗口组件_取HWND.
//
// hXCGUI: 对象句柄.
func X窗口组件_取HWND(对象句柄 int) uintptr {
	r, _, _ := xWidget_GetHWND.Call(uintptr(对象句柄))
	return r
}

// 窗口组件_取HWINDOW.
//
// hXCGUI: 对象句柄.
func X窗口组件_取HWINDOW(对象句柄 int) int {
	r, _, _ := xWidget_GetHWINDOW.Call(uintptr(对象句柄))
	return int(r)
}

// 窗口组件_布局项_启用换行, 强制换行.
//
// hXCGUI: UI对象句柄.
//
// bWrap: 是否换行.
func X窗口组件_布局项_启用换行(对象句柄 int, 是否换行 bool) int {
	r, _, _ := xWidget_LayoutItem_EnableWrap.Call(uintptr(对象句柄), 炫彩工具类.BoolPtr(是否换行))
	return int(r)
}

// 窗口组件_布局项_启用交换, 根据水平垂直布局变换, 交换属性(宽度,高度,最小宽度,最小高度).
//
// hXCGUI: UI对象句柄.
//
// bEnable: 是否启用.
func X窗口组件_布局项_启用交换(对象句柄 int, 是否启用 bool) int {
	r, _, _ := xWidget_LayoutItem_EnableSwap.Call(uintptr(对象句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 窗口组件_布局项_启用浮动, 向反方向对齐.
//
// hXCGUI: UI对象句柄.
//
// bFloat: 是否浮动.
func X窗口组件_布局项_启用浮动(对象句柄 int, 是否浮动 bool) int {
	r, _, _ := xWidget_LayoutItem_EnableFloat.Call(uintptr(对象句柄), 炫彩工具类.BoolPtr(是否浮动))
	return int(r)
}

// 窗口组件_布局项_置宽度.
//
// hXCGUI: UI对象句柄.
//
// nType: 类型: xcc.Layout_Size_.
//
// nWidth: 宽度.
func X窗口组件_布局项_置宽度(对象句柄 int, 类型 炫彩常量类.Layout_Size_, 宽度 int32) int {
	r, _, _ := xWidget_LayoutItem_SetWidth.Call(uintptr(对象句柄), uintptr(类型), uintptr(宽度))
	return int(r)
}

// 窗口组件_布局项_置高度.
//
// hXCGUI: UI对象句柄.
//
// nType: 类型: xcc.Layout_Size_.
//
// nHeight: 高度.
func X窗口组件_布局项_置高度(对象句柄 int, 类型 炫彩常量类.Layout_Size_, 高度 int32) int {
	r, _, _ := xWidget_LayoutItem_SetHeight.Call(uintptr(对象句柄), uintptr(类型), uintptr(高度))
	return int(r)
}

// 窗口组件_布局项_取宽度.
//
// hXCGUI: UI对象句柄.
//
// pType: 返回类型.
//
// pWidth: 返回宽度.
func X窗口组件_布局项_取宽度(对象句柄 int, 返回类型 *炫彩常量类.Layout_Size_, 返回宽度 *int32) int {
	r, _, _ := xWidget_LayoutItem_GetWidth.Call(uintptr(对象句柄), uintptr(unsafe.Pointer(返回类型)), uintptr(unsafe.Pointer(返回宽度)))
	return int(r)
}

// 窗口组件_布局项_取高度.
//
// hXCGUI: UI对象句柄.
//
// pType: 返回类型.
//
// pHeight: 返回高度.
func X窗口组件_布局项_取高度(对象句柄 int, 返回类型 *炫彩常量类.Layout_Size_, 返回高度 *int32) int {
	r, _, _ := xWidget_LayoutItem_GetHeight.Call(uintptr(对象句柄), uintptr(unsafe.Pointer(返回类型)), uintptr(unsafe.Pointer(返回高度)))
	return int(r)
}

// 窗口组件_布局项_置对齐, 根据水平垂直轴变化对齐.
//
// hXCGUI: UI对象句柄.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func X窗口组件_布局项_置对齐(对象句柄 int, 对齐方式 炫彩常量类.Layout_Align_Axis_) int {
	r, _, _ := xWidget_LayoutItem_SetAlign.Call(uintptr(对象句柄), uintptr(对齐方式))
	return int(r)
}

// 窗口组件_布局项_置外间距.
//
// hXCGUI: UI对象句柄.
func X窗口组件_布局项_置外间距(对象句柄, left, top, right, bottom int) int {
	r, _, _ := xWidget_LayoutItem_SetMargin.Call(uintptr(对象句柄), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(r)
}

// 窗口组件_布局项_置对齐.
//
// hXCGUI: UI对象句柄.
//
// pMargin: 接收返回.
func X窗口组件_布局项_取外间距(对象句柄 int, 接收返回 *RECT) int {
	r, _, _ := xWidget_LayoutItem_GetMargin.Call(uintptr(对象句柄), uintptr(unsafe.Pointer(接收返回)))
	return int(r)
}

// 窗口组件_布局项_置最小大小, 限制大小仅针对缩放有效(自动, 填充父, 比例, 百分比).
//
// hXCGUI: UI对象句柄.
//
// width: 最小宽度.
//
// height: 最小高度.
func X窗口组件_布局项_置最小大小(对象句柄, 最小宽度, 最小高度 int) int {
	r, _, _ := xWidget_LayoutItem_SetMinSize.Call(uintptr(对象句柄), uintptr(最小宽度), uintptr(最小高度))
	return int(r)
}

// 窗口组件_布局项_置位置, 相对位置, 值大于等于0有效.
//
// hXCGUI: UI对象句柄.
//
// left: 左边距离.
//
// top: 上边距离.
//
// right: 右边距离.
//
// bottom: 下边距离.
func X窗口组件_布局项_置位置(对象句柄, 左边距离, 上边距离, 右边距离, 下边距离 int) int {
	r, _, _ := xWidget_LayoutItem_SetPosition.Call(uintptr(对象句柄), uintptr(左边距离), uintptr(上边距离), uintptr(右边距离), uintptr(下边距离))
	return int(r)
}

// 窗口组件_置ID, 设置元素ID.
//
// hXCGUI: UI对象句柄.
//
// nID: ID值.
func X窗口组件_置ID(对象句柄 int, ID值 int32) int {
	r, _, _ := xWidget_SetID.Call(uintptr(对象句柄), uintptr(ID值))
	return int(r)
}

// 窗口组件_置UID, 设置元素UID, 全局唯一标识符.
//
// hXCGUI: UI对象句柄.
//
// nUID: UID值.
func X窗口组件_置UID(对象句柄 int, UID值 int32) int {
	r, _, _ := xWidget_SetUID.Call(uintptr(对象句柄), uintptr(UID值))
	return int(r)
}

// 窗口组件_置名称 设置元素name.
//
// hXCGUI: UI对象句柄.
//
// pName: name值.
func X窗口组件_置名称(对象句柄 int, 名称 string) int {
	r, _, _ := xWidget_SetName.Call(uintptr(对象句柄), 炫彩工具类.StrPtr(名称))
	return int(r)
}

// 窗口组件_取ID, 获取元素ID.
//
// hXCGUI: UI对象句柄.
func X窗口组件_取ID(对象句柄 int) int32 {
	r, _, _ := xWidget_GetID.Call(uintptr(对象句柄))
	return int32(r)
}

// 窗口组件_取UID, 获取元素UID, 全局唯一标识符.
//
// hXCGUI: UI对象句柄.
func X窗口组件_取UID(对象句柄 int) int32 {
	r, _, _ := xWidget_GetUID.Call(uintptr(对象句柄))
	return int32(r)
}

// 窗口组件_取名称 获取元素name.
//
// hXCGUI: UI对象句柄.
func X窗口组件_取名称(对象句柄 int) string {
	r, _, _ := xWidget_GetName.Call(uintptr(对象句柄))
	return 炫彩工具类.UintPtrToString(r)
}
