
# <翻译开始>
func NewEdit(x int, y int, cx int, cy int, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewEdit(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func NewEdit(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func NewEdit(x int, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewEdit(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewEdit
X创建编辑框
# <翻译结束>


# <翻译开始>
func NewEditEx(x int, y int, cx int, cy int, nType xcc.Edit_Type_, hParent
父窗口句柄或元素句柄
# <翻译结束>

# <翻译开始>
func NewEditEx(x int, y int, cx int, cy int, nType
类型
# <翻译结束>

# <翻译开始>
func NewEditEx(x int, y int, cx int, cy
高度
# <翻译结束>

# <翻译开始>
func NewEditEx(x int, y int, cx
宽度
# <翻译结束>

# <翻译开始>
func NewEditEx(x int, y
元素y坐标
# <翻译结束>

# <翻译开始>
func NewEditEx(x
元素x坐标
# <翻译结束>

# <翻译开始>
func NewEditEx
X创建编辑框EX
# <翻译结束>


# <翻译开始>
func NewEditByHandle(handle
句柄
# <翻译结束>

# <翻译开始>
func NewEditByHandle
X创建编辑框并按句柄
# <翻译结束>

# <翻译开始>
func NewEditByName(name
名称
# <翻译结束>

# <翻译开始>
func NewEditByName
X创建编辑框并按名称
# <翻译结束>


# <翻译开始>
func NewEditByUID
X创建编辑框并按UID
# <翻译结束>


# <翻译开始>
func NewEditByUIDName(name
UID名称
# <翻译结束>


# <翻译开始>
func NewEditByUIDName
X创建编辑框并按UID名称
# <翻译结束>


# <翻译开始>
func (e *Edit) EnableAutoWrap(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Edit) EnableAutoWrap
X启用自动换行
# <翻译结束>


# <翻译开始>
func (e *Edit) EnableReadOnly(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Edit) EnableReadOnly
X启用只读
# <翻译结束>

# <翻译开始>
func (e *Edit) EnableMultiLine(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Edit) EnableMultiLine
X启用多行
# <翻译结束>


# <翻译开始>
func (e *Edit) EnablePassword(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Edit) EnablePassword
X启用密码
# <翻译结束>


# <翻译开始>
func (e *Edit) EnableAutoSelAll(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Edit) EnableAutoSelAll
X启用自动选择
# <翻译结束>


# <翻译开始>
func (e *Edit) EnableAutoCancelSel(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func (e *Edit) EnableAutoCancelSel
X启用自动取消选择
# <翻译结束>


# <翻译开始>
func (e *Edit) IsReadOnly
X是否只读
# <翻译结束>


# <翻译开始>
func (e *Edit) IsMultiLine
X是否多行
# <翻译结束>


# <翻译开始>
func (e *Edit) IsPassword
X是否密码
# <翻译结束>


# <翻译开始>
func (e *Edit) IsAutoWrap
X是否自动换行
# <翻译结束>


# <翻译开始>
func (e *Edit) IsEmpty
X判断为空
# <翻译结束>


# <翻译开始>
func (e *Edit) IsInSelect(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) IsInSelect(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) IsInSelect
X是否在选择区域
# <翻译结束>


# <翻译开始>
func (e *Edit) GetRowCount
X取总行数
# <翻译结束>


# <翻译开始>
func (e *Edit) GetData
X取数据
# <翻译结束>


# <翻译开始>
func (e *Edit) AddData(pData *xc.Edit_Data_Copy_, styleTable []uint16, nStyleCount
样式数量
# <翻译结束>

# <翻译开始>
func (e *Edit) AddData(pData *xc.Edit_Data_Copy_, styleTable
样式表
# <翻译结束>

# <翻译开始>
func (e *Edit) AddData(pData
数据结构
# <翻译结束>

# <翻译开始>
func (e *Edit) AddData
X添加数据
# <翻译结束>


# <翻译开始>
func (e *Edit) FreeData
X释放数据
# <翻译结束>


# <翻译开始>
func (e *Edit) SetDefaultText(pString
文本内容
# <翻译结束>

# <翻译开始>
func (e *Edit) SetDefaultText
X置默认文本
# <翻译结束>


# <翻译开始>
func (e *Edit) SetDefaultTextColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func (e *Edit) SetDefaultTextColor
X置默认文本颜色
# <翻译结束>


# <翻译开始>
func (e *Edit) SetPasswordCharacter(ch
字符
# <翻译结束>

# <翻译开始>
func (e *Edit) SetPasswordCharacter
X置密码字符
# <翻译结束>


# <翻译开始>
func (e *Edit) SetTextAlign(align
对齐方式
# <翻译结束>

# <翻译开始>
func (e *Edit) SetTextAlign
X置文本对齐
# <翻译结束>


# <翻译开始>
func (e *Edit) SetTabSpace(nSpace
空格数量
# <翻译结束>

# <翻译开始>
func (e *Edit) SetTabSpace
X置TAB空格
# <翻译结束>


# <翻译开始>
func (e *Edit) SetText(pString
字符串
# <翻译结束>

# <翻译开始>
func (e *Edit) SetText
X置文本
# <翻译结束>


# <翻译开始>
func (e *Edit) SetTextInt(nValue
整数值
# <翻译结束>

# <翻译开始>
func (e *Edit) SetTextInt
X置文本整数
# <翻译结束>


# <翻译开始>
func (e *Edit) GetText(pOut *string, nOutlen
内存大小
# <翻译结束>

# <翻译开始>
func (e *Edit) GetText(pOut
接收文本指针
# <翻译结束>

# <翻译开始>
func (e *Edit) GetText
X取文本
# <翻译结束>


# <翻译开始>
func (e *Edit) GetTextEx
X取文本Ex
# <翻译结束>


# <翻译开始>
func (e *Edit) GetSelectTextEx
X取选择文本Ex
# <翻译结束>


# <翻译开始>
func (e *Edit) GetTextRow(iRow int, pOut *string, nOutlen
接收文本内存块长度
# <翻译结束>

# <翻译开始>
func (e *Edit) GetTextRow(iRow int, pOut
接收文本指针
# <翻译结束>

# <翻译开始>
func (e *Edit) GetTextRow(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetTextRow
X取文本行
# <翻译结束>


# <翻译开始>
func (e *Edit) GetTextRowEx(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetTextRowEx
X取文本行Ex
# <翻译结束>


# <翻译开始>
func (e *Edit) GetLength
X取内容长度
# <翻译结束>


# <翻译开始>
func (e *Edit) GetLengthRow(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetLengthRow
X取内容长度行
# <翻译结束>


# <翻译开始>
func (e *Edit) GetAt(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetAt(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetAt
X取字符
# <翻译结束>


# <翻译开始>
func (e *Edit) InsertText(iRow int, iCol int, pString
字符串
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertText(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertText(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertText
X插入文本
# <翻译结束>


# <翻译开始>
func (e *Edit) AddTextUser(pString
字符串
# <翻译结束>

# <翻译开始>
func (e *Edit) AddTextUser
X插入文本模拟用户操作
# <翻译结束>


# <翻译开始>
func (e *Edit) AddText(pString
字符串
# <翻译结束>

# <翻译开始>
func (e *Edit) AddText
X添加文本
# <翻译结束>


# <翻译开始>
func (e *Edit) AddTextEx(pString string, iStyle
样式索引
# <翻译结束>

# <翻译开始>
func (e *Edit) AddTextEx(pString
字符串
# <翻译结束>

# <翻译开始>
func (e *Edit) AddTextEx
X添加文本EX
# <翻译结束>


# <翻译开始>
func (e *Edit) AddObject(hObj
对象句柄
# <翻译结束>

# <翻译开始>
func (e *Edit) AddObject
X添加对象
# <翻译结束>


# <翻译开始>
func (e *Edit) AddByStyle(iStyle
样式索引
# <翻译结束>

# <翻译开始>
func (e *Edit) AddByStyle
X添加对象并按样式
# <翻译结束>


# <翻译开始>
func (e *Edit) AddStyle(hFont_image_Obj int, color int, bColor
是否使用颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyle(hFont_image_Obj int, color
颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyle(hFont_image_Obj
字体
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyle
X添加样式
# <翻译结束>


# <翻译开始>
func (e *Edit) AddStyleEx(fontName string, fontSize int, fontStyle xcc.FontStyle_, color int, bColor
是否使用颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyleEx(fontName string, fontSize int, fontStyle xcc.FontStyle_, color
颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyleEx(fontName string, fontSize int, fontStyle
字体样式
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyleEx(fontName string, fontSize
字体大小
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyleEx(fontName
字体名称
# <翻译结束>

# <翻译开始>
func (e *Edit) AddStyleEx
X添加样式EX
# <翻译结束>


# <翻译开始>
func (e *Edit) GetStyleInfo(iStyle int, info
返回样式信息
# <翻译结束>

# <翻译开始>
func (e *Edit) GetStyleInfo(iStyle
样式索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetStyleInfo
X取样式信息
# <翻译结束>


# <翻译开始>
func (e *Edit) SetCurStyle(iStyle
样式索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCurStyle
X置当前样式
# <翻译结束>


# <翻译开始>
func (e *Edit) SetCaretColor(color
颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCaretColor
X置插入符颜色
# <翻译结束>


# <翻译开始>
func (e *Edit) SetCaretWidth(nWidth
宽度
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCaretWidth
X置插入符宽度
# <翻译结束>


# <翻译开始>
func (e *Edit) SetSelectBkColor(color
ABGR颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSelectBkColor
X置选择背景颜色
# <翻译结束>


# <翻译开始>
func (e *Edit) SetRowHeight(nHeight
行高
# <翻译结束>

# <翻译开始>
func (e *Edit) SetRowHeight
X置默认行高
# <翻译结束>


# <翻译开始>
func (e *Edit) SetRowHeightEx(iRow int, nHeight
高度
# <翻译结束>

# <翻译开始>
func (e *Edit) SetRowHeightEx(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetRowHeightEx
X置指定行高度
# <翻译结束>


# <翻译开始>
func (e *Edit) SetCurPos(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCurPos
X置当前位置
# <翻译结束>


# <翻译开始>
func (e *Edit) GetCurPos
X取当前位置点
# <翻译结束>


# <翻译开始>
func (e *Edit) GetCurRow
X取当前行
# <翻译结束>


# <翻译开始>
func (e *Edit) GetCurCol
X取当前列
# <翻译结束>


# <翻译开始>
func (e *Edit) GetPoint(iRow int, iCol int, pOut
接收返回坐标点
# <翻译结束>

# <翻译开始>
func (e *Edit) GetPoint(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetPoint(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetPoint
X取坐标点
# <翻译结束>


# <翻译开始>
func (e *Edit) AutoScroll
X自动滚动
# <翻译结束>


# <翻译开始>
func (e *Edit) AutoScrollEx(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) AutoScrollEx(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) AutoScrollEx
X自动滚动EX
# <翻译结束>


# <翻译开始>
func (e *Edit) PosToRowCol(iPos int, pInfo
行列
# <翻译结束>

# <翻译开始>
func (e *Edit) PosToRowCol(iPos
位置点
# <翻译结束>

# <翻译开始>
func (e *Edit) PosToRowCol
X转换位置
# <翻译结束>


# <翻译开始>
func (e *Edit) SelectAll
X选择全部
# <翻译结束>


# <翻译开始>
func (e *Edit) CancelSelect
X取消选择
# <翻译结束>


# <翻译开始>
func (e *Edit) DeleteSelect
X删除选择内容
# <翻译结束>


# <翻译开始>
func (e *Edit) SetSelect(iStartRow int, iStartCol int, iEndRow int, iEndCol
结束行列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSelect(iStartRow int, iStartCol int, iEndRow
结束行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSelect(iStartRow int, iStartCol
起始行列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSelect(iStartRow
起始行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSelect
X置选择
# <翻译结束>


# <翻译开始>
func (e *Edit) GetSelectText(pOut *string, nOutLen
接收内存大小
# <翻译结束>

# <翻译开始>
func (e *Edit) GetSelectText(pOut
接收返回文本
# <翻译结束>

# <翻译开始>
func (e *Edit) GetSelectText
X取选择文本
# <翻译结束>


# <翻译开始>
func (e *Edit) GetSelectRange(pBegin *xc.Position_, pEnd
结束位置
# <翻译结束>

# <翻译开始>
func (e *Edit) GetSelectRange(pBegin
起始位置
# <翻译结束>

# <翻译开始>
func (e *Edit) GetSelectRange
X取选择内容范围
# <翻译结束>


# <翻译开始>
func (e *Edit) GetVisibleRowRange(piStart *int32, piEnd
结束行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetVisibleRowRange(piStart
起始行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetVisibleRowRange
X取可视行范围
# <翻译结束>


# <翻译开始>
func (e *Edit) Delete(iStartRow int, iStartCol int, iEndRow int, iEndCol
结束行列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) Delete(iStartRow int, iStartCol int, iEndRow
结束行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) Delete(iStartRow int, iStartCol
起始行列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) Delete(iStartRow
起始行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) Delete
X删除
# <翻译结束>


# <翻译开始>
func (e *Edit) DeleteRow(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) DeleteRow
X删除行
# <翻译结束>


# <翻译开始>
func (e *Edit) ClipboardCut
X剪贴板剪切
# <翻译结束>


# <翻译开始>
func (e *Edit) ClipboardCopy
X剪贴板复制
# <翻译结束>


# <翻译开始>
func (e *Edit) ClipboardPaste
X剪贴板粘贴
# <翻译结束>


# <翻译开始>
func (e *Edit) Undo
X撤销
# <翻译结束>


# <翻译开始>
func (e *Edit) Redo
X恢复
# <翻译结束>


# <翻译开始>
func (e *Edit) AddChatBegin(hImageAvatar int, hImageBubble int, nFlag
标志
# <翻译结束>

# <翻译开始>
func (e *Edit) AddChatBegin(hImageAvatar int, hImageBubble
气泡背景
# <翻译结束>

# <翻译开始>
func (e *Edit) AddChatBegin(hImageAvatar
头像
# <翻译结束>

# <翻译开始>
func (e *Edit) AddChatBegin
X添加气泡开始
# <翻译结束>


# <翻译开始>
func (e *Edit) AddChatEnd
X添加气泡结束
# <翻译结束>


# <翻译开始>
func (e *Edit) SetChatIndentation(nIndentation
缩进值
# <翻译结束>

# <翻译开始>
func (e *Edit) SetChatIndentation
X置气泡缩进
# <翻译结束>


# <翻译开始>
func (e *Edit) SetRowSpace(nSpace
行间隔大小
# <翻译结束>

# <翻译开始>
func (e *Edit) SetRowSpace
X置行间隔
# <翻译结束>


# <翻译开始>
func (e *Edit) SetCurPosEx(iRow, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCurPosEx(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCurPosEx
X置当前位置EX
# <翻译结束>


# <翻译开始>
func (e *Edit) GetCurPosEx(iRow, iCol
返回列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetCurPosEx(iRow
返回行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetCurPosEx
X取当前位置EX
# <翻译结束>


# <翻译开始>
func (e *Edit) MoveEnd
X移动到末尾
# <翻译结束>


# <翻译开始>
func (e *Edit) RowColToPos(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) RowColToPos(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) RowColToPos
X行列到位置
# <翻译结束>


# <翻译开始>
func (e *Edit) SetBackFont(hFont
字体
# <翻译结束>

# <翻译开始>
func (e *Edit) SetBackFont
X置后备字体
# <翻译结束>


# <翻译开始>
func (e *Edit) ReleaseStyle(iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Edit) ReleaseStyle
X释放样式
# <翻译结束>


# <翻译开始>
func (e *Edit) ModifyStyle(iStyle int, hFont int, color int, bColor
是否使用颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) ModifyStyle(iStyle int, hFont int, color
ABGR颜色
# <翻译结束>

# <翻译开始>
func (e *Edit) ModifyStyle(iStyle int, hFont
字体句柄
# <翻译结束>

# <翻译开始>
func (e *Edit) ModifyStyle(iStyle
样式索引
# <翻译结束>

# <翻译开始>
func (e *Edit) ModifyStyle
X修改样式
# <翻译结束>


# <翻译开始>
func (e *Edit) SetSpaceSize(size
空格大小
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSpaceSize
X置空格大小
# <翻译结束>


# <翻译开始>
func (e *Edit) SetCharSpaceSize(size int, sizeZh
中文字符间距
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCharSpaceSize(size
英文字符间距
# <翻译结束>

# <翻译开始>
func (e *Edit) SetCharSpaceSize
X置字符间距
# <翻译结束>


# <翻译开始>
func (e *Edit) GetSelectTextLength
X取选择文本长度
# <翻译结束>


# <翻译开始>
func (e *Edit) SetSelectTextStyle(iStyle
样式索引
# <翻译结束>

# <翻译开始>
func (e *Edit) SetSelectTextStyle
X置选择文本样式
# <翻译结束>


# <翻译开始>
func (e *Edit) GetText_Temp
X取文本Tmp
# <翻译结束>


# <翻译开始>
func (e *Edit) GetTextRow_Temp(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetTextRow_Temp
X取行文本Tmp
# <翻译结束>


# <翻译开始>
func (e *Edit) GetSelectText_Temp
X取选择文本Tmp
# <翻译结束>


# <翻译开始>
func (e *Edit) InsertChatBegin(hImageAvatar int, hImageBubble int, nFlag
对齐方式
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertChatBegin(hImageAvatar int, hImageBubble
背景图片句柄
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertChatBegin(hImageAvatar
头像图片句柄
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertChatBegin
X插入气泡开始
# <翻译结束>


# <翻译开始>
func (e *Edit) GetChatFlags(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) GetChatFlags
X取指定行气泡标识
# <翻译结束>


# <翻译开始>
func (e *Edit) InsertTextEx(iRow int, iCol int, pString string, iStyle
样式
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertTextEx(iRow int, iCol int, pString
字符串
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertTextEx(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertTextEx(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertTextEx
X插入文本EX
# <翻译结束>


# <翻译开始>
func (e *Edit) InsertObject(iRow int, iCol int, hObj
对象句柄
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertObject(iRow int, iCol
列索引
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertObject(iRow
行索引
# <翻译结束>

# <翻译开始>
func (e *Edit) InsertObject
X插入对象
# <翻译结束>


# <翻译开始>
func (e *Edit) SetChatMaxWidth(nWidth
最大宽度
# <翻译结束>

# <翻译开始>
func (e *Edit) SetChatMaxWidth
X置气泡最大宽度
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_COLOR_CHANGE
X事件_颜色被改变
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_COLOR_CHANGE1
X事件_颜色被改变1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_SWAPROW
X事件_交换行
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_SWAPROW1
X事件_交换行1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_SET
X事件_编辑框设置
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_SET1
X事件_编辑框设置1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_DRAWROW
X暂未使用Event_EDIT_DRAWROW
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_DRAWROW1
X暂未使用Event_EDIT_DRAWROW1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_CHANGED
X事件_内容被改变
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_CHANGED1
X事件_内容被改变1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_POS_CHANGED
X事件_光标位置被改变
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_POS_CHANGED1
X事件_光标位置被改变1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_STYLE_CHANGED
X事件_样式被改变
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_STYLE_CHANGED1
X事件_样式被改变1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN
X事件_回车TAB对齐
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN1
X事件_回车TAB对齐1
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_ROW_CHANGED
X事件_行被改变
# <翻译结束>


# <翻译开始>
func (e *Edit) Event_EDIT_ROW_CHANGED1
X事件_行被改变1
# <翻译结束>

