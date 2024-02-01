package widget

import (
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// 编辑框(常规, 富文本, 聊天气泡).
type Edit struct {
	ScrollView
}

// I编辑框_创建.
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
func I编辑框_创建(x int, y int, cx int, cy int, hParent int) *Edit {
	p := &Edit{}
	p.SetHandle(xc.XEdit_Create(x, y, cx, cy, hParent))
	return p
}

// I编辑框_创建扩展.
//
// x: 元素x坐标.
//
// y: 元素y坐标.
//
// cx: 宽度.
//
// cy: 高度.
//
// nType: 类型, I常量_编辑框类型_.
//
// hParent: 父为窗口句柄或元素句柄.
func I编辑框_创建扩展(x int, y int, cx int, cy int, nType xcc.I常量_编辑框类型_, hParent int) *Edit {
	p := &Edit{}
	p.I句柄 = xc.XEdit_CreateEx(x, y, cx, cy, nType, hParent)
	return p
}

// I编辑框_从句柄创建对象.
func I编辑框_从句柄创建(handle int) *Edit {
	p := &Edit{}
	p.SetHandle(handle)
	return p
}

// I编辑框_从name创建对象, 失败返回nil.
func I编辑框_从name创建(name string) *Edit {
	handle := xc.XC_GetObjectByName(name)
	if handle > 0 {
		p := &Edit{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I编辑框_从UID创建对象, 失败返回nil.
func I编辑框_从UID创建(nUID int) *Edit {
	handle := xc.XC_GetObjectByUID(nUID)
	if handle > 0 {
		p := &Edit{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// I编辑框_从UID名称创建对象, 失败返回nil.
func I编辑框_从UID名称创建(name string) *Edit {
	handle := xc.XC_GetObjectByUIDName(name)
	if handle > 0 {
		p := &Edit{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// 编辑框_启用自动换行.
//
// bEnable: 是否启用.
func (e *Edit) I启用自动换行(bEnable bool) int {
	return xc.XEdit_EnableAutoWrap(e.I句柄, bEnable)
}

// 编辑框_启用只读.
//
// bEnable: 是否启用.
func (e *Edit) I启用只读(bEnable bool) int {
	return xc.XEdit_EnableReadOnly(e.I句柄, bEnable)
}

// 编辑框_启用多行.
//
// bEnable:.
func (e *Edit) I启用多行(bEnable bool) int {
	return xc.XEdit_EnableMultiLine(e.I句柄, bEnable)
}

// 编辑框_启用密码, 启用密码模式(只支持默认类型编辑框).
//
// bEnable: 是否启用.
func (e *Edit) I启用密码(bEnable bool) int {
	return xc.XEdit_EnablePassword(e.I句柄, bEnable)
}

// 编辑框_启用自动选择, 当获得焦点时,自动选择所有内容.
//
// bEnable: 是否启用.
func (e *Edit) I启用自动选择(bEnable bool) int {
	return xc.XEdit_EnableAutoSelAll(e.I句柄, bEnable)
}

// 编辑框_启用自动取消选择, 当失去焦点时自动取消选择.
//
// bEnable: 是否启用.
func (e *Edit) I启用自动取消选择(bEnable bool) int {
	return xc.XEdit_EnableAutoCancelSel(e.I句柄, bEnable)
}

// 编辑框_是否只读.
func (e *Edit) I是否只读() bool {
	return xc.XEdit_IsReadOnly(e.I句柄)
}

// 编辑框_是否多行.
func (e *Edit) I是否多行() bool {
	return xc.XEdit_IsMultiLine(e.I句柄)
}

// 编辑框_是否密码.
func (e *Edit) I是否密码() bool {
	return xc.XEdit_IsPassword(e.I句柄)
}

// 编辑框_是否自动换行.
func (e *Edit) I是否自动换行() bool {
	return xc.XEdit_IsAutoWrap(e.I句柄)
}

// 编辑框_判断为空.
func (e *Edit) I判断为空() bool {
	return xc.XEdit_IsEmpty(e.I句柄)
}

// 编辑框_是否在选择区域.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) I是否在选择区域(iRow int, iCol int) bool {
	return xc.XEdit_IsInSelect(e.I句柄, iRow, iCol)
}

// 编辑框_取总行数.
func (e *Edit) I取总行数() int {
	return xc.XEdit_GetRowCount(e.I句柄)
}

// 编辑框_取数据, 包含文本或非文本内容.
func (e *Edit) I取数据() xc.Edit_Data_Copy_ {
	return xc.XEdit_GetData(e.I句柄)
}

// 编辑框_添加数据.
//
// pData: 数据结构.
//
// styleTable: 样式表.
//
// nStyleCount: 样式数量.
func (e *Edit) I添加数据(pData *xc.Edit_Data_Copy_, styleTable []uint16, nStyleCount int) int {
	return xc.XEdit_AddData(e.I句柄, pData, styleTable, nStyleCount)
}

// 编辑框_释放数据.
func (e *Edit) I释放数据(pData *xc.Edit_Data_Copy_) int {
	return xc.XEdit_FreeData(pData)
}

// 编辑框_置默认文本, 当内容为空时, 显示默认文本.
//
// pString: 文本内容.
func (e *Edit) I置默认文本(pString string) int {
	return xc.XEdit_SetDefaultText(e.I句柄, pString)
}

// 编辑框_置默认文本颜色.
//
// color: ABGR 颜色值.
func (e *Edit) I置默认文本颜色(color int) int {
	return xc.XEdit_SetDefaultTextColor(e.I句柄, color)
}

// 编辑框_置密码字符.
//
// ch: 字符.
func (e *Edit) I置密码字符(ch int) int {
	return xc.XEdit_SetPasswordCharacter(e.I句柄, ch)
}

// 编辑框_置文本对齐, 单行模式下有效.
//
// align: 对齐方式, I常量_编辑框文本对齐_.
func (e *Edit) I置文本对齐(align xcc.I常量_编辑框文本对齐_) int {
	return xc.XEdit_SetTextAlign(e.I句柄, align)
}

// 编辑框_置TAB空格.
//
// nSpace: 空格数量.
func (e *Edit) I置TAB空格(nSpace int) int {
	return xc.XEdit_SetTabSpace(e.I句柄, nSpace)
}

// 编辑框_置文本.
//
// pString: 字符串.
func (e *Edit) I置文本(pString string) int {
	return xc.XEdit_SetText(e.I句柄, pString)
}

// 编辑框_置文本整数.
//
// nValue: 整数值.
func (e *Edit) I置文本整数(nValue int) int {
	return xc.XEdit_SetTextInt(e.I句柄, nValue)
}

// 编辑框_取文本, 不包含非文本内容, 返回实际接收文本长度.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 内存大小.
func (e *Edit) I取文本(pOut *string, nOutlen int) int {
	return xc.XEdit_GetText(e.I句柄, pOut, nOutlen)
}

// 编辑框_取文本行.
//
// iRow: 行索引.
//
// pOut: 接收文本内存指针.
//
// nOutlen: 接收文本内存块长度.
func (e *Edit) I取文本行(iRow int, pOut *string, nOutlen int) int {
	return xc.XEdit_GetTextRow(e.I句柄, iRow, pOut, nOutlen)
}

// 编辑框_取内容长度, 包含非文本内容.
func (e *Edit) I取内容长度() int {
	return xc.XEdit_GetLength(e.I句柄)
}

// 编辑框_取内容长度行, 包含非文本内容.
//
// iRow: 行索引.
func (e *Edit) I取内容长度行(iRow int) int {
	return xc.XEdit_GetLengthRow(e.I句柄, iRow)
}

// 编辑框_取字符, 返回指定位置字符.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) I取字符(iRow int, iCol int) int {
	return xc.XEdit_GetAt(e.I句柄, iRow, iCol)
}

// 编辑框_插入文本.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pString: 字符串.
func (e *Edit) I插入文本(iRow int, iCol int, pString string) int {
	return xc.XEdit_InsertText(e.I句柄, iRow, iCol, pString)
}

// I插入文本模拟用户操作 编辑框_插入文本模拟用户操作, 自动刷新UI, 支持撤销/恢复.
//	@param pString 字符串.
//	@return int
//
func (e *Edit) I插入文本模拟用户操作(pString string) int {
	return xc.XEdit_AddTextUser(e.I句柄, pString)
}

// 编辑框_添加文本.
//
// pString: 字符串.
func (e *Edit) I添加文本(pString string) int {
	return xc.XEdit_AddText(e.I句柄, pString)
}

// 编辑框_添加文本扩展.
//
// pString: 字符串.
//
// iStyle: 样式索引.
func (e *Edit) I添加文本扩展(pString string, iStyle int) int {
	return xc.XEdit_AddTextEx(e.I句柄, pString, iStyle)
}

// 编辑框_添加对象, 例如: 字体, 图片, UI对象, 返回样式索引.
//
// hObj: 对象句柄.
func (e *Edit) I添加对象(hObj int) int {
	return xc.XEdit_AddObject(e.I句柄, hObj)
}

// 编辑框_添加对象从样式, 当样式为图片时有效.
//
// iStyle: 样式索引.
func (e *Edit) I添加对象从样式(iStyle int) int {
	return xc.XEdit_AddByStyle(e.I句柄, iStyle)
}

// 编辑框_添加样式, 返回样式索引.
//
// hFont_image_Obj: 字体.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) I添加样式(hFont_image_Obj int, color int, bColor bool) int {
	return xc.XEdit_AddStyle(e.I句柄, hFont_image_Obj, color, bColor)
}

// 编辑框_添加样式扩展, 返回样式索引.
//
// fontName: 字体名称.
//
// fontSize: 字体大小.
//
// fontStyle: 字体样式, I常量_字体样式_.
//
// color: 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) I添加样式扩展(fontName string, fontSize int, fontStyle xcc.I常量_字体样式_, color int, bColor bool) int {
	return xc.XEdit_AddStyleEx(e.I句柄, fontName, fontSize, fontStyle, color, bColor)
}

// 编辑框_取样式信息.
//
// iStyle: 样式索引.
//
// info: 返回样式信息.
func (e *Edit) I取样式信息(iStyle int, info *xc.Edit_Style_Info_) bool {
	return xc.XEdit_GetStyleInfo(e.I句柄, iStyle, info)
}

// 编辑框_置当前样式.
//
// iStyle: 样式索引.
func (e *Edit) I置当前样式(iStyle int) int {
	return xc.XEdit_SetCurStyle(e.I句柄, iStyle)
}

// 编辑框_置插入符颜色.
//
// color: 颜色.
func (e *Edit) I置插入符颜色(color int) int {
	return xc.XEdit_SetCaretColor(e.I句柄, color)
}

// 编辑框_置插入符宽度.
//
// nWidth: 宽度.
func (e *Edit) I置插入符宽度(nWidth int) int {
	return xc.XEdit_SetCaretWidth(e.I句柄, nWidth)
}

// 编辑框_置选择背景颜色.
//
// color: ABGR 颜色.
func (e *Edit) I置选择背景颜色(color int) int {
	return xc.XEdit_SetSelectBkColor(e.I句柄, color)
}

// 编辑框_置默认行高.
//
// nHeight: 行高.
func (e *Edit) I置默认行高(nHeight int) int {
	return xc.XEdit_SetRowHeight(e.I句柄, nHeight)
}

// 编辑框_置指定行高度, 类型为 I常量_编辑框类型_富文本 支持指定不同行高.
//
// iRow: 行索引.
//
// nHeight: 高度.
func (e *Edit) I置指定行高度(iRow int, nHeight int) int {
	return xc.XEdit_SetRowHeightEx(e.I句柄, iRow, nHeight)
}

// I置当前位置 编辑框_置当前位置.
//	@param iRow 行索引.
//	@return int
//
func (e *Edit) I置当前位置(iRow int) int {
	return xc.XEdit_SetCurPos(e.I句柄, iRow)
}

// 编辑框_取当前位置点, 返回范围位置点.
func (e *Edit) I取当前位置点() int {
	return xc.XEdit_GetCurPos(e.I句柄)
}

// 编辑框_取当前行, 返回行索引.
func (e *Edit) I取当前行() int {
	return xc.XEdit_GetCurRow(e.I句柄)
}

// 编辑框_取当前列, 返回列索引.
func (e *Edit) I取当前列() int {
	return xc.XEdit_GetCurCol(e.I句柄)
}

// 编辑框_取坐标点.
//
// iRow: 行索引.
//
// iCol: 列索引.
//
// pOut: 接收返回坐标点.
func (e *Edit) I取坐标点(iRow int, iCol int, pOut *xc.POINT) int {
	return xc.XEdit_GetPoint(e.I句柄, iRow, iCol, pOut)
}

// 编辑框_自动滚动, 视图自动滚动到当前插入符位置.
func (e *Edit) I自动滚动() bool {
	return xc.XEdit_AutoScroll(e.I句柄)
}

// 编辑框_自动滚动扩展, 视图自动滚动到指定位置.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) I自动滚动扩展(iRow int, iCol int) bool {
	return xc.XEdit_AutoScrollEx(e.I句柄, iRow, iCol)
}

// I转换位置 编辑框_转换位置, 转换位置点到行列.
//	@param iPos 位置点.
//	@param pInfo 行列.
//	@return int
//
func (e *Edit) I转换位置(iPos int, pInfo *xc.Position_) int {
	return xc.XEdit_PosToRowCol(e.I句柄, iPos, pInfo)
}

// 编辑框_选择全部.
func (e *Edit) I选择全部() bool {
	return xc.XEdit_SelectAll(e.I句柄)
}

// 编辑框_取消选择.
func (e *Edit) I取消选择() bool {
	return xc.XEdit_CancelSelect(e.I句柄)
}

// 编辑框_删除选择内容.
func (e *Edit) I删除选择内容() bool {
	return xc.XEdit_DeleteSelect(e.I句柄)
}

// 编辑框_置选择.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func (e *Edit) I置选择(iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {
	return xc.XEdit_SetSelect(e.I句柄, iStartRow, iStartCol, iEndRow, iEndCol)
}

// 编辑框_取选择文本, 不包括非文本内容, 返回接收文本内容实际长度.
//
// pOut: 接收返回文本内容.
//
// nOutLen: 接收内存大小.
func (e *Edit) I取选择文本(pOut *string, nOutLen int) int {
	return xc.XEdit_GetSelectText(e.I句柄, pOut, nOutLen)
}

// 编辑框_取选择内容范围.
//
// pBegin: 起始位置.
//
// pEnd: 结束位置.
func (e *Edit) I取选择内容范围(pBegin *xc.Position_, pEnd *xc.Position_) bool {
	return xc.XEdit_GetSelectRange(e.I句柄, pBegin, pEnd)
}

// 编辑框_取可视行范围.
//
// piStart: 起始行索引.
//
// piEnd: 结束行索引.
func (e *Edit) I取可视行范围(piStart *int, piEnd *int) int {
	return xc.XEdit_GetVisibleRowRange(e.I句柄, piStart, piEnd)
}

// 编辑框_删除, 删除指定范围内容.
//
// iStartRow: 起始行索引.
//
// iStartCol: 起始行列索引.
//
// iEndRow: 结束行索引.
//
// iEndCol: 结束行列索引.
func (e *Edit) I删除(iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {
	return xc.XEdit_Delete(e.I句柄, iStartRow, iStartCol, iEndRow, iEndCol)
}

// 编辑框_删除行.
//
// iRow: 行索引.
func (e *Edit) I删除行(iRow int) bool {
	return xc.XEdit_DeleteRow(e.I句柄, iRow)
}

// 编辑框_剪贴板剪切.
func (e *Edit) I剪贴板剪切() bool {
	return xc.XEdit_ClipboardCut(e.I句柄)
}

// 编辑框_剪贴板复制.
func (e *Edit) I剪贴板复制() bool {
	return xc.XEdit_ClipboardCopy(e.I句柄)
}

// 编辑框_剪贴板粘贴.
func (e *Edit) I剪贴板粘贴() bool {
	return xc.XEdit_ClipboardPaste(e.I句柄)
}

// 编辑框_撤销.
func (e *Edit) I撤销() bool {
	return xc.XEdit_Undo(e.I句柄)
}

// 编辑框_恢复.
func (e *Edit) I恢复() bool {
	return xc.XEdit_Redo(e.I句柄)
}

// 编辑框_添加气泡开始, 当前行开始.
//
// hImageAvatar: 头像.
//
// hImageBubble: 气泡背景.
//
// nFlag: 标志, Chat_Flag_.
func (e *Edit) I添加气泡开始(hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {
	return xc.XEdit_AddChatBegin(e.I句柄, hImageAvatar, hImageBubble, nFlag)
}

// 编辑框_添加气泡结束, 当前行结束.
func (e *Edit) I添加气泡结束() int {
	return xc.XEdit_AddChatEnd(e.I句柄)
}

// 编辑框_置气泡缩进, 设置聊天气泡内容缩进.
//
// nIndentation: 缩进值.
func (e *Edit) I置气泡缩进(nIndentation int) int {
	return xc.XEdit_SetChatIndentation(e.I句柄, nIndentation)
}

// 编辑框_置行间隔, 设置行间隔大小, 多行模式有效.
//
// nSpace: 行间隔大小.
func (e *Edit) I置行间隔(nSpace int) int {
	return xc.XEdit_SetRowSpace(e.I句柄, nSpace)
}

// 编辑框_置当前位置扩展.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) I置当前位置扩展(iRow int, iCol int) int {
	return xc.XEdit_SetCurPosEx(e.I句柄, iRow, iCol)
}

// 编辑框_取当前位置扩展.
//
// iRow: 返回行索引.
//
// iCol: 返回列索引.
func (e *Edit) I取当前位置扩展(iRow *int, iCol *int) int {
	return xc.XEdit_GetCurPosEx(e.I句柄, iRow, iCol)
}

// 编辑框_移动到末尾.
func (e *Edit) I移动到末尾() int {
	return xc.XEdit_MoveEnd(e.I句柄)
}

// 编辑框_行列到位置, 返回位置点.
//
// iRow: 行索引.
//
// iCol: 列索引.
func (e *Edit) I行列到位置(iRow int, iCol int) int {
	return xc.XEdit_RowColToPos(e.I句柄, iRow, iCol)
}

// 编辑框_置后备字体, 置中文字体. 如果已设置, 当遇到中文字符时使用后备字体, 解决不支持中文的字体的问题
//
// hFont: 字体.
func (e *Edit) I置后备字体(hFont int) int {
	return xc.XEdit_SetBackFont(e.I句柄, hFont)
}

// 编辑框_释放样式.
//
// iStyle: 样式.
func (e *Edit) I释放样式(iStyle int) bool {
	return xc.XEdit_ReleaseStyle(e.I句柄, iStyle)
}

// 编辑框_修改样式.
//
// iStyle: 样式索引.
//
// hFont: 字体句柄.
//
// color: ABGR 颜色.
//
// bColor: 是否使用颜色.
func (e *Edit) I修改样式(iStyle int, hFont int, color int, bColor bool) bool {
	return xc.XEdit_ModifyStyle(e.I句柄, iStyle, hFont, color, bColor)
}

// 编辑框_置空格大小.
//
// size: 空格大小.
func (e *Edit) I置空格大小(size int) int {
	return xc.XEdit_SetSpaceSize(e.I句柄, size)
}

// 编辑框_置字符间距.
//
// size: 英文字符间距大小.
//
// sizeZh: 中文字符间距大小.
func (e *Edit) I置字符间距(size int, sizeZh int) int {
	return xc.XEdit_SetCharSpaceSize(e.I句柄, size, sizeZh)
}

// 编辑框_取选择文本长度, 不包括非文本内容, 返回文本内容长度.
func (e *Edit) I取选择文本长度() int {
	return xc.XEdit_GetSelectTextLength(e.I句柄)
}

// 编辑框_置选择文本样式.
//
// iStyle: 样式索引.
func (e *Edit) I置选择文本样式(iStyle int) int {
	return xc.XEdit_SetSelectTextStyle(e.I句柄, iStyle)
}

/*
以下都是事件
*/

type XE_EDIT_SET func(pbHandled *bool) int                                               // 编辑框_置文本.
type XE_EDIT_SET1 func(hEle int, pbHandled *bool) int                                    // 编辑框_置文本.
type XE_EDIT_DRAWROW func(hDraw int, iRow int, pbHandled *bool) int                      // 和XE_EDIT_CHANGED的对换.
type XE_EDIT_DRAWROW1 func(hEle int, hDraw int, iRow int, pbHandled *bool) int           // 和XE_EDIT_CHANGED的对换.
type XE_EDIT_CHANGED func(pbHandled *bool) int                                           // 编辑框_内容被改变.
type XE_EDIT_CHANGED1 func(hEle int, pbHandled *bool) int                                // 编辑框_内容被改变.
type XE_EDIT_POS_CHANGED func(iPos int, pbHandled *bool) int                             // 编辑框_光标位置_被改变.
type XE_EDIT_POS_CHANGED1 func(hEle int, iPos int, pbHandled *bool) int                  // 编辑框_光标位置_被改变.
type XE_EDIT_STYLE_CHANGED func(iStyle int, pbHandled *bool) int                         // 编辑框_样式_被改变.
type XE_EDIT_STYLE_CHANGED1 func(hEle int, iStyle int, pbHandled *bool) int              // 编辑框_样式_被改变.
type XE_EDIT_ENTER_GET_TABALIGN func(pbHandled *bool) int                                // 编辑框_回车_获取标签?.
type XE_EDIT_ENTER_GET_TABALIGN1 func(hEle int, pbHandled *bool) int                     // 编辑框_回车_获取标签?.
type XE_EDIT_ROW_CHANGED func(iRow int, nChangeRows int, pbHandled *bool) int            // 编辑框_行_被改变.
type XE_EDIT_ROW_CHANGED1 func(hEle int, iRow int, nChangeRows int, pbHandled *bool) int // 编辑框_行_被改变.

// 编辑框_ I事件_置文本.
func (e *Edit) I事件_置文本(pFun XE_EDIT_SET) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_SET, pFun)
}

// 编辑框_ I事件_置文本.
func (e *Edit) I事件_置文本1(pFun XE_EDIT_SET1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_SET, pFun)
}

// 和XE_EDIT_CHANGED的对换.
func (e *Edit) Event_EDIT_DRAWROW(pFun XE_EDIT_DRAWROW) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_DRAWROW, pFun)
}

