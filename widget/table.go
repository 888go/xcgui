package widget

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Table 表格.
type Table struct {
	Shape
}

// 创建表格
//
// x: 按钮x坐标.
//
// y: 按钮y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父窗口句柄或元素句柄.
func NewTable(x int, y int, cx int, cy int, hParent int) *Table {
	p := &Table{}
	p.SetHandle(xc.XTable_Create(x, y, cx, cy, hParent))
	return p
}

// 创建表格并按句柄
func NewTableByHandle(handle int) *Table {
	p := &Table{}
	p.SetHandle(handle)
	return p
}

// 创建表格并按名称, 失败返回nil.
func NewTableByName(name string) *Table {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Table{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建表格并按UID 失败返回nil.
func NewTableByUID(nUID int) *Table {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Table{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 创建表格并按UID名称 失败返回nil.
func NewTableByUIDName(name string) *Table {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Table{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 重置.
//
// nRow: 行数量.
//
// nCol: 列数量.
func (t *Table) Reset(nRow int, nCol int) int {
	return xc.XTable_Reset(t.Handle, nRow, nCol)
}

// 组合行.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) ComboRow(iRow int, iCol int, count int) int {
	return xc.XTable_ComboRow(t.Handle, iRow, iCol, count)
}

// 组合列.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func (t *Table) ComboCol(iRow int, iCol int, count int) int {
	return xc.XTable_ComboCol(t.Handle, iRow, iCol, count)
}

// 置列宽.
//
// iCol: 列索引.
//
// width: 宽度.
func (t *Table) SetColWidth(iCol int, width int) int {
	return xc.XTable_SetColWidth(t.Handle, iCol, width)
}

// 置行高.
//
// iRow: 行索引.
//
// height: 高度.
func (t *Table) SetRowHeight(iRow int, height int) int {
	return xc.XTable_SetRowHeight(t.Handle, iRow, height)
}

// 置边颜色.
//
// color: 颜色.
func (t *Table) SetBorderColor(color int) int {
	return xc.XTable_SetBorderColor(t.Handle, color)
}

// 置文本颜色.
//
// color: 颜色.
func (t *Table) SetTextColor(color int) int {
	return xc.XTable_SetTextColor(t.Handle, color)
}

// 置字体.
//
// hFont: 字体句柄.
func (t *Table) SetFont(hFont int) int {
	return xc.XTable_SetFont(t.Handle, hFont)
}

// 置项内填充.
//
// leftSize: 左.
//
// topSize: 上.
//
// rightSize: 右.
//
// bottomSize: 下.
func (t *Table) SetItemPadding(leftSize int, topSize int, rightSize int, bottomSize int) int {
	return xc.XTable_SetItemPadding(t.Handle, leftSize, topSize, rightSize, bottomSize)
}

// 置项文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
func (t *Table) SetItemText(iRow int, iCol int, pText string) int {
	return xc.XTable_SetItemText(t.Handle, iRow, iCol, pText)
}

// 置项字体.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hFont: 字体句柄.
func (t *Table) SetItemFont(iRow int, iCol int, hFont int) int {
	return xc.XTable_SetItemFont(t.Handle, iRow, iCol, hFont)
}

// 置项文本对齐.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// nAlign: 对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func (t *Table) SetItemTextAlign(iRow int, iCol int, nAlign xcc.TextFormatFlag_) int {
	return xc.XTable_SetItemTextAlign(t.Handle, iRow, iCol, nAlign)
}

// 置项文本色.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func (t *Table) SetItemTextColor(iRow int, iCol int, color int, bColor bool) int {
	return xc.XTable_SetItemTextColor(t.Handle, iRow, iCol, color, bColor)
}

// 置项背景色.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func (t *Table) SetItemBkColor(iRow int, iCol int, color int, bColor bool) int {
	return xc.XTable_SetItemBkColor(t.Handle, iRow, iCol, color, bColor)
}

// 置项线.
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
func (t *Table) SetItemLine(iRow1 int, iCol1 int, iRow2 int, iCol2 int, nFlag int, color int) int {
	return xc.XTable_SetItemLine(t.Handle, iRow1, iCol1, iRow2, iCol2, nFlag, color)
}

// 置项标识.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// flag: 标识, Table_Flag_.
func (t *Table) SetItemFlag(iRow int, iCol int, flag xcc.Table_Flag_) int {
	return xc.XTable_SetItemFlag(t.Handle, iRow, iCol, flag)
}

// 取项坐标.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pRect: 接收返回坐标.
func (t *Table) GetItemRect(iRow int, iCol int, pRect *xc.RECT) bool {
	return xc.XTable_GetItemRect(t.Handle, iRow, iCol, pRect)
}
