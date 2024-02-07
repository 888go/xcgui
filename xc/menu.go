package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/xcc"
)

// 菜单_创建, 创建菜单, 默认弹出菜单窗口关闭后自动销毁.
func X菜单_创建() int {
	r, _, _ := xMenu_Create.Call()
	return int(r)
}

// 菜单_添加项, 添加菜单项.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// nFlags: 标识, Menu_Item_Flag_.
func X菜单_添加项(菜单句柄 int, 项ID int32, 文本内容 string, 父项ID int32, 标识 炫彩常量类.Menu_Item_Flag_) {
	xMenu_AddItem.Call(uintptr(菜单句柄), uintptr(项ID), 炫彩工具类.StrPtr(文本内容), uintptr(父项ID), uintptr(标识))
}

// 菜单_添加项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nParentID: 父项ID.
//
// hIcon: 菜单项图标句柄.
//
// nFlags: 标识, Menu_Item_Flag_.
func X菜单_添加项图标(菜单句柄 int, 项ID int32, 文本内容 string, 父项ID int32, 菜单项图标句柄 int, 标识 炫彩常量类.Menu_Item_Flag_) {
	xMenu_AddItemIcon.Call(uintptr(菜单句柄), uintptr(项ID), 炫彩工具类.StrPtr(文本内容), uintptr(父项ID), uintptr(菜单项图标句柄), uintptr(标识))
}

// 菜单_插入项.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nFlags: 标识, Menu_Item_Flag_.
//
// insertID: 插入位置ID.
func X菜单_插入项(菜单句柄 int, 项ID int32, 文本内容 string, 标识 炫彩常量类.Menu_Item_Flag_, 插入位置ID int32) {
	xMenu_InsertItem.Call(uintptr(菜单句柄), uintptr(项ID), 炫彩工具类.StrPtr(文本内容), uintptr(标识), uintptr(插入位置ID))
}

// 菜单_插入项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// hIcon: 菜单项图标句柄.
//
// nFlags: 标识, Menu_Item_Flag_.
//
// insertID: 插入位置ID.
func X菜单_插入项图标(菜单句柄 int, 项ID int32, 文本内容 string, 菜单项图标句柄 int, 标识 炫彩常量类.Menu_Item_Flag_, 插入位置ID int32) {
	xMenu_InsertItemIcon.Call(uintptr(菜单句柄), uintptr(项ID), 炫彩工具类.StrPtr(文本内容), uintptr(菜单项图标句柄), uintptr(标识), uintptr(插入位置ID))
}

// 菜单_取第一个子项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取第一个子项(菜单句柄 int, 项ID int32) int32 {
	r, _, _ := xMenu_GetFirstChildItem.Call(uintptr(菜单句柄), uintptr(项ID))
	return int32(r)
}

