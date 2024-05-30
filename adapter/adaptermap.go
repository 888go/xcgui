package adapter

import (
	"github.com/twgh/xcgui/xc"
)

// AdapterMap 数据适配器-单列Map-列表头(listHeader).
type AdapterMap struct {
	adapter
}

// 数据适配器MAP_创建, 创建数据适配器, 单列数据.

// ff:创建适配器MAP
func NewAdapterMap() *AdapterMap {
	p := &AdapterMap{}
	p.SetHandle(xc.XAdMap_Create())
	return p
}

// 从句柄创建对象.

// ff:创建适配器MAP并按句柄
// handle:
func NewAdapterMapByHandle(handle int) *AdapterMap {
	p := &AdapterMap{}
	p.SetHandle(handle)
	return p
}

// 数据适配器MAP_添加项文本, 增加数据项.
//
// pName: 字段称.
//
// pValue: 值.

// ff:添加项文本
// pValue:值
// pName:字段称
func (a *AdapterMap) AddItemText(pName string, pValue string) bool {
	return xc.XAdMap_AddItemText(a.Handle, pName, pValue)
}

// 数据适配器MAP_添加项图片, 增加数据项.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:添加项图片
// hImage:图片句柄
// pName:字段称
func (a *AdapterMap) AddItemImage(pName string, hImage int) bool {
	return xc.XAdMap_AddItemImage(a.Handle, pName, hImage)
}

// 数据适配器MAP_删除项, 删除数据项.
//
// pName: 字段称.

// ff:删除项
// pName:字段称
func (a *AdapterMap) DeleteItem(pName string) bool {
	return xc.XAdMap_DeleteItem(a.Handle, pName)
}

// 数据适配器MAP_取项数量, 返回项数量.

// ff:取项数量
func (a *AdapterMap) GetCount() int {
	return xc.XAdMap_GetCount(a.Handle)
}

// 数据适配器MAP_取项文本, 获取项内容, 如果内容为文本.
//
// pName: 字段称.

// ff:取项文本
// pName:字段称
func (a *AdapterMap) GetItemText(pName string) string {
	return xc.XAdMap_GetItemText(a.Handle, pName)
}

// 数据适配器MAP_取项图片, 获取项内容, 如果内容为图片句柄, 返回图片句柄.
//
// pName: 字段称.

// ff:取项图片
// pName:字段称
func (a *AdapterMap) GetItemImage(pName string) int {
	return xc.XAdMap_GetItemImage(a.Handle, pName)
}

// 数据适配器MAP_置项文本, 设置项内容.
//
// pName: 字段称.
//
// pValue: 值.

// ff:置项文本
// pValue:值
// pName:字段称
func (a *AdapterMap) SetItemText(pName string, pValue string) bool {
	return xc.XAdMap_SetItemText(a.Handle, pName, pValue)
}

// 数据适配器MAP_置项图片, 设置项内容.
//
// pName: 字段称.
//
// hImage: 值.

// ff:置项图片
// hImage:值
// pName:字段称
func (a *AdapterMap) SetItemImage(pName string, hImage int) bool {
	return xc.XAdMap_SetItemImage(a.Handle, pName, hImage)
}
