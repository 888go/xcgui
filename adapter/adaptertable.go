package 炫彩数据适配器类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// AdapterTable 数据适配器-XList-XListBox.
type AdapterTable struct {
	adapter
}

// 数据适配器表_创建, 创建列表框元素数据适配器.

// ff:创建适配器表
func X创建适配器表() *AdapterTable {
	p := &AdapterTable{}
	p.X设置句柄(炫彩基类.X数据适配器表_创建())
	return p
}

// 从句柄创建对象.

// ff:创建适配器表并按句柄
// handle:
func X创建适配器表并按句柄(handle int) *AdapterTable {
	p := &AdapterTable{}
	p.X设置句柄(handle)
	return p
}

// 数据适配器表_排序, 对内容进行排序.
//
// iColumn: 要排序的列索引。.
//
// bAscending: 是否按照升序方式排序.

// ff:排序
// bAscending:升序排序
// iColumn:要排序的列索引
func (a *AdapterTable) X排序(要排序的列索引 int, 升序排序 bool) int {
	return 炫彩基类.X数据适配器表_排序(a.Handle, 要排序的列索引, 升序排序)
}

// 数据适配器表_取项数据类型, 获取项数据类型, 返回: Adapter_Date_Type_.
//
// iItem: 项索引.
//
// iColumn: 列索引.

// ff:取项数据类型
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X取项数据类型(项索引 int, 列索引 int) 炫彩常量类.Adapter_Date_Type_ {
	return 炫彩基类.X数据适配器表_取项数据类型(a.Handle, 项索引, 列索引)
}

// 数据适配器表_取项数据类型扩展, 返回: Adapter_Date_Type_.
//
// iItem: 项索引.
//
// pName: 字段称.

// ff:取项数据类型EX
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X取项数据类型EX(项索引 int, 字段称 string) 炫彩常量类.Adapter_Date_Type_ {
	return 炫彩基类.X数据适配器表_取项数据类型EX(a.Handle, 项索引, 字段称)
}

// 数据适配器表_添加列, 添加数据列.
//
// pName: 字段称.

// ff:添加列
// pName:字段称
func (a *AdapterTable) X添加列(字段称 string) int {
	return 炫彩基类.X数据适配器表_添加列(a.Handle, 字段称)
}

// 数据适配器表_置列, 设置列, 返回列数量.
//
// pColName: 列名, 多个列名用逗号分开.

// ff:置列
// pColName:列名
func (a *AdapterTable) X置列(列名 string) int {
	return 炫彩基类.X数据适配器表_置列(a.Handle, 列名)
}

// 数据适配器表_添加项文本, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// pValue: 值.

// ff:添加项文本
// pValue:值
func (a *AdapterTable) X添加项文本(值 string) int {
	return 炫彩基类.X数据适配器表_添加项文本(a.Handle, 值)
}

// 数据适配器表_添加项文本扩展, 添加数据项, 填充指定列数据, 返回项索引.
//
// pName: 字段称.
//
// pValue: 值.

// ff:添加项文本EX
// pValue:值
// pName:字段称
func (a *AdapterTable) X添加项文本EX(字段称 string, 值 string) int {
	return 炫彩基类.X数据适配器表_添加项文本EX(a.Handle, 字段称, 值)
}

// 数据适配器表_添加项图片, 添加数据项, 默认值放到第一列, 返回项索引值.
//
// hImage: 图片句柄.

// ff:添加项图片
// hImage:图片句柄
func (a *AdapterTable) X添加项图片(图片句柄 int) int {
	return 炫彩基类.X数据适配器表_添加项图片(a.Handle, 图片句柄)
}

// 数据适配器表_添加项图片扩展, 添加数据项, 并填充指定列数据.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:添加项图片EX
// hImage:图片句柄
// pName:字段称
func (a *AdapterTable) X添加项图片EX(字段称 string, 图片句柄 int) int {
	return 炫彩基类.X数据适配器表_添加项图片EX(a.Handle, 字段称, 图片句柄)
}

// 数据适配器表_插入项文本, 插入数据项, 填充第一列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// pValue: 值.

