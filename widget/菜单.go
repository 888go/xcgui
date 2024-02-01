package widget

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 菜单.
type Menu struct {
	objectbase.ObjectBase
}

// I菜单_创建, 创建菜单, 默认弹出菜单窗口关闭后自动销毁.
func I菜单_创建() *Menu {
	p := &Menu{}
	p.SetHandle(xc.XMenu_Create())
	return p
}

// I菜单_从句柄创建对象.
func I菜单_从句柄创建(handle int) *Menu {
	p := &Menu{}
	p.SetHandle(handle)
	return p
}

// I菜单_从name创建对象, 失败返回nil.
func I菜单_从name创建(name string) *Menu {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Menu{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I菜单_从UID创建对象, 失败返回nil.
func I菜单_从UID创建(nUID int) *Menu {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Menu{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I菜单_从UID名称创建对象, 失败返回nil.
func I菜单_从UID名称创建(name string) *Menu {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Menu{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 菜单_添加项, 添加菜单项.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// nFlags: 标识, I常量_菜单_标识_.
func (m *Menu) I添加项(nID int, pText string, nParentID int, nFlags xcc.I常量_菜单_标识_) int {
	return xc.XMenu_AddItem(m.I句柄, nID, pText, nParentID, nFlags)
}

// 菜单_添加项图标.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// hIcon: 菜单项图标句柄.
//
// nFlags: 标识, I常量_菜单_标识_.
func (m *Menu) I添加项图标(nID int, pText string, nParentID int, hIcon int, nFlags xcc.I常量_菜单_标识_) int {
	return xc.XMenu_AddItemIcon(m.I句柄, nID, pText, nParentID, hIcon, nFlags)
}

// 菜单_插入项.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nFlags: 标识, I常量_菜单_标识_.
//
// insertID: 插入位置ID.
func (m *Menu) I插入项(nID int, pText string, nFlags xcc.I常量_菜单_标识_, insertID int) int {
	return xc.XMenu_InsertItem(m.I句柄, nID, pText, nFlags, insertID)
}

// 菜单_插入项图标.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// hIcon: 菜单项图标句柄.
//
// nFlags: 标识, I常量_菜单_标识_.
//
// insertID: 插入位置ID.
func (m *Menu) I插入项图标(nID int, pText string, hIcon int, nFlags xcc.I常量_菜单_标识_, insertID int) int {
	return xc.XMenu_InsertItemIcon(m.I句柄, nID, pText, hIcon, nFlags, insertID)
}

// 菜单_取第一个子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) I取第一个子项(nID int) int {
	return xc.XMenu_GetFirstChildItem(m.I句柄, nID)
}

// 菜单_取末尾子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) I取末尾子项(nID int) int {
	return xc.XMenu_GetEndChildItem(m.I句柄, nID)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) I取上一个兄弟项(nID int) int {
	return xc.XMenu_GetPrevSiblingItem(m.I句柄, nID)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) I取下一个兄弟项(nID int) int {
	return xc.XMenu_GetNextSiblingItem(m.I句柄, nID)
}

// 菜单_取父项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) I取父项(nID int) int {
	return xc.XMenu_GetParentItem(m.I句柄, nID)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// bAuto: 是否自动销毁.
func (m *Menu) I置自动销毁(bAuto bool) int {
	return xc.XMenu_SetAutoDestroy(m.I句柄, bAuto)
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// bEnable: 是否启用.
func (m *Menu) I启用用户绘制背景(bEnable bool) int {
	return xc.XMenu_EnableDrawBackground(m.I句柄, bEnable)
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// bEnable: 是否启用.
func (m *Menu) I启用用户绘制项(bEnable bool) int {
	return xc.XMenu_EnableDrawItem(m.I句柄, bEnable)
}

// 菜单_弹出.
//
// hParentWnd: 父窗口句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// hParentEle: 父元素句柄, 如果该值不为NULL, hParentEle元素将接收菜单消息事件, 否则将由hParentWnd窗口接收菜单的消息事件.
//
// nPosition: 弹出位置, I常量_菜单弹出方向_.
func (m *Menu) I弹出(hParentWnd int, x int, y int, hParentEle int, nPosition xcc.I常量_菜单弹出方向_) bool {
	return xc.XMenu_Popup(m.I句柄, hParentWnd, x, y, hParentEle, nPosition)
}

// 菜单_销毁.
func (m *Menu) I销毁() int {
	return xc.XMenu_DestroyMenu(m.I句柄)
}

// 菜单_关闭.
func (m *Menu) I关闭() int {
	return xc.XMenu_CloseMenu(m.I句柄)
}

// 菜单_置背景图片.
//
// hImage: 图片句柄.
func (m *Menu) I置背景图片(hImage int) int {
	return xc.XMenu_SetBkImage(m.I句柄, hImage)
}

// 菜单_置项文本.
//
// nID: 项ID.
//
// pText: 文本内容.
func (m *Menu) I置项文本(nID int, pText string) bool {
	return xc.XMenu_SetItemText(m.I句柄, nID, pText)
}

// 菜单_取项文本.
//
// nID: 项ID.
func (m *Menu) I取项文本(nID int) string {
	return xc.XMenu_GetItemText(m.I句柄, nID)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// nID: 项ID.
func (m *Menu) I取项文本长度(nID int) int {
	return xc.XMenu_GetItemTextLength(m.I句柄, nID)
}

// 菜单_置项图标.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func (m *Menu) I置项图标(nID int, hIcon int) bool {
	return xc.XMenu_SetItemIcon(m.I句柄, nID, hIcon)
}

// 菜单_置项标志.
//
// nID: 项ID.
//
// uFlags: 标识, I常量_菜单_标识_.
func (m *Menu) I置项标志(nID int, uFlags int) bool {
	return xc.XMenu_SetItemFlags(m.I句柄, nID, uFlags)
}

// 菜单_置项高度.
//
// height: 高度.
func (m *Menu) I置项高度(height int) int {
	return xc.XMenu_SetItemHeight(m.I句柄, height)
}

// 菜单_取项高度.
func (m *Menu) I取项高度() int {
	return xc.XMenu_GetItemHeight(m.I句柄)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// crColor: ABGR 颜色.
func (m *Menu) I置边框颜色(crColor int) int {
	return xc.XMenu_SetBorderColor(m.I句柄, crColor)
}

// 菜单_置边框大小, 设置弹出菜单窗口边框大小.
//
// nLeft: 边大小.
//
// nTop: 边大小.
//
// nRight: 边大小.
//
// nBottom: 边大小.
func (m *Menu) I置边框大小(nLeft int, nTop int, nRight int, nBottom int) int {
	return xc.XMenu_SetBorderSize(m.I句柄, nLeft, nTop, nRight, nBottom)
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
func (m *Menu) I取左侧宽度() int {
	return xc.XMenu_GetLeftWidth(m.I句柄)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
func (m *Menu) I取左侧文本间隔() int {
	return xc.XMenu_GetLeftSpaceText(m.I句柄)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
func (m *Menu) I取项数量() int {
	return xc.XMenu_GetItemCount(m.I句柄)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func (m *Menu) I置项勾选(nID int, bCheck bool) bool {
	return xc.XMenu_SetItemCheck(m.I句柄, nID, bCheck)
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// nID: 菜单项ID.
func (m *Menu) I判断项勾选(nID int) bool {
	return xc.XMenu_IsItemCheck(m.I句柄, nID)
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.
func (m *Menu) I置项宽度(nID int, nWidth int) bool {
	return xc.XMenu_SetItemWidth(m.I句柄, nID, nWidth)
}
