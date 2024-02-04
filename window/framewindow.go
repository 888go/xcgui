package window

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// FrameWindow 框架窗口.
type FrameWindow struct {
	windowBase
}

// 创建框架窗口
// // x: 窗口左上角x坐标.
// // y: 窗口左上角y坐标.
// // cx: 窗口宽度.
// // cy: 窗口高度.
// // pTitle: 窗口标题.
// // hWndParent: 父窗口真实句柄.
// // XCStyle: GUI库窗口样式: Window_Style_.
func NewFrameWindow(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *FrameWindow {
	p := &FrameWindow{}
	p.SetHandle(xc.XFrameWnd_Create(x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// 创建框架窗口EX
// // dwExStyle: 窗口扩展样式.
// // dwStyle: 窗口样式.
// // lpClassName: 窗口类名.
// // x: 窗口左上角x坐标.
// // y: 窗口左上角y坐标.
// // cx: 窗口宽度.
// // cy: 窗口高度.
// // pTitle: 窗口名.
// // hWndParent: 父窗口.
// // XCStyle: GUI库窗口样式: Window_Style_.
func NewFrameWindowEx(dwExStyle int, dwStyle int, lpClassName string, x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) *FrameWindow {
	p := &FrameWindow{}
	p.SetHandle(xc.XFrameWnd_CreateEx(dwExStyle, dwStyle, lpClassName, x, y, cx, cy, pTitle, hWndParent, XCStyle))
	return p
}

// 创建框架窗口并按句柄
func NewFrameWindowByHandle(handle int) *FrameWindow {
	p := &FrameWindow{}
	p.SetHandle(handle)
	return p
}

// 创建框架窗口并按名称, 失败返回nil.
func NewFrameWindowByName(name string) *FrameWindow {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &FrameWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建框架窗口并按UID  失败返回nil.
func NewFrameWindowByUID(nUID int) *FrameWindow {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &FrameWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建框架窗口并按UID名称 失败返回nil.
func NewFrameWindowByUIDName(name string) *FrameWindow {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &FrameWindow{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 取布局区域坐标, 用来布局窗格的区域坐标, 不包含码头.
// // pRect: 返回坐标.
func (fw *FrameWindow) GetLayoutAreaRect(pRect *xc.RECT) int {
	return xc.XFrameWnd_GetLayoutAreaRect(fw.Handle, pRect)
}

// 置视图, 设置窗格组TabBar高度.
// // hEle: 元素句柄.
func (fw *FrameWindow) SetView(hEle int) int {
	return xc.XFrameWnd_SetView(fw.Handle, hEle)
}

// 置窗格分隔条颜色.
// // color: ABGR 颜色值.
func (fw *FrameWindow) SetPaneSplitBarColor(color int) int {
	return xc.XFrameWnd_SetPaneSplitBarColor(fw.Handle, color)
}

// 置TabBar条高度, 设置窗格组TabBar高度.
// // nHeight: 高度.
func (fw *FrameWindow) SetTabBarHeight(nHeight int32) int {
	return xc.XFrameWnd_SetTabBarHeight(fw.Handle, nHeight)
}

// 保存布局到文件, 保存布局信息到文件.
// // pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func (fw *FrameWindow) SaveLayoutToFile(pFileName string) bool {
	return xc.XFrameWnd_SaveLayoutToFile(fw.Handle, pFileName)
}

// 加载布局信息文件, 加载布局信息文件.
// // aPaneList: 窗格句柄数组.
// // nPaneCount: 窗格数量.
// // pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func (fw *FrameWindow) LoadLayoutFile(aPaneList []int, nPaneCount int32, pFileName string) bool {
	return xc.XFrameWnd_LoadLayoutFile(fw.Handle, aPaneList, nPaneCount, pFileName)
}

// 添加窗格, 添加窗格到框架窗口.
// // hPaneDest: 目标窗格.
// // hPaneNew: 当前窗格.
// // align: 对齐方式, Pane_Align_.
func (fw *FrameWindow) AddPane(hPaneDest int, hPaneNew int, align xcc.Pane_Align_) bool {
	return xc.XFrameWnd_AddPane(fw.Handle, hPaneDest, hPaneNew, align)
}

// 合并窗格.
// // hPaneDest: 目标窗格.
// // hPaneNew: 当前窗格.
func (fw *FrameWindow) MergePane(hPaneDest int, hPaneNew int) bool {
	return xc.XFrameWnd_MergePane(fw.Handle, hPaneDest, hPaneNew)
}

// 取拖动浮动窗格停留位置标识, 返回 拖动窗格 所处框架窗口单元格标识: xcc.FrameWnd_Cell_Type_.
func (fw *FrameWindow) GetDragFloatWndTopFlag() xcc.FrameWnd_Cell_Type_ {
	return xc.XFrameWnd_GetDragFloatWndTopFlag(fw.Handle)
}

// 附加窗口, 返回窗口对象.
// // hWnd: 要附加的外部窗口句柄.
// // XCStyle: 炫彩窗口样式: Window_Style_.
func FrameWnd_Attach(hWnd uintptr, XCStyle int) *Window {
	p := &Window{}
	p.SetHandle(xc.XFrameWnd_Attach(hWnd, XCStyle))
	return p
}

// 取主视图坐标. 获取框架窗口主视图区域坐标.
// // pRect: 返回坐标.
func (fw *FrameWindow) GetViewRect(pRect *xc.RECT) {
	xc.XFrameWnd_GetViewRect(fw.Handle, pRect)
}

// 置窗格分隔条宽度.
// // nWidth: 宽度.
func (fw *FrameWindow) SetPaneSplitBarWidth(nWidth int32) {
	xc.XFrameWnd_SetPaneSplitBarWidth(fw.Handle, nWidth)
}

// 取窗格分隔条宽度.
func (fw *FrameWindow) GetPaneSplitBarWidth() int32 {
	return xc.XFrameWnd_GetPaneSplitBarWidth(fw.Handle)
}