// ff:插入项文本
// pValue:值
// iItem:插入位置索引
func (a *AdapterTable) X插入项文本(插入位置索引 int, 值 string) int {
	return 炫彩基类.X数据适配器表_插入项文本(a.Handle, 插入位置索引, 值)
}

// 数据适配器表_插入项文本扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// pValue: 值.

// ff:插入项文本EX
// pValue:值
// pName:字段称
// iItem:插入位置索引
func (a *AdapterTable) X插入项文本EX(插入位置索引 int, 字段称 string, 值 string) int {
	return 炫彩基类.X数据适配器表_插入项文本EX(a.Handle, 插入位置索引, 字段称, 值)
}

// 数据适配器表_插入项图片, 插入数据项, 填充第一列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// hImage: 图片句柄.

// ff:插入项图片
// hImage:图片句柄
// iItem:插入位置索引
func (a *AdapterTable) X插入项图片(插入位置索引 int, 图片句柄 int) int {
	return 炫彩基类.X数据适配器表_插入项图片(a.Handle, 插入位置索引, 图片句柄)
}

// 数据适配器表_插入项图片扩展, 插入数据项, 并填充指定列数据, 返回项索引.
//
// iItem: 插入位置索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:插入项图片EX
// hImage:图片句柄
// pName:字段称
// iItem:插入位置索引
func (a *AdapterTable) X插入项图片EX(插入位置索引 int, 字段称 string, 图片句柄 int) int {
	return 炫彩基类.X数据适配器表_插入项图片EX(a.Handle, 插入位置索引, 字段称, 图片句柄)
}

// 数据适配器表_置项文本, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.

// ff:置项文本
// pValue:值
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X置项文本(项索引 int, 列索引 int, 值 string) bool {
	return 炫彩基类.X数据适配器表_置项文本(a.Handle, 项索引, 列索引, 值)
}

// 数据适配器表_置项文本扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.

// ff:置项文本EX
// pValue:值
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X置项文本EX(项索引 int, 字段称 string, 值 string) bool {
	return 炫彩基类.X数据适配器表_置项文本EX(a.Handle, 项索引, 字段称, 值)
}

// 数据适配器表_置项整数值, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// nValue: 值.

// ff:置项整数值
// nValue:值
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X置项整数值(项索引 int, 列索引 int, 值 int32) bool {
	return 炫彩基类.X数据适配器表_置项整数值(a.Handle, 项索引, 列索引, 值)
}

// 数据适配器表_置项整数值扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// nValue: 值.

// ff:置项整数值EX
// nValue:值
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X置项整数值EX(项索引 int, 字段称 string, 值 int32) bool {
	return 炫彩基类.X数据适配器表_置项整数值EX(a.Handle, 项索引, 字段称, 值)
}

// 数据适配器表_置项浮点值, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// fValue: 值.

// ff:置项浮点值
// fValue:值
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X置项浮点值(项索引 int, 列索引 int, 值 float32) bool {
	return 炫彩基类.X数据适配器表_置项浮点值(a.Handle, 项索引, 列索引, 值)
}

// 数据适配器表_置项浮点值扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// fValue: 值.

// ff:置项浮点值EX
// fValue:值
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X置项浮点值EX(项索引 int, 字段称 string, 值 float32) bool {
	return 炫彩基类.X数据适配器表_置项浮点值EX(a.Handle, 项索引, 字段称, 值)
}

// 数据适配器表_置项图片, 设置项数据.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.

// ff:置项图片
// hImage:图片句柄
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X置项图片(项索引 int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器表_置项图片(a.Handle, 项索引, 列索引, 图片句柄)
}

// 数据适配器表_置项图片扩展, 设置项数据.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.

// ff:置项图片EX
// hImage:图片句柄
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X置项图片EX(项索引 int, 字段称 string, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器表_置项图片EX(a.Handle, 项索引, 字段称, 图片句柄)
}

// 数据适配器表_删除项, 删除项.
//
// iItem: 项索引.

// ff:删除项
// iItem:项索引
func (a *AdapterTable) X删除项(项索引 int) bool {
	return 炫彩基类.X数据适配器表_删除项(a.Handle, 项索引)
}

