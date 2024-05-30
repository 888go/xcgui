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
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func XEdit_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=编辑框_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func XEdit_CreateEx(x int, y int, cx int, cy int, nType xcc.Edit_Type_, hParent int) int {]
ff=编辑框_创建EX
nType=类型
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func XEdit_EnableAutoWrap(hEle int, bEnable bool) int {]
ff=编辑框_启用自动换行
bEnable=是否启用
hEle=元素句柄

[func XEdit_EnableReadOnly(hEle int, bEnable bool) int {]
ff=编辑框_启用只读
bEnable=是否启用
hEle=元素句柄

[func XEdit_EnableMultiLine(hEle int, bEnable bool) int {]
ff=编辑框_启用多行
bEnable=是否启用
hEle=元素句柄

[func XEdit_EnablePassword(hEle int, bEnable bool) int {]
ff=编辑框_启用密码
bEnable=是否启用
hEle=元素句柄

[func XEdit_EnableAutoSelAll(hEle int, bEnable bool) int {]
ff=编辑框_启用自动选择
bEnable=是否启用
hEle=元素句柄

[func XEdit_EnableAutoCancelSel(hEle int, bEnable bool) int {]
ff=编辑框_启用自动取消选择
bEnable=是否启用
hEle=元素句柄

[func XEdit_IsReadOnly(hEle int) bool {]
ff=编辑框_是否只读
hEle=元素句柄

[func XEdit_IsMultiLine(hEle int) bool {]
ff=编辑框_是否多行
hEle=元素句柄

[func XEdit_IsPassword(hEle int) bool {]
ff=编辑框_是否密码
hEle=元素句柄

[func XEdit_IsAutoWrap(hEle int) bool {]
ff=编辑框_是否自动换行
hEle=元素句柄

[func XEdit_IsEmpty(hEle int) bool {]
ff=编辑框_判断为空
hEle=元素句柄

[func XEdit_IsInSelect(hEle int, iRow int, iCol int) bool {]
ff=编辑框_是否在选择区域
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_GetRowCount(hEle int) int {]
ff=编辑框_取总行数
hEle=元素句柄

[func XEdit_GetData(hEle int) Edit_Data_Copy_ {]
ff=编辑框_取数据
hEle=元素句柄

[func XEdit_AddData(hEle int, pData *Edit_Data_Copy_, styleTable #左中括号##右中括号#uint16, nStyleCount int) int {]
ff=编辑框_添加数据
nStyleCount=样式数量
styleTable=样式表
pData=数据结构
hEle=元素句柄

[func XEdit_FreeData(pData *Edit_Data_Copy_) int {]
ff=编辑框_释放数据
pData=数据结构

[func XEdit_SetDefaultText(hEle int, pString string) int {]
ff=编辑框_置默认文本
pString=文本内容
hEle=元素句柄

[func XEdit_SetDefaultTextColor(hEle int, color int) int {]
ff=编辑框_置默认文本颜色
color=ABGR颜色值
hEle=元素句柄

[func XEdit_SetPasswordCharacter(hEle int, ch int) int {]
ff=编辑框_置密码字符
ch=字符
hEle=元素句柄

[func XEdit_SetTextAlign(hEle int, align xcc.Edit_TextAlign_Flag_) int {]
ff=编辑框_置文本对齐
align=对齐方式
hEle=元素句柄

[func XEdit_SetTabSpace(hEle int, nSpace int) int {]
ff=编辑框_置TAB空格
nSpace=空格数量
hEle=元素句柄

[func XEdit_SetText(hEle int, pString string) int {]
ff=编辑框_置文本
pString=字符串
hEle=元素句柄

[func XEdit_SetTextInt(hEle int, nValue int) int {]
ff=编辑框_置文本整数
nValue=值
hEle=元素句柄

[func XEdit_GetText(hEle int, pOut *string, nOutlen int) int {]
ff=编辑框_取文本
nOutlen=内存大小
pOut=接收文本内存指针
hEle=元素句柄

[func XEdit_GetTextRow(hEle int, iRow int, pOut *string, nOutlen int) int {]
ff=编辑框_取文本行
nOutlen=接收文本内存块长度
pOut=接收文本内存指针
iRow=行索引
hEle=元素句柄

[func XEdit_GetLength(hEle int) int {]
ff=编辑框_取内容长度
hEle=元素句柄

[func XEdit_GetLengthRow(hEle int, iRow int) int {]
ff=编辑框_取内容长度行
iRow=行索引
hEle=元素句柄

[func XEdit_GetAt(hEle int, iRow int, iCol int) int {]
ff=编辑框_取字符
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_InsertText(hEle int, iRow int, iCol int, pString string) int {]
ff=编辑框_插入文本
pString=字符串
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_AddTextUser(hEle int, pString string) int {]
ff=编辑框_插入文本模拟用户操作
pString=字符串
hEle=元素句柄

[func XEdit_AddText(hEle int, pString string) int {]
ff=编辑框_添加文本
pString=字符串
hEle=元素句柄

[func XEdit_AddTextEx(hEle int, pString string, iStyle int) int {]
ff=编辑框_添加文本EX
iStyle=样式索引
pString=字符串
hEle=元素句柄

[func XEdit_AddObject(hEle int, hObj int) int {]
ff=编辑框_添加对象
hObj=对象句柄
hEle=元素句柄

[func XEdit_AddByStyle(hEle int, iStyle int) int {]
ff=编辑框_添加对象从样式
iStyle=样式索引
hEle=元素句柄

[func XEdit_AddStyle(hEle int, hFont_image_Obj int, color int, bColor bool) int {]
ff=编辑框_添加样式
bColor=是否使用颜色
color=颜色
hFont_image_Obj=字体
hEle=元素句柄

[func XEdit_AddStyleEx(hEle int, fontName string, fontSize int, fontStyle xcc.FontStyle_, color int, bColor bool) int {]
ff=编辑框_添加样式EX
fontStyle=字体样式
fontSize=字体大小
fontName=字体名称
hEle=元素句柄

[func XEdit_GetStyleInfo(hEle int, iStyle int, info *Edit_Style_Info_) bool {]
ff=编辑框_取样式信息
info=返回样式信息
iStyle=样式索引
hEle=元素句柄

[func XEdit_SetCurStyle(hEle int, iStyle int) int {]
ff=编辑框_置当前样式
iStyle=样式索引
hEle=元素句柄

[func XEdit_SetCaretColor(hEle int, color int) int {]
ff=编辑框_置插入符颜色
color=颜色
hEle=元素句柄

[func XEdit_SetCaretWidth(hEle int, nWidth int) int {]
ff=编辑框_置插入符宽度
nWidth=宽度
hEle=元素句柄

[func XEdit_SetSelectBkColor(hEle int, color int) int {]
ff=编辑框_置选择背景颜色
color=ABGR颜色
hEle=元素句柄

[func XEdit_SetRowHeight(hEle int, nHeight int) int {]
ff=编辑框_置默认行高
nHeight=行高
hEle=元素句柄

[func XEdit_SetRowHeightEx(hEle int, iRow int, nHeight int) int {]
ff=编辑框_置指定行高度
nHeight=高度
iRow=行索引
hEle=元素句柄

[func XEdit_SetCurPos(hEle int, iRow int) int {]
ff=编辑框_置当前位置
iRow=行索引
hEle=元素句柄

[func XEdit_GetCurPos(hEle int) int {]
ff=编辑框_取当前位置点
hEle=元素句柄

[func XEdit_GetCurRow(hEle int) int {]
ff=编辑框_取当前行
hEle=元素句柄

[func XEdit_GetCurCol(hEle int) int {]
ff=编辑框_取当前列
hEle=元素句柄

[func XEdit_GetPoint(hEle int, iRow int, iCol int, pOut *POINT) int {]
ff=编辑框_取坐标点
pOut=接收返回坐标点
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_AutoScroll(hEle int) bool {]
ff=编辑框_自动滚动
hEle=元素句柄

[func XEdit_AutoScrollEx(hEle int, iRow int, iCol int) bool {]
ff=编辑框_自动滚动EX
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_PosToRowCol(hEle int, iPos int, pInfo *Position_) int {]
ff=编辑框_转换位置
pInfo=行列
iPos=位置点
hEle=元素句柄

[func XEdit_SelectAll(hEle int) bool {]
ff=编辑框_选择全部
hEle=元素句柄

[func XEdit_CancelSelect(hEle int) bool {]
ff=编辑框_取消选择
hEle=元素句柄

[func XEdit_DeleteSelect(hEle int) bool {]
ff=编辑框_删除选择内容
hEle=元素句柄

[func XEdit_SetSelect(hEle int, iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {]
ff=编辑框_置选择
iEndCol=结束行列索引
iEndRow=结束行索引
iStartCol=起始行列索引
iStartRow=起始行索引
hEle=元素句柄

[func XEdit_GetSelectText(hEle int, pOut *string, nOutLen int) int {]
ff=编辑框_取选择文本
nOutLen=接收内存大小
pOut=接收返回文本内容
hEle=元素句柄

[func XEdit_GetSelectRange(hEle int, pBegin *Position_, pEnd *Position_) bool {]
ff=编辑框_取选择内容范围
pEnd=结束位置
pBegin=起始位置
hEle=元素句柄

[func XEdit_GetVisibleRowRange(hEle int, piStart *int32, piEnd *int32) int {]
ff=编辑框_取可视行范围
piEnd=结束行索引
piStart=起始行索引
hEle=元素句柄

[func XEdit_Delete(hEle int, iStartRow int, iStartCol int, iEndRow int, iEndCol int) bool {]
ff=编辑框_删除
iEndCol=结束行列索引
iEndRow=结束行索引
iStartCol=起始行列索引
iStartRow=起始行索引
hEle=元素句柄

[func XEdit_DeleteRow(hEle int, iRow int) bool {]
ff=编辑框_删除行
iRow=行索引
hEle=元素句柄

[func XEdit_ClipboardCut(hEle int) bool {]
ff=编辑框_剪贴板剪切
hEle=元素句柄

[func XEdit_ClipboardCopy(hEle int) bool {]
ff=编辑框_剪贴板复制
hEle=元素句柄

[func XEdit_ClipboardPaste(hEle int) bool {]
ff=编辑框_剪贴板粘贴
hEle=元素句柄

[func XEdit_Undo(hEle int) bool {]
ff=编辑框_撤销
hEle=元素句柄

[func XEdit_Redo(hEle int) bool {]
ff=编辑框_恢复
hEle=元素句柄

[func XEdit_AddChatBegin(hEle int, hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {]
ff=编辑框_添加气泡开始
nFlag=标志
hImageBubble=气泡背景
hImageAvatar=头像
hEle=元素句柄

[func XEdit_AddChatEnd(hEle int) int {]
ff=编辑框_添加气泡结束
hEle=元素句柄

[func XEdit_SetChatIndentation(hEle int, nIndentation int) int {]
ff=编辑框_置气泡缩进
nIndentation=缩进值
hEle=元素句柄

[func XEdit_SetRowSpace(hEle int, nSpace int) int {]
ff=编辑框_置行间隔
nSpace=行间隔大小
hEle=元素句柄

[func XEdit_SetCurPosEx(hEle int, iRow, iCol int32) int {]
ff=编辑框_置当前位置EX
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_GetCurPosEx(hEle int, iRow, iCol *int32) int {]
ff=编辑框_取当前位置EX
iCol=返回列索引
iRow=返回行索引
hEle=元素句柄

[func XEdit_MoveEnd(hEle int) int {]
ff=编辑框_移动到末尾
hEle=元素句柄

[func XEdit_RowColToPos(hEle int, iRow int, iCol int) int {]
ff=编辑框_行列到位置
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_SetBackFont(hEle int, hFont int) int {]
ff=编辑框_置后备字体
hFont=字体
hEle=元素句柄

[func XEdit_ReleaseStyle(hEle int, iStyle int) bool {]
ff=编辑框_释放样式
iStyle=样式
hEle=元素句柄

[func XEdit_ModifyStyle(hEle int, iStyle int, hFont int, color int, bColor bool) bool {]
ff=编辑框_修改样式
bColor=是否使用颜色
color=ABGR颜色
hFont=字体句柄
iStyle=样式索引
hEle=元素句柄

[func XEdit_SetSpaceSize(hEle int, size int) int {]
ff=编辑框_置空格大小
size=空格大小
hEle=元素句柄

[func XEdit_SetCharSpaceSize(hEle int, size int, sizeZh int) int {]
ff=编辑框_置字符间距
sizeZh=中文字符
size=英文字符
hEle=元素句柄

[func XEdit_GetSelectTextLength(hEle int) int {]
ff=编辑框_取选择文本长度
hEle=元素句柄

[func XEdit_SetSelectTextStyle(hEle int, iStyle int) int {]
ff=编辑框_置选择文本样式
iStyle=样式索引
hEle=元素句柄

[func XEdit_GetText_Temp(hEle int) string {]
ff=编辑框_取文本Tmp
hEle=元素句柄

[func XEdit_GetTextRow_Temp(hEle int, iRow int) string {]
ff=编辑框_取文本行Tmp
iRow=行索引
hEle=元素句柄

[func XEdit_GetSelectText_Temp(hEle int) string {]
ff=编辑框_取选择文本Tmp
hEle=元素句柄

[func XEdit_InsertChatBegin(hEle int, hImageAvatar int, hImageBubble int, nFlag xcc.Chat_Flag_) int {]
ff=编辑框_插入气泡开始
nFlag=标志
hImageBubble=气泡背景图片句柄
hImageAvatar=头像图片句柄
hEle=元素句柄

[func XEdit_GetChatFlags(hEle int, iRow int) xcc.Chat_Flag_ {]
ff=编辑框_取指定行气泡标识
iRow=行索引
hEle=元素句柄

[func XEdit_InsertTextEx(hEle int, iRow int, iCol int, pString string, iStyle int) int {]
ff=编辑框_插入文本EX
iStyle=样式
pString=字符串
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_InsertObject(hEle int, iRow int, iCol int, hObj int) int {]
ff=编辑框_插入对象
hObj=对象句柄
iCol=列索引
iRow=行索引
hEle=元素句柄

[func XEdit_SetChatMaxWidth(hEle int, nWidth int32) {]
ff=编辑框_置气泡最大宽度
nWidth=最大宽度
hEle=元素句柄
