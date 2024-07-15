package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"

	"github.com/888go/xcgui/xcc"
)

// 数据适配器表_创建, 创建列表框元素数据适配器, 返回数据适配器句柄.

// ff:数据适配器表_创建
func X数据适配器表_创建() int {
	r, _, _ := xAdTable_Create.Call()
	return int(r)
}

// 数据适配器表_排序, 对内容进行排序.
//
// hAdapter: 数据适配器句柄.
//
// iColumn: 要排序的列索引。.
//
// bAscending: 是否按照升序方式排序.

// ff:数据适配器表_排序
// bAscending:是否按照升序排序
// iColumn:列索引
// hAdapter:数据适配器句柄
func X数据适配器表_排序(数据适配器句柄 int, 列索引 int, 是否按照升序排序 bool) int {
	r, _, _ := xAdTable_Sort.Call(uintptr(数据适配器句柄), uintptr(列索引), 炫彩工具类.BoolPtr(是否按照升序排序))
	return int(r)
}

// 数据适配器表_取项数据类型, 获取项数据类型, 返回: Adapter_Date_Type_.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.

// ff:数据适配器表_取项数据类型
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项数据类型(数据适配器句柄 int, 项索引 int, 列索引 int) 炫彩常量类.Adapter_Date_Type_ {
	r, _, _ := xAdTable_GetItemDataType.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引))
	return 炫彩常量类.Adapter_Date_Type_(r)
}

// 数据适配器表_取项数据类型扩展, 返回: Adapter_Date_Type_.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.

// ff:数据适配器表_取项数据类型EX
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项数据类型EX(数据适配器句柄 int, 项索引 int, 字段称 string) 炫彩常量类.Adapter_Date_Type_ {
	r, _, _ := xAdTable_GetItemDataTypeEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return 炫彩常量类.Adapter_Date_Type_(r)
}

// 数据适配器表_添加列, 添加数据列.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.

// ff:数据适配器表_添加列
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器表_添加列(数据适配器句柄 int, 字段称 string) int {
	r, _, _ := xAdTable_AddColumn.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器表_置列, 设置列, 返回列数量.
//
// hAdapter: 数据适配器句柄.
//
// pColName: 列名, 多个列名用逗号分开.

// ff:数据适配器表_置列
// pColName:列名
// hAdapter:数据适配器句柄
func X数据适配器表_置列(数据适配器句柄 int, 列名 string) int {
	r, _, _ := xAdTable_SetColumn.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(列名))
	return int(r)
}

// 数据适配器表_添加项文本, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// hAdapter: 数据适配器句柄.
//
// pValue: 值.

// ff:数据适配器表_添加项文本
// pValue:值
// hAdapter:数据适配器句柄
func X数据适配器表_添加项文本(数据适配器句柄 int, 值 string) int {
	r, _, _ := xAdTable_AddItemText.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 数据适配器表_添加项文本扩展, 添加数据项, 填充指定列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.

// ff:数据适配器表_添加项文本EX
// pValue:值
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器表_添加项文本EX(数据适配器句柄 int, 字段称 string, 值 string) int {
	r, _, _ := xAdTable_AddItemTextEx.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 数据适配器表_添加项图片, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// hAdapter: 数据适配器句柄.
//
// hImage: 图片句柄.

// ff:数据适配器表_添加项图片
// hImage:图片句柄
// hAdapter:数据适配器句柄
func X数据适配器表_添加项图片(数据适配器句柄 int, 图片句柄 int) int {
	r, _, _ := xAdTable_AddItemImage.Call(uintptr(数据适配器句柄), uintptr(图片句柄))
	return int(r)
}

// 数据适配器表_添加项图片扩展, 添加数据项, 并填充指定列数据.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:数据适配器表_添加项图片EX
// hImage:图片句柄
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器表_添加项图片EX(数据适配器句柄 int, 字段称 string, 图片句柄 int) int {
	r, _, _ := xAdTable_AddItemImageEx.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄))
	return int(r)
}

// 数据适配器表_插入项文本, 插入数据项, 填充第一列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// pValue: 值.

