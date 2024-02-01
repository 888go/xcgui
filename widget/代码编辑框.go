package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 代码编辑框.
type Editor struct {
	Edit
}

// I代码编辑框_创建.
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
func I代码编辑框_创建(x int, y int, cx int, cy int, hParent int) *Editor {
	p := &Editor{}
	p.SetHandle(xc.XEditor_Create(x, y, cx, cy, hParent))
	return p
}

// I代码编辑框_从句柄创建对象.
func I代码编辑框_从句柄创建(handle int) *Editor {
	p := &Editor{}
	p.SetHandle(handle)
	return p
}

// I代码编辑框_从name创建对象, 失败返回nil.
func I代码编辑框_从name创建(name string) *Editor {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Editor{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I代码编辑框_从UID创建对象, 失败返回nil.
func I代码编辑框_从UID创建(nUID int) *Editor {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Editor{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I代码编辑框_从UID名称创建对象, 失败返回nil.
func I代码编辑框_从UID名称创建(name string) *Editor {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Editor{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 代码编辑框_启用空格选择自动匹配项.
//
// bEnable: 是否启用.
func (e *Editor) I启用空格选择自动匹配项(bEnable bool) int {
	return xc.XEidtor_EnableAutoMatchSpaseSelect(e.I句柄, bEnable)
}

// 代码编辑框_判断断点.
//
// iRow: 行索引.
func (e *Editor) I判断断点(iRow int) bool {
	return xc.XEditor_IsBreakpoint(e.I句柄, iRow)
}

// 代码编辑框_置断点.
//
// iRow: 行索引.
//
// bActivate: 是否激活.
func (e *Editor) I置断点(iRow int, bActivate bool) bool {
	return xc.XEditor_SetBreakpoint(e.I句柄, iRow, bActivate)
}

// 代码编辑框_移除断点.
//
// iRow: 行索引.
func (e *Editor) I移除断点(iRow int) bool {
	return xc.XEditor_RemoveBreakpoint(e.I句柄, iRow)
}

// 代码编辑框_清空断点.
func (e *Editor) I清空断点() int {
	return xc.XEditor_ClearBreakpoint(e.I句柄)
}

// 代码编辑框_置当前运行.
//
// iRow: 行索引.
func (e *Editor) I置当前运行(iRow int) bool {
	return xc.XEditor_SetRunRow(e.I句柄, iRow)
}

// 代码编辑框_取颜色信息.
//
// pInfo: 颜色信息结构体指针.
func (e *Editor) I取颜色信息(pInfo *xc.Editor_Color_) int {
	return xc.XEditor_GetColor(e.I句柄, pInfo)
}

// 代码编辑框_置颜色.
//
// pInfo: 颜色信息结构体指针.
func (e *Editor) I置颜色(pInfo *xc.Editor_Color_) int {
	return xc.XEditor_SetColor(e.I句柄, pInfo)
}

// 代码编辑框_置常量样式.
//
// iStyle: 样式.
func (e *Editor) I置常量样式(iStyle int) int {
	return xc.XEditor_SetStyleKeyword(e.I句柄, iStyle)
}

// 代码编辑框_置函数样式.
//
// iStyle: 样式.
func (e *Editor) I置函数样式(iStyle int) int {
	return xc.XEditor_SetStyleFunction(e.I句柄, iStyle)
}

// 代码编辑框_置变量样式.
//
// iStyle: 样式.
func (e *Editor) I置变量样式(iStyle int) int {
	return xc.XEditor_SetStyleVar(e.I句柄, iStyle)
}

// 代码编辑框_置数据类型样式.
//
// iStyle: 样式.
func (e *Editor) I置数据类型样式(iStyle int) int {
	return xc.XEditor_SetStyleDataType(e.I句柄, iStyle)
}

// 代码编辑框_置类样式.
//
// iStyle: 样式.
func (e *Editor) I置类样式(iStyle int) int {
	return xc.XEditor_SetStyleClass(e.I句柄, iStyle)
}

// 代码编辑框_置宏样式.
//
// iStyle: 样式.
func (e *Editor) I置宏样式(iStyle int) int {
	return xc.XEditor_SetStyleMacro(e.I句柄, iStyle)
}

// 代码编辑框_置字符串样式.
//
// iStyle: 样式.
func (e *Editor) I置字符串样式(iStyle int) int {
	return xc.XEditor_SetStyleString(e.I句柄, iStyle)
}

// 代码编辑框_置注释样式.
//
// iStyle: 样式.
func (e *Editor) I置注释样式(iStyle int) int {
	return xc.XEditor_SetStyleComment(e.I句柄, iStyle)
}

// 代码编辑框_置数字样式.
//
// iStyle: 样式.
func (e *Editor) I置数字样式(iStyle int) int {
	return xc.XEditor_SetStyleNumber(e.I句柄, iStyle)
}

// 代码编辑框_取断点数量.
func (e *Editor) I取断点数量() int {
	return xc.XEditor_GetBreakpointCount(e.I句柄)
}

// 代码编辑框_取全部断点, 返回实际获取断点数量.
//
// aPoints: 接收断点数组.
//
// nCount: 数组大小.
func (e *Editor) I取全部断点(aPoints *[]int32, nCount int) int {
	return xc.XEditor_GetBreakpoints(e.I句柄, aPoints, nCount)
}

// 代码编辑框_设置当前行, 跳过收缩行.
//
// iRow: 行索引.
func (e *Editor) I设置当前行(iRow int) int {
	return xc.XEditor_SetCurRow(e.I句柄, iRow)
}

// 代码编辑框_获取深度.
//
// iRow: 行索引.
func (e *Editor) I获取深度(iRow int) int {
	return xc.XEditor_GetDepth(e.I句柄, iRow)
}

// 代码编辑框_转换到展开行, 跳过收缩行.
//
// iRow: 行索引.
func (e *Editor) I转换到展开行(iRow int) int {
	return xc.XEditor_ToExpandRow(e.I句柄, iRow)
}

// 代码编辑框_展开扩展, 完全展开指定行, 例如:行包含在折叠内容中, 将其展开.
//
// iRow: 行索引.
func (e *Editor) I展开扩展(iRow int) int {
	return xc.XEditor_ExpandEx(e.I句柄, iRow)
}

// 代码编辑框_展开全部.
//
// bExpand: 是否展开.
func (e *Editor) I展开全部(bExpand bool) int {
	return xc.XEditor_ExpandAll(e.I句柄, bExpand)
}

// 代码编辑框_展开指定行.
//
// iRow: 行索引.
//
// bExpand: 是否展开.
func (e *Editor) I展开指定行(iRow int, bExpand bool) int {
	return xc.XEditor_Expand(e.I句柄, iRow, bExpand)
}

// 代码编辑框_添加关键字.
//
// pKey: 字符串.
//
// iStyle: 样式.
func (e *Editor) I添加关键字(pKey string, iStyle int) int {
	return xc.XEditor_AddKeyword(e.I句柄, pKey, iStyle)
}

// 代码编辑框_添加自动匹配字符串.
//
// pKey: 字符串.
func (e *Editor) I添加自动匹配字符串(pKey string) int {
	return xc.XEditor_AddConst(e.I句柄, pKey)
}

// 代码编辑框_添加自动匹配函数.
//
// pKey: 字符串.
func (e *Editor) I添加自动匹配函数(pKey string) int {
	return xc.XEditor_AddFunction(e.I句柄, pKey)
}

// 代码编辑框_添加排除定义变量关键字, 排除定义变量的关键字, 用于排除定义变量, 因为定义变量禁用自动匹配; 此关键字不加入自动匹配,仅用于排除定义变量.
//
// pKeyword: 字符串.
func (e *Editor) I添加排除定义变量关键字(pKeyword string) int {
	return xc.XEditor_AddExcludeDefVarKeyword(e.I句柄, pKeyword)
}

// 代码编辑框_获取折叠状态.
func (Editor *Edit) I获取折叠状态() string {
	return xc.XEditor_GetExpandState(Editor.I句柄)
}

// 代码编辑框_设置折叠状态.
//
// pString: .
func (Editor *Edit) I设置折叠状态(pString string) int {
	return xc.XEditor_SetExpandState(Editor.I句柄, pString)
}

// 代码编辑框_获取缩进.
//
// iRow: 行.
func (Editor *Edit) I获取缩进(iRow int) int {
	return xc.XEditor_GetIndentation(Editor.I句柄, iRow)
}

// 代码编辑框_是否为空行.
//
// iRow: 行.
func (Editor *Edit) I是否为空行(iRow int) int {
	return xc.XEidtor_IsEmptyRow(Editor.I句柄, iRow)
}

// 代码编辑框_置自动匹配结果显示模式.
//
// mode: 0:中英文, 1:英文, 3:中文.
func (Editor *Edit) I置自动匹配结果显示模式(mode int) int {
	return xc.XEditor_SetAutoMatchMode(Editor.I句柄, mode)
}

/*
以下都是事件
*/

type XE_EDITOR_MODIFY_ROWS func(iRow int, nRows int, pbHandled *bool) int                 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_MODIFY_ROWS1 func(hEle int, iRow int, nRows int, pbHandled *bool) int      // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_SETBREAKPOINT func(iRow int, bCheck bool, pbHandled *bool) int             // 代码编辑框_设置断点.
type XE_EDITOR_SETBREAKPOINT1 func(hEle int, iRow int, bCheck bool, pbHandled *bool) int  // 代码编辑框_设置断点.
type XE_EDITOR_REMOVEBREAKPOINT func(iRow int, pbHandled *bool) int                       // 代码编辑框_移除断点.
type XE_EDITOR_REMOVEBREAKPOINT1 func(hEle int, iRow int, pbHandled *bool) int            // 代码编辑框_移除断点.
type XE_EDITOR_AUTOMATCH_SELECT func(iRow int, nRows int, pbHandled *bool) int            // 代码编辑框_自动匹配选择.
type XE_EDITOR_AUTOMATCH_SELECT1 func(hEle int, iRow int, nRows int, pbHandled *bool) int // 代码编辑框_自动匹配选择.

// I事件_多行内容改变 事件 例如:区块注释操作, 区块缩进操作, 代码格式化.
func (e *Editor) I事件_多行内容改变(pFun XE_EDITOR_MODIFY_ROWS) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDITOR_MODIFY_ROWS, pFun)
}

// I事件_多行内容改变 事件 例如:区块注释操作, 区块缩进操作, 代码格式化.
func (e *Editor) I事件_多行内容改变1(pFun XE_EDITOR_MODIFY_ROWS1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDITOR_MODIFY_ROWS, pFun)
}

// I事件_设置断点.
func (e *Editor) I事件_设置断点(pFun XE_EDITOR_SETBREAKPOINT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDITOR_SETBREAKPOINT, pFun)
}

// 代码编辑框_  I事件_设置断点.
func (e *Editor) I事件_设置断点1(pFun XE_EDITOR_SETBREAKPOINT1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDITOR_SETBREAKPOINT, pFun)
}

// 代码编辑框_  I事件_移除断点.
func (e *Editor) I事件_移除断点(pFun XE_EDITOR_REMOVEBREAKPOINT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDITOR_REMOVEBREAKPOINT, pFun)
}

// 代码编辑框_  I事件_移除断点.
func (e *Editor) I事件_移除断点1(pFun XE_EDITOR_REMOVEBREAKPOINT1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDITOR_REMOVEBREAKPOINT, pFun)
}

// 代码编辑框_  I事件_自动匹配选择.
func (e *Editor) I事件_自动匹配选择(pFun XE_EDITOR_AUTOMATCH_SELECT) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDITOR_AUTOMATCH_SELECT, pFun)
}

// 代码编辑框_  I事件_自动匹配选择 .
func (e *Editor) I事件_自动匹配选择1(pFun XE_EDITOR_AUTOMATCH_SELECT1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDITOR_AUTOMATCH_SELECT, pFun)
}
