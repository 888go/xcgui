package widget

import "e.coding.net/gogit/go/xcgui/xc"

// 菜单条.
type MenuBar struct {
	Element
}

// I菜单条_创建, 创建菜单条元素; 如果指定了父为窗口, 默认调用XWnd_AddMenuBar()函数, 将菜单条添加到窗口非客户区.
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
func I菜单条_创建(x int, y int, cx int, cy int, hParent int) *MenuBar {
	p := &MenuBar{}
	p.SetHandle(xc.XMenuBar_Create(x, y, cx, cy, hParent))
	return p
}

// I菜单条_从句柄创建对象.
func I菜单条_从句柄创建(handle int) *MenuBar {
	p := &MenuBar{}
	p.SetHandle(handle)
	return p
}

// I菜单条_从name创建对象, 失败返回nil.
func I菜单条_从name创建(name string) *MenuBar {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &MenuBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I菜单条_从UID创建对象, 失败返回nil.
func I菜单条_从UID创建(nUID int) *MenuBar {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &MenuBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I菜单条_从UID名称创建对象, 失败返回nil.
func I菜单条_从UID名称创建(name string) *MenuBar {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &MenuBar{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 菜单条_添加按钮, 添加弹出菜单按钮, 返回菜单按钮索引.
//
// pText: 文本内容.
func (m *MenuBar) I添加按钮(pText string) int {
	return xc.XMenuBar_AddButton(m.I句柄, pText)
}

// 菜单条_置按钮高度, 根据内容自动调整宽度.
//
// height: 高度.
func (m *MenuBar) I置按钮高度(height int) int {
	return xc.XMenuBar_SetButtonHeight(m.I句柄, height)
}

// 菜单条_取菜单, 返回菜单句柄.
//
// nIndex: 菜单条上菜单按钮的索引.
func (m *MenuBar) I取菜单(nIndex int) int {
	return xc.XMenuBar_GetMenu(m.I句柄, nIndex)
}

// 菜单条_删除按钮, 删除菜单条上的菜单按钮, 同时该按钮下的弹出菜单也被销毁.
//
// nIndex: 菜单条按钮索引.
func (m *MenuBar) I删除按钮(nIndex int) bool {
	return xc.XMenuBar_DeleteButton(m.I句柄, nIndex)
}

// 菜单条_启用自动宽度, 根据内容自动调整宽度.
//
// bEnable: 是否启用.
func (m *MenuBar) I启用自动宽度(bEnable bool) int {
	return xc.XMenuBar_EnableAutoWidth(m.I句柄, bEnable)
}

// 菜单条_取菜单按钮. 返回按钮句柄.
//
// nIndex: 菜单条按钮索引.
func (m *MenuBar) I取菜单按钮(nIndex int) bool {
	return xc.XMenuBar_GetButton(m.I句柄, nIndex)
}
