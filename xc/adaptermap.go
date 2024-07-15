package 炫彩基类

import "github.com/888go/xcgui/common"

// 数据适配器MAP_创建, 创建数据适配器, 单列数据, 返回数据适配器句柄.

// ff:数据适配器MAP_创建
func X数据适配器MAP_创建() int {
	r, _, _ := xAdMap_Create.Call()
	return int(r)
}

// 数据适配器MAP_添加项文本, 增加数据项.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.

// ff:数据适配器MAP_添加项文本
// pValue:值
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_添加项文本(数据适配器句柄 int, 字段称 string, 值 string) bool {
	r, _, _ := xAdMap_AddItemText.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器MAP_添加项图片, 增加数据项.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:数据适配器MAP_添加项图片
// hImage:图片句柄
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_添加项图片(数据适配器句柄 int, 字段称 string, 图片句柄 int) bool {
	r, _, _ := xAdMap_AddItemImage.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), uintptr(图片句柄))
	return r != 0
}

// 数据适配器MAP_删除项, 删除数据项.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.

// ff:数据适配器MAP_删除项
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_删除项(数据适配器句柄 int, 字段称 string) bool {
	r, _, _ := xAdMap_DeleteItem.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return r != 0
}

// 数据适配器MAP_取项数量, 返回项数量.
//
// hAdapter: 数据适配器句柄.

// ff:数据适配器MAP_取项数量
// hAdapter:数据适配器句柄
func X数据适配器MAP_取项数量(数据适配器句柄 int) int {
	r, _, _ := xAdMap_GetCount.Call(uintptr(数据适配器句柄))
	return int(r)
}

// 数据适配器MAP_取项文本, 获取项内容, 如果内容为文本.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.

// ff:数据适配器MAP_取项文本
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_取项文本(数据适配器句柄 int, 字段称 string) string {
	r, _, _ := xAdMap_GetItemText.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return 炫彩工具类.UintPtrToString(r)
}

// 数据适配器MAP_取项图片, 获取项内容, 如果内容为图片句柄, 返回图片句柄.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.

// ff:数据适配器MAP_取项图片
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_取项图片(数据适配器句柄 int, 字段称 string) int {
	r, _, _ := xAdMap_GetItemImage.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称))
	return int(r)
}

// 数据适配器MAP_置项文本, 设置项内容.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// pValue: 值.

// ff:数据适配器MAP_置项文本
// pValue:值
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_置项文本(数据适配器句柄 int, 字段称 string, 值 string) bool {
	r, _, _ := xAdMap_SetItemText.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), 炫彩工具类.StrPtr(值))
	return r != 0
}

// 数据适配器MAP_置项图片, 设置项内容.
//
// hAdapter: 数据适配器句柄.
//
// pName: 字段称.
//
// hImage: 值.

// ff:数据适配器MAP_置项图片
// hImage:值
// pName:字段称
// hAdapter:数据适配器句柄
func X数据适配器MAP_置项图片(数据适配器句柄 int, 字段称 string, 值 int) bool {
	r, _, _ := xAdMap_SetItemImage.Call(uintptr(数据适配器句柄), 炫彩工具类.StrPtr(字段称), uintptr(值))
	return r != 0
}
