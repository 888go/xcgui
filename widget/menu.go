package 炫彩组件类

import (
	"github.com/888go/xcgui/objectbase"
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Menu 弹出菜单.
type Menu struct {
	炫彩对象基类.ObjectBase
}

// 菜单_创建, 创建菜单, 默认弹出菜单窗口关闭后自动销毁.
func X创建菜单() *Menu {
	p := &Menu{}
	p.X设置句柄(炫彩基类.X菜单_创建())
	return p
}

// 从句柄创建对象.
func X创建菜单并按句柄(handle int) *Menu {
	p := &Menu{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建菜单并按名称(name string) *Menu {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Menu{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建菜单并按UID(nUID int) *Menu {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Menu{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建菜单并按UID名称(name string) *Menu {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Menu{}
		p.X设置句柄(handle)
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
// nFlags: 标识, Menu_Item_Flag_.
func (m *Menu) X添加项(项ID int32, 文本内容 string, 父项ID int32, 标识 炫彩常量类.Menu_Item_Flag_) {
	炫彩基类.X菜单_添加项(m.Handle, 项ID, 文本内容, 父项ID, 标识)
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
// nFlags: 标识, Menu_Item_Flag_.
func (m *Menu) X添加项图标(项ID int32, 文本内容 string, 父项ID int32, 图标句柄 int, 标识 炫彩常量类.Menu_Item_Flag_) {
	炫彩基类.X菜单_添加项图标(m.Handle, 项ID, 文本内容, 父项ID, 图标句柄, 标识)
}

// 菜单_插入项.
//
// nID: 项ID.
//
// pText: 文本内容.
//
// nFlags: 标识, Menu_Item_Flag_.
//
// insertID: 插入位置ID.
func (m *Menu) X插入项(项ID int32, 文本内容 string, 标识 炫彩常量类.Menu_Item_Flag_, 插入位置ID int32) {
	炫彩基类.X菜单_插入项(m.Handle, 项ID, 文本内容, 标识, 插入位置ID)
}

// 菜单_插入项图标.
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
func (m *Menu) X插入项图标(项ID int32, 文本内容 string, 图标句柄 int, 标识 炫彩常量类.Menu_Item_Flag_, 插入位置ID int32) {
	炫彩基类.X菜单_插入项图标(m.Handle, 项ID, 文本内容, 图标句柄, 标识, 插入位置ID)
}

// 菜单_取第一个子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) X取第一个子项(项ID int32) int32 {
	return 炫彩基类.X菜单_取第一个子项(m.Handle, 项ID)
}

// 菜单_取末尾子项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) X取末尾子项(项ID int32) int32 {
	return 炫彩基类.X菜单_取末尾子项(m.Handle, 项ID)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) X取上一个兄弟项(项ID int32) int32 {
	return 炫彩基类.X菜单_取上一个兄弟项(m.Handle, 项ID)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) X取下一个兄弟项(项ID int32) int32 {
	return 炫彩基类.X菜单_取下一个兄弟项(m.Handle, 项ID)
}

// 菜单_取父项, 返回项ID.
//
// nID: 项ID.
func (m *Menu) X取父项(项ID int32) int32 {
	return 炫彩基类.X菜单_取父项(m.Handle, 项ID)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// bAuto: 是否自动销毁.
func (m *Menu) X置自动销毁(是否自动销毁 bool) {
	炫彩基类.X菜单_置自动销毁(m.Handle, 是否自动销毁)
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// bEnable: 是否启用.
func (m *Menu) X启用用户绘制背景(是否启用 bool) {
	炫彩基类.X菜单_启用用户绘制背景(m.Handle, 是否启用)
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// bEnable: 是否启用.
func (m *Menu) X启用用户绘制项(是否启用 bool) {
	炫彩基类.X菜单_启用用户绘制项(m.Handle, 是否启用)
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
// nPosition: 弹出位置, Menu_Popup_Position_.
func (m *Menu) X弹出(父窗口句柄 uintptr, x坐标, y坐标 int32, 父元素句柄 int, 弹出位置 炫彩常量类.Menu_Popup_Position_) bool {
	return 炫彩基类.X菜单_弹出(m.Handle, 父窗口句柄, x坐标, y坐标, 父元素句柄, 弹出位置)
}

// 菜单_销毁.
func (m *Menu) X销毁() {
	炫彩基类.X菜单_销毁(m.Handle)
}

// 菜单_关闭.
func (m *Menu) X关闭() {
	炫彩基类.X菜单_关闭(m.Handle)
}

// 菜单_置背景图片.
//
// hImage: 图片句柄.
func (m *Menu) X置背景图片(图片句柄 int) {
	炫彩基类.X菜单_置背景图片(m.Handle, 图片句柄)
}

// 菜单_置项文本.
//
// nID: 项ID.
//
// pText: 文本内容.
func (m *Menu) X置项文本(项ID int32, 文本内容 string) bool {
	return 炫彩基类.X菜单_置项文本(m.Handle, 项ID, 文本内容)
}

// 菜单_取项文本.
//
// nID: 项ID.
func (m *Menu) X取项文本(项ID int32) string {
	return 炫彩基类.X菜单_取项文本(m.Handle, 项ID)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// nID: 项ID.
func (m *Menu) X取项文本长度(项ID int32) int32 {
	return 炫彩基类.X菜单_取项文本长度(m.Handle, 项ID)
}

// 菜单_置项图标.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.
func (m *Menu) X置项图标(项ID int32, 菜单项图标句柄 int) bool {
	return 炫彩基类.X菜单_置项图标(m.Handle, 项ID, 菜单项图标句柄)
}

// 菜单_置项标志.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.
func (m *Menu) X置项标志(项ID int32, 标识 炫彩常量类.Menu_Item_Flag_) bool {
	return 炫彩基类.X菜单_置项标志(m.Handle, 项ID, 标识)
}

// 菜单_置项高度.
//
// height: 高度.
func (m *Menu) X置项高度(高度 int32) {
	炫彩基类.X菜单_置项高度(m.Handle, 高度)
}

// 菜单_取项高度.
func (m *Menu) X取项高度() int32 {
	return 炫彩基类.X菜单_取项高度(m.Handle)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// crColor: ABGR 颜色.
func (m *Menu) X置边框颜色(ABGR颜色 int) {
	炫彩基类.X菜单_置边框颜色(m.Handle, ABGR颜色)
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
func (m *Menu) X置边框大小(左, 上, 右, 下 int32) {
	炫彩基类.X菜单_置边框大小(m.Handle, 左, 上, 右, 下)
}

// 菜单_取左侧宽度, 获取左侧区域宽度.
func (m *Menu) X取左侧宽度() int32 {
	return 炫彩基类.X菜单_取左侧宽度(m.Handle)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.
func (m *Menu) X取左侧文本间隔() int32 {
	return 炫彩基类.X菜单_取左侧文本间隔(m.Handle)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.
func (m *Menu) X取项数量() int32 {
	return 炫彩基类.X菜单_取项数量(m.Handle)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.
func (m *Menu) X置项勾选(菜单项ID int32, 勾选TRUE bool) bool {
	return 炫彩基类.X菜单_置项勾选(m.Handle, 菜单项ID, 勾选TRUE)
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// nID: 菜单项ID.
func (m *Menu) X判断项勾选(菜单项ID int32) bool {
	return 炫彩基类.X菜单_判断项勾选(m.Handle, 菜单项ID)
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.
func (m *Menu) X置项宽度(项ID, 宽度 int32) bool {
	return 炫彩基类.X菜单_置项宽度(m.Handle, 项ID, 宽度)
}

// 菜单_取菜单条, 返回菜单条句柄.
func (m *Menu) X取菜单条() int {
	return 炫彩基类.X菜单_取菜单条(m.Handle)
}
