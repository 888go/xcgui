package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// Table 表格.
type Table struct {
	Shape
}

// 表格_创建.
//
// x: 按钮x坐标.
//
// y: 按钮y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func X创建表格(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *Table {
	p := &Table{}
	p.X设置句柄(炫彩基类.X表格_创建(x坐标, y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建表格并按句柄(handle int) *Table {
	p := &Table{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建表格并按名称(name string) *Table {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Table{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建表格并按UID(nUID int) *Table {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Table{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建表格并按UID名称(name string) *Table {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Table{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 表格_重置.
//
// nRow: 行数量.
//
// nCol: 列数量.
func (t *Table) X重置(行数量 int, 列数量 int) int {
	return 炫彩基类.X表格_重置(t.Handle, 行数量, 列数量)
}

// 表格_组合行.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) X组合行(行索引 int, 列索引 int, 数量 int) int {
	return 炫彩基类.X表格_组合行(t.Handle, 行索引, 列索引, 数量)
}

// 表格_组合列.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) X组合列(行索引 int, 列索引 int, 数量 int) int {
	return 炫彩基类.X表格_组合列(t.Handle, 行索引, 列索引, 数量)
}

// 表格_置列宽.
//
// iCol: 列索引.
//
// width: 宽度.
func (t *Table) X置列宽(列索引 int, 宽度 int) int {
	return 炫彩基类.X表格_置列宽(t.Handle, 列索引, 宽度)
}

// 表格_置行高.
//
// iRow: 行索引.
//
// height: 高度.
func (t *Table) X置行高(行索引 int, 高度 int) int {
	return 炫彩基类.X表格_置行高(t.Handle, 行索引, 高度)
}

// 表格_置边颜色.
//
// color: 颜色.
func (t *Table) X置边颜色(颜色 int) int {
	return 炫彩基类.X表格_置边颜色(t.Handle, 颜色)
}

// 表格_置文本颜色.
//
// color: 颜色.
func (t *Table) X置文本颜色(颜色 int) int {
	return 炫彩基类.X表格_置文本颜色(t.Handle, 颜色)
}

// 表格_置字体.
//
// hFont: 字体句柄.
func (t *Table) X置字体(字体句柄 int) int {
	return 炫彩基类.X表格_置字体(t.Handle, 字体句柄)
}

// 表格_置项内填充.
//
// leftSize: 内填充大小.
//
// topSize: 内填充大小.
//
// rightSize: 内填充大小.
//
// bottomSize: 内填充大小.
func (t *Table) X置项内填充(左 int, 上 int, 右 int, 下 int) int {
	return 炫彩基类.X表格_置项内填充(t.Handle, 左, 上, 右, 下)
}

// 表格_置项文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
func (t *Table) X置项文本(行索引 int, 列索引 int, 文本 string) int {
	return 炫彩基类.X表格_置项文本(t.Handle, 行索引, 列索引, 文本)
}

// 表格_置项字体.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hFont: 字体句柄.
func (t *Table) X置项字体(行索引 int, 列索引 int, 字体句柄 int) int {
	return 炫彩基类.X表格_置项字体(t.Handle, 行索引, 列索引, 字体句柄)
}

// 表格_置项文本对齐.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// nAlign: 对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (t *Table) X置项文本对齐(行索引 int, 列索引 int, 对齐方式 炫彩常量类.TextFormatFlag_) int {
	return 炫彩基类.X表格_置项文本对齐(t.Handle, 行索引, 列索引, 对齐方式)
}

// 表格_置项文本色.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func (t *Table) X置项文本色(行索引 int, 列索引 int, 颜色 int, 是否使用 bool) int {
	return 炫彩基类.X表格_置项文本色(t.Handle, 行索引, 列索引, 颜色, 是否使用)
}

// 表格_置项背景色.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func (t *Table) X置项背景色(行索引 int, 列索引 int, 颜色 int, 是否使用 bool) int {
	return 炫彩基类.X表格_置项背景色(t.Handle, 行索引, 列索引, 颜色, 是否使用)
}

// 表格_置项线.
//
// iRow1: 行索引1.
//
// iCol1: 列索引1.
//
// iRow2: 行索引2.
//
// iCol2: 列索引2.
//
// nFlag: 标识, Table_Line_Flag_, 暂时没有, 填0.
//
// color: 颜色.
func (t *Table) X置项线(行索引1 int, 列索引1 int, 行索引2 int, 列索引2 int, 标识 int, 颜色 int) int {
	return 炫彩基类.X表格_置项线(t.Handle, 行索引1, 列索引1, 行索引2, 列索引2, 标识, 颜色)
}

// 表格_置项标识.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// flag: 标识, Table_Flag_.
func (t *Table) X置项标识(行索引 int, 列索引 int, 标识 炫彩常量类.Table_Flag_) int {
	return 炫彩基类.X表格_置项标识(t.Handle, 行索引, 列索引, 标识)
}

// 表格_取项坐标.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pRect: 接收返回坐标.
func (t *Table) X取项坐标(行索引 int, 列索引 int, 接收返回坐标 *炫彩基类.RECT) bool {
	return 炫彩基类.X表格_取项坐标(t.Handle, 行索引, 列索引, 接收返回坐标)
}
