package 炫彩基类

import (
	"github.com/888go/xcgui/common"
)

// 数据适配器列表视_创建, 创建列表视元素数据适配器, 返回数据适配器句柄.
func X数据适配器列表视_创建() int {
	r, _, _ := xAdListView_Create.Call()
	return int(r)
}

// 数据适配器列表视_组添加列, 组操作, 添加数据列, 返回列索引.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func X数据适配器列表视_组添加列(数据适配器句柄 int, 字段称 string) int {
	r, _, _ := xAdListView_Group_AddColumn.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器列表视_添加组文本, 组操作, 添加组, 数据默认填充到第一列, 返回组索引.
//
// hAdapter: 数据适配器句柄.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加组文本(数据适配器句柄 int, 值 string, 插入位置 int) int {
	r, _, _ := xAdListView_Group_AddItemText.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_添加组文本扩展, 组操作, 添加组, 数据填充指定列, 返回组索引.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加组文本EX(数据适配器句柄 int, 字段称 string, 值 string, 插入位置 int) int {
	r, _, _ := xAdListView_Group_AddItemTextEx.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_添加组图片, 组操作, 添加组, 数据默认填充第一列.
//
// hAdapter: 数据适配器句柄.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加组图片(数据适配器句柄 int, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xAdListView_Group_AddItemImage.Call(uintptr(数据适配器句柄), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_添加组图片扩展, 组操作, 添加组, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加组图片EX(数据适配器句柄 int, 字段称 string, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xAdListView_Group_AddItemImageEx.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_组设置文本, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func X数据适配器列表视_组设置文本(数据适配器句柄 int, 组索引 int, 列索引 int, 值 string) bool {
	r, _, _ := xAdListView_Group_SetText.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器列表视_组设置文本扩展, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func X数据适配器列表视_组设置文本EX(数据适配器句柄 int, 组索引 int, 字段名 string, 值 string) bool {
	r, _, _ := xAdListView_Group_SetTextEx.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器列表视_组设置图片, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func X数据适配器列表视_组设置图片(数据适配器句柄 int, 组索引 int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xAdListView_Group_SetImage.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 数据适配器列表视_组设置图片扩展, 组操作, 设置指定项数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func X数据适配器列表视_组设置图片EX(数据适配器句柄 int, 组索引 int, 字段名 string, 图片句柄 int) bool {
	r, _, _ := xAdListView_Group_SetImageEx.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名), uintptr(图片句柄))
	return r != 0
}

// 数据适配器列表视_项添加列, 项操作, 添加列.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
func X数据适配器列表视_项添加列(数据适配器句柄 int, 字段称 string) int {
	r, _, _ := xAdListView_Item_AddColumn.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器列表视_取组数量, 组操作, 获取组数量.
//
// hAdapter: 数据适配器句柄.
func X数据适配器列表视_取组数量(数据适配器句柄 int) int {
	r, _, _ := xAdListView_Group_GetCount.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器列表视_取组中项数量, 项操作, 获取指定组中项数量.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
func X数据适配器列表视_取组中项数量(数据适配器句柄 int, 组索引 int) int {
	r, _, _ := xAdListView_Item_GetCount.Call(uintptr(数据适配器句柄), uintptr(组索引))
	return int(r)
}

// 数据适配器列表视_添加项文本, 项操作, 添加项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加项文本(数据适配器句柄 int, 组索引 int, 值 string, 插入位置 int) int {
	r, _, _ := xAdListView_Item_AddItemText.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_添加项文本扩展, 项操作, 添加项, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加项文本EX(数据适配器句柄 int, 组索引 int, 字段称 string, 值 string, 插入位置 int) int {
	r, _, _ := xAdListView_Item_AddItemTextEx.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_添加项图片, 项操作, 添加项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加项图片(数据适配器句柄 int, 组索引 int, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xAdListView_Item_AddItemImage.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_添加项图片扩展, 项操作, 添加项, 填充指定列数据.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func X数据适配器列表视_添加项图片EX(数据适配器句柄 int, 组索引 int, 字段称 string, 图片句柄 int, 插入位置 int) int {
	r, _, _ := xAdListView_Item_AddItemImageEx.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄), uintptr(插入位置))
	return int(r)
}

// 数据适配器列表视_组删除项, 删除组, 自动删除子项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
func X数据适配器列表视_组删除项(数据适配器句柄 int, 组索引 int) bool {
	r, _, _ := xAdListView_Group_DeleteItem.Call(uintptr(数据适配器句柄), uintptr(组索引))
	return r != 0
}

// 数据适配器列表视_删除全部子项, 删除指定组的所有子项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
func X数据适配器列表视_删除全部子项(数据适配器句柄 int, 组索引 int) int {
	r, _, _ := xAdListView_Group_DeleteAllChildItem.Call(uintptr(数据适配器句柄), uintptr(组索引))
	return int(r)
}

// 数据适配器列表视_删除项, 删除项.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func X数据适配器列表视_删除项(数据适配器句柄 int, 组索引 int, 项索引 int) bool {
	r, _, _ := xAdListView_Item_DeleteItem.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引))
	return r != 0
}

