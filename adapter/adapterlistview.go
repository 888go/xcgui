package 炫彩数据适配器类

import (
	"github.com/888go/xcgui/xc"
)

// 数据适配器-列表视元素.
type AdapterListView struct {
	adapter
}

// 数据适配器列表视_创建, 创建列表视元素数据适配器.
func X创建适配器列表视() *AdapterListView {
	p := &AdapterListView{}
	p.X设置句柄(炫彩基类.X数据适配器列表视_创建())
	return p
}

// 从句柄创建对象.
func X创建适配器列表视并按句柄(handle int) *AdapterListView {
	p := &AdapterListView{}
	p.X设置句柄(handle)
	return p
}

// 数据适配器列表视_组添加列, 组操作, 添加数据列, 返回列索引.
//
// pName: 字段称.
func (a *AdapterListView) X组添加列(字段称 string) int {
	return 炫彩基类.X数据适配器列表视_组添加列(a.Handle, 字段称)
}

// 数据适配器列表视_添加组文本, 组操作, 添加组, 数据默认填充到第一列, 返回组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加组文本(值 string, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加组文本(a.Handle, 值, 插入位置)
}

// 数据适配器列表视_添加组文本扩展, 组操作, 添加组, 数据填充指定列, 返回组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加组文本EX(字段称 string, 值 string, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加组文本EX(a.Handle, 字段称, 值, 插入位置)
}

// 数据适配器列表视_添加组图片, 组操作, 添加组, 数据默认填充第一列.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加组图片(图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加组图片(a.Handle, 图片句柄, 插入位置)
}

// 数据适配器列表视_添加组图片扩展, 组操作, 添加组, 数据填充指定列.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加组图片EX(字段称 string, 图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加组图片EX(a.Handle, 字段称, 图片句柄, 插入位置)
}

// 数据适配器列表视_组设置文本, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterListView) X组设置文本(组索引 int, 列索引 int, 值 string) bool {
	return 炫彩基类.X数据适配器列表视_组设置文本(a.Handle, 组索引, 列索引, 值)
}

// 数据适配器列表视_组设置文本扩展, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// pValue: 值.
func (a *AdapterListView) X组设置文本EX(组索引 int, 字段名 string, 值 string) bool {
	return 炫彩基类.X数据适配器列表视_组设置文本EX(a.Handle, 组索引, 字段名, 值)
}

// 数据适配器列表视_组设置图片, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterListView) X组设置图片(组索引 int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器列表视_组设置图片(a.Handle, 组索引, 列索引, 图片句柄)
}

// 数据适配器列表视_组设置图片扩展, 组操作, 设置指定项数据.
//
// iGroup: 组索引.
//
// pName: 字段名.
//
// hImage: 图片句柄.
func (a *AdapterListView) X组设置图片EX(组索引 int, 字段名 string, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器列表视_组设置图片EX(a.Handle, 组索引, 字段名, 图片句柄)
}

// 数据适配器列表视_项添加列, 项操作, 添加列.
//
// pName: 字段称.
func (a *AdapterListView) X项添加列(字段称 string) int {
	return 炫彩基类.X数据适配器列表视_项添加列(a.Handle, 字段称)
}

// 数据适配器列表视_取组数量, 组操作, 获取组数量.
func (a *AdapterListView) X取组数量() int {
	return 炫彩基类.X数据适配器列表视_取组数量(a.Handle)
}

// 数据适配器列表视_取组中项数量, 项操作, 获取指定组中项数量.
//
// iGroup: 组索引.
func (a *AdapterListView) X取组中项数量(组索引 int) int {
	return 炫彩基类.X数据适配器列表视_取组中项数量(a.Handle, 组索引)
}

// 数据适配器列表视_添加项文本, 项操作, 添加项.
//
// iGroup: 组索引.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加项文本(组索引 int, 值 string, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加项文本(a.Handle, 组索引, 值, 插入位置)
}

// 数据适配器列表视_添加项文本扩展, 项操作, 添加项, 数据填充指定列.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// pValue: 值.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加项文本EX(组索引 int, 字段称 string, 值 string, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加项文本EX(a.Handle, 组索引, 字段称, 值, 插入位置)
}

// 数据适配器列表视_添加项图片, 项操作, 添加项.
//
// iGroup: 组索引.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加项图片(组索引 int, 图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加项图片(a.Handle, 组索引, 图片句柄, 插入位置)
}

// 数据适配器列表视_添加项图片扩展, 项操作, 添加项, 填充指定列数据.
//
// iGroup: 组索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
//
// iPos: 插入位置, 可为-1.
func (a *AdapterListView) X添加项图片EX(组索引 int, 字段称 string, 图片句柄 int, 插入位置 int) int {
	return 炫彩基类.X数据适配器列表视_添加项图片EX(a.Handle, 组索引, 字段称, 图片句柄, 插入位置)
}

// 数据适配器列表视_组删除项, 删除组, 自动删除子项.
//
// iGroup: 组索引.
func (a *AdapterListView) X组删除项(组索引 int) bool {
	return 炫彩基类.X数据适配器列表视_组删除项(a.Handle, 组索引)
}

