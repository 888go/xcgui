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

// ff:创建菜单
func X创建菜单() *Menu {
	p := &Menu{}
	p.X设置句柄(炫彩基类.X菜单_创建())
	return p
}

// 从句柄创建对象.

// ff:创建菜单并按句柄
// handle:
func X创建菜单并按句柄(handle int) *Menu {
	p := &Menu{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.

// ff:创建菜单并按名称
// name:
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

// ff:创建菜单并按UID
// nUID:
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

// ff:创建菜单并按UID名称
// name:
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

// ff:添加项
// nFlags:标识
// nParentID:父项ID
// pText:文本内容
// nID:项ID
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

// ff:添加项图标
// nFlags:标识
// hIcon:图标句柄
// nParentID:父项ID
// pText:文本内容
// nID:项ID
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

// ff:插入项
// insertID:
// nFlags:标识
// pText:文本内容
// nID:项ID
func (m *Menu) X插入项(项ID int32, 文本内容 string, 标识 炫彩常量类.Menu_Item_Flag_, insertID int32) {
	炫彩基类.X菜单_插入项(m.Handle, 项ID, 文本内容, 标识, insertID)
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

// ff:插入项图标
// insertID:
// nFlags:标识
// hIcon:图标句柄
// pText:文本内容
// nID:项ID
func (m *Menu) X插入项图标(项ID int32, 文本内容 string, 图标句柄 int, 标识 炫彩常量类.Menu_Item_Flag_, insertID int32) {
	炫彩基类.X菜单_插入项图标(m.Handle, 项ID, 文本内容, 图标句柄, 标识, insertID)
}

// 菜单_取第一个子项, 返回项ID.
//
// nID: 项ID.

// ff:取第一个子项
// nID:项ID
func (m *Menu) X取第一个子项(项ID int32) int32 {
	return 炫彩基类.X菜单_取第一个子项(m.Handle, 项ID)
}

// 菜单_取末尾子项, 返回项ID.
//
// nID: 项ID.

// ff:取末尾子项
// nID:项ID
func (m *Menu) X取末尾子项(项ID int32) int32 {
	return 炫彩基类.X菜单_取末尾子项(m.Handle, 项ID)
}

// 菜单_取上一个兄弟项, 返回项ID.
//
// nID: 项ID.

// ff:取上一个兄弟项
// nID:项ID
func (m *Menu) X取上一个兄弟项(项ID int32) int32 {
	return 炫彩基类.X菜单_取上一个兄弟项(m.Handle, 项ID)
}

// 菜单_取下一个兄弟项, 返回项ID.
//
// nID: 项ID.

// ff:取下一个兄弟项
// nID:项ID
func (m *Menu) X取下一个兄弟项(项ID int32) int32 {
	return 炫彩基类.X菜单_取下一个兄弟项(m.Handle, 项ID)
}

// 菜单_取父项, 返回项ID.
//
// nID: 项ID.

// ff:取父项
// nID:项ID
func (m *Menu) X取父项(项ID int32) int32 {
	return 炫彩基类.X菜单_取父项(m.Handle, 项ID)
}

// 菜单_置自动销毁, 设置是否自动销毁菜单.
//
// bAuto: 是否自动销毁.

// ff:置自动销毁
// bAuto:是否自动销毁
func (m *Menu) X置自动销毁(是否自动销毁 bool) {
	炫彩基类.X菜单_置自动销毁(m.Handle, 是否自动销毁)
}

// 菜单_启用用户绘制背景, 是否有用户绘制菜单背景, 如果启用XWM_MENU_DRAW_BACKGROUND和XE_MENU_DRAW_BACKGROUND事件有效.
//
// bEnable: 是否启用.

// ff:启用用户绘制背景
// bEnable:是否启用
func (m *Menu) X启用用户绘制背景(是否启用 bool) {
	炫彩基类.X菜单_启用用户绘制背景(m.Handle, 是否启用)
}

// 菜单_启用用户绘制项, 是否有用户绘制菜单项, 如果启用XWM_MENU_DRAWITEM和XE_MENU_DRAWITEM事件有效.
//
// bEnable: 是否启用.

// ff:启用用户绘制项
// bEnable:是否启用
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

// ff:弹出
// nPosition:弹出位置
// hParentEle:父元素句柄
// y:y坐标
// x:x坐标
// hParentWnd:父窗口句柄
func (m *Menu) X弹出(父窗口句柄 uintptr, x坐标, y坐标 int32, 父元素句柄 int, 弹出位置 炫彩常量类.Menu_Popup_Position_) bool {
	return 炫彩基类.X菜单_弹出(m.Handle, 父窗口句柄, x坐标, y坐标, 父元素句柄, 弹出位置)
}

// 菜单_销毁.

// ff:销毁
func (m *Menu) X销毁() {
	炫彩基类.X菜单_销毁(m.Handle)
}

// 菜单_关闭.

// ff:关闭
func (m *Menu) X关闭() {
	炫彩基类.X菜单_关闭(m.Handle)
}

// 菜单_置背景图片.
//
// hImage: 图片句柄.

// ff:置背景图片
// hImage:图片句柄
func (m *Menu) X置背景图片(图片句柄 int) {
	炫彩基类.X菜单_置背景图片(m.Handle, 图片句柄)
}

// 菜单_置项文本.
//
// nID: 项ID.
//
// pText: 文本内容.

// ff:置项文本
// pText:文本内容
// nID:项ID
func (m *Menu) X置项文本(项ID int32, 文本内容 string) bool {
	return 炫彩基类.X菜单_置项文本(m.Handle, 项ID, 文本内容)
}

// 菜单_取项文本.
//
// nID: 项ID.

// ff:取项文本
// nID:项ID
func (m *Menu) X取项文本(项ID int32) string {
	return 炫彩基类.X菜单_取项文本(m.Handle, 项ID)
}

// 菜单_取项文本长度, 获取项文本长度, 不包含字符串空终止符.
//
// nID: 项ID.

// ff:取项文本长度
// nID:项ID
func (m *Menu) X取项文本长度(项ID int32) int32 {
	return 炫彩基类.X菜单_取项文本长度(m.Handle, 项ID)
}

// 菜单_置项图标.
//
// nID: 项ID.
//
// hIcon: 菜单项图标句柄.

// ff:置项图标
// hIcon:菜单项图标句柄
// nID:项ID
func (m *Menu) X置项图标(项ID int32, 菜单项图标句柄 int) bool {
	return 炫彩基类.X菜单_置项图标(m.Handle, 项ID, 菜单项图标句柄)
}

// 菜单_置项标志.
//
// nID: 项ID.
//
// uFlags: 标识, Menu_Item_Flag_.

// ff:置项标志
// uFlags:标识
// nID:项ID
func (m *Menu) X置项标志(项ID int32, 标识 炫彩常量类.Menu_Item_Flag_) bool {
	return 炫彩基类.X菜单_置项标志(m.Handle, 项ID, 标识)
}

// 菜单_置项高度.
//
// height: 高度.

// ff:置项高度
// height:高度
func (m *Menu) X置项高度(高度 int32) {
	炫彩基类.X菜单_置项高度(m.Handle, 高度)
}

// 菜单_取项高度.

// ff:取项高度
func (m *Menu) X取项高度() int32 {
	return 炫彩基类.X菜单_取项高度(m.Handle)
}

// 菜单_置边框颜色, 设置菜单边框颜色.
//
// crColor: ABGR 颜色.

// ff:置边框颜色
// crColor:ABGR颜色
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

// ff:置边框大小
// nBottom:下
// nRight:右
// nTop:上
// nLeft:左
func (m *Menu) X置边框大小(左, 上, 右, 下 int32) {
	炫彩基类.X菜单_置边框大小(m.Handle, 左, 上, 右, 下)
}

// 菜单_取左侧宽度, 获取左侧区域宽度.

// ff:取左侧宽度
func (m *Menu) X取左侧宽度() int32 {
	return 炫彩基类.X菜单_取左侧宽度(m.Handle)
}

// 菜单_取左侧文本间隔, 获取菜单项文本左间隔.

// ff:取左侧文本间隔
func (m *Menu) X取左侧文本间隔() int32 {
	return 炫彩基类.X菜单_取左侧文本间隔(m.Handle)
}

// 菜单_取项数量, 获取菜单项数量, 包含子菜单项.

// ff:取项数量
func (m *Menu) X取项数量() int32 {
	return 炫彩基类.X菜单_取项数量(m.Handle)
}

// 菜单_置项勾选, 设置菜单项勾选状态.
//
// nID: 菜单项ID.
//
// bCheck: 勾选TRUE.

// ff:置项勾选
// bCheck:勾选TRUE
// nID:菜单项ID
func (m *Menu) X置项勾选(菜单项ID int32, 勾选TRUE bool) bool {
	return 炫彩基类.X菜单_置项勾选(m.Handle, 菜单项ID, 勾选TRUE)
}

// 菜单_判断项勾选, 判断菜单项是否勾选.
//
// nID: 菜单项ID.

// ff:判断项勾选
// nID:菜单项ID
func (m *Menu) X判断项勾选(菜单项ID int32) bool {
	return 炫彩基类.X菜单_判断项勾选(m.Handle, 菜单项ID)
}

// 菜单_置项宽度, 此宽度为文本显示区域宽度, 不包含侧边条和与文本间隔.
//
// nID: 项ID.
//
// nWidth: 指定文本区域宽度.

// ff:置项宽度
// nWidth:宽度
// nID:项ID
func (m *Menu) X置项宽度(项ID, 宽度 int32) bool {
	return 炫彩基类.X菜单_置项宽度(m.Handle, 项ID, 宽度)
}

// 菜单_取菜单条, 返回菜单条句柄.

// ff:取菜单条
func (m *Menu) X取菜单条() int {
	return 炫彩基类.X菜单_取菜单条(m.Handle)
}
