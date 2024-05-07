package 炫彩窗口基类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// FrameWindow 框架窗口.
type FrameWindow struct {
	windowBase
}

// 框架窗口_创建.
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
func X创建框架窗口(x坐标, y坐标, 宽度, 高度 int32, 标题 string, 父窗口句柄 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) *FrameWindow {
	p := &FrameWindow{}
	p.X设置句柄(炫彩基类.X框架窗口_创建(x坐标, y坐标, 宽度, 高度, 标题, 父窗口句柄, GUI库窗口样式))
	return p
}

// 框架窗口_创建扩展.
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
func X创建框架窗口EX(扩展样式 int, 样式 int, 类名 string, x坐标, y坐标, 宽度, 高度 int32, 窗口名 string, 父窗口 uintptr, GUI库窗口样式 炫彩常量类.Window_Style_) *FrameWindow {
	p := &FrameWindow{}
	p.X设置句柄(炫彩基类.X框架窗口_创建EX(扩展样式, 样式, 类名, x坐标, y坐标, 宽度, 高度, 窗口名, 父窗口, GUI库窗口样式))
	return p
}

// 从句柄创建对象.
func X创建框架窗口并按句柄(handle int) *FrameWindow {
	p := &FrameWindow{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建框架窗口并按名称(name string) *FrameWindow {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &FrameWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建框架窗口并按UID(nUID int) *FrameWindow {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &FrameWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建框架窗口并按UID名称(name string) *FrameWindow {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &FrameWindow{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 框架窗口_取布局区域坐标, 用来布局窗格的区域坐标, 不包含码头.
//
// pRect: 返回坐标.
func (fw *FrameWindow) X取布局区域坐标(返回坐标 *炫彩基类.RECT) int {
	return 炫彩基类.X框架窗口_取布局区域坐标(fw.Handle, 返回坐标)
}

// 框架窗口_置视图, 设置窗格组TabBar高度.
//
// hEle: 元素句柄.
func (fw *FrameWindow) X置视图(元素句柄 int) int {
	return 炫彩基类.X框架窗口_置视图(fw.Handle, 元素句柄)
}

// 框架窗口_置窗格分隔条颜色.
//
// color: ABGR 颜色值.
func (fw *FrameWindow) X置窗格分隔条颜色(ABGR颜色值 int) int {
	return 炫彩基类.X框架窗口_置窗格分隔条颜色(fw.Handle, ABGR颜色值)
}

// 框架窗口_置TabBar条高度, 设置窗格组TabBar高度.
//
// nHeight: 高度.
func (fw *FrameWindow) X置TabBar条高度(高度 int32) int {
	return 炫彩基类.X框架窗口_置TabBar条高度(fw.Handle, 高度)
}

// 框架窗口_保存布局到文件, 保存布局信息到文件.
//
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func (fw *FrameWindow) X保存布局到文件(文件名 string) bool {
	return 炫彩基类.X框架窗口_保存布局到文件(fw.Handle, 文件名)
}

// 框架窗口_加载布局信息文件, 加载布局信息文件.
//
// aPaneList: 窗格句柄切片.
//
// nPaneCount: 窗格数量.
//
// pFileName: 文件名，如果文件名为空，将使用默认文件名frameWnd_layout.xml.
func (fw *FrameWindow) X加载布局信息文件(窗格句柄切片 []int, 窗格数量 int32, 文件名 string) bool {
	return 炫彩基类.X框架窗口_加载布局信息文件(fw.Handle, 窗格句柄切片, 窗格数量, 文件名)
}

// 框架窗口_添加窗格, 添加窗格到框架窗口.
//
// hPaneDest: 目标窗格.
//
// hPaneNew: 当前窗格.
//
// align: 对齐方式, Pane_Align_.
func (fw *FrameWindow) X添加窗格(目标窗格 int, 当前窗格 int, 对齐方式 炫彩常量类.Pane_Align_) bool {
	return 炫彩基类.X框架窗口_添加窗格(fw.Handle, 目标窗格, 当前窗格, 对齐方式)
}

// 框架窗口_合并窗格.
//
// hPaneDest: 目标窗格.
//
// hPaneNew: 当前窗格.
func (fw *FrameWindow) X合并窗格(目标窗格 int, 当前窗格 int) bool {
	return 炫彩基类.X框架窗口_合并窗格(fw.Handle, 目标窗格, 当前窗格)
}

// 框架窗口_取拖动浮动窗格停留位置标识, 返回 拖动窗格 所处框架窗口单元格标识: xcc.FrameWnd_Cell_Type_.
func (fw *FrameWindow) X取拖动浮动窗格停留位置标识() 炫彩常量类.FrameWnd_Cell_Type_ {
	return 炫彩基类.X框架窗口_取拖动浮动窗格停留位置标识(fw.Handle)
}

// 框架窗口_附加窗口, 返回窗口对象.
//
// hWnd: 要附加的外部窗口句柄.
//
// XCStyle: 炫彩窗口样式: Window_Style_.
func X附加窗口(外部窗口句柄 uintptr, 炫彩窗口样式 int) *Window {
	p := &Window{}
	p.X设置句柄(炫彩基类.X框架窗口_附加窗口(外部窗口句柄, 炫彩窗口样式))
	return p
}

// 框架窗口_取主视图坐标. 获取框架窗口主视图区域坐标.
//
// pRect: 返回坐标.
func (fw *FrameWindow) X取主视图坐标(返回坐标 *炫彩基类.RECT) {
	炫彩基类.X框架窗口_取主视图坐标(fw.Handle, 返回坐标)
}

// 框架窗口_置窗格分隔条宽度.
//
// nWidth: 宽度.
func (fw *FrameWindow) X置窗格分隔条宽度(宽度 int32) {
	炫彩基类.X框架窗口_置窗格分隔条宽度(fw.Handle, 宽度)
}

// 框架窗口_取窗格分隔条宽度.
func (fw *FrameWindow) X取窗格分隔条宽度() int32 {
	return 炫彩基类.X框架窗口_取窗格分隔条宽度(fw.Handle)
}