// 和XE_EDIT_CHANGED的对换.
func (e *Edit) Event_EDIT_DRAWROW1(pFun XE_EDIT_DRAWROW1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_DRAWROW, pFun)
}

// 编辑框_ I事件_内容被改变.
func (e *Edit) I事件_内容被改变(pFun XE_EDIT_CHANGED) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_CHANGED, pFun)
}

// 编辑框_ I事件_内容被改变.
func (e *Edit) I事件_内容被改变1(pFun XE_EDIT_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_CHANGED, pFun)
}

// 编辑框_ I事件_光标位置被改变.
func (e *Edit) I事件_光标位置被改变(pFun XE_EDIT_POS_CHANGED) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_ I事件_光标位置被改变.
func (e *Edit) I事件_光标位置被改变1(pFun XE_EDIT_POS_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_POS_CHANGED, pFun)
}

// 编辑框_ I事件_样式被改变.
func (e *Edit) I事件_样式被改变(pFun XE_EDIT_STYLE_CHANGED) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_STYLE_CHANGED, pFun)
}

// 编辑框_ I事件_样式被改变.
func (e *Edit) I事件_样式被改变1(pFun XE_EDIT_STYLE_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_STYLE_CHANGED, pFun)
}

// 编辑框_ I事件_回车_获取标签 ?.
func (e *Edit) I事件_回车_获取标签(pFun XE_EDIT_ENTER_GET_TABALIGN) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 编辑框_ I事件_回车_获取标签 ?.
func (e *Edit) I事件_回车_获取标签1(pFun XE_EDIT_ENTER_GET_TABALIGN1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_ENTER_GET_TABALIGN, pFun)
}

// 编辑框_ I事件_行被改变.
func (e *Edit) I事件_行被改变(pFun XE_EDIT_ROW_CHANGED) bool {
	return xc.XEle_RegEventC(e.I句柄, xcc.XE_EDIT_ROW_CHANGED, pFun)
}

// 编辑框_ I事件_行被改变.
func (e *Edit) I事件_行被改变1(pFun XE_EDIT_ROW_CHANGED1) bool {
	return xc.XEle_RegEventC1(e.I句柄, xcc.XE_EDIT_ROW_CHANGED, pFun)
}
