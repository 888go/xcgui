package 炫彩基类

import (
	"github.com/888go/xcgui/common"
	"unsafe"
)

// 代码编辑框_创建, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// hParent: 父为窗口句柄或元素句柄.
func X代码编辑框_创建(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xEditor_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 代码编辑框_启用空格选择自动匹配项.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X代码编辑框_启用空格选择自动匹配项(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEidtor_EnableAutoMatchSpaseSelect.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 代码编辑框_判断断点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_判断断点(元素句柄 int, 行索引 int) bool {
	r, _, _ := xEditor_IsBreakpoint.Call(uintptr(元素句柄), uintptr(行索引))
	return r != 0
}

// 代码编辑框_置断点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// bActivate: 是否激活.
func X代码编辑框_置断点(元素句柄 int, 行索引 int, 是否激活 bool) bool {
	r, _, _ := xEditor_SetBreakpoint.Call(uintptr(元素句柄), uintptr(行索引), 炫彩工具类.BoolPtr(是否激活))
	return r != 0
}

// 代码编辑框_移除断点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_移除断点(元素句柄 int, 行索引 int) bool {
	r, _, _ := xEditor_RemoveBreakpoint.Call(uintptr(元素句柄), uintptr(行索引))
	return r != 0
}

// 代码编辑框_清空断点.
//
// hEle: 元素句柄.
func X代码编辑框_清空断点(元素句柄 int) int {
	r, _, _ := xEditor_ClearBreakpoint.Call(uintptr(元素句柄))
	return int(r)
}

// 代码编辑框_置当前运行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_置当前运行(元素句柄 int, 行索引 int) bool {
	r, _, _ := xEditor_SetRunRow.Call(uintptr(元素句柄), uintptr(行索引))
	return r != 0
}

// 代码编辑框_取颜色信息.
//
// hEle: 元素句柄.
//
// pInfo: 颜色信息结构体指针.
func X代码编辑框_取颜色信息(元素句柄 int, 颜色信息结构体指针 *Editor_Color_) int {
	r, _, _ := xEditor_GetColor.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(颜色信息结构体指针)))
	return int(r)
}

// 代码编辑框_置颜色.
//
// hEle: 元素句柄.
//
// pInfo: 颜色信息结构体指针.
func X代码编辑框_置颜色(元素句柄 int, 颜色信息结构体指针 *Editor_Color_) int {
	r, _, _ := xEditor_SetColor.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(颜色信息结构体指针)))
	return int(r)
}

// 代码编辑框_置常量样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置常量样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleKeyword.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置函数样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置函数样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleFunction.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置变量样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置变量样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleVar.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置数据类型样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置数据类型样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleDataType.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置类样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置类样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleClass.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置宏样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置宏样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleMacro.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置字符串样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置字符串样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleString.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置注释样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置注释样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleComment.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_置数字样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X代码编辑框_置数字样式(元素句柄 int, 样式 int) int {
	r, _, _ := xEditor_SetStyleNumber.Call(uintptr(元素句柄), uintptr(样式))
	return int(r)
}

// 代码编辑框_取断点数量.
//
// hEle: 元素句柄.
func X代码编辑框_取断点数量(元素句柄 int) int {
	r, _, _ := xEditor_GetBreakpointCount.Call(uintptr(元素句柄))
	return int(r)
}

// 代码编辑框_取全部断点, 返回实际获取断点数量.
//
// hEle: 元素句柄.
//
// aPoints: 接收断点切片.
//
// nCount: 切片大小.
func X代码编辑框_取全部断点(元素句柄 int, 接收断点切片 *[]int32, 切片大小 int) int {
	if 切片大小 < 1 {
		return 0
	}
	*接收断点切片 = make([]int32, 切片大小)
	r, _, _ := xEditor_GetBreakpoints.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(&(*接收断点切片)[0])), uintptr(切片大小))
	return int(r)
}

// 代码编辑框_设置当前行, 跳过收缩行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_设置当前行(元素句柄 int, 行索引 int) int {
	r, _, _ := xEditor_SetCurRow.Call(uintptr(元素句柄), uintptr(行索引))
	return int(r)
}

// 代码编辑框_获取深度.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_获取深度(元素句柄 int, 行索引 int) int {
	r, _, _ := xEditor_GetDepth.Call(uintptr(元素句柄), uintptr(行索引))
	return int(r)
}