// 数据适配器列表视_删除全部子项, 删除指定组的所有子项.
//
// iGroup: 组索引.
func (a *AdapterListView) X删除全部子项(组索引 int) int {
	return 炫彩基类.X数据适配器列表视_删除全部子项(a.Handle, 组索引)
}

// 数据适配器列表视_删除项, 删除项.
//
// iGroup: 组索引.
//
// iItem: 项索引.
func (a *AdapterListView) X删除项(组索引 int, 项索引 int) bool {
	return 炫彩基类.X数据适配器列表视_删除项(a.Handle, 组索引, 项索引)
}

// 数据适配器列表视_删除全部, 删除所有(组,项,列).
func (a *AdapterListView) X删除全部() int {
	return 炫彩基类.X数据适配器列表视_删除全部(a.Handle)
}

// 数据适配器列表视_删除全部组, 删除所有的组.
func (a *AdapterListView) X删除全部组() int {
	return 炫彩基类.X数据适配器列表视_删除全部组(a.Handle)
}

// 数据适配器列表视_删除全部项, 删除所有的项.
func (a *AdapterListView) X删除全部项() int {
	return 炫彩基类.X数据适配器列表视_删除全部项(a.Handle)
}

// 数据适配器列表视_组删除列, 删除组的列.
//
// iColumn: 列索引.
func (a *AdapterListView) X组删除列(列索引 int) int {
	return 炫彩基类.X数据适配器列表视_组删除列(a.Handle, 列索引)
}

// 数据适配器列表视_项删除列, 删除项的列.
//
// iColumn: 列索引.
func (a *AdapterListView) X项删除列(列索引 int) int {
	return 炫彩基类.X数据适配器列表视_项删除列(a.Handle, 列索引)
}

// 数据适配器列表视_项获取文本扩展, 项操作, 获取项文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterListView) X项获取文本EX(组索引 int, 项索引 int, 字段称 string) string {
	return 炫彩基类.X数据适配器列表视_项获取文本EX(a.Handle, 组索引, 项索引, 字段称)
}

// 数据适配器列表视_项获取图片扩展, 项操作, 获取项图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
func (a *AdapterListView) X项获取图片EX(组索引 int, 项索引 int, 字段称 string) int {
	return 炫彩基类.X数据适配器列表视_项获取图片EX(a.Handle, 组索引, 项索引, 字段称)
}

// 数据适配器列表视_项置文本, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// pValue: 值.
func (a *AdapterListView) X项置文本(组索引 int, 项索引 int, 列索引 int, 值 string) bool {
	return 炫彩基类.X数据适配器列表视_项置文本(a.Handle, 组索引, 项索引, 列索引, 值)
}

// 数据适配器列表视_项置文本扩展, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterListView) X项置文本EX(组索引 int, 项索引 int, 字段称 string, 值 string) bool {
	return 炫彩基类.X数据适配器列表视_项置文本EX(a.Handle, 组索引, 项索引, 字段称, 值)
}

// 数据适配器列表视_项置图片, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
//
// hImage: 图片句柄.
func (a *AdapterListView) X项置图片(组索引 int, 项索引 int, 列索引 int, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器列表视_项置图片(a.Handle, 组索引, 项索引, 列索引, 图片句柄)
}

// 数据适配器列表视_项置图片扩展, 项操作, 数据填充指定列.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterListView) X项置图片EX(组索引 int, 项索引 int, 字段称 string, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器列表视_项置图片EX(a.Handle, 组索引, 项索引, 字段称, 图片句柄)
}

// 数据适配器列表视_组取文本, 返回文本内容.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (a *AdapterListView) X组取文本(组索引 int, 列索引 int) string {
	return 炫彩基类.X数据适配器列表视_组取文本(a.Handle, 组索引, 列索引)
}

// 数据适配器列表视_组取文本扩展, 返回文本内容.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (a *AdapterListView) X组取文本EX(组索引 int, 字段名称 string) string {
	return 炫彩基类.X数据适配器列表视_组取文本EX(a.Handle, 组索引, 字段名称)
}

// 数据适配器列表视_组取图片, 返回图片句柄.
//
// iGroup: 组索引.
//
// iColumn: 列索引.
func (a *AdapterListView) X组取图片(组索引 int, 列索引 int) int {
	return 炫彩基类.X数据适配器列表视_组取图片(a.Handle, 组索引, 列索引)
}

// 数据适配器列表视_组取图片扩展, 返回图片句柄.
//
// iGroup: 组索引.
//
// pName: 字段名称.
func (a *AdapterListView) X组取图片EX(组索引 int, 字段名称 string) int {
	return 炫彩基类.X数据适配器列表视_组取图片EX(a.Handle, 组索引, 字段名称)
}

// 数据适配器列表视_项取文本. 项操作, 返回项文本内容.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterListView) X项取文本(组索引 int, 项索引 int, 列索引 int) string {
	return 炫彩基类.X数据适配器列表视_项取文本(a.Handle, 组索引, 项索引, 列索引)
}

// 数据适配器列表视_项取图片. 项操作, 返回项图片句柄.
//
// iGroup: 组索引.
//
// iItem: 项索引.
//
// iColumn: 列索引.
func (a *AdapterListView) X项取图片(组索引 int, 项索引 int, 列索引 int) int {
	return 炫彩基类.X数据适配器列表视_项取图片(a.Handle, 组索引, 项索引, 列索引)
}