// ff:数据适配器表_插入项文本
// pValue:值
// iItem:插入位置索引
// hAdapter:数据适配器句柄
func X数据适配器表_插入项文本(数据适配器句柄 int, 插入位置索引 int, 值 string) int {
	r, _, _ := xAdTable_InsertItemText.Call(uintptr(数据适配器句柄), uintptr(插入位置索引), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 数据适配器表_插入项文本扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// pValue: 值.

// ff:数据适配器表_插入项文本EX
// pValue:值
// pName:字段称
// iItem:插入位置索引
// hAdapter:数据适配器句柄
func X数据适配器表_插入项文本EX(数据适配器句柄 int, 插入位置索引 int, 字段称 string, 值 string) int {
	r, _, _ := xAdTable_InsertItemTextEx.Call(uintptr(数据适配器句柄), uintptr(插入位置索引), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return int(r)
}

// 数据适配器表_插入项图片, 插入数据项, 填充第一列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// hImage: 图片句柄.

// ff:数据适配器表_插入项图片
// hImage:图片句柄
// iItem:插入位置索引
// hAdapter:数据适配器句柄
func X数据适配器表_插入项图片(数据适配器句柄 int, 插入位置索引 int, 图片句柄 int) int {
	r, _, _ := xAdTable_InsertItemImage.Call(uintptr(数据适配器句柄), uintptr(插入位置索引), uintptr(图片句柄))
	return int(r)
}

// 数据适配器表_插入项图片扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:数据适配器表_插入项图片EX
// hImage:图片句柄
// pName:字段称
// iItem:插入位置索引
// hAdapter:数据适配器句柄
func X数据适配器表_插入项图片EX(数据适配器句柄 int, 插入位置索引 int, 字段称 string, 图片句柄 int) int {
	r, _, _ := xAdTable_InsertItemImageEx.Call(uintptr(数据适配器句柄), uintptr(插入位置索引), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄))
	return int(r)
}

// 数据适配器表_置项文本, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.

// ff:数据适配器表_置项文本
// pValue:值
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项文本(数据适配器句柄 int, 项索引 int, 列索引 int, 值 string) bool {
	r, _, _ := xAdTable_SetItemText.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器表_置项文本扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.

// ff:数据适配器表_置项文本EX
// pValue:值
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项文本EX(数据适配器句柄 int, 项索引 int, 字段称 string, 值 string) bool {
	r, _, _ := xAdTable_SetItemTextEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器表_置项整数值, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 值.

// ff:数据适配器表_置项整数值
// nValue:值
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项整数值(数据适配器句柄 int, 项索引 int, 列索引 int, 值 int32) bool {
	r, _, _ := xAdTable_SetItemInt.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引), uintptr(值))
	return r != 0
}

// 数据适配器表_置项整数值扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// nValue: 值.

// ff:数据适配器表_置项整数值EX
// nValue:值
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项整数值EX(数据适配器句柄 int, 项索引 int, 字段称 string, 值 int32) bool {
	r, _, _ := xAdTable_SetItemIntEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称), uintptr(值))
	return r != 0
}

// 数据适配器表_置项浮点值, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fValue: 值.

// ff:数据适配器表_置项浮点值
// fValue:值
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项浮点值(数据适配器句柄 int, 项索引 int, 列索引 int, 值 float32) bool {
	r, _, _ := xAdTable_SetItemFloat.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引), 炫彩工具类.Float32Ptr(值))
	return r != 0
}

// 数据适配器表_置项浮点值扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// fValue: 值.

// ff:数据适配器表_置项浮点值EX
// fValue:值
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项浮点值EX(数据适配器句柄 int, 项索引 int, 字段称 string, 值 float32) bool {
	r, _, _ := xAdTable_SetItemFloatEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称), 炫彩工具类.Float32Ptr(值))
	return r != 0
}

// 数据适配器表_置项图片, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.

// ff:数据适配器表_置项图片
// hImage:图片句柄
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项图片(数据适配器句柄 int, 项索引 int, 列索引 int, 图片句柄 int) bool {
	r, _, _ := xAdTable_SetItemImage.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引), uintptr(图片句柄))
	return r != 0
}

// 数据适配器表_置项图片扩展, 设置项数据.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:数据适配器表_置项图片EX
// hImage:图片句柄
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_置项图片EX(数据适配器句柄 int, 项索引 int, 字段称 string, 图片句柄 int) bool {
	r, _, _ := xAdTable_SetItemImageEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄))
	return r != 0
}

// 数据适配器表_删除项, 删除项.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.