// 数据适配器列表视_删除全部, 删除所有(组,项,列).
//
// hAdapter: 数据适配器句柄.
func X数据适配器列表视_删除全部(数据适配器句柄 int) int {
	r, _, _ := xAdListView_DeleteAll.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器列表视_删除全部组, 删除所有的组.
//
// hAdapter: 数据适配器句柄.
func X数据适配器列表视_删除全部组(数据适配器句柄 int) int {
	r, _, _ := xAdListView_DeleteAllGroup.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器列表视_删除全部项, 删除所有的项.
//
// hAdapter: 数据适配器句柄.
func X数据适配器列表视_删除全部项(数据适配器句柄 int) int {
	r, _, _ := xAdListView_DeleteAllItem.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器列表视_组删除列, 删除组的列.
//
// hAdapter: 数据适配器句柄.
//
// iColumn: 列索引.
func X数据适配器列表视_组删除列(数据适配器句柄 int, 列索引 int) int {
	r, _, _ := xAdListView_DeleteColumnGroup.Call(uintptr(数据适配器句柄), uintptr(列索引))
	return int(r)
}

// 数据适配器列表视_项删除列, 删除项的列.
//
// hAdapter: 数据适配器句柄.
//
// iColumn: 列索引.
func X数据适配器列表视_项删除列(数据适配器句柄 int, 列索引 int) int {
	r, _, _ := xAdListView_DeleteColumnItem.Call(uintptr(数据适配器句柄), uintptr(列索引))
	return int(r)
}

// 数据适配器列表视_项获取文本扩展, 项操作, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func X数据适配器列表视_项获取文本EX(数据适配器句柄 int, 组索引 int, 项索引 int, 字段称 string) string {
	r, _, _ := xAdListView_Item_GetTextEx.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器列表视_项获取图片扩展, 项操作, 获取项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func X数据适配器列表视_项获取图片EX(数据适配器句柄 int, 组索引 int, 项索引 int, 字段称 string) int {
	r, _, _ := xAdListView_Item_GetImageEx.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器列表视_项置文本, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func X数据适配器列表视_项置文本(数据适配器句柄 int, 组索引 int, 项索引 int, 列索引 int, 值 string) bool {
	r, _, _ := xAdListView_Item_SetText.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器列表视_项置文本扩展, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func X数据适配器列表视_项置文本EX(数据适配器句柄 int, 组索引 int, 项索引 int, 字段称 string, 值 string) bool {
	r, _, _ := xAdListView_Item_SetTextEx.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器列表视_项置图片, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func X数据适配器列表视_项置图片(数据适配器句柄 int, 组索引 int, 项索引 int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xAdListView_Item_SetImage.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 数据适配器列表视_项置图片扩展, 项操作, 数据填充指定列.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func X数据适配器列表视_项置图片EX(数据适配器句柄 int, 组索引 int, 项索引 int, 字段称 string, 图片句柄 int) bool {
	r, _, _ := xAdListView_Item_SetImageEx.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄))
	return r != 0
}

// 数据适配器列表视_组取文本, 返回文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func X数据适配器列表视_组取文本(数据适配器句柄 int, 组索引 int, 列索引 int) string {
	r, _, _ := xAdListView_Group_GetText.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器列表视_组取文本扩展, 返回文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func X数据适配器列表视_组取文本EX(数据适配器句柄 int, 组索引 int, 字段名称 string) string {
	r, _, _ := xAdListView_Group_GetTextEx.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名称))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器列表视_组取图片, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func X数据适配器列表视_组取图片(数据适配器句柄 int, 组索引 int, 列索引 int) int {
	r, _, _ := xAdListView_Group_GetImage.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(列索引))
	return int(r)
}

// 数据适配器列表视_组取图片扩展, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func X数据适配器列表视_组取图片EX(数据适配器句柄 int, 组索引 int, 字段名称 string) int {
	r, _, _ := xAdListView_Group_GetImageEx.Call(uintptr(数据适配器句柄), uintptr(组索引), 炫彩工具类.StrPtr(字段名称))
	return int(r)
}

// 数据适配器列表视_项取文本. 项操作, 返回项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func X数据适配器列表视_项取文本(数据适配器句柄 int, 组索引 int, 项索引 int, 列索引 int) string {
	r, _, _ := xAdListView_Item_GetText.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器列表视_项取图片. 项操作, 返回项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func X数据适配器列表视_项取图片(数据适配器句柄 int, 组索引 int, 项索引 int, 列索引 int) int {
	r, _, _ := xAdListView_Item_GetImage.Call(uintptr(数据适配器句柄), uintptr(组索引), uintptr(项索引), uintptr(列索引))
	return int(r)
}
