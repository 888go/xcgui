package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

type Table struct {
	Shape
}

// I表格_创建.
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
func I表格_创建(x int, y int, cx int, cy int, hParent int) *Table {
	p := &Table{}
	p.SetHandle(xc.XTable_Create(x, y, cx, cy, hParent))
	return p
}

// I表格_从句柄创建对象.
func I表格_从句柄创建(handle int) *Table {
	p := &Table{}
	p.SetHandle(handle)
	return p
}

// I表格_从name创建对象, 失败返回nil.
func I表格_从name创建(name string) *Table {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Table{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I表格_从UID创建对象, 失败返回nil.
func I表格_从UID创建(nUID int) *Table {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Table{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I表格_从UID名称创建对象, 失败返回nil.
func I表格_从UID名称创建(name string) *Table {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Table{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 表格_重置.
//
// nRow: 行数量.
//
// nCol: 列数量.
func (t *Table) I重置(nRow int, nCol int) int {
	return xc.XTable_Reset(t.I句柄, nRow, nCol)
}

// 表格_组合行.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) I组合行(iRow int, iCol int, count int) int {
	return xc.XTable_ComboRow(t.I句柄, iRow, iCol, count)
}

// 表格_组合列.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) I组合列(iRow int, iCol int, count int) int {
	return xc.XTable_ComboCol(t.I句柄, iRow, iCol, count)
}

// 表格_置列宽.
//
// iCol: 列索引.
//
// width: 宽度.
func (t *Table) I置列宽(iCol int, width int) int {
	return xc.XTable_SetColWidth(t.I句柄, iCol, width)
}

// 表格_置行高.
//
// iRow: 行索引.
//
// height: 高度.
func (t *Table) I置行高(iRow int, height int) int {
	return xc.XTable_SetRowHeight(t.I句柄, iRow, height)
}

// 表格_置边颜色.
//
// color: 颜色.
func (t *Table) I置边颜色(color int) int {
	return xc.XTable_SetBorderColor(t.I句柄, color)
}

// 表格_置文本颜色.
//
// color: 颜色.
func (t *Table) I置文本颜色(color int) int {
	return xc.XTable_SetTextColor(t.I句柄, color)
}

// 表格_置字体.
//
// hFont: 字体句柄.
func (t *Table) I置字体(hFont int) int {
	return xc.XTable_SetFont(t.I句柄, hFont)
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
func (t *Table) I置项内填充(leftSize int, topSize int, rightSize int, bottomSize int) int {
	return xc.XTable_SetItemPadding(t.I句柄, leftSize, topSize, rightSize, bottomSize)
}

// 表格_置项文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
func (t *Table) I置项文本(iRow int, iCol int, pText string) int {
	return xc.XTable_SetItemText(t.I句柄, iRow, iCol, pText)
}

// 表格_置项字体.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hFont: 字体句柄.
func (t *Table) I置项字体(iRow int, iCol int, hFont int) int {
	return xc.XTable_SetItemFont(t.I句柄, iRow, iCol, hFont)
}

// 表格_置项文本对齐.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// nAlign: 对齐方式, I常量_文本对齐_, TextAlignFlag_, TextTrimming_.
func (t *Table) I置项文本对齐(iRow int, iCol int, nAlign xcc.I常量_文本对齐_) int {
	return xc.XTable_SetItemTextAlign(t.I句柄, iRow, iCol, nAlign)
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
func (t *Table) I置项文本色(iRow int, iCol int, color int, bColor bool) int {
	return xc.XTable_SetItemTextColor(t.I句柄, iRow, iCol, color, bColor)
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
func (t *Table) I置项背景色(iRow int, iCol int, color int, bColor bool) int {
	return xc.XTable_SetItemBkColor(t.I句柄, iRow, iCol, color, bColor)
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
func (t *Table) I置项线(iRow1 int, iCol1 int, iRow2 int, iCol2 int, nFlag int, color int) int {
	return xc.XTable_SetItemLine(t.I句柄, iRow1, iCol1, iRow2, iCol2, nFlag, color)
}

// 表格_置项标识.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// flag: 标识, I常量_表格标识_.
func (t *Table) I置项标识(iRow int, iCol int, flag xcc.I常量_表格标识_) int {
	return xc.XTable_SetItemFlag(t.I句柄, iRow, iCol, flag)
}

// 表格_取项坐标.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pRect: 接收返回坐标.
func (t *Table) I取项坐标(iRow int, iCol int, pRect *xc.RECT) bool {
	return xc.XTable_GetItemRect(t.I句柄, iRow, iCol, pRect)
}
