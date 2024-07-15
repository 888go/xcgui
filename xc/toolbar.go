package 炫彩基类

import "github.com/888go/xcgui/common"

// 工具条_创建, 创建工具条元素, 返回元素句柄; 如果指定了父为窗口, 默认调用XWnd_AddToolBar()函数, 将工具条添加到窗口非客户区.
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

// ff:工具条_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func X工具条_创建(元素x坐标, 元素y坐标, 宽度, 高度 int32, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xToolBar_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 工具条_插入元素, 插入元素到工具条, 返回插入位置索引.
//
// hEle: 元素句柄.
//
// hNewEle: 将要插入的元素.
//
// index: 插入位置索引, (-1)插入末尾.

// ff:工具条_插入元素
// index:插入位置索引
// hNewEle:将要插入的元素
// hEle:元素句柄
func X工具条_插入元素(元素句柄 int, 将要插入的元素 int, 插入位置索引 int) int {
	r, _, _ := xToolBar_InsertEle.Call(uintptr(元素句柄), uintptr(将要插入的元素), uintptr(插入位置索引))
	return int(r)
}

// 工具条_插入分割栏, 插入分隔符到工具条, 返回插入位置索引.
//
// hEle: 元素句柄.
//
// index: 插入位置索引, (-1)插入末尾.
//
// color: ABGR 颜色.

// ff:工具条_插入分割栏
// color:ABGR颜色
// index:插入位置索引
// hEle:元素句柄
func X工具条_插入分割栏(元素句柄 int, 插入位置索引 int, ABGR颜色 int) int {
	r, _, _ := xToolBar_InsertSeparator.Call(uintptr(元素句柄), uintptr(插入位置索引), uintptr(ABGR颜色))
	return int(r)
}

// 工具条_启用下拉菜单, 启用下拉菜单, 显示隐藏的项.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:工具条_启用下拉菜单
// bEnable:是否启用
// hEle:元素句柄
func X工具条_启用下拉菜单(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xToolBar_EnableButtonMenu.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 工具条_取元素, 获取工具条上指定元素, 返回元素句柄.
//
// hEle: 元素句柄.
//
// index: 索引值.

// ff:工具条_取元素
// index:索引值
// hEle:元素句柄
func X工具条_取元素(元素句柄 int, 索引值 int) int {
	r, _, _ := xToolBar_GetEle.Call(uintptr(元素句柄), uintptr(索引值))
	return int(r)
}

// 工具条_取左滚动按钮, 获取左滚动按钮句柄.
//
// hEle: 元素句柄.

// ff:工具条_取左滚动按钮
// hEle:元素句柄
func X工具条_取左滚动按钮(元素句柄 int) int {
	r, _, _ := xToolBar_GetButtonLeft.Call(uintptr(元素句柄))
	return int(r)
}

// 工具条_取右滚动按钮, 获取右滚动按钮句柄.
//
// hEle: 元素句柄.

// ff:工具条_取右滚动按钮
// hEle:元素句柄
func X工具条_取右滚动按钮(元素句柄 int) int {
	r, _, _ := xToolBar_GetButtonRight.Call(uintptr(元素句柄))
	return int(r)
}

// 工具条_取菜单按钮, 获取菜单按钮句柄.
//
// hEle: 元素句柄.

// ff:工具条_取菜单按钮
// hEle:元素句柄
func X工具条_取菜单按钮(元素句柄 int) int {
	r, _, _ := xToolBar_GetButtonMenu.Call(uintptr(元素句柄))
	return int(r)
}

// 工具条_置间距, 设置对象之间的间距.
//
// hEle: 元素句柄.
//
// nSize: 间距大小.

// ff:工具条_置间距
// nSize:间距大小
// hEle:元素句柄
func X工具条_置间距(元素句柄 int, 间距大小 int) int {
	r, _, _ := xToolBar_SetSpace.Call(uintptr(元素句柄), uintptr(间距大小))
	return int(r)
}

// 工具条_删除元素, 删除元素, 并且销毁.
//
// hEle: 元素句柄.
//
// index: 索引值.

// ff:工具条_删除元素
// index:索引值
// hEle:元素句柄
func X工具条_删除元素(元素句柄 int, 索引值 int) int {
	r, _, _ := xToolBar_DeleteEle.Call(uintptr(元素句柄), uintptr(索引值))
	return int(r)
}

// 工具条_删除全部, 删除所有元素, 并且销毁.
//
// hEle: 元素句柄.

// ff:工具条_删除全部
// hEle:元素句柄
func X工具条_删除全部(元素句柄 int) int {
	r, _, _ := xToolBar_DeleteAllEle.Call(uintptr(元素句柄))
	return int(r)
}
