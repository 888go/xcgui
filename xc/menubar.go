package 炫彩基类

import "github.com/888go/xcgui/common"

// 菜单条_创建, 创建菜单条元素; 如果指定了父为窗口, 默认调用XWnd_AddMenuBar()函数, 将菜单条添加到窗口非客户区. 返回元素句柄.
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

// ff:菜单条_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X菜单条_创建(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xMenuBar_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 菜单条_添加按钮, 添加弹出菜单按钮, 返回菜单按钮索引.
//
// hEle: 元素句柄.
//
// pText: 文本内容.

// ff:菜单条_添加按钮
// pText:文本内容
// hEle:元素句柄
func X菜单条_添加按钮(元素句柄 int, 文本内容 string) int32 {
	r, _, _ := xMenuBar_AddButton.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本内容))
	return int32(r)
}

// 菜单条_置按钮高度, 根据内容自动调整宽度. (已废弃)请使用内填充限制高度.
//
// hEle: 元素句柄.
//
// height: 高度.

// ff:菜单条_置按钮高度
// height:高度
// hEle:元素句柄
func X菜单条_置按钮高度(元素句柄 int, 高度 int32) {
	xMenuBar_SetButtonHeight.Call(uintptr(元素句柄), uintptr(高度))
}

// 菜单条_取菜单, 返回菜单句柄.
//
// hEle: 元素句柄.
//
// nIndex: 菜单条上菜单按钮的索引.

// ff:菜单条_取菜单
// nIndex:菜单条按钮索引
// hEle:元素句柄
func X菜单条_取菜单(元素句柄 int, 菜单条按钮索引 int32) int {
	r, _, _ := xMenuBar_GetMenu.Call(uintptr(元素句柄), uintptr(菜单条按钮索引))
	return int(r)
}

// 菜单条_删除按钮, 删除菜单条上的菜单按钮, 同时该按钮下的弹出菜单也被销毁.
//
// hEle: 元素句柄.
//
// nIndex: 菜单条按钮索引.

// ff:菜单条_删除按钮
// nIndex:菜单条按钮索引
// hEle:元素句柄
func X菜单条_删除按钮(元素句柄 int, 菜单条按钮索引 int32) bool {
	r, _, _ := xMenuBar_DeleteButton.Call(uintptr(元素句柄), uintptr(菜单条按钮索引))
	return r != 0
}

// 菜单条_启用自动宽度, 根据内容自动调整宽度.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:菜单条_启用自动宽度
// bEnable:是否启用
// hEle:元素句柄
func X菜单条_启用自动宽度(元素句柄 int, 是否启用 bool) {
	xMenuBar_EnableAutoWidth.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
}

// 菜单条_取菜单按钮. 返回按钮句柄.
//
// hEle: 元素句柄.
//
// nIndex: 菜单条按钮索引.

// ff:菜单条_取菜单按钮
// nIndex:菜单条按钮索引
// hEle:元素句柄
func X菜单条_取菜单按钮(元素句柄 int, 菜单条按钮索引 int32) bool {
	r, _, _ := xMenuBar_GetButton.Call(uintptr(元素句柄), uintptr(菜单条按钮索引))
	return r != 0
}

// 菜单条_取选择项. 菜单条当前选择项, 当前弹出菜单的按钮索引.
//
// hEle: 元素句柄.

// ff:菜单条_取选择项
// hEle:元素句柄
func X菜单条_取选择项(元素句柄 int) int32 {
	r, _, _ := xMenuBar_GetSelect.Call(uintptr(元素句柄))
	return int32(r)
}
