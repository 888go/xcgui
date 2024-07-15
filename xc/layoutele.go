package 炫彩基类

import "github.com/888go/xcgui/common"

// 布局_创建, 创建布局元素, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.

// ff:布局_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X布局_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xLayout_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 布局_创建扩展, 创建布局元素, 返回元素句柄.
//
// hParent: 父为窗口句柄或元素句柄.

// ff:布局_创建EX
// hParent:父窗口句柄或元素句柄
func X布局_创建EX(父窗口句柄或元素句柄 int) int {
	r, _, _ := xLayout_CreateEx.Call(uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 布局_判断启用, 是否已经启用布局功能.
//
// hEle: 元素句柄.

// ff:布局_判断启用
// hEle:元素句柄
func X布局_判断启用(元素句柄 int) bool {
	r, _, _ := xLayout_IsEnableLayout.Call(uintptr(元素句柄))
	return r != 0
}

// 布局_启用, 启用布局功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:布局_启用
// bEnable:是否启用
// hEle:元素句柄
func X布局_启用(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xLayout_EnableLayout.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 布局_显示布局边界, 显示布局边界.
//
// hEle: 元素句柄.
//
// bEnable: 是否显示.

// ff:布局_显示布局边界
// bEnable:是否显示
// hEle:元素句柄
func X布局_显示布局边界(元素句柄 int, 是否显示 bool) int {
	r, _, _ := xLayout_ShowLayoutFrame.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 布局_取内宽度, 获取宽度,不包含内边距大小.
//
// hEle: 元素句柄.

// ff:布局_取内宽度
// hEle:元素句柄
func X布局_取内宽度(元素句柄 int) int {
	r, _, _ := xLayout_GetWidthIn.Call(uintptr(元素句柄))
	return int(r)
}

// 布局_取内高度, 获取高度,不包含内边距大小.
//
// hEle: 元素句柄.

// ff:布局_取内高度
// hEle:元素句柄
func X布局_取内高度(元素句柄 int) int {
	r, _, _ := xLayout_GetHeightIn.Call(uintptr(元素句柄))
	return int(r)
}
