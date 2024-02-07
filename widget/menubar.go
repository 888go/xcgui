package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
)

// MenuBar 菜单条.
type MenuBar struct {
	Element
}

// 菜单条_创建, 创建菜单条元素; 如果指定了父为窗口, 默认调用XWnd_AddMenuBar()函数, 将菜单条添加到窗口非客户区.
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
func X创建菜单条(元素x坐标 int32, 元素y坐标 int32, 宽度 int32, 高度 int32, 父窗口句柄或元素句柄 int) *MenuBar {
	p := &MenuBar{}
	p.X设置句柄(炫彩基类.X菜单条_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建菜单条并按句柄(handle int) *MenuBar {
	p := &MenuBar{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建菜单条并按名称(name string) *MenuBar {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &MenuBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建菜单条并按UID(nUID int) *MenuBar {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &MenuBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建菜单条并按UID名称(name string) *MenuBar {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &MenuBar{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 菜单条_添加按钮, 添加弹出菜单按钮, 返回菜单按钮索引.
//
// pText: 文本内容.
func (m *MenuBar) X添加按钮(文本内容 string) int32 {
	return 炫彩基类.X菜单条_添加按钮(m.Handle, 文本内容)
}

// 菜单条_置按钮高度, 根据内容自动调整宽度.
//
// height: 高度.
func (m *MenuBar) X置按钮高度(高度 int32) {
	炫彩基类.X菜单条_置按钮高度(m.Handle, 高度)
}

// 菜单条_取菜单, 返回菜单句柄.
//
// nIndex: 菜单条上菜单按钮的索引.
func (m *MenuBar) X取菜单(按钮索引 int32) int {
	return 炫彩基类.X菜单条_取菜单(m.Handle, 按钮索引)
}

// 菜单条_删除按钮, 删除菜单条上的菜单按钮, 同时该按钮下的弹出菜单也被销毁.
//
// nIndex: 菜单条按钮索引.
func (m *MenuBar) X删除按钮(按钮索引 int32) bool {
	return 炫彩基类.X菜单条_删除按钮(m.Handle, 按钮索引)
}

// 菜单条_启用自动宽度, 根据内容自动调整宽度. (已废弃)请使用内填充限制高度
//
// bEnable: 是否启用.
func (m *MenuBar) X启用自动宽度(是否启用 bool) {
	炫彩基类.X菜单条_启用自动宽度(m.Handle, 是否启用)
}

// 菜单条_取菜单按钮. 返回按钮句柄.
//
// nIndex: 菜单条按钮索引.
func (m *MenuBar) X取菜单按钮(按钮索引 int32) bool {
	return 炫彩基类.X菜单条_取菜单按钮(m.Handle, 按钮索引)
}

// 菜单条_取选择项. 菜单条当前选择项, 当前弹出菜单的按钮索引.
func (m *MenuBar) X取选择项() int32 {
	return 炫彩基类.X菜单条_取选择项(m.Handle)
}