// 数据适配器表_删除项扩展, 删除多个项.
//
// iItem: 项起始索引.
//
// nCount: 删除项数量.

// ff:删除项EX
// nCount:删除项数量
// iItem:项起始索引
func (a *AdapterTable) X删除项EX(项起始索引 int, 删除项数量 int) bool {
	return 炫彩基类.X数据适配器表_删除项EX(a.Handle, 项起始索引, 删除项数量)
}

// 数据适配器表_删除项全部, 删除所有项.

// ff:删除项全部
func (a *AdapterTable) X删除项全部() int {
	return 炫彩基类.X数据适配器表_删除项全部(a.Handle)
}

// 数据适配器表_删除列全部, 删除所有列, 并且清空所有数据.

// ff:删除列全部
func (a *AdapterTable) X删除列全部() int {
	return 炫彩基类.X数据适配器表_删除列全部(a.Handle)
}

// 数据适配器表_取项数量, 获取项数量.

// ff:取项数量
func (a *AdapterTable) X取项数量() int {
	return 炫彩基类.X数据适配器表_取项数量(a.Handle)
}

// 数据适配器表_取列数量, 获取列数量.

// ff:取列数量
func (a *AdapterTable) X取列数量() int {
	return 炫彩基类.X数据适配器表_取列数量(a.Handle)
}

// 数据适配器表_取项文本, 获取项文本内容.
//
// iItem: 项索引.
//
// iColumn: 列索引.

// ff:取项文本
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X取项文本(项索引 int, 列索引 int) string {
	return 炫彩基类.X数据适配器表_取项文本(a.Handle, 项索引, 列索引)
}

// 数据适配器表_取项文本扩展, 获取项文本内容.
//
// iItem: 项索引.
//
// pName: 字段称.

// ff:取项文本EX
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X取项文本EX(项索引 int, 字段称 string) string {
	return 炫彩基类.X数据适配器表_取项文本EX(a.Handle, 项索引, 字段称)
}

// 数据适配器表_取项图片, 获取项图片句柄.
//
// iItem: 项索引.
//
// iColumn: 列索引.

// ff:取项图片
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X取项图片(项索引 int, 列索引 int) int {
	return 炫彩基类.X数据适配器表_取项图片(a.Handle, 项索引, 列索引)
}

// 数据适配器表_取项图片扩展, 获取项图片句柄.
//
// iItem: 项索引.
//
// pName: 字段称.

// ff:取项图片EX
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X取项图片EX(项索引 int, 字段称 string) int {
	return 炫彩基类.X数据适配器表_取项图片EX(a.Handle, 项索引, 字段称)
}

// 数据适配器表_取项整数值, 获取项值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.

// ff:取项整数值
// pOutValue:接收返还值
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X取项整数值(项索引 int, 列索引 int, 接收返还值 *int32) bool {
	return 炫彩基类.X数据适配器表_取项整数值(a.Handle, 项索引, 列索引, 接收返还值)
}

// 数据适配器表_取项整数值扩展, 获取项值.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.

// ff:取项整数值EX
// pOutValue:接收返还值
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X取项整数值EX(项索引 int, 字段称 string, 接收返还值 *int32) bool {
	return 炫彩基类.X数据适配器表_取项整数值EX(a.Handle, 项索引, 字段称, 接收返还值)
}

// 数据适配器表_取项浮点值, 获取项值.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pOutValue: 接收返还值.

// ff:取项浮点值
// pOutValue:接收返还值
// iColumn:列索引
// iItem:项索引
func (a *AdapterTable) X取项浮点值(项索引 int, 列索引 int, 接收返还值 *float32) bool {
	return 炫彩基类.X数据适配器表_取项浮点值(a.Handle, 项索引, 列索引, 接收返还值)
}

// 数据适配器表_取项浮点值扩展, 获取项值.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pOutValue: 接收返还值.

// ff:取项浮点值EX
// pOutValue:接收返还值
// pName:字段称
// iItem:项索引
func (a *AdapterTable) X取项浮点值EX(项索引 int, 字段称 string, 接收返还值 *float32) bool {
	return 炫彩基类.X数据适配器表_取项浮点值EX(a.Handle, 项索引, 字段称, 接收返还值)
}
