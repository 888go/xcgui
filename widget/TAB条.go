package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

type TabBar struct {
	Element
}

// TAB条_创建, 创建tabBar元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父是窗口资源句柄或UI元素资源句柄. 如果是窗口资源句柄将被添加到窗口, 如果是元素资源句柄将被添加到元素.
func TAB条_创建(x int, y int, cx int, cy int, hParent int) *TabBar {
	p := &TabBar{}
	p.SetHandle(xc.XTabBar_Create(x, y, cx, cy, hParent))
	return p
}

// TAB条_从句柄创建对象.
func TAB条_从句柄创建(handle int) *TabBar {
	p := &TabBar{}
	p.SetHandle(handle)
	return p
}

// TAB条_从name创建对象, 失败返回nil.
func TAB条_从name创建(name string) *TabBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &TabBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// TAB条_从UID创建对象, 失败返回nil.
func TAB条_从UID创建(nUID int) *TabBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &TabBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// TAB条_从UID名称创建对象, 失败返回nil.
func TAB条_从UID名称创建(name string) *TabBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &TabBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// TAB条_添加标签, 添加一个标签, 返回标签索引.
//
// pName: 标签文本内容.
func (t *TabBar) I添加标签(pName string) int {
	return xc.XTabBar_AddLabel(t.I句柄, pName)
}

// TAB条插入_标签, 插入一个标签, 返回标签索引.
//
// index: 插入位置.
//
// pName: 标签文本内容.
func (t *TabBar) I插入标签(index int, pName string) int {
	return xc.XTabBar_InsertLabel(t.I句柄, index, pName)
}

// TAB条_移动标签.
//
// iSrc: 源位置索引.
//
// iDest: 目标位置索引.
func (t *TabBar) I移动标签(iSrc int, iDest int) bool {
	return xc.XTabBar_MoveLabel(t.I句柄, iSrc, iDest)
}

// TAB条_删除标签, 删除一个标签.
//
// index: 位置索引.
func (t *TabBar) I删除标签(index int) bool {
	return xc.XTabBar_DeleteLabel(t.I句柄, index)
}

// TAB条_删除全部, 删除所有标签.
func (t *TabBar) I删除全部() int {
	return xc.XTabBar_DeleteLabelAll(t.I句柄)
}

// TAB条_取标签, 获取标签按钮句柄.
//
// index: 位置索引.
func (t *TabBar) I取标签(index int) int {
	return xc.XTabBar_GetLabel(t.I句柄, index)
}

// TAB条_取标签上的关闭按钮, 获取标签上关闭按钮句柄.
//
// index: 位置索引.
func (t *TabBar) I取标签上的关闭按钮(index int) int {
	return xc.XTabBar_GetLabelClose(t.I句柄, index)
}

// TAB条_取左滚动按钮, 获取左滚动按钮句柄.
func (t *TabBar) I取左滚动按钮() int {
	return xc.XTabBar_GetButtonLeft(t.I句柄)
}

// TAB条_取右滚动按钮, 获取右滚动按钮句柄.
func (t *TabBar) I取右滚动按钮() int {
	return xc.XTabBar_GetButtonRight(t.I句柄)
}

// TAB条_取下拉菜单按钮句柄.
func (t *TabBar) I取下拉菜单按钮句柄() int {
	return xc.XTabBar_GetButtonDropMenu(t.I句柄)
}

// TAB条_取当前选择, 获取选择的标签索引.
func (t *TabBar) I取当前选择() int {
	return xc.XTabBar_GetSelect(t.I句柄)
}

// TAB条_取间隔, 获取标签间距, 0没有间距.
func (t *TabBar) I取间隔() int {
	return xc.XTabBar_GetLabelSpacing(t.I句柄)
}

// TAB条_取标签数量, 获取标签项数量.
func (t *TabBar) I取标签数量() int {
	return xc.XTabBar_GetLabelCount(t.I句柄)
}

// TAB条_取标签位置索引, 获取标签按钮位置索引, 成功返回索引值, 否则返回 I常量_错误.
//
// hLabel: 标签按钮句柄.
func (t *TabBar) I取标签位置索引(hLabel int) int {
	return xc.XTabBar_GetindexByEle(t.I句柄, hLabel)
}

// TAB条_置间隔, 设置标签间距, 0没有间距.
//
// spacing: 标签间隔大小.
func (t *TabBar) I置间隔(spacing int) int {
	return xc.XTabBar_SetLabelSpacing(t.I句柄, spacing)
}

