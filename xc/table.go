package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
	
	"github.com/888go/xcgui/xcc"
)

// 表格_创建, 返回句柄.
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
func X表格_创建(按钮x坐标 int, 按钮y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xTable_Create.Call(uintptr(按钮x坐标), uintptr(按钮y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 表格_重置.
//
// hShape: 形状对象句柄.
//
// nRow: 行数量.
//
// nCol: 列数量.
func X表格_重置(形状对象句柄 int, 行数量 int, 列数量 int) int {
	r, _, _ := xTable_Reset.Call(uintptr(形状对象句柄), uintptr(行数量), uintptr(列数量))
	return int(r)
}

// 表格_组合行.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func X表格_组合行(形状对象句柄 int, 行索引 int, 列索引 int, 数量 int) int {
	r, _, _ := xTable_ComboRow.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(数量))
	return int(r)
}

// 表格_组合列.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// count: 数量.
func X表格_组合列(形状对象句柄 int, 行索引 int, 列索引 int, 数量 int) int {
	r, _, _ := xTable_ComboCol.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(数量))
	return int(r)
}

// 表格_置列宽.
//
// hShape: 形状对象句柄.
//
// iCol: 列索引.
//
// width: 宽度.
func X表格_置列宽(形状对象句柄 int, 列索引 int, 宽度 int) int {
	r, _, _ := xTable_SetColWidth.Call(uintptr(形状对象句柄), uintptr(列索引), uintptr(宽度))
	return int(r)
}

// 表格_置行高.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// height: 高度.
func X表格_置行高(形状对象句柄 int, 行索引 int, 高度 int) int {
	r, _, _ := xTable_SetRowHeight.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(高度))
	return int(r)
}

// 表格_置边颜色.
//
// hShape: 形状对象句柄.
//
// color: 颜色.
func X表格_置边颜色(形状对象句柄 int, 颜色 int) int {
	r, _, _ := xTable_SetBorderColor.Call(uintptr(形状对象句柄), uintptr(颜色))
	return int(r)
}

// 表格_置文本颜色.
//
// hShape: 形状对象句柄.
//
// color: 颜色.
func X表格_置文本颜色(形状对象句柄 int, 颜色 int) int {
	r, _, _ := xTable_SetTextColor.Call(uintptr(形状对象句柄), uintptr(颜色))
	return int(r)
}

// 表格_置字体.
//
// hShape: 形状对象句柄.
//
// hFont: 字体句柄.
func X表格_置字体(形状对象句柄 int, 字体句柄 int) int {
	r, _, _ := xTable_SetFont.Call(uintptr(形状对象句柄), uintptr(字体句柄))
	return int(r)
}

// 表格_置项内填充.
//
// hShape: 形状对象句柄.
//
// leftSize: 内填充大小.
//
// topSize: 内填充大小.
//
// rightSize: 内填充大小.
//
// bottomSize: 内填充大小.
func X表格_置项内填充(形状对象句柄 int, 左 int, 上 int, 右 int, 下 int) int {
	r, _, _ := xTable_SetItemPadding.Call(uintptr(形状对象句柄), uintptr(左), uintptr(上), uintptr(右), uintptr(下))
	return int(r)
}

// 表格_置项文本.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pText: 文本.
func X表格_置项文本(形状对象句柄 int, 行索引 int, 列索引 int, 文本 string) int {
	r, _, _ := xTable_SetItemText.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), 炫彩工具类.StrPtr(文本))
	return int(r)
}

// 表格_置项字体.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hFont: 字体句柄.
func X表格_置项字体(形状对象句柄 int, 行索引 int, 列索引 int, 字体句柄 int) int {
	r, _, _ := xTable_SetItemFont.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(字体句柄))
	return int(r)
}

// 表格_置项文本对齐.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// nAlign: 对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
func X表格_置项文本对齐(形状对象句柄 int, 行索引 int, 列索引 int, 对齐方式 炫彩常量类.TextFormatFlag_) int {
	r, _, _ := xTable_SetItemTextAlign.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(对齐方式))
	return int(r)
}

// 表格_置项文本色.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func X表格_置项文本色(形状对象句柄 int, 行索引 int, 列索引 int, 颜色 int, 是否使用 bool) int {
	r, _, _ := xTable_SetItemTextColor.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(颜色), 炫彩工具类.BoolPtr(是否使用))
	return int(r)
}

// 表格_置项背景色.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// color: 颜色.
//
// bColor: 是否使用.
func X表格_置项背景色(形状对象句柄 int, 行索引 int, 列索引 int, 颜色 int, 是否使用 bool) int {
	r, _, _ := xTable_SetItemBkColor.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(颜色), 炫彩工具类.BoolPtr(是否使用))
	return int(r)
}

// 表格_置项线.
//
// hShape: 形状对象句柄.
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
func X表格_置项线(形状对象句柄 int, 行索引1 int, 列索引1 int, 行索引2 int, 列索引2 int, 标识 int, 颜色 int) int {
	r, _, _ := xTable_SetItemLine.Call(uintptr(形状对象句柄), uintptr(行索引1), uintptr(列索引1), uintptr(行索引2), uintptr(列索引2), uintptr(标识), uintptr(颜色))
	return int(r)
}

// 表格_置项标识.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// flag: 标识, Table_Flag_.
func X表格_置项标识(形状对象句柄 int, 行索引 int, 列索引 int, 标识 炫彩常量类.Table_Flag_) int {
	r, _, _ := xTable_SetItemFlag.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(标识))
	return int(r)
}

// 表格_取项坐标.
//
// hShape: 形状对象句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pRect: 接收返回坐标.
func X表格_取项坐标(形状对象句柄 int, 行索引 int, 列索引 int, 接收返回坐标 *RECT) bool {
	r, _, _ := xTable_GetItemRect.Call(uintptr(形状对象句柄), uintptr(行索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返回坐标)))
	return r != 0
}
