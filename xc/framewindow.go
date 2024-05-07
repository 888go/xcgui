package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"

	"github.com/888go/xcgui/xcc"
)

// 框架窗口_创建, 返回GUI库窗口资源句柄.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口标题.
//
// hWndParent: 父窗口真实句柄.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func X框架窗口_创建(左上角x坐标, 左上角y坐标, 宽度, 高度 int32, 标题 string, 父窗口真实句柄 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xFrameWnd_Create.Call(uintptr(左上角x坐标), uintptr(左上角y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(标题), 父窗口真实句柄, uintptr(GUI库窗口样式))
	return int(r)
}

// 框架窗口_创建扩展, 返回GUI库窗口资源句柄.
//
// dwExStyle: 窗口扩展样式.
//
// dwStyle: 窗口样式.
//
// lpClassName: 窗口类名.
//
// x: 窗口左上角x坐标.
//
// y: 窗口左上角y坐标.
//
// cx: 窗口宽度.
//
// cy: 窗口高度.
//
// pTitle: 窗口名.
//
// hWndParent: 父窗口.
//
// XCStyle: GUI库窗口样式: Window_Style_.
func X框架窗口_创建EX(扩展样式 int, 样式 int, 类名 string, 左上角x坐标, 左上角y坐标, 宽度, 高度 int32, 窗口名 string, 父窗口 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xFrameWnd_CreateEx.Call(uintptr(扩展样式), uintptr(样式), 炫彩工具类.StrPtr(类名), uintptr(左上角x坐标), uintptr(左上角y坐标), uintptr(宽度), uintptr(高度), 炫彩工具类.StrPtr(窗口名), 父窗口, uintptr(GUI库窗口样式))
	return int(r)
}

// 框架窗口_取布局区域坐标, 用来布局窗格的区域坐标, 不包含码头.
//
// hWindow: 窗口句柄.
//
// pRect: 返回坐标.
func X框架窗口_取布局区域坐标(窗口句柄 int, 返回坐标 *RECT) int {
	r, _, _ := xFrameWnd_GetLayoutAreaRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(返回坐标)))
	return int(r)
}

// 框架窗口_置视图, 设置窗格组TabBar高度.
//
// hWindow: 窗口句柄.
//
// hEle: 元素句柄.
func X框架窗口_置视图(窗口句柄 int, 元素句柄 int) int {
	r, _, _ := xFrameWnd_SetView.Call(uintptr(窗口句柄), uintptr(元素句柄))
	return int(r)
}

// 框架窗口_置窗格分隔条颜色.
//
// hWindow: 窗口句柄.
//
// color: ABGR 颜色值.
func X框架窗口_置窗格分隔条颜色(窗口句柄 int, ABGR颜色值 int) int {
	r, _, _ := xFrameWnd_SetPaneSplitBarColor.Call(uintptr(窗口句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 框架窗口_置TabBar条高度, 设置窗格组TabBar高度.
//
// hWindow: 窗口句柄.
//
// nHeight: 高度.
func X框架窗口_置TabBar条高度(窗口句柄 int, 高度 int32) int {
	r, _, _ := xFrameWnd_SetTabBarHeight.Call(uintptr(窗口句柄), uintptr(高度))
	return int(r)
}

// 框架窗口_保存布局到文件, 保存布局信息到文件.
//
// hWindow: 窗口句柄.
//
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func X框架窗口_保存布局到文件(窗口句柄 int, 文件名 string) bool {
	r, _, _ := xFrameWnd_SaveLayoutToFile.Call(uintptr(窗口句柄), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 框架窗口_加载布局信息文件, 加载布局信息文件.
//
// hWindow: 窗口句柄.
//
// aPaneList: 窗格句柄切片.
//
// nPaneCount: 窗格数量.
//
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func X框架窗口_加载布局信息文件(窗口句柄 int, 窗格句柄切片 []int, 窗格数量 int32, 文件名 string) bool {
	r, _, _ := xFrameWnd_LoadLayoutFile.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(&窗格句柄切片[0])), uintptr(窗格数量), 炫彩工具类.StrPtr(文件名))
	return r != 0
}

// 框架窗口_添加窗格, 添加窗格到框架窗口.
//
// hWindow: 窗口句柄.
//
// hPaneDest: 目标窗格.
//
// hPaneNew: 当前窗格.
//
// align: 对齐方式, Pane_Align_.
func X框架窗口_添加窗格(窗口句柄 int, 目标窗格 int, 当前窗格 int, 对齐方式 炫彩常量类.Pane_Align_) bool {
	r, _, _ := xFrameWnd_AddPane.Call(uintptr(窗口句柄), uintptr(目标窗格), uintptr(当前窗格), uintptr(对齐方式))
	return r != 0
}

// 框架窗口_合并窗格.
//
// hWindow: 窗口句柄.
//
// hPaneDest: 目标窗格.
//
// hPaneNew: 当前窗格.
func X框架窗口_合并窗格(窗口句柄 int, 目标窗格 int, 当前窗格 int) bool {
	r, _, _ := xFrameWnd_MergePane.Call(uintptr(窗口句柄), uintptr(目标窗格), uintptr(当前窗格))
	return r != 0
}

// 框架窗口_附加窗口, 返回窗口资源句柄.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func X框架窗口_附加窗口(外部窗口句柄 uintptr, 炫彩窗口样式 int) int {
	r, _, _ := xFrameWnd_Attach.Call(外部窗口句柄, uintptr(炫彩窗口样式))
	return int(r)
}

// 框架窗口_取拖动浮动窗格停留位置标识, 返回 拖动窗格 所处框架窗口单元格标识: xcc.FrameWnd_Cell_Type_.
//
// hWindow: 窗口句柄.
func X框架窗口_取拖动浮动窗格停留位置标识(窗口句柄 int) 炫彩常量类.FrameWnd_Cell_Type_ {
	r, _, _ := xFrameWnd_GetDragFloatWndTopFlag.Call(uintptr(窗口句柄))
	return 炫彩常量类.FrameWnd_Cell_Type_(r)
}

// 框架窗口_取主视图坐标. 获取框架窗口主视图区域坐标.
//
// hWindow: 窗口句柄.
//
// pRect: 返回坐标.
func X框架窗口_取主视图坐标(窗口句柄 int, 返回坐标 *RECT) {
	xFrameWnd_GetViewRect.Call(uintptr(窗口句柄), uintptr(unsafe.Pointer(返回坐标)))
}

// 框架窗口_置窗格分隔条宽度.
//
// hWindow: 窗口句柄.
//
// nWidth: 宽度.
func X框架窗口_置窗格分隔条宽度(窗口句柄 int, 宽度 int32) {
	xFrameWnd_SetPaneSplitBarWidth.Call(uintptr(窗口句柄), uintptr(宽度))
}

// 框架窗口_取窗格分隔条宽度.
//
// hWindow: 窗口句柄.
func X框架窗口_取窗格分隔条宽度(窗口句柄 int) int32 {
	r, _, _ := xFrameWnd_GetPaneSplitBarWidth.Call(uintptr(窗口句柄))
	return int32(r)
}
