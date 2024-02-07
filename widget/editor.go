package 炫彩组件类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// 代码编辑框.
type Editor struct {
	Edit
}

// 代码编辑框_创建.
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
func X创建代码编辑框(元素x坐标 int, 元素y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) *Editor {
	p := &Editor{}
	p.X设置句柄(炫彩基类.X代码编辑框_创建(元素x坐标, 元素y坐标, 宽度, 高度, 父窗口句柄或元素句柄))
	return p
}

// 从句柄创建对象.
func X创建代码编辑框并按句柄(handle int) *Editor {
	p := &Editor{}
	p.X设置句柄(handle)
	return p
}

// 从name创建对象, 失败返回nil.
func X创建代码编辑框并按名称(name string) *Editor {
	handle := 炫彩基类.X取对象从名称(name)
	if handle > 0 {
		p := &Editor{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID创建对象, 失败返回nil.
func X创建代码编辑框并按UID(nUID int) *Editor {
	handle := 炫彩基类.X取对象从UID(nUID)
	if handle > 0 {
		p := &Editor{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 从UID名称创建对象, 失败返回nil.
func X创建代码编辑框并按UID名称(name string) *Editor {
	handle := 炫彩基类.X取对象从UID名称(name)
	if handle > 0 {
		p := &Editor{}
		p.X设置句柄(handle)
		return p
	}
	return nil
}

// 代码编辑框_启用空格选择自动匹配项.
//
// bEnable: 是否启用.
func (e *Editor) X启用空格选择自动匹配项(是否启用 bool) int {
	return 炫彩基类.X代码编辑框_启用空格选择自动匹配项(e.Handle, 是否启用)
}

// 代码编辑框_判断断点.
//
// iRow: 行索引.
func (e *Editor) X判断断点(行索引 int) bool {
	return 炫彩基类.X代码编辑框_判断断点(e.Handle, 行索引)
}

// 代码编辑框_置断点.
//
// iRow: 行索引.
//
// bActivate: 是否激活.
func (e *Editor) X置断点(行索引 int, 是否激活 bool) bool {
	return 炫彩基类.X代码编辑框_置断点(e.Handle, 行索引, 是否激活)
}

// 代码编辑框_移除断点.
//
// iRow: 行索引.
func (e *Editor) X移除断点(行索引 int) bool {
	return 炫彩基类.X代码编辑框_移除断点(e.Handle, 行索引)
}

// 代码编辑框_清空断点.
func (e *Editor) X清空断点() int {
	return 炫彩基类.X代码编辑框_清空断点(e.Handle)
}

// 代码编辑框_置当前运行.
//
// iRow: 行索引.
func (e *Editor) X置当前运行(行索引 int) bool {
	return 炫彩基类.X代码编辑框_置当前运行(e.Handle, 行索引)
}

// 代码编辑框_取颜色信息.
//
// pInfo: 颜色信息结构体指针.
func (e *Editor) X取颜色信息(颜色信息结构体指针 *炫彩基类.Editor_Color_) int {
	return 炫彩基类.X代码编辑框_取颜色信息(e.Handle, 颜色信息结构体指针)
}

// 代码编辑框_置颜色.
//
// pInfo: 颜色信息结构体指针.
func (e *Editor) X置颜色(颜色信息结构体指针 *炫彩基类.Editor_Color_) int {
	return 炫彩基类.X代码编辑框_置颜色(e.Handle, 颜色信息结构体指针)
}

// 代码编辑框_置常量样式.
//
// iStyle: 样式.
func (e *Editor) X置常量样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置常量样式(e.Handle, 样式)
}

// 代码编辑框_置函数样式.
//
// iStyle: 样式.
func (e *Editor) X置函数样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置函数样式(e.Handle, 样式)
}

// 代码编辑框_置变量样式.
//
// iStyle: 样式.
func (e *Editor) X置变量样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置变量样式(e.Handle, 样式)
}

// 代码编辑框_置数据类型样式.
//
// iStyle: 样式.
func (e *Editor) X置数据类型样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置数据类型样式(e.Handle, 样式)
}

// 代码编辑框_置类样式.
//
// iStyle: 样式.
func (e *Editor) X置类样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置类样式(e.Handle, 样式)
}

// 代码编辑框_置宏样式.
//
// iStyle: 样式.
func (e *Editor) X置宏样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置宏样式(e.Handle, 样式)
}

