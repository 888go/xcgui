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

[func XEditor_Create(x int, y int, cx int, cy int, hParent int) int {]
ff=代码编辑框_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func XEidtor_EnableAutoMatchSpaseSelect(hEle int, bEnable bool) int {]
ff=代码编辑框_启用空格选择自动匹配项
bEnable=是否启用
hEle=元素句柄

[func XEditor_IsBreakpoint(hEle int, iRow int) bool {]
ff=代码编辑框_判断断点
iRow=行索引
hEle=元素句柄

[func XEditor_SetBreakpoint(hEle int, iRow int, bActivate bool) bool {]
ff=代码编辑框_置断点
bActivate=是否激活
iRow=行索引
hEle=元素句柄

[func XEditor_RemoveBreakpoint(hEle int, iRow int) bool {]
ff=代码编辑框_移除断点
iRow=行索引
hEle=元素句柄

[func XEditor_ClearBreakpoint(hEle int) int {]
ff=代码编辑框_清空断点
hEle=元素句柄

[func XEditor_SetRunRow(hEle int, iRow int) bool {]
ff=代码编辑框_置当前运行
iRow=行索引
hEle=元素句柄

[func XEditor_GetColor(hEle int, pInfo *Editor_Color_) int {]
ff=代码编辑框_取颜色信息
pInfo=颜色信息结构体指针
hEle=元素句柄

[func XEditor_SetColor(hEle int, pInfo *Editor_Color_) int {]
ff=代码编辑框_置颜色
pInfo=颜色信息结构体指针
hEle=元素句柄

[func XEditor_SetStyleKeyword(hEle int, iStyle int) int {]
ff=代码编辑框_置常量样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleFunction(hEle int, iStyle int) int {]
ff=代码编辑框_置函数样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleVar(hEle int, iStyle int) int {]
ff=代码编辑框_置变量样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleDataType(hEle int, iStyle int) int {]
ff=代码编辑框_置数据类型样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleClass(hEle int, iStyle int) int {]
ff=代码编辑框_置类样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleMacro(hEle int, iStyle int) int {]
ff=代码编辑框_置宏样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleString(hEle int, iStyle int) int {]
ff=代码编辑框_置字符串样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleComment(hEle int, iStyle int) int {]
ff=代码编辑框_置注释样式
iStyle=样式
hEle=元素句柄

[func XEditor_SetStyleNumber(hEle int, iStyle int) int {]
ff=代码编辑框_置数字样式
iStyle=样式
hEle=元素句柄

[func XEditor_GetBreakpointCount(hEle int) int {]
ff=代码编辑框_取断点数量
hEle=元素句柄

[func XEditor_GetBreakpoints(hEle int, aPoints *#左中括号##右中括号#int32, nCount int) int {]
ff=代码编辑框_取全部断点
nCount=切片大小
aPoints=接收断点切片
hEle=元素句柄

[func XEditor_SetCurRow(hEle int, iRow int) int {]
ff=代码编辑框_设置当前行
iRow=行索引
hEle=元素句柄

[func XEditor_GetDepth(hEle int, iRow int) int {]
ff=代码编辑框_获取深度
iRow=行索引
hEle=元素句柄

[func XEditor_ToExpandRow(hEle int, iRow int) int {]
ff=代码编辑框_转换到展开行
iRow=行索引
hEle=元素句柄

[func XEditor_ExpandEx(hEle int, iRow int) int {]
ff=代码编辑框_展开EX
iRow=行索引
hEle=元素句柄

[func XEditor_ExpandAll(hEle int, bExpand bool) int {]
ff=代码编辑框_展开全部
bExpand=是否展开
hEle=元素句柄

[func XEditor_Expand(hEle int, iRow int, bExpand bool) int {]
ff=代码编辑框_展开指定行
bExpand=是否展开
iRow=行索引
hEle=元素句柄

[func XEditor_AddKeyword(hEle int, pKey string, iStyle int) int {]
ff=代码编辑框_添加关键字
iStyle=样式
pKey=字符串
hEle=元素句柄

[func XEditor_AddConst(hEle int, pKey string) int {]
ff=代码编辑框_添加自动匹配字符串
pKey=字符串
hEle=元素句柄

[func XEditor_AddFunction(hEle int, pKey string) int {]
ff=代码编辑框_添加自动匹配函数
pKey=字符串
hEle=元素句柄

[func XEditor_AddExcludeDefVarKeyword(hEle int, pKeyword string) int {]
ff=代码编辑框_添加排除定义变量关键字
pKeyword=字符串
hEle=元素句柄

[func XEditor_GetExpandState(hEle int) string {]
ff=代码编辑框_获取折叠状态
hEle=元素句柄

[func XEditor_SetExpandState(hEle int, pString string) int {]
ff=代码编辑框_设置折叠状态
hEle=元素句柄

[func XEditor_GetIndentation(hEle int, iRow int) int {]
ff=代码编辑框_获取缩进
iRow=行
hEle=元素句柄

[func XEidtor_IsEmptyRow(hEle int, iRow int) int {]
ff=代码编辑框_是否为空行
iRow=行
hEle=元素句柄

[func XEditor_SetAutoMatchMode(hEle int, mode int) int {]
ff=代码编辑框_置自动匹配结果显示模式
mode=模式
hEle=元素句柄
