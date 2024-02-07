package 炫彩基类

import (
	"syscall"
	"unsafe"

	"github.com/888go/xcgui/common"

	"github.com/888go/xcgui/xcc"
)

// 编辑框_创建, 返回元素句柄.
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
func X编辑框_创建(x坐标 int, y坐标 int, 宽度 int, 高度 int, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xEdit_Create.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 编辑框_创建扩展, 返回元素句柄.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// nType: 类型, Edit_Type_.
//
// hParent: 父为窗口句柄或元素句柄.
func X编辑框_创建EX(x坐标 int, y坐标 int, 宽度 int, 高度 int, 类型 炫彩常量类.Edit_Type_, 父窗口句柄或元素句柄 int) int {
	r, _, _ := xEdit_CreateEx.Call(uintptr(x坐标), uintptr(y坐标), uintptr(宽度), uintptr(高度), uintptr(类型), uintptr(父窗口句柄或元素句柄))
	return int(r)
}

// 编辑框_启用自动换行.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X编辑框_启用自动换行(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEdit_EnableAutoWrap.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 编辑框_启用只读.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X编辑框_启用只读(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEdit_EnableReadOnly.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 编辑框_启用多行.
//
// hEle:.
//
// bEnable:.
func X编辑框_启用多行(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEdit_EnableMultiLine.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 编辑框_启用密码, 启用密码模式(只支持默认类型编辑框).
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X编辑框_启用密码(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEdit_EnablePassword.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 编辑框_启用自动选择, 当获得焦点时,自动选择所有内容.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X编辑框_启用自动选择(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEdit_EnableAutoSelAll.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 编辑框_启用自动取消选择, 当失去焦点时自动取消选择.
//
// hEle: 元素句柄.
//
// bEnable: 是否启用.
func X编辑框_启用自动取消选择(元素句柄 int, 是否启用 bool) int {
	r, _, _ := xEdit_EnableAutoCancelSel.Call(uintptr(元素句柄), 炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 编辑框_是否只读.
//
// hEle: 元素句柄.
func X编辑框_是否只读(元素句柄 int) bool {
	r, _, _ := xEdit_IsReadOnly.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_是否多行.
//
// hEle: 元素句柄.
func X编辑框_是否多行(元素句柄 int) bool {
	r, _, _ := xEdit_IsMultiLine.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_是否密码.
//
// hEle: 元素句柄.
func X编辑框_是否密码(元素句柄 int) bool {
	r, _, _ := xEdit_IsPassword.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_是否自动换行.
//
// hEle: 元素句柄.
func X编辑框_是否自动换行(元素句柄 int) bool {
	r, _, _ := xEdit_IsAutoWrap.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_判断为空.
//
// hEle: 元素句柄.
func X编辑框_判断为空(元素句柄 int) bool {
	r, _, _ := xEdit_IsEmpty.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_是否在选择区域.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func X编辑框_是否在选择区域(元素句柄 int, 行索引 int, 列索引 int) bool {
	r, _, _ := xEdit_IsInSelect.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引))
	return r != 0
}

// 编辑框_取总行数.
//
// hEle: 元素句柄.
func X编辑框_取总行数(元素句柄 int) int {
	r, _, _ := xEdit_GetRowCount.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_取数据, 包含文本或非文本内容.
//
// hEle: 元素句柄.
func X编辑框_取数据(元素句柄 int) Edit_Data_Copy_ {
	r, _, _ := xEdit_GetData.Call(uintptr(元素句柄))
	return *(*Edit_Data_Copy_)(unsafe.Pointer(&r))
}

// 编辑框_添加数据.
//
// hEle: 元素句柄.
//
// pData: 数据结构.
//
// styleTable: 样式表.
//
// nStyleCount: 样式数量.
func X编辑框_添加数据(元素句柄 int, 数据结构 *Edit_Data_Copy_, 样式表 []uint16, 样式数量 int) int {
	r, _, _ := xEdit_AddData.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(数据结构)), uintptr(unsafe.Pointer(&样式表[0])), uintptr(样式数量))
	return int(r)
}

// 编辑框_释放数据.
//
// pData: 数据结构.
func X编辑框_释放数据(数据结构 *Edit_Data_Copy_) int {
	r, _, _ := xEdit_FreeData.Call(uintptr(unsafe.Pointer(数据结构)))
	return int(r)
}

// 编辑框_置默认文本, 当内容为空时, 显示默认文本.
//
// hEle: 元素句柄.
//
// pString: 文本内容.
func X编辑框_置默认文本(元素句柄 int, 文本内容 string) int {
	r, _, _ := xEdit_SetDefaultText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(文本内容))
	return int(r)
}

// 编辑框_置默认文本颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色值.
func X编辑框_置默认文本颜色(元素句柄 int, ABGR颜色值 int) int {
	r, _, _ := xEdit_SetDefaultTextColor.Call(uintptr(元素句柄), uintptr(ABGR颜色值))
	return int(r)
}

// 编辑框_置密码字符.
//
// hEle: 元素句柄.
//
// ch: 字符.
func X编辑框_置密码字符(元素句柄 int, 字符 int) int {
	r, _, _ := xEdit_SetPasswordCharacter.Call(uintptr(元素句柄), uintptr(字符))
	return int(r)
}

// 编辑框_置文本对齐, 单行模式下有效.
//
// hEle: 元素句柄.
//
// align: 对齐方式, Edit_TextAlign_Flag_.
func X编辑框_置文本对齐(元素句柄 int, 对齐方式 炫彩常量类.Edit_TextAlign_Flag_) int {
	r, _, _ := xEdit_SetTextAlign.Call(uintptr(元素句柄), uintptr(对齐方式))
	return int(r)
}

// 编辑框_置TAB空格.
//
// hEle: 元素句柄.
//
// nSpace: 空格数量.
func X编辑框_置TAB空格(元素句柄 int, 空格数量 int) int {
	r, _, _ := xEdit_SetTabSpace.Call(uintptr(元素句柄), uintptr(空格数量))
	return int(r)
}

// 编辑框_置文本.
//
// hEle: 元素句柄.
//
// pString: 字符串.
func X编辑框_置文本(元素句柄 int, 字符串 string) int {
	r, _, _ := xEdit_SetText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 编辑框_置文本整数.
//
// hEle: 元素句柄.
//
// nValue: 整数值.
func X编辑框_置文本整数(元素句柄 int, 值 int) int {
	r, _, _ := xEdit_SetTextInt.Call(uintptr(元素句柄), uintptr(值))
	return int(r)
}

// 编辑框_取文本, 不包含非文本内容, 返回实际接收文本长度.
//
// hEle: 元素句柄.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 内存大小. 例: xc.XEdit_GetLength()+1 .
func X编辑框_取文本(元素句柄 int, 接收文本内存指针 *string, 内存大小 int) int {
	buf := make([]uint16, 内存大小)
	r, _, _ := xEdit_GetText.Call(uintptr(元素句柄), 炫彩工具类.Uint16SliceDataPtr(&buf), uintptr(内存大小))
	*接收文本内存指针 = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 编辑框_取文本行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 接收文本内存块长度. 例: xc.XEdit_GetLengthRow()+1 .
func X编辑框_取文本行(元素句柄 int, 行索引 int, 接收文本内存指针 *string, 接收文本内存块长度 int) int {
	buf := make([]uint16, 接收文本内存块长度)
	r, _, _ := xEdit_GetTextRow.Call(uintptr(元素句柄), uintptr(行索引), 炫彩工具类.Uint16SliceDataPtr(&buf), uintptr(接收文本内存块长度))
	*接收文本内存指针 = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 编辑框_取内容长度, 包含非文本内容.
//
// hEle: 元素句柄.
func X编辑框_取内容长度(元素句柄 int) int {
	r, _, _ := xEdit_GetLength.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_取内容长度行, 包含非文本内容.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X编辑框_取内容长度行(元素句柄 int, 行索引 int) int {
	r, _, _ := xEdit_GetLengthRow.Call(uintptr(元素句柄), uintptr(行索引))
	return int(r)
}

// 编辑框_取字符, 返回指定位置字符.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func X编辑框_取字符(元素句柄 int, 行索引 int, 列索引 int) int {
	r, _, _ := xEdit_GetAt.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引))
	return int(r)
}

// 编辑框_插入文本.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
func X编辑框_插入文本(元素句柄 int, 行索引 int, 列索引 int, 字符串 string) int {
	r, _, _ := xEdit_InsertText.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// XEdit_AddTextUser 编辑框_插入文本模拟用户操作, 自动刷新UI, 支持撤销/恢复.
//
//	@param hEle 元素句柄.
//	@param pString 字符串.
//	@return int
func X编辑框_插入文本模拟用户操作(元素句柄 int, 字符串 string) int {
	r, _, _ := xEdit_AddTextUser.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 编辑框_添加文本.
//
// hEle: 元素句柄.
//
// pString: 字符串.
func X编辑框_添加文本(元素句柄 int, 字符串 string) int {
	r, _, _ := xEdit_AddText.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串))
	return int(r)
}

// 编辑框_添加文本扩展.
//
// hEle: 元素句柄.
//
// pString: 字符串.
//
// iStyle: 样式索引.
func X编辑框_添加文本EX(元素句柄 int, 字符串 string, 样式索引 int) int {
	r, _, _ := xEdit_AddTextEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字符串), uintptr(样式索引))
	return int(r)
}

// 编辑框_添加对象, 例如: 字体, 图片, UI对象, 返回样式索引.
//
// hEle: 元素句柄.
//
// hObj: 对象句柄.
func X编辑框_添加对象(元素句柄 int, 对象句柄 int) int {
	r, _, _ := xEdit_AddObject.Call(uintptr(元素句柄), uintptr(对象句柄))
	return int(r)
}

// 编辑框_添加对象从样式, 当样式为图片时有效.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
func X编辑框_添加对象从样式(元素句柄 int, 样式索引 int) int {
	r, _, _ := xEdit_AddByStyle.Call(uintptr(元素句柄), uintptr(样式索引))
	return int(r)
}

// 编辑框_添加样式, 返回样式索引.
//
// hEle: 元素句柄.
//
// hFont_image_Obj: 字体.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func X编辑框_添加样式(元素句柄 int, 字体 int, 颜色 int, 是否使用颜色 bool) int {
	r, _, _ := xEdit_AddStyle.Call(uintptr(元素句柄), uintptr(字体), uintptr(颜色), 炫彩工具类.BoolPtr(是否使用颜色))
	return int(r)
}

// 编辑框_添加样式扩展, 返回样式索引.
//
// hEle: 元素句柄.
//
// fontName: 字体名称.
//
// fontSize: 字体大小.
//
// fontStyle: 字体样式, FontStyle_.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func X编辑框_添加样式EX(元素句柄 int, 字体名称 string, 字体大小 int, 字体样式 炫彩常量类.FontStyle_, 颜色 int, 是否使用颜色 bool) int {
	r, _, _ := xEdit_AddStyleEx.Call(uintptr(元素句柄), 炫彩工具类.StrPtr(字体名称), uintptr(字体大小), uintptr(字体样式), uintptr(颜色), 炫彩工具类.BoolPtr(是否使用颜色))
	return int(r)
}

// 编辑框_取样式信息.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
//
// info: 返回样式信息.
func X编辑框_取样式信息(元素句柄 int, 样式索引 int, 返回样式信息 *Edit_Style_Info_) bool {
	r, _, _ := xEdit_GetStyleInfo.Call(uintptr(元素句柄), uintptr(样式索引), uintptr(unsafe.Pointer(返回样式信息)))
	return r != 0
}

// 编辑框_置当前样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
func X编辑框_置当前样式(元素句柄 int, 样式索引 int) int {
	r, _, _ := xEdit_SetCurStyle.Call(uintptr(元素句柄), uintptr(样式索引))
	return int(r)
}

// 编辑框_置插入符颜色.
//
// hEle: 元素句柄.
//
// color: 颜色.
func X编辑框_置插入符颜色(元素句柄 int, 颜色 int) int {
	r, _, _ := xEdit_SetCaretColor.Call(uintptr(元素句柄), uintptr(颜色))
	return int(r)
}

// 编辑框_置插入符宽度.
//
// hEle: 元素句柄.
//
// nWidth: 宽度.
func X编辑框_置插入符宽度(元素句柄 int, 宽度 int) int {
	r, _, _ := xEdit_SetCaretWidth.Call(uintptr(元素句柄), uintptr(宽度))
	return int(r)
}

// 编辑框_置选择背景颜色.
//
// hEle: 元素句柄.
//
// color: ABGR 颜色.
func X编辑框_置选择背景颜色(元素句柄 int, ABGR颜色 int) int {
	r, _, _ := xEdit_SetSelectBkColor.Call(uintptr(元素句柄), uintptr(ABGR颜色))
	return int(r)
}

// 编辑框_置默认行高.
//
// hEle: 元素句柄.
//
// nHeight: 行高.
func X编辑框_置默认行高(元素句柄 int, 行高 int) int {
	r, _, _ := xEdit_SetRowHeight.Call(uintptr(元素句柄), uintptr(行高))
	return int(r)
}

// 编辑框_置指定行高度, 类型为 Edit_Type_Richedit 支持指定不同行高.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// nHeight: 高度.
func X编辑框_置指定行高度(元素句柄 int, 行索引 int, 高度 int) int {
	r, _, _ := xEdit_SetRowHeightEx.Call(uintptr(元素句柄), uintptr(行索引), uintptr(高度))
	return int(r)
}

// XEdit_SetCurPos 编辑框_置当前位置.
//
//	@param hEle: 元素句柄.
//	@param iRow: 行索引.
//	@return int
func X编辑框_置当前位置(元素句柄 int, 行索引 int) int {
	r, _, _ := xEdit_SetCurPos.Call(uintptr(元素句柄), uintptr(行索引))
	return int(r)
}

// 编辑框_取当前位置点, 返回范围位置点.
//
// hEle: 元素句柄.
func X编辑框_取当前位置点(元素句柄 int) int {
	r, _, _ := xEdit_GetCurPos.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_取当前行, 返回行索引.
//
// hEle: 元素句柄.
func X编辑框_取当前行(元素句柄 int) int {
	r, _, _ := xEdit_GetCurRow.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_取当前列, 返回列索引.
//
// hEle: 元素句柄.
func X编辑框_取当前列(元素句柄 int) int {
	r, _, _ := xEdit_GetCurCol.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_取坐标点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pOut: 接收返回坐标点.
func X编辑框_取坐标点(元素句柄 int, 行索引 int, 列索引 int, 接收返回坐标点 *POINT) int {
	r, _, _ := xEdit_GetPoint.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引), uintptr(unsafe.Pointer(接收返回坐标点)))
	return int(r)
}

// 编辑框_自动滚动, 视图自动滚动到当前插入符位置.
//
// hEle: 元素句柄.
func X编辑框_自动滚动(元素句柄 int) bool {
	r, _, _ := xEdit_AutoScroll.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_自动滚动扩展, 视图自动滚动到指定位置.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func X编辑框_自动滚动EX(元素句柄 int, 行索引 int, 列索引 int) bool {
	r, _, _ := xEdit_AutoScrollEx.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引))
	return r != 0
}

// XEdit_PosToRowCol 编辑框_转换位置, 转换位置点到行列.
//
//	@param hEle 元素句柄.
//	@param iPos 位置点.
//	@param pInfo 行列.
//	@return int
func X编辑框_转换位置(元素句柄 int, 位置点 int, 行列 *Position_) int {
	r, _, _ := xEdit_PosToRowCol.Call(uintptr(元素句柄), uintptr(位置点), uintptr(unsafe.Pointer(行列)))
	return int(r)
}

// 编辑框_选择全部.
//
// hEle: 元素句柄.
func X编辑框_选择全部(元素句柄 int) bool {
	r, _, _ := xEdit_SelectAll.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_取消选择.
//
// hEle: 元素句柄.
func X编辑框_取消选择(元素句柄 int) bool {
	r, _, _ := xEdit_CancelSelect.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_删除选择内容.
//
// hEle: 元素句柄.
func X编辑框_删除选择内容(元素句柄 int) bool {
	r, _, _ := xEdit_DeleteSelect.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_置选择.
//
// hEle: 元素句柄.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func X编辑框_置选择(元素句柄 int, 起始行索引 int, 起始行列索引 int, 结束行索引 int, 结束行列索引 int) bool {
	r, _, _ := xEdit_SetSelect.Call(uintptr(元素句柄), uintptr(起始行索引), uintptr(起始行列索引), uintptr(结束行索引), uintptr(结束行列索引))
	return r != 0
}

// 编辑框_取选择文本, 不包括非文本内容, 返回接收文本内容实际长度.
//
// hEle: 元素句柄.
//
// pOut: 接收返回文本内容.
//
// nOutLen: 接收内存大小. xc.XEdit_GetSelectTextLength()+1 .
func X编辑框_取选择文本(元素句柄 int, 接收返回文本内容 *string, 接收内存大小 int) int {
	buf := make([]uint16, 接收内存大小)
	r, _, _ := xEdit_GetSelectText.Call(uintptr(元素句柄), 炫彩工具类.Uint16SliceDataPtr(&buf), uintptr(接收内存大小))
	*接收返回文本内容 = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 编辑框_取选择内容范围.
//
// hEle: 元素句柄.
//
// pBegin: 起始位置.
//
// pEnd: 结束位置.
func X编辑框_取选择内容范围(元素句柄 int, 起始位置 *Position_, 结束位置 *Position_) bool {
	r, _, _ := xEdit_GetSelectRange.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(起始位置)), uintptr(unsafe.Pointer(结束位置)))
	return r != 0
}

// 编辑框_取可视行范围.
//
// hEle: 元素句柄.
//
// piStart: 起始行索引.
//
// piEnd: 结束行索引.
func X编辑框_取可视行范围(元素句柄 int, 起始行索引 *int32, 结束行索引 *int32) int {
	r, _, _ := xEdit_GetVisibleRowRange.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(起始行索引)), uintptr(unsafe.Pointer(结束行索引)))
	return int(r)
}

// 编辑框_删除, 删除指定范围内容.
//
// hEle: 元素句柄.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func X编辑框_删除(元素句柄 int, 起始行索引 int, 起始行列索引 int, 结束行索引 int, 结束行列索引 int) bool {
	r, _, _ := xEdit_Delete.Call(uintptr(元素句柄), uintptr(起始行索引), uintptr(起始行列索引), uintptr(结束行索引), uintptr(结束行列索引))
	return r != 0
}

// 编辑框_删除行.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X编辑框_删除行(元素句柄 int, 行索引 int) bool {
	r, _, _ := xEdit_DeleteRow.Call(uintptr(元素句柄), uintptr(行索引))
	return r != 0
}

// 编辑框_剪贴板剪切.
//
// hEle: 元素句柄.
func X编辑框_剪贴板剪切(元素句柄 int) bool {
	r, _, _ := xEdit_ClipboardCut.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_剪贴板复制.
//
// hEle: 元素句柄.
func X编辑框_剪贴板复制(元素句柄 int) bool {
	r, _, _ := xEdit_ClipboardCopy.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_剪贴板粘贴.
//
// hEle: 元素句柄.
func X编辑框_剪贴板粘贴(元素句柄 int) bool {
	r, _, _ := xEdit_ClipboardPaste.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_撤销.
//
// hEle: 元素句柄.
func X编辑框_撤销(元素句柄 int) bool {
	r, _, _ := xEdit_Undo.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_恢复.
//
// hEle: 元素句柄.
func X编辑框_恢复(元素句柄 int) bool {
	r, _, _ := xEdit_Redo.Call(uintptr(元素句柄))
	return r != 0
}

// 编辑框_添加气泡开始, 当前行开始.
//
// hEle: 元素句柄.
//
// hImageAvatar: 头像.
//
// hImageBubble: 气泡背景.
//
// nFlag: 标志, Chat_Flag_.
func X编辑框_添加气泡开始(元素句柄 int, 头像 int, 气泡背景 int, 标志 炫彩常量类.Chat_Flag_) int {
	r, _, _ := xEdit_AddChatBegin.Call(uintptr(元素句柄), uintptr(头像), uintptr(气泡背景), uintptr(标志))
	return int(r)
}

// 编辑框_添加气泡结束, 当前行结束.
//
// hEle: 元素句柄.
func X编辑框_添加气泡结束(元素句柄 int) int {
	r, _, _ := xEdit_AddChatEnd.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_置气泡缩进, 设置聊天气泡内容缩进.
//
// hEle: 元素句柄.
//
// nIndentation: 缩进值.
func X编辑框_置气泡缩进(元素句柄 int, 缩进值 int) int {
	r, _, _ := xEdit_SetChatIndentation.Call(uintptr(元素句柄), uintptr(缩进值))
	return int(r)
}

// 编辑框_置行间隔, 设置行间隔大小, 多行模式有效.
//
// hEle: 元素句柄.
//
// nSpace: 行间隔大小.
func X编辑框_置行间隔(元素句柄 int, 行间隔大小 int) int {
	r, _, _ := xEdit_SetRowSpace.Call(uintptr(元素句柄), uintptr(行间隔大小))
	return int(r)
}

// 编辑框_置当前位置扩展.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func X编辑框_置当前位置EX(元素句柄 int, 行索引, 列索引 int32) int {
	r, _, _ := xEdit_SetCurPosEx.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引))
	return int(r)
}

// 编辑框_取当前位置扩展.
//
// hEle: 元素句柄.
//
// iRow: 返回行索引.
//
// iCol: 返回列索引.
func X编辑框_取当前位置EX(元素句柄 int, 返回行索引, 返回列索引 *int32) int {
	r, _, _ := xEdit_GetCurPosEx.Call(uintptr(元素句柄), uintptr(unsafe.Pointer(返回行索引)), uintptr(unsafe.Pointer(返回列索引)))
	return int(r)
}

// 编辑框_移动到末尾.
//
// hEle: 元素句柄.
func X编辑框_移动到末尾(元素句柄 int) int {
	r, _, _ := xEdit_MoveEnd.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_行列到位置, 返回位置点.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
func X编辑框_行列到位置(元素句柄 int, 行索引 int, 列索引 int) int {
	r, _, _ := xEdit_RowColToPos.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引))
	return int(r)
}

// 编辑框_置后备字体, 置中文字体. 如果已设置, 当遇到中文字符时使用后备字体, 解决不支持中文的字体的问题
//
// hEle: 元素句柄.
//
// hFont: 字体.
func X编辑框_置后备字体(元素句柄 int, 字体 int) int {
	r, _, _ := xEdit_SetBackFont.Call(uintptr(元素句柄), uintptr(字体))
	return int(r)
}

// 编辑框_释放样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式.
func X编辑框_释放样式(元素句柄 int, 样式 int) bool {
	r, _, _ := xEdit_ReleaseStyle.Call(uintptr(元素句柄), uintptr(样式))
	return r != 0
}

// 编辑框_修改样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
//
// hFont: 字体句柄.
//
// color: ABGR 颜色.
//
// bColor: 是否使用颜色.
func X编辑框_修改样式(元素句柄 int, 样式索引 int, 字体句柄 int, ABGR颜色 int, 是否使用颜色 bool) bool {
	r, _, _ := xEdit_ModifyStyle.Call(uintptr(元素句柄), uintptr(样式索引), uintptr(字体句柄), uintptr(ABGR颜色), 炫彩工具类.BoolPtr(是否使用颜色))
	return r != 0
}

// 编辑框_置空格大小.
//
// hEle: 元素句柄.
//
// size: 空格大小.
func X编辑框_置空格大小(元素句柄 int, 空格大小 int) int {
	r, _, _ := xEdit_SetSpaceSize.Call(uintptr(元素句柄), uintptr(空格大小))
	return int(r)
}

// 编辑框_置字符间距.
//
// hEle: 元素句柄.
//
// size: 英文字符间距大小.
//
// sizeZh: 中文字符间距大小.
func X编辑框_置字符间距(元素句柄 int, 英文字符 int, 中文字符 int) int {
	r, _, _ := xEdit_SetCharSpaceSize.Call(uintptr(元素句柄), uintptr(英文字符), uintptr(中文字符))
	return int(r)
}

// 编辑框_取选择文本长度, 不包括非文本内容, 返回文本内容长度.
//
// hEle: 元素句柄.
func X编辑框_取选择文本长度(元素句柄 int) int {
	r, _, _ := xEdit_GetSelectTextLength.Call(uintptr(元素句柄))
	return int(r)
}

// 编辑框_置选择文本样式.
//
// hEle: 元素句柄.
//
// iStyle: 样式索引.
func X编辑框_置选择文本样式(元素句柄 int, 样式索引 int) int {
	r, _, _ := xEdit_SetSelectTextStyle.Call(uintptr(元素句柄), uintptr(样式索引))
	return int(r)
}

// 编辑框_取文本_临时, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// hEle: 元素句柄.
func X编辑框_取文本Tmp(元素句柄 int) string {
	r, _, _ := xEdit_GetText_Temp.Call(uintptr(元素句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 编辑框_取文本行_临时, 获取指定行文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X编辑框_取文本行Tmp(元素句柄 int, 行索引 int) string {
	r, _, _ := xEdit_GetTextRow_Temp.Call(uintptr(元素句柄), uintptr(行索引))
	return 炫彩工具类.UintPtrToString(r)
}

// 编辑框_取选择文本, 不包含非文本内容. 返回临时文本, 临时缓存区大小: xcc.Text_Buffer_Size .
//
// hEle: 元素句柄.
func X编辑框_取选择文本tmp(元素句柄 int) string {
	r, _, _ := xEdit_GetSelectText_Temp.Call(uintptr(元素句柄))
	return 炫彩工具类.UintPtrToString(r)
}

// 编辑框_插入气泡开始, 当前行开始.
//
// hEle: 元素句柄.
//
// hImageAvatar: 头像图片句柄.
//
// hImageBubble: 气泡背景图片句柄.
//
// nFlag: 聊天气泡对齐方式: xcc.Chat_Flag_ .
func X编辑框_插入气泡开始(元素句柄 int, 头像图片句柄 int, 气泡背景图片句柄 int, 标志 炫彩常量类.Chat_Flag_) int {
	r, _, _ := xEdit_InsertChatBegin.Call(uintptr(元素句柄), uintptr(头像图片句柄), uintptr(气泡背景图片句柄), uintptr(标志))
	return int(r)
}

// 编辑框_取指定行气泡标识. 返回行标识: xcc.Chat_Flag_
//
// hEle: 元素句柄.
//
// iRow: 行索引.
func X编辑框_取指定行气泡标识(元素句柄 int, 行索引 int) 炫彩常量类.Chat_Flag_ {
	r, _, _ := xEdit_GetChatFlags.Call(uintptr(元素句柄), uintptr(行索引))
	return 炫彩常量类.Chat_Flag_(r)
}

// 编辑框_插入文本扩展.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
//
// iStyle: 样式.
func X编辑框_插入文本EX(元素句柄 int, 行索引 int, 列索引 int, 字符串 string, 样式 int) int {
	r, _, _ := xEdit_InsertTextEx.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引), 炫彩工具类.StrPtr(字符串), uintptr(样式))
	return int(r)
}

// 编辑框_插入对象.
//
// hEle: 元素句柄.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// hObj: 对象句柄.
func X编辑框_插入对象(元素句柄 int, 行索引 int, 列索引 int, 对象句柄 int) int {
	r, _, _ := xEdit_InsertObject.Call(uintptr(元素句柄), uintptr(行索引), uintptr(列索引), uintptr(对象句柄))
	return int(r)
}

// 编辑框_置气泡最大宽度. 当值为0时代表不限制宽度.
//
// hEle: 元素句柄.
//
// nWidth: 最大宽度.
func X编辑框_置气泡最大宽度(元素句柄 int, 最大宽度 int32) {
	xEdit_SetChatMaxWidth.Call(uintptr(元素句柄), uintptr(最大宽度))
}
