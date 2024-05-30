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

[func NewEditor(x int, y int, cx int, cy int, hParent int) *Editor {]
ff=创建代码编辑框
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=元素y坐标
x=元素x坐标

[func NewEditorByHandle(handle int) *Editor {]
ff=创建代码编辑框并按句柄

[func NewEditorByName(name string) *Editor {]
ff=创建代码编辑框并按名称

[func NewEditorByUID(nUID int) *Editor {]
ff=创建代码编辑框并按UID

[func NewEditorByUIDName(name string) *Editor {]
ff=创建代码编辑框并按UID名称

[func (e *Editor) EnableAutoMatchSpaseSelect(bEnable bool) int {]
ff=启用空格选择自动匹配项
bEnable=是否启用

[func (e *Editor) IsBreakpoint(iRow int) bool {]
ff=判断断点
iRow=行索引

[func (e *Editor) SetBreakpoint(iRow int, bActivate bool) bool {]
ff=置断点
bActivate=是否激活
iRow=行索引

[func (e *Editor) RemoveBreakpoint(iRow int) bool {]
ff=移除断点
iRow=行索引

[func (e *Editor) ClearBreakpoint() int {]
ff=清空断点

[func (e *Editor) SetRunRow(iRow int) bool {]
ff=置当前运行
iRow=行索引

[func (e *Editor) GetColor(pInfo *xc.Editor_Color_) int {]
ff=取颜色信息
pInfo=颜色信息结构体指针

[func (e *Editor) SetColor(pInfo *xc.Editor_Color_) int {]
ff=置颜色
pInfo=颜色信息结构体指针

[func (e *Editor) SetStyleKeyword(iStyle int) int {]
ff=置常量样式
iStyle=样式

[func (e *Editor) SetStyleFunction(iStyle int) int {]
ff=置函数样式
iStyle=样式

[func (e *Editor) SetStyleVar(iStyle int) int {]
ff=置变量样式
iStyle=样式

[func (e *Editor) SetStyleDataType(iStyle int) int {]
ff=置数据类型样式
iStyle=样式

[func (e *Editor) SetStyleClass(iStyle int) int {]
ff=置类样式
iStyle=样式

[func (e *Editor) SetStyleMacro(iStyle int) int {]
ff=置宏样式
iStyle=样式

[func (e *Editor) SetStyleString(iStyle int) int {]
ff=置字符串样式
iStyle=样式

[func (e *Editor) SetStyleComment(iStyle int) int {]
ff=置注释样式
iStyle=样式

[func (e *Editor) SetStyleNumber(iStyle int) int {]
ff=置数字样式
iStyle=样式

[func (e *Editor) GetBreakpointCount() int {]
ff=取断点数量

[func (e *Editor) GetBreakpoints(aPoints *#左中括号##右中括号#int32, nCount int) int {]
ff=取全部断点
nCount=切片大小
aPoints=接收断点切片

[func (e *Editor) SetCurRow(iRow int) int {]
ff=设置当前行
iRow=行索引

[func (e *Editor) GetDepth(iRow int) int {]
ff=获取深度
iRow=行索引

[func (e *Editor) ToExpandRow(iRow int) int {]
ff=转换到展开行
iRow=行索引

[func (e *Editor) ExpandEx(iRow int) int {]
ff=展开EX
iRow=行索引

[func (e *Editor) ExpandAll(bExpand bool) int {]
ff=展开全部
bExpand=是否展开

[func (e *Editor) Expand(iRow int, bExpand bool) int {]
ff=展开指定行
bExpand=是否展开
iRow=行索引

[func (e *Editor) AddKeyword(pKey string, iStyle int) int {]
ff=添加关键字
iStyle=样式
pKey=字符串

[func (e *Editor) AddConst(pKey string) int {]
ff=添加自动匹配字符串
pKey=字符串

[func (e *Editor) AddFunction(pKey string) int {]
ff=添加自动匹配函数
pKey=字符串

[func (e *Editor) AddExcludeDefVarKeyword(pKeyword string) int {]
ff=添加排除定义变量关键字
pKeyword=字符串

[func (e *Editor) GetExpandState() string {]
ff=获取折叠状态

[func (e *Editor) SetExpandState(pString string) int {]
ff=设置折叠状态

[func (e *Editor) GetIndentation(iRow int) int {]
ff=获取缩进
iRow=行

[func (e *Editor) IsEmptyRow(iRow int) int {]
ff=是否为空行
iRow=行

[func (e *Editor) SetAutoMatchMode(mode int) int {]
ff=置自动匹配结果显示模式
mode=模式

[func (e *Editor) Event_EDITOR_MODIFY_ROWS(pFun XE_EDITOR_MODIFY_ROWS) bool {]
ff=事件_多行内容改变

[func (e *Editor) Event_EDITOR_MODIFY_ROWS1(pFun XE_EDITOR_MODIFY_ROWS1) bool {]
ff=事件_多行内容改变1

[func (e *Editor) Event_EDITOR_SETBREAKPOINT(pFun XE_EDITOR_SETBREAKPOINT) bool {]
ff=事件_设置断点

[func (e *Editor) Event_EDITOR_SETBREAKPOINT1(pFun XE_EDITOR_SETBREAKPOINT1) bool {]
ff=事件_设置断点1

[func (e *Editor) Event_EDITOR_REMOVEBREAKPOINT(pFun XE_EDITOR_REMOVEBREAKPOINT) bool {]
ff=事件_移除断点

[func (e *Editor) Event_EDITOR_REMOVEBREAKPOINT1(pFun XE_EDITOR_REMOVEBREAKPOINT1) bool {]
ff=事件_移除断点1

[func (e *Editor) Event_EDITOR_AUTOMATCH_SELECT(pFun XE_EDITOR_AUTOMATCH_SELECT) bool {]
ff=事件_自动匹配选择
[func (e *Editor) Event_EDITOR_AUTOMATCH_SELECT1(pFun XE_EDITOR_AUTOMATCH_SELECT1) bool {]
ff=事件_自动匹配选择1