// 代码编辑框_置字符串样式.
//
// iStyle: 样式.
func (e *Editor) X置字符串样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置字符串样式(e.Handle, 样式)
}

// 代码编辑框_置注释样式.
//
// iStyle: 样式.
func (e *Editor) X置注释样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置注释样式(e.Handle, 样式)
}

// 代码编辑框_置数字样式.
//
// iStyle: 样式.
func (e *Editor) X置数字样式(样式 int) int {
	return 炫彩基类.X代码编辑框_置数字样式(e.Handle, 样式)
}

// 代码编辑框_取断点数量.
func (e *Editor) X取断点数量() int {
	return 炫彩基类.X代码编辑框_取断点数量(e.Handle)
}

// 代码编辑框_取全部断点, 返回实际获取断点数量.
//
// aPoints: 接收断点数组.
//
// nCount: 数组大小.
func (e *Editor) X取全部断点(接收断点数组 *[]int32, 数组大小 int) int {
	return 炫彩基类.X代码编辑框_取全部断点(e.Handle, 接收断点数组, 数组大小)
}

// 代码编辑框_设置当前行, 跳过收缩行.
//
// iRow: 行索引.
func (e *Editor) X设置当前行(行索引 int) int {
	return 炫彩基类.X代码编辑框_设置当前行(e.Handle, 行索引)
}

// 代码编辑框_获取深度.
//
// iRow: 行索引.
func (e *Editor) X获取深度(行索引 int) int {
	return 炫彩基类.X代码编辑框_获取深度(e.Handle, 行索引)
}

// 代码编辑框_转换到展开行, 跳过收缩行.
//
// iRow: 行索引.
func (e *Editor) X转换到展开行(行索引 int) int {
	return 炫彩基类.X代码编辑框_转换到展开行(e.Handle, 行索引)
}

// 代码编辑框_展开扩展, 完全展开指定行, 例如:行包含在折叠内容中, 将其展开.
//
// iRow: 行索引.
func (e *Editor) X展开EX(行索引 int) int {
	return 炫彩基类.X代码编辑框_展开EX(e.Handle, 行索引)
}

// 代码编辑框_展开全部.
//
// bExpand: 是否展开.
func (e *Editor) X展开全部(是否展开 bool) int {
	return 炫彩基类.X代码编辑框_展开全部(e.Handle, 是否展开)
}

// 代码编辑框_展开指定行.
//
// iRow: 行索引.
//
// bExpand: 是否展开.
func (e *Editor) X展开指定行(行索引 int, 是否展开 bool) int {
	return 炫彩基类.X代码编辑框_展开指定行(e.Handle, 行索引, 是否展开)
}

// 代码编辑框_添加关键字.
//
// pKey: 字符串.
//
// iStyle: 样式.
func (e *Editor) X添加关键字(字符串 string, 样式 int) int {
	return 炫彩基类.X代码编辑框_添加关键字(e.Handle, 字符串, 样式)
}

// 代码编辑框_添加自动匹配字符串.
//
// pKey: 字符串.
func (e *Editor) X添加自动匹配字符串(字符串 string) int {
	return 炫彩基类.X代码编辑框_添加自动匹配字符串(e.Handle, 字符串)
}

// 代码编辑框_添加自动匹配函数.
//
// pKey: 字符串.
func (e *Editor) X添加自动匹配函数(字符串 string) int {
	return 炫彩基类.X代码编辑框_添加自动匹配函数(e.Handle, 字符串)
}

// 代码编辑框_添加排除定义变量关键字, 排除定义变量的关键字, 用于排除定义变量, 因为定义变量禁用自动匹配; 此关键字不加入自动匹配,仅用于排除定义变量.
//
// pKeyword: 字符串.
func (e *Editor) X添加排除定义变量关键字(字符串 string) int {
	return 炫彩基类.X代码编辑框_添加排除定义变量关键字(e.Handle, 字符串)
}

// 代码编辑框_获取折叠状态.
func (e *Editor) X获取折叠状态() string {
	return 炫彩基类.X代码编辑框_获取折叠状态(e.Handle)
}

