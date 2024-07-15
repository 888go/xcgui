package 炫彩基类

import "github.com/888go/xcgui/common"

// 数据适配器树_创建, 创建树元素数据适配器, 返回数据适配器句柄.

// ff:数据适配器树_创建
func X数据适配器树_创建() int {
	r, _, _ := xAdTree_Create.Call()
	return int(r)
}

// 数据适配器树_添加列, 添加列, 返回索引值.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.

// ff:数据适配器树_添加列
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器树_添加列(数据适配器句柄 int, 字段称 string) int {
	r, _, _ := xAdTree_AddColumn.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器树_置列, 设置列, 返回列数量.
//
// hAdapter: 数据适配器句柄.
//
// pColName: 列名, 列名, 多个列名用逗号分开.

// ff:数据适配器树_置列
// pColName:列名
// hAdapter:数据适配器句柄
func X数据适配器树_置列(数据适配器句柄 int, 列名 string) int {
	r, _, _ := xAdTree_SetColumn.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(列名))
	return int(r)
}

// 数据适配器树_插入项文本, 插入项, 数据填充到第一列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.

// ff:数据适配器树_插入项文本
// insertID:插入位置ID
// nParentID:父ID
// pValue:值
// hAdapter:数据适配器句柄
func X数据适配器树_插入项文本(数据适配器句柄 int, 值 string, 父ID int, 插入位置ID int) int {
	r, _, _ := xAdTree_InsertItemText.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(值), uintptr(父ID), uintptr(插入位置ID))
	return int(r)
}

// 数据适配器树_插入项文本扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.

// ff:数据适配器树_插入项文本EX
// insertID:插入位置ID
// nParentID:父ID
// pValue:值
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器树_插入项文本EX(数据适配器句柄 int, 字段称 string, 值 string, 父ID int, 插入位置ID int) int {
	r, _, _ := xAdTree_InsertItemTextEx.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值), uintptr(父ID), uintptr(插入位置ID))
	return int(r)
}

// 数据适配器树_插入项图片, 插入项, 数据填充到第一列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.

// ff:数据适配器树_插入项图片
// insertID:插入位置ID
// nParentID:父ID
// hImage:图片句柄
// hAdapter:数据适配器句柄
func X数据适配器树_插入项图片(数据适配器句柄 int, 图片句柄 int, 父ID int, 插入位置ID int) int {
	r, _, _ := xAdTree_InsertItemImage.Call(uintptr(数据适配器句柄), uintptr(图片句柄), uintptr(父ID), uintptr(插入位置ID))
	return int(r)
}

// 数据适配器树_插入项图片扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.

// ff:数据适配器树_插入项图片EX
// insertID:插入位置ID
// nParentID:父ID
// hImage:图片句柄
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器树_插入项图片EX(数据适配器句柄 int, 字段称 string, 图片句柄 int, 父ID int, 插入位置ID int) int {
	r, _, _ := xAdTree_InsertItemImageEx.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄), uintptr(父ID), uintptr(插入位置ID))
	return int(r)
}

// 数据适配器树_取项数量, 获取项数量.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器树_取项数量
// hAdapter:数据适配器句柄
func X数据适配器树_取项数量(数据适配器句柄 int) int {
	r, _, _ := xAdTree_GetCount.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器树_取列数量, 获取列数量.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器树_取列数量
// hAdapter:数据适配器句柄
func X数据适配器树_取列数量(数据适配器句柄 int) int {
	r, _, _ := xAdTree_GetCountColumn.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器树_置项文本, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// pValue: 值.

// ff:数据适配器树_置项文本
// pValue:值
// iColumn:列索引
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_置项文本(数据适配器句柄 int, 项ID int, 列索引 int, 值 string) bool {
	r, _, _ := xAdTree_SetItemText.Call(uintptr(数据适配器句柄), uintptr(项ID), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器树_置项文本扩展, 设置项文件内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.
//
// pValue: 值.

// ff:数据适配器树_置项文本EX
// pValue:值
// pName:字段称
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_置项文本EX(数据适配器句柄 int, 项ID int, 字段称 string, 值 string) bool {
	r, _, _ := xAdTree_SetItemTextEx.Call(uintptr(数据适配器句柄), uintptr(项ID), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器树_置项图片, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.

// ff:数据适配器树_置项图片
// hImage:图片句柄
// iColumn:列索引
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_置项图片(数据适配器句柄 int, 项ID int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xAdTree_SetItemImage.Call(uintptr(数据适配器句柄), uintptr(项ID), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 数据适配器树_置项图片扩展, 设置项内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:数据适配器树_置项图片EX
// hImage:图片句柄
// pName:字段称
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_置项图片EX(数据适配器句柄 int, 项ID int, 字段称 string, 图片句柄 int) bool {
	r, _, _ := xAdTree_SetItemImageEx.Call(uintptr(数据适配器句柄), uintptr(项ID), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄))
	return r != 0
}

// 数据适配器树_取项文本, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.

// ff:数据适配器树_取项文本
// iColumn:列索引
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_取项文本(数据适配器句柄 int, 项ID int, 列索引 int) string {
	r, _, _ := xAdTree_GetItemText.Call(uintptr(数据适配器句柄), uintptr(项ID), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器树_取项文本扩展, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.

// ff:数据适配器树_取项文本EX
// pName:字段称
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_取项文本EX(数据适配器句柄 int, 项ID int, 字段称 string) string {
	r, _, _ := xAdTree_GetItemTextEx.Call(uintptr(数据适配器句柄), uintptr(项ID), 炫彩工具类.StrPtr(字段称))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器树_取项图片, 获取项内容, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.

// ff:数据适配器树_取项图片
// iColumn:列索引
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_取项图片(数据适配器句柄 int, 项ID int, 列索引 int) int {
	r, _, _ := xAdTree_GetItemImage.Call(uintptr(数据适配器句柄), uintptr(项ID), uintptr(列索引))
	return int(r)
}

// 数据适配器树_取项图片扩展, 获取项内容, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.
//
// pName: 字段称.

// ff:数据适配器树_取项图片EX
// pName:字段称
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_取项图片EX(数据适配器句柄 int, 项ID int, 字段称 string) int {
	r, _, _ := xAdTree_GetItemImageEx.Call(uintptr(数据适配器句柄), uintptr(项ID), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器树_删除项, 删除项.
//
// hAdapter: 数据适配器句柄.
//
// nID: 项ID.

// ff:数据适配器树_删除项
// nID:项ID
// hAdapter:数据适配器句柄
func X数据适配器树_删除项(数据适配器句柄 int, 项ID int) bool {
	r, _, _ := xAdTree_DeleteItem.Call(uintptr(数据适配器句柄), uintptr(项ID))
	return r != 0
}

// 数据适配器树_删除项全部, 删除所有项.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器树_删除项全部
// hAdapter:数据适配器句柄
func X数据适配器树_删除项全部(数据适配器句柄 int) int {
	r, _, _ := xAdTree_DeleteItemAll.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器树_删除列全部, 删除所有列, 并且清空数据.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器树_删除列全部
// hAdapter:数据适配器句柄
func X数据适配器树_删除列全部(数据适配器句柄 int) int {
	r, _, _ := xAdTree_DeleteColumnAll.Call(uintptr(数据适配器句柄))
	return int(r)
}
