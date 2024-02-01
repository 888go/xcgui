package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// Pane元素.
type Pane struct {
	Element
}

// I窗格_创建, 创建窗格元素, 返回元素句柄.
//
// pName: 窗格标题.
//
// nWidth: 宽度.
//
// nHeight: 高度.
//
// hFrameWnd: 框架窗口.
func I窗格_创建(pName string, nWidth int, nHeight int, hFrameWnd int) *Pane {
	p := &Pane{}
	p.SetHandle(xc.XPane_Create(pName, nWidth, nHeight, hFrameWnd))
	return p
}

// I窗格_从句柄创建对象.
func I窗格_从句柄创建(handle int) *Pane {
	p := &Pane{}
	p.SetHandle(handle)
	return p
}

// I窗格_从name创建对象, 失败返回nil.
func I窗格_从name创建(name string) *Pane {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Pane{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗格_从UID创建对象, 失败返回nil.
func I窗格_从UID创建(nUID int) *Pane {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Pane{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I窗格_从UID名称创建对象, 失败返回nil.
func I窗格_从UID名称创建(name string) *Pane {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Pane{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 窗格_置视图, 设置窗格视图元素.
//
// hView: 绑定视图元素.
func (p *Pane) I置视图(hView int) int {
	return xc.XPane_SetView(p.I句柄, hView)
}

// 窗格_置标题, 设置标题文本.
//
// pTitle: 文本内容.
func (p *Pane) I置标题(pTitle string) int {
	return xc.XPane_SetTitle(p.I句柄, pTitle)
}

// 窗格_取标题, 获取标题文本.
func (p *Pane) I取标题() string {
	return xc.XPane_GetTitle(p.I句柄)
}

// 窗格_置标题栏高度, 设置标题栏高度.
//
// nHeight: 高度.
func (p *Pane) I置标题栏高度(nHeight int) int {
	return xc.XPane_SetCaptionHeight(p.I句柄, nHeight)
}

// 窗格_取标题栏高度, 获取标题栏高度.
func (p *Pane) I取标题栏高度() int {
	return xc.XPane_GetCaptionHeight(p.I句柄)
}

// 窗格_判断显示, 判断窗格是否显示.
func (p *Pane) I判断显示() bool {
	return xc.XPane_IsShowPane(p.I句柄)
}

// 窗格_置大小, 设置窗格大小.
//
// nWidth: 宽度.
//
// nHeight: 高度.
func (p *Pane) I置大小(nWidth int, nHeight int) int {
	return xc.XPane_SetSize(p.I句柄, nWidth, nHeight)
}

// 窗格_取状态, 获取窗格停靠状态, 返回: I常量_窗格状态_.
func (p *Pane) I取状态() xcc.I常量_窗格状态_ {
	return xc.XPane_GetState(p.I句柄)
}

// 窗格_取视图坐标, 获取窗格视图坐标.
//
// pRect: 接收返回的坐标值.
func (p *Pane) I取视图坐标(pRect *xc.RECT) int {
	return xc.XPane_GetViewRect(p.I句柄, pRect)
}

// I隐藏 窗格_隐藏.
//	@param bGroupActivate 当为窗格组成员时, 延迟处理窗格组成员激活的切换.
//	@return int
//
func (p *Pane) I隐藏(bGroupActivate bool) int {
	return xc.XPane_HidePane(p.I句柄, bGroupActivate)
}

// I显示 窗格_显示.
//	@param bGroupActivate 如果是窗格组成员, 那么窗格组切换当前窗格为显示状态.
//	@return int
//
func (p *Pane) I显示(bGroupActivate bool) int {
	return xc.XPane_ShowPane(p.I句柄, bGroupActivate)
}

// 窗格_停靠, 窗格停靠到码头.
func (p *Pane) I停靠() int {
	return xc.XPane_DockPane(p.I句柄)
}

// 窗格_锁定, 锁定窗格.
func (p *Pane) I锁定() int {
	return xc.XPane_LockPane(p.I句柄)
}

// 窗格_浮动.
func (p *Pane) I浮动() int {
	return xc.XPane_FloatPane(p.I句柄)
}

// 窗格_绘制, 手动调用该函数绘制窗格, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.
func (p *Pane) I绘制(hDraw int) int {
	return xc.XPane_DrawPane(p.I句柄, hDraw)
}

// 窗口_置选中, 如果窗格是组成员, 设置选中当前窗格可见.
func (p *Pane) I置选中() bool {
	return xc.XPane_SetSelect(p.I句柄)
}

// 窗格_是否激活. 判断窗格是否激活, 当为组成员时有效.
func (p *Pane) I是否激活() bool {
	return xc.XPane_IsGroupActivate(p.I句柄)
}