// 代码编辑框_设置折叠状态.
//
// pString: .
func (e *Editor) X设置折叠状态(pString string) int {
	return 炫彩基类.X代码编辑框_设置折叠状态(e.Handle, pString)
}

// 代码编辑框_获取缩进.
//
// iRow: 行.
func (e *Editor) X获取缩进(行 int) int {
	return 炫彩基类.X代码编辑框_获取缩进(e.Handle, 行)
}

// 代码编辑框_是否为空行.
//
// iRow: 行.
func (e *Editor) X是否为空行(行 int) int {
	return 炫彩基类.X代码编辑框_是否为空行(e.Handle, 行)
}

// 代码编辑框_置自动匹配结果显示模式.
//
// mode: 0:中英文, 1:英文, 3:中文.
func (e *Editor) X置自动匹配结果显示模式(模式 int) int {
	return 炫彩基类.X代码编辑框_置自动匹配结果显示模式(e.Handle, 模式)
}

/*
以下都是事件
*/

type XE_EDITOR_MODIFY_ROWS func(iRow int32, nRows int32, pbHandled *bool) int                 // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_MODIFY_ROWS1 func(hEle int, iRow int32, nRows int32, pbHandled *bool) int      // 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化. iRow: 开始行. nRows: 改变行数量.
type XE_EDITOR_SETBREAKPOINT func(iRow int32, bCheck bool, pbHandled *bool) int               // 代码编辑框_设置断点.
type XE_EDITOR_SETBREAKPOINT1 func(hEle int, iRow int32, bCheck bool, pbHandled *bool) int    // 代码编辑框_设置断点.
type XE_EDITOR_REMOVEBREAKPOINT func(iRow int32, pbHandled *bool) int                         // 代码编辑框_移除断点.
type XE_EDITOR_REMOVEBREAKPOINT1 func(hEle int, iRow int32, pbHandled *bool) int              // 代码编辑框_移除断点.
type XE_EDITOR_AUTOMATCH_SELECT func(iRow int32, nRows int32, pbHandled *bool) int            // 代码编辑框_自动匹配选择.
type XE_EDITOR_AUTOMATCH_SELECT1 func(hEle int, iRow int32, nRows int32, pbHandled *bool) int // 代码编辑框_自动匹配选择.

// 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化.
func (e *Editor) X事件_多行内容改变(pFun XE_EDITOR_MODIFY_ROWS) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDITOR_MODIFY_ROWS, pFun)
}

// 多行内容改变事件 例如:区块注释操作, 区块缩进操作, 代码格式化.
func (e *Editor) X事件_多行内容改变1(pFun XE_EDITOR_MODIFY_ROWS1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDITOR_MODIFY_ROWS, pFun)
}

// 代码编辑框_设置断点.
func (e *Editor) X事件_设置断点(pFun XE_EDITOR_SETBREAKPOINT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDITOR_SETBREAKPOINT, pFun)
}

// 代码编辑框_设置断点.
func (e *Editor) X事件_设置断点1(pFun XE_EDITOR_SETBREAKPOINT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDITOR_SETBREAKPOINT, pFun)
}

// 代码编辑框_移除断点.
func (e *Editor) X事件_移除断点(pFun XE_EDITOR_REMOVEBREAKPOINT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDITOR_REMOVEBREAKPOINT, pFun)
}

// 代码编辑框_移除断点.
func (e *Editor) X事件_移除断点1(pFun XE_EDITOR_REMOVEBREAKPOINT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDITOR_REMOVEBREAKPOINT, pFun)
}

// 代码编辑框_自动匹配选择.
func (e *Editor) X事件_自动匹配选择(pFun XE_EDITOR_AUTOMATCH_SELECT) bool {
	return 炫彩基类.X元素_注册事件C(e.Handle, 炫彩常量类.XE_EDITOR_AUTOMATCH_SELECT, pFun)
}

// 代码编辑框_自动匹配选择.
func (e *Editor) X事件_自动匹配选择1(pFun XE_EDITOR_AUTOMATCH_SELECT1) bool {
	return 炫彩基类.X元素_注册事件C1(e.Handle, 炫彩常量类.XE_EDITOR_AUTOMATCH_SELECT, pFun)
}