// TAB条_置边距, 设置内容与边框的间隔大小.
//
// left: 左边间隔大小.
//
// top: 上边间隔大小.
//
// right: 右边间隔大小.
//
// bottom: 下边间隔大小.
func (t *TabBar) I置边距(left int, top int, right int, bottom int) int {
	return xc.XTabBar_SetPadding(t.I句柄, left, top, right, bottom)
}

// TAB条_置选择, 设置选择标签.
//
// index: 标签位置索引.
func (t *TabBar) I置选择(index int) int {
	return xc.XTabBar_SetSelect(t.I句柄, index)
}

// TAB条_左滚动, 左按钮滚动.
func (t *TabBar) I左滚动() int {
	return xc.XTabBar_SetUp(t.I句柄)
}

// TAB条_右滚动, 右按钮滚动.
func (t *TabBar) I右滚动() int {
	return xc.XTabBar_SetDown(t.I句柄)
}

// TAB条_启用平铺, 平铺标签, 每个标签显示相同大小.
//
// bTile: 是否启用.
func (t *TabBar) I启用平铺(bTile bool) int {
	return xc.XTabBar_EnableTile(t.I句柄, bTile)
}

// TAB条_启用下拉菜单按钮.
//
// bEnable:.
func (t *TabBar) I启用下拉菜单按钮(bEnable bool) int {
	return xc.XTabBar_EnableDropMenu(t.I句柄, bEnable)
}

// TAB条_启用标签带关闭按钮, 启用关闭标签功能.
//
// bEnable: 是否启用.
func (t *TabBar) I启用标签带关闭按钮(bEnable bool) int {
	return xc.XTabBar_EnableClose(t.I句柄, bEnable)
}

// TAB条_置关闭按钮大小, 设置关闭按钮大小.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.
func (t *TabBar) I置关闭按钮大小(pSize *xc.SIZE) int {
	return xc.XTabBar_SetCloseSize(t.I句柄, pSize)
}

// TAB条_置滚动按钮大小.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.
func (t *TabBar) I置滚动按钮大小(pSize *xc.SIZE) int {
	return xc.XTabBar_SetTurnButtonSize(t.I句柄, pSize)
}

// TAB条_置指定标签固定宽度.
//
// index: 索引.
//
// nWidth: 宽度, , 如果值为-1, 那么自动计算宽度.
func (t *TabBar) I置指定标签固定宽度(index int, nWidth int) int {
	return xc.XTabBar_SetLabelWidth(t.I句柄, index, nWidth)
}

// TAB条_显示标签, 显示或隐藏指定标签.
//
// index: 标签索引.
//
// bShow: 是否显示.
func (t *TabBar) I显示标签(index int, bShow bool) bool {
	return xc.XTabBar_ShowLabel(t.I句柄, index, bShow)
}

/*
以下都是事件
*/

type XE_TABBAR_SELECT func(iItem int, pbHandled *bool) int            // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_SELECT1 func(hEle int, iItem int, pbHandled *bool) int // TabBar标签按钮选择改变事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE func(iItem int, pbHandled *bool) int            // TabBar标签按钮删除事件. iItem: 标签位置索引.
type XE_TABBAR_DELETE1 func(hEle int, iItem int, pbHandled *bool) int // TabBar标签按钮删除事件. iItem: 标签位置索引.

// TabBar  I事件_标签按钮选择改变 事件.
func (t *TabBar) I事件_标签按钮选择改变(pFun XE_TABBAR_SELECT) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TABBAR_SELECT, pFun)
}

// TabBar I事件_标签按钮选择改变 事件.
func (t *TabBar) I事件_标签按钮选择改变1(pFun XE_TABBAR_SELECT1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TABBAR_SELECT, pFun)
}

// TabBar I事件_标签按钮删除  事件.
func (t *TabBar) I事件_标签按钮删除(pFun XE_TABBAR_DELETE) bool {
	return xc.XEle_RegEventC(t.I句柄, xcc.XE_TABBAR_DELETE, pFun)
}

// TabBar  I事件_标签按钮删除 事件.
func (t *TabBar) I事件_标签按钮删除1(pFun XE_TABBAR_DELETE1) bool {
	return xc.XEle_RegEventC1(t.I句柄, xcc.XE_TABBAR_DELETE, pFun)
}