// 菜单_取末尾子项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取末尾子项(菜单句柄 int, 项ID int32) int32 {
	r, _, _ := xMenu_GetEndChildItem.Call(uintptr(菜单句柄), uintptr(项ID))
	return int32(r)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取上一个兄弟项(菜单句柄 int, 项ID int32) int32 {
	r, _, _ := xMenu_GetPrevSiblingItem.Call(uintptr(菜单句柄), uintptr(项ID))
	return int32(r)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取下一个兄弟项(菜单句柄 int, 项ID int32) int32 {
	r, _, _ := xMenu_GetNextSiblingItem.Call(uintptr(菜单句柄), uintptr(项ID))
	return int32(r)
}

// 菜单_取父项, 返回项ID.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取父项(菜单句柄 int, 项ID int32) int32 {
	r, _, _ := xMenu_GetParentItem.Call(uintptr(菜单句柄), uintptr(项ID))
	return int32(r)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// hMenu: 菜单句柄.
//
// bAuto: 是否自动销毁.
func X菜单_置自动销毁(菜单句柄 int, 是否自动销毁 bool) {
	xMenu_SetAutoDestroy.Call(uintptr(菜单句柄), 炫彩工具类.BoolPtr(是否自动销毁))
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// hMenu: 菜单句柄.
//
// bEnable: 是否启用.
func X菜单_启用用户绘制背景(菜单句柄 int, 是否启用 bool) {
	xMenu_EnableDrawBackground.Call(uintptr(菜单句柄), 炫彩工具类.BoolPtr(是否启用))
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// hMenu: 菜单句柄.
//
// bEnable: 是否启用.
func X菜单_启用用户绘制项(菜单句柄 int, 是否启用 bool) {
	xMenu_EnableDrawItem.Call(uintptr(菜单句柄), 炫彩工具类.BoolPtr(是否启用))
}

// 菜单_弹出.
//
// hMenu: 菜单句柄.
//
// hParentWnd: 父窗口真实句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// hParentEle: 父元素句柄, 如果该值不为NULL, hParentEle元素将接收菜单消息事件, 否则将由hParentWnd窗口接收菜单的消息事件.
//
// nPosition: 弹出位置, Menu_Popup_Position_.
func X菜单_弹出(菜单句柄 int, 父窗口真实句柄 uintptr, x坐标, y坐标 int32, 父元素句柄 int, 弹出位置 炫彩常量类.Menu_Popup_Position_) bool {
	r, _, _ := xMenu_Popup.Call(uintptr(菜单句柄), 父窗口真实句柄, uintptr(x坐标), uintptr(y坐标), uintptr(父元素句柄), uintptr(弹出位置))
	return r != 0
}

// 菜单_销毁.
//
// hMenu: 菜单句柄.
func X菜单_销毁(菜单句柄 int) {
	xMenu_DestroyMenu.Call(uintptr(菜单句柄))
}

// 菜单_关闭.
//
// hMenu: 菜单句柄.
func X菜单_关闭(菜单句柄 int) {
	xMenu_CloseMenu.Call(uintptr(菜单句柄))
}

// 菜单_置背景图片.
//
// hMenu: 菜单句柄.
//
// hImage: 图片句柄.
func X菜单_置背景图片(菜单句柄 int, 图片句柄 int) {
	xMenu_SetBkImage.Call(uintptr(菜单句柄), uintptr(图片句柄))
}

// 菜单_置项文本.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// pText: 文本内容.
func X菜单_置项文本(菜单句柄 int, 项ID int32, 文本内容 string) bool {
	r, _, _ := xMenu_SetItemText.Call(uintptr(菜单句柄), uintptr(项ID), 炫彩工具类.StrPtr(文本内容))
	return r != 0
}

// 菜单_取项文本.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取项文本(菜单句柄 int, 项ID int32) string {
	r, _, _ := xMenu_GetItemText.Call(uintptr(菜单句柄), uintptr(项ID))
	return 炫彩工具类.UintPtrToString(r)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
func X菜单_取项文本长度(菜单句柄 int, 项ID int32) int32 {
	r, _, _ := xMenu_GetItemTextLength.Call(uintptr(菜单句柄), uintptr(项ID))
	return int32(r)
}

// 菜单_置项图标.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func X菜单_置项图标(菜单句柄 int, 项ID int32, 菜单项图标句柄 int) bool {
	r, _, _ := xMenu_SetItemIcon.Call(uintptr(菜单句柄), uintptr(项ID), uintptr(菜单项图标句柄))
	return r != 0
}

// 菜单_置项标志.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.
func X菜单_置项标志(菜单句柄 int, 项ID int32, 标识 炫彩常量类.Menu_Item_Flag_) bool {
	r, _, _ := xMenu_SetItemFlags.Call(uintptr(菜单句柄), uintptr(项ID), uintptr(标识))
	return r != 0
}

// 菜单_置项高度.
//
// hMenu: 菜单句柄.
//
// height: 高度.
func X菜单_置项高度(菜单句柄 int, 高度 int32) {
	xMenu_SetItemHeight.Call(uintptr(菜单句柄), uintptr(高度))
}

// 菜单_取项高度.
//
// hMenu: 菜单句柄.
func X菜单_取项高度(菜单句柄 int) int32 {
	r, _, _ := xMenu_GetItemHeight.Call(uintptr(菜单句柄))
	return int32(r)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// hMenu: 菜单句柄.
//
// crColor: ABGR 颜色.
func X菜单_置边框颜色(菜单句柄 int, ABGR颜色 int) {
	xMenu_SetBorderColor.Call(uintptr(菜单句柄), uintptr(ABGR颜色))
}

// 菜单_置边框大小, 设置弹出菜单窗口边框大小.
//
// hMenu: 菜单句柄.
//
// nLeft: 边大小.
//
// nTop: 边大小.
//
// nRight: 边大小.
//
// nBottom: 边大小.
func X菜单_置边框大小(菜单句柄 int, 左, 上, 右, 下 int32) {
	xMenu_SetBorderSize.Call(uintptr(菜单句柄), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
//
// hMenu: 菜单句柄.
func X菜单_取左侧宽度(菜单句柄 int) int32 {
	r, _, _ := xMenu_GetLeftWidth.Call(uintptr(菜单句柄))
	return int32(r)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
//
// hMenu: 菜单句柄.
func X菜单_取左侧文本间隔(菜单句柄 int) int32 {
	r, _, _ := xMenu_GetLeftSpaceText.Call(uintptr(菜单句柄))
	return int32(r)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
//
// hMenu: 菜单句柄.
func X菜单_取项数量(菜单句柄 int) int32 {
	r, _, _ := xMenu_GetItemCount.Call(uintptr(菜单句柄))
	return int32(r)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// hMenu: 菜单句柄.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func X菜单_置项勾选(菜单句柄 int, 菜单项ID int32, 勾选TRUE bool) bool {
	r, _, _ := xMenu_SetItemCheck.Call(uintptr(菜单句柄), uintptr(菜单项ID), 炫彩工具类.BoolPtr(勾选TRUE))
	return r != 0
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// hMenu: 菜单句柄.
//
// nID: 菜单项ID.
func X菜单_判断项勾选(菜单句柄 int, 菜单项ID int32) bool {
	r, _, _ := xMenu_IsItemCheck.Call(uintptr(菜单句柄), uintptr(菜单项ID))
	return r != 0
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// hMenu: 菜单句柄.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.
func X菜单_置项宽度(菜单句柄 int, 项ID, 指定文本区域宽度 int32) bool {
	r, _, _ := xMenu_SetItemWidth.Call(uintptr(菜单句柄), uintptr(项ID), uintptr(指定文本区域宽度))
	return r != 0
}

// 菜单_取菜单条, 返回菜单条句柄.
//
// hMenu: 菜单句柄.
func X菜单_取菜单条(菜单句柄 int) int {
	r, _, _ := xMenu_GetMenuBar.Call(uintptr(菜单句柄))
	return int(r)
}