// 代码编辑框_转换到展开行, 跳过收缩行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_转换到展开行(元素句柄 int, 行索引 int) int {
	r, _, _ := xEditor_ToExpandRow.Call(uintptr(元素句柄), uintptr(行索引))
	return int(r)
}

// 代码编辑框_展开扩展, 完全展开指定行, 例如:行包含在折叠内容中, 将其展开.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X代码编辑框_展开EX(元素句柄 int, 行索引 int) int {
	r, _, _ := xEditor_ExpandEx.Call(uintptr(元素句柄), uintptr(行索引))
	return int(r)
}

// 代码编辑框_展开全部.
//
// hEle: 元素句柄.
//
// bExpand: 是否展开.
func X代码编辑框_展开全部(元素句柄 int, 是否展开 bool) int {
	r, _, _ := xEditor_ExpandAll.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否展开))
	return int(r)
}

// 代码编辑框_展开指定行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// bExpand: 是否展开.
func X代码编辑框_展开指定行(元素句柄 int, 行索引 int, 是否展开 bool) int {
	r, _, _ := xEditor_Expand.Call(uintptr(元素句柄), uintptr(行索引), 炫彩工具类.BoolPtr(是否展开))
	return int(r)
}

// 代码编辑框_添加关键字.
//
// hEle: 元素句柄.
//
// pKey: 字符串.
//
// iStyle: 样式.
func X代码编辑框_添加关键字(元素句柄 int, 字符串 string, 样式 int) int {
	r, _, _ := xEditor_AddKeyword.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串), uintptr(样式))
	return int(r)
}

// 代码编辑框_添加自动匹配字符串.
//
// hEle: 元素句柄.
//
// pKey: 字符串.
func X代码编辑框_添加自动匹配字符串(元素句柄 int, 字符串 string) int {
	r, _, _ := xEditor_AddConst.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 代码编辑框_添加自动匹配函数.
//
// hEle: 元素句柄.
//
// pKey: 字符串.
func X代码编辑框_添加自动匹配函数(元素句柄 int, 字符串 string) int {
	r, _, _ := xEditor_AddFunction.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 代码编辑框_添加排除定义变量关键字, 排除定义变量的关键字, 用于排除定义变量, 因为定义变量禁用自动匹配; 此关键字不加入自动匹配,仅用于排除定义变量.
//
// hEle: 元素句柄.
//
// pKeyword: 字符串.
func X代码编辑框_添加排除定义变量关键字(元素句柄 int, 字符串 string) int {
	r, _, _ := xEditor_AddExcludeDefVarKeyword.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 代码编辑框_获取折叠状态.
//
// hEle: 元素句柄.
func X代码编辑框_获取折叠状态(元素句柄 int) string {
	r, _, _ := xEditor_GetExpandState.Call(uintptr(元素句柄))
	return A2W(r)
}

// 代码编辑框_设置折叠状态.
//
// hEle: 元素句柄.
//
// pString: .
func X代码编辑框_设置折叠状态(元素句柄 int, pString string) int {
	r, _, _ := xEditor_SetExpandState.Call(uintptr(元素句柄), W2A(pString))
	return int(r)
}

// 代码编辑框_获取缩进.
//
// hEle: 元素句柄.
//
// iRow: 行.
func X代码编辑框_获取缩进(元素句柄 int, 行 int) int {
	r, _, _ := xEditor_GetIndentation.Call(uintptr(元素句柄), uintptr(行))
	return int(r)
}

// 代码编辑框_是否为空行.
//
// hEle: 元素句柄.
//
// iRow: 行.
func X代码编辑框_是否为空行(元素句柄 int, 行 int) int {
	r, _, _ := xEidtor_IsEmptyRow.Call(uintptr(元素句柄), uintptr(行))
	return int(r)
}

// 代码编辑框_置自动匹配结果显示模式.
//
// hEle: 元素句柄.
//
// mode: 0:中英文, 1:英文, 3:中文.
func X代码编辑框_置自动匹配结果显示模式(元素句柄 int, 模式 int) int {
	r, _, _ := xEditor_SetAutoMatchMode.Call(uintptr(元素句柄), uintptr(模式))
	return int(r)
}
