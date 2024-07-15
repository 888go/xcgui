# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func NewEdit(x int, y int, cx int, cy int, hParent int) *Edit {]
ff=创建编辑框
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewEditEx(x int, y int, cx int, cy int, nType xcc.Edit_Type_, hParent int) *Edit {]
ff=创建编辑框EX
nType=类型
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewEditByHandle(handle int) *Edit {]
ff=创建编辑框并按句柄
handle=句柄

[func NewEditByName(name string) *Edit {]
ff=创建编辑框并按名称
name=名称

[func NewEditByUID(nUID int) *Edit {]
ff=创建编辑框并按UID

[func NewEditByUIDName(name string) *Edit {]
ff=创建编辑框并按UID名称
name=UID名称

[func (e *Edit) EnableAutoWrap(bEnable bool) int {]
ff=启用自动换行
bEnable=是否启用

[func (e *Edit) EnableReadOnly(bEnable bool) int {]
ff=启用只读
bEnable=是否启用

[func (e *Edit) EnableMultiLine(bEnable bool) int {]
ff=启用多行
bEnable=是否启用

[func (e *Edit) EnablePassword(bEnable bool) int {]
ff=启用密码
bEnable=是否启用

[func (e *Edit) EnableAutoSelAll(bEnable bool) int {]
ff=启用自动选择
bEnable=是否启用

[func (e *Edit) EnableAutoCancelSel(bEnable bool) int {]
ff=启用自动取消选择
bEnable=是否启用

[func (e *Edit) IsReadOnly() bool {]
ff=是否只读

[func (e *Edit) IsMultiLine() bool {]
ff=是否多行

[func (e *Edit) IsPassword() bool {]
ff=是否密码

[func (e *Edit) IsAutoWrap() bool {]
ff=是否自动换行

[func (e *Edit) IsEmpty() bool {]
ff=判断为空

[func (e *Edit) IsInSelect(iRow int, iCol int) bool {]
ff=是否在选择区域
iCol=列索引
iRow=行索引

[func (e *Edit) GetRowCount() int {]
ff=取总行数

[func (e *Edit) GetData() xc.Edit_Data_Copy_ {]
ff=取数据

[func (e *Edit) AddData(pData *xc.Edit_Data_Copy_, styleTable #左中括号##右中括号#uint16, nStyleCount int) int {]
ff=添加数据
pData=数据结构

[func (e *Edit) FreeData(pData *xc.Edit_Data_Copy_) int {]
ff=释放数据

[func (e *Edit) SetDefaultText(pString string) int {]
ff=置默认文本
pString=文本内容

[func (e *Edit) SetDefaultTextColor(color int) int {]
ff=置默认文本颜色
color=ABGR颜色值

[func (e *Edit) SetPasswordCharacter(ch int) int {]
ff=置密码字符
ch=字符

[func (e *Edit) SetTextAlign(align xcc.Edit_TextAlign_Flag_) int {]
ff=置文本对齐
align=对齐方式

[func (e *Edit) SetTabSpace(nSpace int) int {]
ff=置TAB空格
nSpace=空格数量

[func (e *Edit) SetText(pString string) int {]
ff=置文本
pString=字符串

[func (e *Edit) SetTextInt(nValue int) int {]
ff=置文本整数
nValue=整数值

[func (e *Edit) GetText(pOut *string, nOutlen int) int {]
ff=取文本
nOutlen=内存大小
pOut=接收文本指针

[func (e *Edit) GetTextEx() string {]
ff=取文本Ex

[func (e *Edit) GetSelectTextEx() string {]
ff=取选择文本Ex

[func (e *Edit) GetTextRow(iRow int, pOut *string, nOutlen int) int {]
ff=取文本行
nOutlen=接收文本内存块长度
pOut=接收文本指针
iRow=行索引

[func (e *Edit) GetTextRowEx(iRow int) string {]
ff=取文本行Ex
iRow=行索引

[func (e *Edit) GetLength() int {]
ff=取内容长度

[func (e *Edit) GetLengthRow(iRow int) int {]
ff=取内容长度行
iRow=行索引

[func (e *Edit) GetAt(iRow int, iCol int) int {]
ff=取字符
iCol=列索引
iRow=行索引

[func (e *Edit) InsertText(iRow int, iCol int, pString string) int {]
ff=插入文本
pString=字符串
iCol=列索引
iRow=行索引

[func (e *Edit) AddTextUser(pString string) int {]
ff=插入文本模拟用户操作
pString=字符串

[func (e *Edit) AddText(pString string) int {]
ff=添加文本
pString=字符串

[func (e *Edit) AddTextEx(pString string, iStyle int) int {]
ff=添加文本EX
iStyle=样式索引
pString=字符串

[func (e *Edit) AddObject(hObj int) int {]
ff=添加对象
hObj=对象句柄

[func (e *Edit) AddByStyle(iStyle int) int {]
ff=添加对象并按样式
iStyle=样式索引

[func (e *Edit) AddStyle(hFont_image_Obj int, color int, bColor bool) int {]
ff=添加样式
bColor=是否使用颜色
color=颜色
hFont_image_Obj=字体

[func (e *Edit) AddStyleEx(fontName string, fontSize int, fontStyle xcc.FontStyle_, color int, bColor bool) int {]
ff=添加样式EX
fontStyle=字体样式
fontSize=字体大小
fontName=字体名称

[func (e *Edit) GetStyleInfo(iStyle int, info *xc.Edit_Style_Info_) bool {]
ff=取样式信息
info=返回样式信息
iStyle=样式索引

[func (e *Edit) SetCurStyle(iStyle int) int {]
ff=置当前样式
iStyle=样式索引

[func (e *Edit) SetCaretColor(color int) int {]
ff=置插入符颜色
color=颜色

[func (e *Edit) SetCaretWidth(nWidth int) int {]
ff=置插入符宽度
nWidth=宽度

[func (e *Edit) SetSelectBkColor(color int) int {]
ff=置选择背景颜色
color=ABGR颜色

[func (e *Edit) SetRowHeight(nHeight int) int {]
ff=置默认行高
nHeight=行高

[func (e *Edit) SetRowHeightEx(iRow int, nHeight int) int {]
ff=置指定行高度
nHeight=高度
iRow=行索引

[func (e *Edit) SetCurPos(iRow int) int {]
ff=置当前位置
iRow=行索引

[func (e *Edit) GetCurPos() int {]
ff=取当前位置点

[func (e *Edit) GetCurRow() int {]
ff=取当前行

[func (e *Edit) GetCurCol() int {]
ff=取当前列

[func (e *Edit) GetPoint(iRow int, iCol int, pOut *xc.POINT) int {]
ff=取坐标点
pOut=接收返回坐标点
iCol=列索引
iRow=行索引

[func (e *Edit) AutoScroll() bool {]
ff=自动滚动

[func (e *Edit) AutoScrollEx(iRow int, iCol int) bool {]
ff=自动滚动EX
iCol=列索引
iRow=行索引

[func (e *Edit) PosToRowCol(iPos int, pInfo *xc.Position_) int {]
ff=转换位置
pInfo=行列
iPos=位置点

[func (e *Edit) SelectAll() bool {]
ff=选择全部

[func (e *Edit) CancelSelect() bool {]
ff=取消选择

[func (e *Edit) DeleteSelect() bool {]
ff=删除选择内容

[func (e *Edit) SetSelect(iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {]
ff=置选择
iEndCol=结束行列索引
iEndRow=结束行索引
iStartCol=起始行列索引
iStartRow=起始行索引

[func (e *Edit) GetSelectText(pOut *string, nOutLen int) int {]
ff=取选择文本
nOutLen=接收内存大小
pOut=接收返回文本

[func (e *Edit) GetSelectRange(pBegin *xc.Position_, pEnd *xc.Position_) bool {]
ff=取选择内容范围
pBegin=起始位置

[func (e *Edit) GetVisibleRowRange(piStart *int32, piEnd *int32) int {]
ff=取可视行范围
piEnd=结束行索引
piStart=起始行索引

[func (e *Edit) Delete(iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {]
ff=删除
iEndCol=结束行列索引
iEndRow=结束行索引
iStartCol=起始行列索引
iStartRow=起始行索引

[func (e *Edit) DeleteRow(iRow int) bool {]
ff=删除行
iRow=行索引

[func (e *Edit) ClipboardCut() bool {]
ff=剪贴板剪切

[func (e *Edit) ClipboardCopy() bool {]
ff=剪贴板复制

[func (e *Edit) ClipboardPaste() bool {]
ff=剪贴板粘贴

[func (e *Edit) Undo() bool {]
ff=撤销

[func (e *Edit) Redo() bool {]
ff=恢复

[func (e *Edit) AddChatBegin(hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {]
ff=添加气泡开始
nFlag=标志
hImageBubble=气泡背景
hImageAvatar=头像

[func (e *Edit) AddChatEnd() int {]
ff=添加气泡结束

[func (e *Edit) SetChatIndentation(nIndentation int) int {]
ff=置气泡缩进
nIndentation=缩进值

[func (e *Edit) SetRowSpace(nSpace int) int {]
ff=置行间隔
nSpace=行间隔大小

[func (e *Edit) SetCurPosEx(iRow, iCol int32) int {]
ff=置当前位置EX
iCol=列索引
iRow=行索引

[func (e *Edit) GetCurPosEx(iRow, iCol *int32) int {]
ff=取当前位置EX
iCol=返回列索引
iRow=返回行索引

[func (e *Edit) MoveEnd() int {]
ff=移动到末尾

[func (e *Edit) RowColToPos(iRow int, iCol int) int {]
ff=行列到位置
iCol=列索引
iRow=行索引

[func (e *Edit) SetBackFont(hFont int) int {]
ff=置后备字体
hFont=字体

[func (e *Edit) ReleaseStyle(iStyle int) bool {]
ff=释放样式
iStyle=样式

[func (e *Edit) ModifyStyle(iStyle int, hFont int, color int, bColor bool) bool {]
ff=修改样式
bColor=是否使用颜色
color=ABGR颜色
hFont=字体句柄
iStyle=样式索引

[func (e *Edit) SetSpaceSize(size int) int {]
ff=置空格大小
size=空格大小

[func (e *Edit) SetCharSpaceSize(size int, sizeZh int) int {]
ff=置字符间距
sizeZh=中文字符间距
size=英文字符间距

[func (e *Edit) GetSelectTextLength() int {]
ff=取选择文本长度

[func (e *Edit) SetSelectTextStyle(iStyle int) int {]
ff=置选择文本样式
iStyle=样式索引

[func (e *Edit) GetText_Temp() string {]
ff=取文本Tmp

[func (e *Edit) GetTextRow_Temp(iRow int) string {]
ff=取行文本Tmp
iRow=行索引

[func (e *Edit) GetSelectText_Temp() string {]
ff=取选择文本Tmp

[func (e *Edit) InsertChatBegin(hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {]
ff=插入气泡开始
nFlag=对齐方式
hImageBubble=背景图片句柄
hImageAvatar=头像图片句柄

[func (e *Edit) GetChatFlags(iRow int) xcc.Chat_Flag_ {]
ff=取指定行气泡标识
iRow=行索引

[func (e *Edit) InsertTextEx(iRow int, iCol int, pString string, iStyle int) int {]
ff=插入文本EX
iStyle=样式
pString=字符串
iCol=列索引
iRow=行索引

[func (e *Edit) InsertObject(iRow int, iCol int, hObj int) int {]
ff=插入对象
hObj=对象句柄
iCol=列索引
iRow=行索引

[func (e *Edit) SetChatMaxWidth(nWidth int32) {]
ff=置气泡最大宽度
nWidth=最大宽度

[func (e *Edit) Event_EDIT_COLOR_CHANGE(pFun XE_EDIT_COLOR_CHANGE) bool {]
ff=事件_颜色被改变

[func (e *Edit) Event_EDIT_COLOR_CHANGE1(pFun XE_EDIT_COLOR_CHANGE1) bool {]
ff=事件_颜色被改变1

[func (e *Edit) Event_EDIT_SWAPROW(pFun XE_EDIT_SWAPROW) bool {]
ff=事件_交换行

[func (e *Edit) Event_EDIT_SWAPROW1(pFun XE_EDIT_SWAPROW1) bool {]
ff=事件_交换行1

[func (e *Edit) Event_EDIT_SET(pFun XE_EDIT_SET) bool {]
ff=事件_编辑框设置

[func (e *Edit) Event_EDIT_SET1(pFun XE_EDIT_SET1) bool {]
ff=事件_编辑框设置1

[func (e *Edit) Event_EDIT_DRAWROW(pFun XE_EDIT_DRAWROW) bool {]
ff=暂未使用Event_EDIT_DRAWROW

[func (e *Edit) Event_EDIT_DRAWROW1(pFun XE_EDIT_DRAWROW1) bool {]
ff=暂未使用Event_EDIT_DRAWROW1

[func (e *Edit) Event_EDIT_CHANGED(pFun XE_EDIT_CHANGED) bool {]
ff=事件_内容被改变

[func (e *Edit) Event_EDIT_CHANGED1(pFun XE_EDIT_CHANGED1) bool {]
ff=事件_内容被改变1

[func (e *Edit) Event_EDIT_POS_CHANGED(pFun XE_EDIT_POS_CHANGED) bool {]
ff=事件_光标位置被改变

[func (e *Edit) Event_EDIT_POS_CHANGED1(pFun XE_EDIT_POS_CHANGED1) bool {]
ff=事件_光标位置被改变1

[func (e *Edit) Event_EDIT_STYLE_CHANGED(pFun XE_EDIT_STYLE_CHANGED) bool {]
ff=事件_样式被改变

[func (e *Edit) Event_EDIT_STYLE_CHANGED1(pFun XE_EDIT_STYLE_CHANGED1) bool {]
ff=事件_样式被改变1

[func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN(pFun XE_EDIT_ENTER_GET_TABALIGN) bool {]
ff=事件_回车TAB对齐

[func (e *Edit) Event_EDIT_ENTER_GET_TABALIGN1(pFun XE_EDIT_ENTER_GET_TABALIGN1) bool {]
ff=事件_回车TAB对齐1

[func (e *Edit) Event_EDIT_ROW_CHANGED(pFun XE_EDIT_ROW_CHANGED) bool {]
ff=事件_行被改变

[func (e *Edit) Event_EDIT_ROW_CHANGED1(pFun XE_EDIT_ROW_CHANGED1) bool {]
ff=事件_行被改变1
