package 炫彩数据适配器类

import (
	"github.com/888go/xcgui/xc"
)

// AdapterMap 数据适配器-单列Map-列表头(listHeader).
type AdapterMap struct {
	adapter
}

// 数据适配器MAP_创建, 创建数据适配器, 单列数据.
func X创建适配器MAP() *AdapterMap {
	p := &AdapterMap{}
	p.X设置句柄(炫彩基类.X数据适配器MAP_创建())
	return p
}

// 从句柄创建对象.
func X创建适配器MAP并按句柄(handle int) *AdapterMap {
	p := &AdapterMap{}
	p.X设置句柄(handle)
	return p
}

// 数据适配器MAP_添加项文本, 增加数据项.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterMap) X添加项文本(字段称 string, 值 string) bool {
	return 炫彩基类.X数据适配器MAP_添加项文本(a.Handle, 字段称, 值)
}

// 数据适配器MAP_添加项图片, 增加数据项.
//
// pName: 字段称.
//
// hImage: 图片句柄.
func (a *AdapterMap) X添加项图片(字段称 string, 图片句柄 int) bool {
	return 炫彩基类.X数据适配器MAP_添加项图片(a.Handle, 字段称, 图片句柄)
}

// 数据适配器MAP_删除项, 删除数据项.
//
// pName: 字段称.
func (a *AdapterMap) X删除项(字段称 string) bool {
	return 炫彩基类.X数据适配器MAP_删除项(a.Handle, 字段称)
}

// 数据适配器MAP_取项数量, 返回项数量.
func (a *AdapterMap) X取项数量() int {
	return 炫彩基类.X数据适配器MAP_取项数量(a.Handle)
}

// 数据适配器MAP_取项文本, 获取项内容, 如果内容为文本.
//
// pName: 字段称.
func (a *AdapterMap) X取项文本(字段称 string) string {
	return 炫彩基类.X数据适配器MAP_取项文本(a.Handle, 字段称)
}

// 数据适配器MAP_取项图片, 获取项内容, 如果内容为图片句柄, 返回图片句柄.
//
// pName: 字段称.
func (a *AdapterMap) X取项图片(字段称 string) int {
	return 炫彩基类.X数据适配器MAP_取项图片(a.Handle, 字段称)
}

// 数据适配器MAP_置项文本, 设置项内容.
//
// pName: 字段称.
//
// pValue: 值.
func (a *AdapterMap) X置项文本(字段称 string, 值 string) bool {
	return 炫彩基类.X数据适配器MAP_置项文本(a.Handle, 字段称, 值)
}

// 数据适配器MAP_置项图片, 设置项内容.
//
// pName: 字段称.
//
// hImage: 值.
func (a *AdapterMap) X置项图片(字段称 string, 值 int) bool {
	return 炫彩基类.X数据适配器MAP_置项图片(a.Handle, 字段称, 值)
}
