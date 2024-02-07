package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 布局盒子_启用水平.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func X布局盒子_启用水平(窗口或布局元素或布局框架句柄 int, 是否启用 bool) int {
	r, _, _ := xLayoutBox_EnableHorizon.Call(uintptr(窗口或布局元素或布局框架句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 布局盒子_启用自动换行.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func X布局盒子_启用自动换行(窗口或布局元素或布局框架句柄 int, 是否启用 bool) int {
	r, _, _ := xLayoutBox_EnableAutoWrap.Call(uintptr(窗口或布局元素或布局框架句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 布局盒子_启用溢出隐藏.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// bEnable: 是否启用.
func X布局盒子_启用溢出隐藏(窗口或布局元素或布局框架句柄 int, 是否启用 bool) int {
	r, _, _ := xLayoutBox_EnableOverflowHide.Call(uintptr(窗口或布局元素或布局框架句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// XLayoutBox_SetAlignH 布局盒子_置水平对齐.
//
//	@param hLayoutBox 窗口或布局元素或布局框架句柄.
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func X布局盒子_置水平对齐(窗口或布局元素或布局框架句柄 int, 对齐方式 炫彩常量类.Layout_Align_) int {
	r, _, _ := xLayoutBox_SetAlignH.Call(uintptr(窗口或布局元素或布局框架句柄), uintptr(对齐方式))
	return int(r)
}

// XLayoutBox_SetAlignV 布局盒子_置垂直对齐.
//
//	@param hLayoutBox 窗口或布局元素或布局框架句柄.
//	@param nAlign 对齐方式: xcc.Layout_Align_.
//	@return int
func X布局盒子_置垂直对齐(窗口或布局元素或布局框架句柄 int, 对齐方式 炫彩常量类.Layout_Align_) int {
	r, _, _ := xLayoutBox_SetAlignV.Call(uintptr(窗口或布局元素或布局框架句柄), uintptr(对齐方式))
	return int(r)
}

// 布局盒子_置对齐基线.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nAlign: 对齐方式: xcc.Layout_Align_Axis_.
func X布局盒子_置对齐基线(窗口或布局元素或布局框架句柄 int, 对齐方式 炫彩常量类.Layout_Align_Axis_) int {
	r, _, _ := xLayoutBox_SetAlignBaseline.Call(uintptr(窗口或布局元素或布局框架句柄), uintptr(对齐方式))
	return int(r)
}

// 布局盒子_置间距.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nSpace: 项间距大小.
func X布局盒子_置间距(窗口或布局元素或布局框架句柄 int, 项间距大小 int) int {
	r, _, _ := xLayoutBox_SetSpace.Call(uintptr(窗口或布局元素或布局框架句柄), uintptr(项间距大小))
	return int(r)
}

// 布局盒子_置行距.
//
// hLayoutBox: 窗口或布局元素或布局框架句柄.
//
// nSpace: 行间距大小.
func X布局盒子_置行距(窗口或布局元素或布局框架句柄 int, 行间距大小 int) int {
	r, _, _ := xLayoutBox_SetSpaceRow.Call(uintptr(窗口或布局元素或布局框架句柄), uintptr(行间距大小))
	return int(r)
}
