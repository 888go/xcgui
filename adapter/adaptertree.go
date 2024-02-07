package 炫彩数据适配器类

import (
	"github.com/888go/xcgui/xc"
)

// AdapterTree 数据适配器-树元素.
type AdapterTree struct {
	adapter
}

// 数据适配器树_创建, 创建树元素数据适配器.
func X创建适配器树() *AdapterTree {
	p := &AdapterTree{}
	p.X设置句柄(炫彩基类.X数据适配器树_创建())
	return p
}

// 从句柄创建对象.
func X创建适配器树并按句柄(handle int) *AdapterTree {
	p := &AdapterTree{}
	p.X设置句柄(handle)
	return p
}

// 数据适配器树_添加列, 添加列, 返回索引值.
//
// pName: 字段称.
func (a *AdapterTree) X添加列(字段称 string) int {
	return 炫彩基类.X数据适配器树_添加列(a.Handle, 字段称)
}

// 数据适配器树_置列, 设置列, 返回列数量.
//
// pColName: 列名, 列名, 多个列名用逗号分开.
func (a *AdapterTree) X置列(列名 string) int {
	return 炫彩基类.X数据适配器树_置列(a.Handle, 列名)
}

// 数据适配器树_插入项文本, 插入项, 数据填充到第一列, 返回项ID值.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) X插入项文本(值 string, 父ID int, 插入位置ID int) int {
	return 炫彩基类.X数据适配器树_插入项文本(a.Handle, 值, 父ID, 插入位置ID)
}

// 数据适配器树_插入项文本扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// pName: 字段称.
//
// pValue: 值.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) X插入项文本EX(字段称 string, 值 string, 父ID int, 插入位置ID int) int {
	return 炫彩基类.X数据适配器树_插入项文本EX(a.Handle, 字段称, 值, 父ID, 插入位置ID)
}

// 数据适配器树_插入项图片, 插入项, 数据填充到第一列, 返回项ID值.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) X插入项图片(图片句柄 int, 父ID int, 插入位置ID int) int {
	return 炫彩基类.X数据适配器树_插入项图片(a.Handle, 图片句柄, 父ID, 插入位置ID)
}

// 数据适配器树_插入项图片扩展, 插入项, 数据填充到指定列, 返回项ID值.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// nParentID: 父ID.
//
// insertID: 插入位置ID.
func (a *AdapterTree) X插入项图片EX(字段称 string, 图片句柄 int, 父ID int, 插入位置ID int) int {
	return 炫彩基类.X数据适配器树_插入项图片EX(a.Handle, 字段称, 图片句柄, 父ID, 插入位置ID)
}

// 数据适配器树_取项数量, 获取项数量.
func (a *AdapterTree) X取项数量() int {
	return 炫彩基类.X数据适配器树_取项数量(a.Handle)
}

// 数据适配器树_取列数量, 获取列数量.
func (a *AdapterTree) X取列数量() int {
	return 炫彩基类.X数据适配器树_取列数量(a.Handle)
}

// 数据适配器树_置项文本, 设置项数据.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterTree) X置项文本(项ID int, 列索引 int, 值 string) bool {
	return 炫彩基类.X数据适配器树_置项文本(a.Handle, 项ID, 列索引, 值)
}

// 数据适配器树_置项文本扩展, 设置项文件内容.
//
// nID: 项ID.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterTree) X置项文本EX(项ID int, 字段称 string, 值 string) bool {
	return 炫彩基类.X数据适配器树_置项文本EX(a.Handle, 项ID, 字段称, 值)
}

// 数据适配器树_置项图片, 设置项数据.
//
// nID: 项ID.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterTree) X置项图片(项ID int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器树_置项图片(a.Handle, 项ID, 列索引, 图片句柄)
}

// 数据适配器树_置项图片扩展, 设置项内容.
//
// nID: 项ID.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterTree) X置项图片EX(项ID int, 字段称 string, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器树_置项图片EX(a.Handle, 项ID, 字段称, 图片句柄)
}

// 数据适配器树_取项文本, 获取项文本内容.
//
// nID: 项ID.
//
// iColumn: 列索引.
func (a *AdapterTree) X取项文本(项ID int, 列索引 int) string {
	return 炫彩基类.X数据适配器树_取项文本(a.Handle, 项ID, 列索引)
}

// 数据适配器树_取项文本扩展, 获取项文本内容.
//
// nID: 项ID.
//
// pName: 字段称.
func (a *AdapterTree) X取项文本EX(项ID int, 字段称 string) string {
	return 炫彩基类.X数据适配器树_取项文本EX(a.Handle, 项ID, 字段称)
}

// 数据适配器树_取项图片, 获取项内容, 返回图片句柄.
//
// nID: 项ID.
//
// iColumn: 列索引.
func (a *AdapterTree) X取项图片(项ID int, 列索引 int) int {
	return 炫彩基类.X数据适配器树_取项图片(a.Handle, 项ID, 列索引)
}

// 数据适配器树_取项图片扩展, 获取项内容, 返回图片句柄.
//
// nID: 项ID.
//
// pName: 字段称.
func (a *AdapterTree) X取项图片EX(项ID int, 字段称 string) int {
	return 炫彩基类.X数据适配器树_取项图片EX(a.Handle, 项ID, 字段称)
}

// 数据适配器树_删除项, 删除项.
//
// nID: 项ID.
func (a *AdapterTree) X删除项(项ID int) bool {
	return 炫彩基类.X数据适配器树_删除项(a.Handle, 项ID)
}

// 数据适配器树_删除项全部, 删除所有项.
func (a *AdapterTree) X删除项全部() int {
	return 炫彩基类.X数据适配器树_删除项全部(a.Handle)
}

// 数据适配器树_删除列全部, 删除所有列, 并且清空数据.
func (a *AdapterTree) X删除列全部() int {
	return 炫彩基类.X数据适配器树_删除列全部(a.Handle)
}