// ff:数据适配器表_删除项
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_删除项(数据适配器句柄 int, 项索引 int) bool {
	r, _, _ := xAdTable_DeleteItem.Call(uintptr(数据适配器句柄), uintptr(项索引))
	return r != 0
}

// 数据适配器表_删除项扩展, 删除多个项.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项起始索引.
//
// nCount: 删除项数量.

// ff:数据适配器表_删除项EX
// nCount:删除项数量
// iItem:项起始索引
// hAdapter:数据适配器句柄
func X数据适配器表_删除项EX(数据适配器句柄 int, 项起始索引 int, 删除项数量 int) bool {
	r, _, _ := xAdTable_DeleteItemEx.Call(uintptr(数据适配器句柄), uintptr(项起始索引), uintptr(删除项数量))
	return r != 0
}

// 数据适配器表_删除项全部, 删除所有项.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器表_删除项全部
// hAdapter:数据适配器句柄
func X数据适配器表_删除项全部(数据适配器句柄 int) int {
	r, _, _ := xAdTable_DeleteItemAll.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器表_删除列全部, 删除所有列, 并且清空所有数据.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器表_删除列全部
// hAdapter:数据适配器句柄
func X数据适配器表_删除列全部(数据适配器句柄 int) int {
	r, _, _ := xAdTable_DeleteColumnAll.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器表_取项数量, 获取项数量.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器表_取项数量
// hAdapter:数据适配器句柄
func X数据适配器表_取项数量(数据适配器句柄 int) int {
	r, _, _ := xAdTable_GetCount.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器表_取列数量, 获取列数量.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器表_取列数量
// hAdapter:数据适配器句柄
func X数据适配器表_取列数量(数据适配器句柄 int) int {
	r, _, _ := xAdTable_GetCountColumn.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器表_取项文本, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.

// ff:数据适配器表_取项文本
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项文本(数据适配器句柄 int, 项索引 int, 列索引 int) string {
	r, _, _ := xAdTable_GetItemText.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器表_取项文本扩展, 获取项文本内容.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.

// ff:数据适配器表_取项文本EX
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项文本EX(数据适配器句柄 int, 项索引 int, 字段称 string) string {
	r, _, _ := xAdTable_GetItemTextEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器表_取项图片, 获取项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.

// ff:数据适配器表_取项图片
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项图片(数据适配器句柄 int, 项索引 int, 列索引 int) int {
	r, _, _ := xAdTable_GetItemImage.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引))
	return int(r)
}

// 数据适配器表_取项图片扩展, 获取项图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.

// ff:数据适配器表_取项图片EX
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项图片EX(数据适配器句柄 int, 项索引 int, 字段称 string) int {
	r, _, _ := xAdTable_GetItemImageEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器表_取项整数值, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.

// ff:数据适配器表_取项整数值
// pOutValue:接收返还值
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项整数值(数据适配器句柄 int, 项索引 int, 列索引 int, 接收返还值 *int32) bool {
	r, _, _ := xAdTable_GetItemInt.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 数据适配器表_取项整数值扩展, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.

// ff:数据适配器表_取项整数值EX
// pOutValue:接收返还值
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项整数值EX(数据适配器句柄 int, 项索引 int, 字段称 string, 接收返还值 *int32) bool {
	r, _, _ := xAdTable_GetItemIntEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 数据适配器表_取项浮点值, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.

// ff:数据适配器表_取项浮点值
// pOutValue:接收返还值
// iColumn:列索引
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项浮点值(数据适配器句柄 int, 项索引 int, 列索引 int, 接收返还值 *float32) bool {
	r, _, _ := xAdTable_GetItemFloat.Call(uintptr(数据适配器句柄), uintptr(项索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}

// 数据适配器表_取项浮点值扩展, 获取项值.
//
// hAdapter: 数据适配器句柄.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.

// ff:数据适配器表_取项浮点值EX
// pOutValue:接收返还值
// pName:字段称
// iItem:项索引
// hAdapter:数据适配器句柄
func X数据适配器表_取项浮点值EX(数据适配器句柄 int, 项索引 int, 字段称 string, 接收返还值 *float32) bool {
	r, _, _ := xAdTable_GetItemFloatEx.Call(uintptr(数据适配器句柄), uintptr(项索引), 炫彩工具类.StrPtr(字段称), uintptr(unsafe.Pointer(接收返还值)))
	return r != 0
}
