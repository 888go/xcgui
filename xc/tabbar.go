package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

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

// ff:TAB条_创建
// hParent:父窗口句柄或元素句柄
// cy:高度
// cx:宽度
// y:元素y坐标
// x:元素x坐标
func TAB条_创建(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xTabBar_Create.Call(uintptr(元素x坐标), uintptr(元素y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// TAB条_添加标签, 添加一个标签, 返回标签索引.
//
// hEle: 元素句柄.
//
// pName: 标签文本内容.

// ff:TAB条_添加标签
// pName:标签文本内容
// hEle:元素句柄
func TAB条_添加标签(元素句柄 int, 标签文本内容 string) int {
	r, _, _ := xTabBar_AddLabel.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(标签文本内容))
	return int(r)
}

// TAB条插入_标签, 插入一个标签, 返回标签索引.
//
// hEle: 元素句柄.
//
// index: 插入位置.
//
// pName: 标签文本内容.

// ff:TAB条插入_标签
// pName:标签文本内容
// index:插入位置
// hEle:元素句柄
func TAB条插入_标签(元素句柄 int, 插入位置 int, 标签文本内容 string) int {
	r, _, _ := xTabBar_InsertLabel.Call(uintptr(元素句柄), uintptr(插入位置), 炫彩工具类.StrPtr(标签文本内容))
	return int(r)
}

// TAB条_移动标签.
//
// hEle: 元素句柄.
//
// iSrc: 源位置索引.
//
// iDest: 目标位置索引.

// ff:TAB条_移动标签
// iDest:目标位置索引
// iSrc:源位置索引
// hEle:元素句柄
func TAB条_移动标签(元素句柄 int, 源位置索引 int, 目标位置索引 int) bool {
	r, _, _ := xTabBar_MoveLabel.Call(uintptr(元素句柄), uintptr(源位置索引), uintptr(目标位置索引))
	return r != 0
}

// TAB条_删除标签, 删除一个标签.
//
// hEle: 元素句柄.
//
// index: 位置索引.

// ff:TAB条_删除标签
// index:位置索引
// hEle:元素句柄
func TAB条_删除标签(元素句柄 int, 位置索引 int) bool {
	r, _, _ := xTabBar_DeleteLabel.Call(uintptr(元素句柄), uintptr(位置索引))
	return r != 0
}

// TAB条_删除全部, 删除所有标签.
//
// hEle: 元素句柄.

// ff:TAB条_删除全部
// hEle:元素句柄
func TAB条_删除全部(元素句柄 int) int {
	r, _, _ := xTabBar_DeleteLabelAll.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取标签, 获取标签按钮句柄.
//
// hEle: 元素句柄.
//
// index: 位置索引.

// ff:TAB条_取标签
// index:位置索引
// hEle:元素句柄
func TAB条_取标签(元素句柄 int, 位置索引 int) int {
	r, _, _ := xTabBar_GetLabel.Call(uintptr(元素句柄), uintptr(位置索引))
	return int(r)
}

// TAB条_取标签上的关闭按钮, 获取标签上关闭按钮句柄.
//
// hEle: 元素句柄.
//
// index: 位置索引.

// ff:TAB条_取标签上的关闭按钮
// index:位置索引
// hEle:元素句柄
func TAB条_取标签上的关闭按钮(元素句柄 int, 位置索引 int) int {
	r, _, _ := xTabBar_GetLabelClose.Call(uintptr(元素句柄), uintptr(位置索引))
	return int(r)
}

// TAB条_取左滚动按钮, 获取左滚动按钮句柄.
//
// hEle: 元素句柄.

// ff:TAB条_取左滚动按钮
// hEle:元素句柄
func TAB条_取左滚动按钮(元素句柄 int) int {
	r, _, _ := xTabBar_GetButtonLeft.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取右滚动按钮, 获取右滚动按钮句柄.
//
// hEle: 元素句柄.

// ff:TAB条_取右滚动按钮
// hEle:元素句柄
func TAB条_取右滚动按钮(元素句柄 int) int {
	r, _, _ := xTabBar_GetButtonRight.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取下拉菜单按钮句柄.
//
// hEle: 元素句柄.

// ff:TAB条_取下拉菜单按钮句柄
// hEle:元素句柄
func TAB条_取下拉菜单按钮句柄(元素句柄 int) int {
	r, _, _ := xTabBar_GetButtonDropMenu.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取当前选择, 获取选择的标签索引.
//
// hEle: 元素句柄.

// ff:TAB条_取当前选择
// hEle:元素句柄
func TAB条_取当前选择(元素句柄 int) int {
	r, _, _ := xTabBar_GetSelect.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取间隔, 获取标签间距, 0没有间距.
//
// hEle: 元素句柄.

// ff:TAB条_取间隔
// hEle:元素句柄
func TAB条_取间隔(元素句柄 int) int {
	r, _, _ := xTabBar_GetLabelSpacing.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取标签数量, 获取标签项数量.
//
// hEle: 元素句柄.

// ff:TAB条_取标签数量
// hEle:元素句柄
func TAB条_取标签数量(元素句柄 int) int {
	r, _, _ := xTabBar_GetLabelCount.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_取标签位置索引, 获取标签按钮位置索引, 成功返回索引值, 否则返回 XC_ID_ERROR.
//
// hEle: 元素句柄.
//
// hLabel: 标签按钮句柄.

// ff:TAB条_取标签位置索引
// hLabel:标签按钮句柄
// hEle:元素句柄
func TAB条_取标签位置索引(元素句柄 int, 标签按钮句柄 int) int {
	r, _, _ := xTabBar_GetindexByEle.Call(uintptr(元素句柄), uintptr(标签按钮句柄))
	return int(r)
}

// TAB条_置间隔, 设置标签间距, 0没有间距.
//
// hEle: 元素句柄.
//
// spacing: 标签间隔大小.

// ff:TAB条_置间隔
// spacing:标签间隔大小
// hEle:元素句柄
func TAB条_置间隔(元素句柄 int, 标签间隔大小 int) int {
	r, _, _ := xTabBar_SetLabelSpacing.Call(uintptr(元素句柄), uintptr(标签间隔大小))
	return int(r)
}

// TAB条_置边距, 设置内容与边框的间隔大小.
//
// hEle: 元素句柄.
//
// left: 左边间隔大小.
//
// top: 上边间隔大小.
//
// right: 右边间隔大小.
//
// bottom: 下边间隔大小.

// ff:TAB条_置边距
// bottom:下
// right:右
// top:上
// left:左
// hEle:元素句柄
func TAB条_置边距(元素句柄 int, 左 int, 上 int, 右 int, 下 int) int {
	r, _, _ := xTabBar_SetPadding.Call(uintptr(元素句柄), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return int(r)
}

// TAB条_置选择, 设置选择标签.
//
// hEle: 元素句柄.
//
// index: 标签位置索引.

// ff:TAB条_置选择
// index:标签位置索引
// hEle:元素句柄
func TAB条_置选择(元素句柄 int, 标签位置索引 int) int {
	r, _, _ := xTabBar_SetSelect.Call(uintptr(元素句柄), uintptr(标签位置索引))
	return int(r)
}

// TAB条_左滚动, 左按钮滚动.
//
// hEle: 元素句柄.

// ff:TAB条_左滚动
// hEle:元素句柄
func TAB条_左滚动(元素句柄 int) int {
	r, _, _ := xTabBar_SetUp.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_右滚动, 右按钮滚动.
//
// hEle: 元素句柄.

// ff:TAB条_右滚动
// hEle:元素句柄
func TAB条_右滚动(元素句柄 int) int {
	r, _, _ := xTabBar_SetDown.Call(uintptr(元素句柄))
	return int(r)
}

// TAB条_启用平铺, 平铺标签, 每个标签显示相同大小.
//
// hEle: 元素句柄.
//
// bTile: 是否启用.

// ff:TAB条_启用平铺
// bTile:是否启用
// hEle:元素句柄
func TAB条_启用平铺(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xTabBar_EnableTile.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// TAB条_启用下拉菜单按钮.
//
// hEle: 元素句柄.
//
// bEnable:.

// ff:TAB条_启用下拉菜单按钮
// bEnable:
// hEle:元素句柄
func TAB条_启用下拉菜单按钮(元素句柄 int, bEnable bool) int {
	r, _, _ := xTabBar_EnableDropMenu.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(bEnable))
	return int(r)
}

// TAB条_启用标签带关闭按钮, 启用关闭标签功能.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.

// ff:TAB条_启用标签带关闭按钮
// bEnable:是否启用
// hEle:元素句柄
func TAB条_启用标签带关闭按钮(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xTabBar_EnableClose.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// TAB条_置关闭按钮大小, 设置关闭按钮大小.
//
// hEle: 元素句柄.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.

// ff:TAB条_置关闭按钮大小
// pSize:大小值
// hEle:元素句柄
func TAB条_置关闭按钮大小(元素句柄 int, 大小值 *SIZE) int {
	r, _, _ := xTabBar_SetCloseSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(大小值)))
	return int(r)
}

// TAB条_置滚动按钮大小.
//
// hEle: 元素句柄.
//
// pSize: 大小值, 宽度和高度可以为-1, -1代表默认值.

// ff:TAB条_置滚动按钮大小
// pSize:大小值
// hEle:元素句柄
func TAB条_置滚动按钮大小(元素句柄 int, 大小值 *SIZE) int {
	r, _, _ := xTabBar_SetTurnButtonSize.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(大小值)))
	return int(r)
}

// TAB条_置指定标签固定宽度.
//
// hEle: 元素句柄.
//
// index: 索引.
//
// nWidth: 宽度, , 如果值为-1, 那么自动计算宽度.

// ff:TAB条_置指定标签固定宽度
// nWidth:宽度
// index:索引
// hEle:元素句柄
func TAB条_置指定标签固定宽度(元素句柄 int, 索引 int, 宽度 int) int {
	r, _, _ := xTabBar_SetLabelWidth.Call(uintptr(元素句柄), uintptr(索引), uintptr(宽度))
	return int(r)
}

// TAB条_显示标签, 显示或隐藏指定标签.
//
// hEle: 元素句柄.
//
// index: 标签索引.
//
// bShow: 是否显示.

// ff:TAB条_显示标签
// bShow:是否显示
// index:标签索引
// hEle:元素句柄
func TAB条_显示标签(元素句柄 int, 标签索引 int, 是否显示 bool) bool {
	r, _, _ := xTabBar_ShowLabel.Call(uintptr(元素句柄), uintptr(标签索引), 炫彩工具类.BoolPtr(是否显示))
	return r != 0
}
